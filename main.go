package main

import (
	"bufio"
	"flag"
	"github.com/samber/oops"
	"github.com/shadowbane/go-logger"
	"go.uber.org/zap"
	"os"
	"strings"
)

var projects []string
var notFoundProjects []string
var errProjects []string
var workingDir = ""
var branchFrom = ""
var branchTo = ""
var push = false
var force = false
var merge = false

var Version = "development"
var Maintainer = "Adli I. Ifkar <adly.shadowbane@gmail.com>"

func main() {
	// set environment variables
	_ = os.Setenv("LOG_FILE_ENABLED", "true")

	logger.Init(logger.LoadEnvForLogger())

	printHeader()
	parseFlags()

	zap.S().Info("Starting application...")

	for _, project := range projects {
		err := os.Chdir(workingDir + "/" + project)
		if err != nil {
			notFoundProjects = append(notFoundProjects, project)
			zap.S().Warnf("%s/%s does not exists! Please make sure the directory exists before continuing", workingDir, project)
		} else {
			switchBranch(project)
		}
	}

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
	var err error

	// get input for branchFrom
	inputBranchFrom := flag.String("source", "", "Branch to merge from")
	inputBranchTo := flag.String("destination", "", "Branch to merge to")
	inputWorkDir := flag.String("workdir", "", "Working directory")
	inputProjects := flag.String("projects", "", "Projects to merge")

	inputLoadFailedOnly := flag.Bool("failedonly", false, "Load failed projects only")
	inputWithPush := flag.Bool("push", false, "Push newly create branch to remote")
	inputForce := flag.Bool("force", false, "Force switch branch (stash changes)")
	inputMerge := flag.Bool("merge", false, "Merge 'source' to 'destination' (default: true)")

	flag.Parse()

	if *inputBranchFrom != "" {
		branchFrom = *inputBranchFrom
	} else {
		if branchFrom == "" {
			err = oops.
				In("ParseFlags").
				Hint("Please use --from to specify branch from").
				Errorf("branch from is required")
		}
	}

	if *inputBranchTo != "" {
		branchTo = *inputBranchTo
	} else {
		if branchTo == "" {
			err = oops.
				In("ParseFlags").
				Hint("Please use --to to specify branch to").
				Errorf("branch to is required")
		}
	}

	if *inputWorkDir != "" {
		workingDir = *inputWorkDir
	} else {
		if workingDir == "" {
			err = oops.
				In("ParseFlags").
				Hint("Please use --workdir to specify working directory").
				Errorf("working directory not specified")
		}
	}

	if *inputProjects != "" {
		// separate projects by comma
		projects = strings.Split(*inputProjects, ",")
	} else {
		if len(projects) == 0 {
			err = oops.
				In("ParseFlags").
				Hint("Please use --projects to specify project. Example: directory-a,some-directory").
				Errorf("projects not specified")
		}
	}

	if *inputLoadFailedOnly {
		// load failed projects from file
		projects = loadFailedProjects(workingDir + "/failed.txt")
	}

	if err != nil {
		zap.S().Errorf("[%s] %s", err.(oops.OopsError).Domain(), err.(oops.OopsError).Stacktrace())
		zap.S().Infof("Hint: %s", err.(oops.OopsError).Hint())

		os.Exit(2)
	}

	push = *inputWithPush
	force = *inputForce
	merge = *inputMerge

	zap.S().Infof("Branch from: %s", branchFrom)
	zap.S().Infof("Branch to: %s", branchTo)
	zap.S().Infof("Working directory: %s", workingDir)
	zap.S().Infof("Projects: %s", projects)

	if push {
		zap.S().Warn("Application set to push to remote repository")
	}

	if force {
		zap.S().Warn("Application set to stash changes before switching branch")
	}

	if merge {
		zap.S().Warn("Application set to merge 'source' to 'destination' branch")
	}
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

func switchBranch(projectName string) {
	defer handleErrors(projectName)

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
		checkoutBranch(projectName, branchFrom)

		pullRepo(projectName)

		checkoutBranch(projectName, branchTo)

		if checkTrackedBranch(projectName, branchTo) {
			setTrackedBranch(projectName, branchTo)
			pullRepo(projectName)
		}

		if merge {
			mergeRepo(projectName, branchFrom, branchTo)
		}

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

func handleErrors(projectName string) {
	if r := recover(); r != nil {
		errProjects = append(errProjects, projectName)
		zap.S().Debugf("Recovered from error")
		//zap.S().Debugf("Recovered from %s", r)
	}
}
