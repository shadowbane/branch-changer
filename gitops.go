package main

import (
	"bytes"
	"go.uber.org/zap"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func getBranch(project string) string {
	zap.S().Infof("Running task: getBranch for %s", project)
	chDir(project)
	zap.S().Debugf("%s Current directory: %s", project, getCwd())
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		zap.S().Panicf("Error getting branch: %s", err)
	}

	// trim new line from destination
	output := strings.TrimSuffix(string(out), "\n")

	return output
}

func fetchRepo(project string) {
	zap.S().Infof("Running task: fetchRepo for %s", project)
	chDir(project)
	zap.S().Debugf("%s Current directory: %s", project, getCwd())
	zap.S().Debugf("%s Fetching repository...", project)
	cmd := exec.Command("git", "fetch", "--all")
	_, err := cmd.Output()
	if err != nil {
		zap.S().Panicf("Error fetching repo: %s", err)
	}
}

func stashChanges(project string) {
	zap.S().Infof("Running task: stashChanges for %s", project)
	stashName := "branch-switcher-" + strconv.FormatInt(time.Now().Unix(), 10)

	chDir(project)
	zap.S().Debugf("%s Stashing repository...", project)
	cmd := exec.Command("git", "stash", "push", "-m", stashName)

	var outputBuffer bytes.Buffer
	cmd.Stderr = &outputBuffer
	_, err := cmd.Output()

	if err != nil {
		zap.S().Errorf("Error stashing repo: %s", err)
		zap.S().Panicf("Error stashing repo: %s", outputBuffer.String())
	}
}

func pullRepo(project string) {
	zap.S().Infof("Running task: pullRepo for %s", project)
	chDir(project)
	zap.S().Debugf("%s Current directory: %s", project, getCwd())
	zap.S().Debugf("%s Pulling repository...", project)
	cmd := exec.Command("git", "pull", "--quiet")
	_, err := cmd.Output()
	if err != nil {
		zap.S().Panicf("Error pulling repo: %s", err)
	}
}

func pushRepo(project string, destinationBranch string, isNew bool) {
	zap.S().Infof("Running task: pushRepo for %s", project)
	chDir(project)
	zap.S().Debugf("%s Current directory: %s", project, getCwd())

	zap.S().Debugf("%s Pushing repository to origin/%s ...", project, destinationBranch)

	var cmd = exec.Cmd{
		Path: "/usr/bin/git",
		Args: []string{"git", "push"},
	}

	if isNew {
		cmd.Args = append(cmd.Args, "--set-upstream")
		cmd.Args = append(cmd.Args, "origin")
		cmd.Args = append(cmd.Args, destinationBranch)
	}

	var output bytes.Buffer
	cmd.Stderr = &output

	_, err := cmd.Output()
	if err != nil {
		zap.S().Debugf("Error pushing repo: %s", err)
		zap.S().Panicf("Error pushing repo: %s", output.String())
	}
}

func branchExist(project string, branchName string) bool {
	zap.S().Infof("Running task: branchExist for %s", project)
	chDir(project)
	zap.S().Debugf("%s Current directory: %s", project, getCwd())

	cmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/heads/"+branchName)
	_, err := cmd.Output()
	if err != nil {
		return false
	}

	return true
}

func checkoutBranch(project string, branchName string) {
	zap.S().Infof("Running task: checkoutBranch for %s", project)
	chDir(project)
	zap.S().Debugf("%s Current directory: %s", project, getCwd())

	cmd := exec.Command("git", "checkout", branchName)
	_, err := cmd.Output()
	if err != nil {
		zap.S().Panicf("Error checking out branch: %s", err)
	}
}

func mergeRepo(project string, branchFrom string, branchTo string) {
	zap.S().Infof("Running task: mergeRepo for %s", project)
	chDir(project)
	zap.S().Debugf("%s Current directory: %s", project, getCwd())
	cmd := exec.Command("git", "merge", branchFrom, branchTo)

	var output bytes.Buffer
	cmd.Stderr = &output

	_, err := cmd.Output()
	if err != nil {
		zap.S().Panicf("%s Error merging repo: %s", project, output.String())
	}
}

func setTrackedBranch(project string, branchName string) {
	zap.S().Infof("Running task: setTrackedBranch for %s", project)
	chDir(project)
	zap.S().Debugf("%s Current directory: %s", project, getCwd())
	cmd := exec.Command("git", "branch", "--set-upstream-to=origin/"+branchName, branchName)

	var output bytes.Buffer
	cmd.Stderr = &output

	_, err := cmd.Output()
	if err != nil {
		zap.S().Panicf("%s Error tracking repo: %s", project, output.String())
	}
}

func createNewBranch(project string, branchName string) {
	zap.S().Infof("Running task: createNewBranch for %s", project)
	chDir(project)
	zap.S().Debugf("%s Current directory: %s", project, getCwd())

	cmd := exec.Command("git", "checkout", "-b", branchName)
	_, err := cmd.Output()
	if err != nil {
		zap.S().Panicf("Error creating new branch: %s", err)
	}
}

func checkTrackedBranch(project string, branchName string) bool {
	zap.S().Infof("Running task: checkTrackedBranch for %s", project)
	chDir(project)
	zap.S().Debugf("%s Current directory: %s", project, getCwd())
	cmd1 := exec.Command("git", "rev-parse", "--quiet", "--abbrev-ref", "--symbolic-full-name", "@{u}")
	cmd2 := exec.Command("wc", "-l")

	var outputBuffer bytes.Buffer
	cmd2.Stdin, _ = cmd1.StdoutPipe()
	cmd2.Stdout = &outputBuffer

	_ = cmd2.Start()
	_ = cmd1.Run()
	_ = cmd2.Wait()

	output := strings.TrimSuffix(outputBuffer.String(), "\n")
	outputInt, _ := strconv.ParseInt(output, 10, 64)

	return outputInt > 0
}
