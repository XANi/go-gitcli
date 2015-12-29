package gitcli

var gitCmd = "git"

type Repo struct {
	workDir string
	gitDir string

}

func New(repoDir string, args ...string) Repo {
	var r Repo
	r.workDir = repoDir
	if (len(args) > 0) {
		r.gitDir = args[0]
	} else {
		r.gitDir = repoDir
		r.gitDir += `./.git`
	}
	return r
}

func (r *Repo) Init() error {
	var err error
	return err
}
