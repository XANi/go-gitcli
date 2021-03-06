package gitcli


import (
	"fmt"
	"errors"
)


// common command wrappers for failed/success ops
func (r *Repo)commonCmd (cmd []string, extra ...string) error{
	cmdArgs := append(cmd,extra...)
	stdout, stderr, err := r.cmd(cmdArgs...)
	if err == nil {
		return err
	} else {
		return errors.New(
			fmt.Sprintf(
				"Error running {%+v} in repo[%s] with extra args {%+v}: error: %s, stdout: %s, stderr: %s",
				cmd,
				r.workDir,
				extra,
				err,
				stdout,
				stderr,			)		)
	}
}
