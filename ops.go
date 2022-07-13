package gitcli

import (
	"fmt"
	"os"
	"strings"
)

func (r *Repo) Push(extraArgs ...string) error {
	args := []string{`push`}
	return r.commonCmd(args, extraArgs...)
}

func (r *Repo) Pull(extraArgs ...string) error {
	args := []string{`pull`}
	return r.commonCmd(args, extraArgs...)
}

func (r *Repo) Fetch(extraArgs ...string) error {
	args := []string{`fetch`}
	return r.commonCmd(args, extraArgs...)
}

func (r *Repo) Checkout(extraArgs ...string) error {
	args := []string{`checkout`}
	return r.commonCmd(args, extraArgs...)
}

func (r *Repo) Log(extraArgs ...string) error {
	args := []string{`log`}
	return r.commonCmd(args, extraArgs...)
}

func (r *Repo) Clean(extraArgs ...string) error {
	args := []string{`clean`}
	return r.commonCmd(args, extraArgs...)
}

// reset and remove index
// also gets rid of "index smaller than expected" problem
func (r *Repo) HardReset() error {
	os.Remove(r.gitDir + "/index")
	return r.commonCmd([]string{"reset", "--hard"})
}

// Add or set a given remote
func (r *Repo) SetRemote(remoteName string, remoteUrl string) error {
	stdout, stderr, err := r.cmd("remote", "set-url", remoteName, remoteUrl)
	if strings.Contains(stderr, "No such remote") {
		err = r.commonCmd([]string{"remote", "add"}, remoteName, remoteUrl)
		return err
	} else if err != nil {
		return fmt.Errorf("set remote failed with stdout:%s , stderr: %s", stdout, stderr)
	}
	return err
}
