package gitcli

import (
	"strings"
	"fmt"
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
	args := []string{`fetch`}
	return r.commonCmd(args, extraArgs...)
}

func (r *Repo)Checkout(extraArgs ...string) error {
	args := []string{`push`}
	return r.commonCmd(args, extraArgs...)
}

// Add or set a given remote
func (r *Repo)SetRemote(remoteName string, remoteUrl string ) error {
	stdout, stderr, err := r.cmd("remote","set-url",remoteName,remoteUrl)
	if strings.Contains(stderr,"No such remote") {
		err = r.commonCmd([]string{"remote","add"},remoteName,remoteUrl)
		return err
	} else if err != nil {
		return fmt.Errorf("set remote failed with stdout:%s , stderr: %s", stdout, stderr)
	}
	return err
}
