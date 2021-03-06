package gitcli

import (
	"regexp"
	"strings"
	"fmt"
)

var gitGpgRegex = regexp.MustCompile(`^(.+)\|(.+)\s*$`)
// verify a given commit's GPG sig
// anything that refers to commit will work (like HEAD~3)
func (r *Repo) GetCommitSignature(commit string) (key  string, signed bool, err error) {
	stdout, stderr, err := r.cmd(`log`, `--format=%G?|%GK`,`-1`, commit, `--`)
	if (err != nil )  { return "", false, fmt.Errorf("Error while running git log: %s|%s|%s", stdout, stderr, err) }
	matches := gitGpgRegex.FindStringSubmatch(stdout)
	if (len(matches) < 3) {
		return "", false, fmt.Errorf("couldn't match git log output: %s | %+v", stdout, matches)
	} else {
		// G - good, U - good, untrusted
		if (matches[1] == "G" || matches[1] == "U") {
			return matches[2], true, nil
		} else {
			return matches[2], false, fmt.Errorf("git returned bad commit state: %s", matches[1])
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
	sigID, _, err := r.GetCommitSignature(commit)
	if err != nil {return false, err}
//	if !correct {return false, fmt.Errorf("Git log said commit sig is incorrect %s", err)}
	// hopefully when git changes it to provide full fingerprint it can be changed to just hash lookup
	for testedSig := range r.trustedSigs {
		if strings.Contains(testedSig, sigID) {
			return true, err
		}
	}
	return false, fmt.Errorf("no commit [%s] with sig [%s] in list: %+v", commit, sigID,r.trustedSigs)
}
