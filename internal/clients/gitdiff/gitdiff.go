package gitdiff

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

const unknownCommitHash = "0000000000000000000000000000000000000000"

// GitDiff client for getting diffs from the command line
type GitDiff struct {
	WorkDir    string
	BaseBranch string
	BaseSHA    string
	Head       string
}

func printFiles(dir string) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("Error in accessing directory:%s", err.Error())
	}

	for _, file := range fileInfos {
		log.Printf(file.Name())
	}
}

// GetDiff uses the command line to compute the diff
func (gd *GitDiff) GetDiff() (string, error) {
	err := os.Chdir(gd.WorkDir)
	var diffCmd *exec.Cmd
	switch {
	case gd.BaseBranch != "":
		// PR event so get diff between base branch and current commit SHA
		gitCmd := exec.Command("git", "-c", "http.sslVerify=false", "fetch", "origin", gd.BaseBranch, "--depth=1")
		err = gitCmd.Run()
		if err != nil {
			return "", err
		}
		diffCmd = exec.Command("git", "diff", fmt.Sprintf("origin/%s", gd.BaseBranch))
	case gd.BaseSHA == "" || gd.BaseSHA == unknownCommitHash:
		// PUSH event for new branch so use git show to get the diff of the most recent commit
		err = exec.Command("git", "fetch", "origin", gd.Head, "--depth=2").Run()
		if err != nil {
			return "", err
		}
		diffCmd = exec.Command("git", "show", gd.Head, "--format=")
	default:
		// PUSH event where last commit action ran on exists
		// use current commit SHA and previous action run commit SHA for diff
		err = exec.Command("git", "fetch", "origin", gd.BaseSHA, "--depth=1").Run()
		if err != nil {
			return "", err
		}
		diffCmd = exec.Command("git", "diff", gd.BaseSHA, gd.Head)
	}

	reader, err := diffCmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	defer reader.Close()

	err = diffCmd.Start()
	if err != nil {
		return "", nil
	}
	buf := new(strings.Builder)
	_, err = io.Copy(buf, reader)
	if err != nil {
		return "", err
	}
	err = diffCmd.Wait()
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
