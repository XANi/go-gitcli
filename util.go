package gitcli

func (r *Repo) Version() (string, error) {
	stdout, _, err := r.cmd(`version`)
	return stdout, err
}
