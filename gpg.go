package gitcli

import (
	"regexp"
	"strings"
)

var gitGpgRegex = regexp.MustCompile(`^(.+)\|(.+)$`)
// verify a given commit's GPG sig
// anything that refers to commit will work (like HEAD~3)
func (r *Repo) GetCommitSignature(commit string) (key  string, signed bool, err error) {
	stdout, _, err := r.cmd(`log`, `--format=%G?|%GK`,`-1`, commit, `--`)
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

// SetTrustedSignatures sets a list of signatures considered valid for VerifyCommit()
// USE LONG FORMAT (full fingerprint), for now git log only passes 16 characters of fingerprint
// but if that ever changes any sig that will be shorter than that won't be matched
func (r *Repo) SetTrustedSignatures(sigs []string) {
	r.trustedSigs = make(map[string]bool)
	for _, sig := range sigs {
		cleanedSig := strings.ToUpper(strings.Replace(sig, " ", "", -1))
		r.trustedSigs[cleanedSig] = true
	}
}

// VerifyCommit checks if given commit is signed by one of sigs set in SetTrustedSignatures
// git log only passes 16 characters of fingerprint so it only checks for substring of that signature
func (r *Repo) VerifyCommit (commit string) (bool, error){
	sigID, correct, err := r.GetCommitSignature(commit)
	if err != nil {return false, err}
	if !correct {return false, err}
	// hopefully when git changes it to provide full fingerprint it can be changed to just hash lookup
	for testedSig := range r.trustedSigs {
		if strings.Contains(testedSig, sigID) {
			return true, err
		}
	}
	return false, err
}
