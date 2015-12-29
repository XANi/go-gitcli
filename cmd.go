package gitcli

import (
//	"io"
	"os/exec"
	"bytes"
//	"fmt"
)


func (r *Repo) cmd(args ...string) ( string, string,  error) {
	var err error
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var cmdArgs []string
	cmdArgs = append(cmdArgs, `--work-tree` , r.workDir, `--git-dir` , r.gitDir)
	cmdArgs = append(cmdArgs, args...)
//	res = exec.Command(gitCmd, cmdArgs...)
//	cmd := exec.Command(`git` ,`version`)
	cmd := exec.Command(`git` ,cmdArgs...)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		return "", "", err
	}
	if err := cmd.Wait(); err != nil {
				return "", "", err
	}
	return stdout.String(), stderr.String(), err
}
