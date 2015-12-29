package gitcli

func (r *Repo) Version() (string, error) {
	stdout, stderr, err := r.cmd(`version`)
	return stdout + stderr, err
}
