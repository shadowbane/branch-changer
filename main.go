package main

import (
	"bufio"
	"flag"
	"github.com/shadowbane/go-logger"
	"go.uber.org/zap"
	"os"
	"strings"
	"sync"
)

var projects = []string{
	"clickargo-docker-compose-config",
	"clickargo-file-manager",
	"clickargo-id-tds-finance",
	"clickargo-microservice-bank-va",
	"clickargo-microservice-company",
	"clickargo-microservice-company-api",
	"clickargo-microservice-container",
	"clickargo-microservice-dashboard",
	"clickargo-microservice-domestic",
	"clickargo-microservice-excel-csv-uploader",
	"clickargo-microservice-export",
	"clickargo-microservice-finance",
	"clickargo-microservice-gateway",
	"clickargo-microservice-import",
	"clickargo-microservice-inttra",
	"clickargo-microservice-logs",
	"clickargo-microservice-order",
	"clickargo-microservice-shell",
	"clickargo-microservice-shipment",
	"clickargo-microservice-tds",
	"clickargo-microservice-terminal-operator",
	"clickargo-microservice-truck",
	"clickargo-view-blade",
	"clickargo-view-trucking",
}
var notFoundProjects []string
var errProjects []string
var workingDir = "/var/www"
var branchFrom = "development-new"
var branchTo = "development-sg"
var push = false
var force = false

var Version = "development"
var Maintainer = "Adli I. Ifkar <adly.shadowbane@gmail.com>"

func main() {
	printHeader()
	parseFlags()

	// set environment variables
	_ = os.Setenv("LOG_FILE_ENABLED", "true")

	logger.Init(logger.LoadEnvForLogger())

	zap.S().Infof("Branch from: %s", branchFrom)
	zap.S().Infof("Branch to: %s", branchTo)
	zap.S().Infof("Working directory: %s", workingDir)
	zap.S().Infof("Projects: %s", projects)

	zap.S().Warn("Application set to push to remote repository")
	zap.S().Warn("Application set to stash changes before switching branch")

	zap.S().Info("Starting application...")

	// create wait group
	wg := sync.WaitGroup{}

	for _, project := range projects {
		err := os.Chdir(workingDir + "/" + project)
		if err != nil {
			notFoundProjects = append(notFoundProjects, project)
			zap.S().Warnf("%s/%s does not exists! Please make sure the directory exists before continuing", workingDir, project)
		} else {
			wg.Add(1)
			switchBranch(project, &wg)
		}
	}

	wg.Wait()

	// Print errors for not found projects
	if len(notFoundProjects) > 0 {
		zap.S().Errorf("%d directory does not exists: %s", len(notFoundProjects), notFoundProjects)
	}

	// Print errors for failed switch
	if len(errProjects) > 0 {
		zap.S().Errorf("Error while switching branch for %d project(s): %s", len(errProjects), errProjects)
	}

	// write to file at workingDir/failed.txt
	err := writeToFile(workingDir + "/failed.txt")
	if err != nil {
		zap.S().Errorf("Failed to write to file: %s", err)
	}

	zap.S().Info("Task Completed!")
}

func writeToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			zap.S().Errorf("Failed to close file: %s", err)
		}
	}(file)

	for _, d := range notFoundProjects {
		_, err := file.WriteString(d + "\n")
		if err != nil {
			return err
		}
	}

	for _, d := range errProjects {
		_, err := file.WriteString(d + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func parseFlags() {
	// get input for branchFrom
	inputBranchFrom := flag.String("from", "", "Branch to merge from")
	inputBranchTo := flag.String("to", "", "Branch to merge to")
	inputWorkDir := flag.String("workdir", "", "Working directory")
	inputProjects := flag.String("projects", "", "Projects to merge")

	inputLoadFailedOnly := flag.Bool("failedonly", false, "Load failed projects only")
	inputWithPush := flag.Bool("push", false, "Push newly create branch to remote")
	inputForce := flag.Bool("force", false, "Force switch branch (stash changes)")

	flag.Parse()

	if *inputBranchFrom != "" {
		branchFrom = *inputBranchFrom
	}

	if *inputBranchTo != "" {
		branchTo = *inputBranchTo
	}

	if *inputWorkDir != "" {
		workingDir = *inputWorkDir
	}

	if *inputProjects != "" {
		// separate projects by comma
		projects = strings.Split(*inputProjects, ",")
	}

	if *inputLoadFailedOnly {
		// load failed projects from file
		projects = loadFailedProjects(workingDir + "/failed.txt")
	}

	push = *inputWithPush
	force = *inputForce
}

func loadFailedProjects(filename string) []string {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		zap.S().Errorf("Failed to open file: %s", err)
		return []string{}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			zap.S().Errorf("Failed to close file: %s", err)
		}
	}(file)

	// read file
	scanner := bufio.NewScanner(file)
	var projects []string
	for scanner.Scan() {
		projects = append(projects, scanner.Text())
	}

	return projects
}

func switchBranch(projectName string, wg *sync.WaitGroup) {
	defer handleErrors(projectName, wg)

	chDir(projectName)
	zap.S().Infof("%s Current directory: %s", projectName, getCwd())

	// print current branch
	currBranch := getBranch(projectName)
	zap.S().Infof("%s Current branch: %s", projectName, currBranch)

	fetchRepo(projectName)

	if force {
		stashChanges(projectName)
	}

	if branchExist(projectName, branchTo) {
		if currBranch != branchTo {
			checkoutBranch(projectName, branchTo)
		}

		if checkTrackedBranch(projectName, branchTo) {
			setTrackedBranch(projectName, branchTo)
			pullRepo(projectName)
		}

		mergeRepo(projectName, branchFrom, branchTo)

		if push {
			pushRepo(projectName, branchTo, !checkTrackedBranch(projectName, branchTo))
		}
	} else {
		checkoutBranch(projectName, branchFrom)

		pullRepo(projectName)

		createNewBranch(projectName, branchTo)

		if push {
			pushRepo(projectName, branchTo, true)
		}
	}

	zap.S().Infof("%s/%s switched to %s", workingDir, projectName, branchTo)

	wg.Done()
}

func chDir(projectName string) {
	// change to the directory
	err := os.Chdir(workingDir + "/" + projectName)
	if err != nil {
		zap.S().Panicf("Error changing directory: %s", err)
	}
}

func getCwd() string {
	// print current directory
	currDir, err := os.Getwd()
	if err != nil {
		zap.S().Panicf("Error getting current directory: %s", err)
	}

	return currDir
}

func handleErrors(projectName string, wg *sync.WaitGroup) {
	if r := recover(); r != nil {
		errProjects = append(errProjects, projectName)
		zap.S().Debugf("Recovered from error")
		//zap.S().Debugf("Recovered from %s", r)
		wg.Done()
	}
}
