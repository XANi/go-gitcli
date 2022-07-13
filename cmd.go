package gitcli

import (
	"bytes"
	"fmt"
	//	"io"
	"os/exec"
)

func (r *Repo) cmd(args ...string) (string, string, error) {
	var err error
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var cmdArgs []string
	cmdArgs = append(cmdArgs, `--work-tree`, r.workDir, `--git-dir`, r.gitDir)
	cmdArgs = append(cmdArgs, args...)
	//	res = exec.Command(gitCmd, cmdArgs...)
	//	cmd := exec.Command(`git` ,`version`)
	// make commands consistent
	if r.debug {
		fmt.Printf("debug: executing git command: %+v\n", cmdArgs)
	}
	cmd := exec.Command(`git`, cmdArgs...)
	// make commands consistent
	cmd.Env = r.filteredEnv
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = r.workDir
	if r.debug {
		fmt.Printf("git command finished. stdout [%s], stderr [%s]\n", stdout.String(), stderr.String())
	}
	if err := cmd.Start(); err != nil {
		return "", "", err
	}
	if err := cmd.Wait(); err != nil {
		return stdout.String(), stderr.String(), err
	}
	return stdout.String(), stderr.String(), err
}
