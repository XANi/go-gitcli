package gitcli

import (
	"strings"
)

func (r *Repo)Push(extraArgs ...string) error {
	args := []string{`push`}
	return r.commonCmd(args, extraArgs...)
}


func (r *Repo)Pull(extraArgs ...string) error {
	args := []string{`pull`}
	return r.commonCmd(args, extraArgs...)
}

func (r *Repo)Fetch(extraArgs ...string) error {
	args := []string{`push`}
	return r.commonCmd(args, extraArgs...)
}

func (r *Repo)Checkout(extraArgs ...string) error {
	args := []string{`push`}
	return r.commonCmd(args, extraArgs...)
}

// Add or set a given remote
func (r *Repo)SetRemote(remoteName string, remoteUrl string ) error {
	_, stderr, err := r.cmd("remote","set-url",remoteName,remoteUrl)
	if strings.Contains(stderr,"No such remote") {
		err = r.commonCmd([]string{"remote","add"},remoteName,remoteUrl)
		return err
	}
	return err
}
