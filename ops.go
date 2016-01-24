package gitcli


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
