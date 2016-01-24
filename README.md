[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/XANi/go-gitcli)

# go-gitcli


Lightweight wrapper around git cli.
This is simpler alternative to libgit2 bindings, it also (will) have builtin GPG signing support (which is main reason for that lib as libgit2 doesnt have it yet).

Requirements: git in the path.

Lib will sanitize git env, removing things like language variables so output of git commands are consistent
