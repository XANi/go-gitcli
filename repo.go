package gitcli

import (
	"errors"
	"fmt"
	"strings"
	"os"
)

var gitCmd = "git"


type Repo struct {
	workDir string
	gitDir string
	filteredEnv []string

}

func New(repoDir string, args ...string) Repo {
	var r Repo
	r.workDir = repoDir
	if (len(args) > 0) {
		r.gitDir = args[0]
	} else {
		r.gitDir = repoDir
		r.gitDir += `/.git`
	}
	for _, v := range os.Environ() {
		// filter locale sillines so command output is always english
		if !strings.HasPrefix(v, "LANG=") && !strings.HasPrefix(v, "LC")  {
			r.filteredEnv = append(r.filteredEnv,v)
		}
	}
	return r
}

func (r *Repo) Init() error {
	var err error
	stdout, stderr, err := r.cmd(`init`)
	if err == nil {
		return err
	} else {
		return errors.New(fmt.Sprintf("Repo init error: %s, stdout: %s, stderr: %s",err,stdout,stderr))
	}

}
