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
	// **full signatures** of allowed commiters
	trustedSigs map[string]bool

}
// Create new repo object.
// First extra argument specifies git dir if it is different than standard ( $repo/.git )
// It will also initialize environment, removing few vars like LC_* or LANG so the output of git commands is consistent regardless of system's language

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
	r.trustedSigs = make(map[string]bool)
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
