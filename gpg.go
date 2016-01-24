package gitcli

import (
	"regexp"
	"strings"
)

var gitGpgRegex = regexp.MustCompile(`^(.+)\|.*using .*key ID (\S+)`)
// verify a given commit's GPG sig
// anything that refers to commit will work (like HEAD~3)
func (r *Repo) GetCommitSignature(commit string) (key  string, signed bool, err error) {
	stdout, _, err := r.cmd(`log`, `--format=%G?|%GG`,`-1`, commit, `--`)
	if (err != nil )  { return "", false, err }
	matches := gitGpgRegex.FindStringSubmatch(stdout)
	if (len(matches) < 3) {
		return "", false, err
	} else {
		// G - good, U - good, untrusted
		if (matches[1] == "G" || matches[1] == "U") {
			return matches[2], true, err
		} else {
			return matches[2], false, err
		}
	}
}

// set trusted sigs list for VerifySignature()
func (r *Repo) SetTrustedSignatures(sigs []string) {
	r.trustedSigs = make(map[string]bool)
	for _, sig := range sigs {
		r.trustedSigs[ strings.ToUpper(sig) ] = true
	}
}

func (r *Repo) VerifyCommit (commit string) (bool, error){
	sigID, correct, err := r.GetCommitSignature(commit)
	if err != nil {return false, err}
	if !correct {return false, err}
	if _, ok := r.trustedSigs[sigID]; ok {
		return true, err
	}
	return false, err
}
