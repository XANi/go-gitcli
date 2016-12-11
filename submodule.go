package gitcli

import (
)



func (r *Repo)SubmoduleInit(extraArgs ...string) error {
	args := []string{`submodule`,`init`}
	return r.commonCmd(args, extraArgs...)
}

func (r *Repo)SubmoduleSync(extraArgs ...string) error {
	args := []string{`submodule`,`sync`}
	return r.commonCmd(args, extraArgs...)
}

func (r *Repo)SubmoduleUpdate(extraArgs ...string) error {
	args := []string{`submodule`,`update`}
	return r.commonCmd(args, extraArgs...)
}
