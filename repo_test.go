package gitcli

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"os"
	"io/ioutil"
)

func TestInit(t *testing.T) {
	tmpdir, err := ioutil.TempDir("tmp","git")
	Convey("create tmp dir" + tmpdir, t, func() {
		So(err, ShouldEqual, nil)
	})
	repo := New(tmpdir)
	err = repo.Init()
	Convey("create repo", t, func() {
		So(err, ShouldEqual, nil)
		repoConfigPath := tmpdir + `/.git/config`
		f, err := os.Stat(repoConfigPath)
		So(err, ShouldEqual, nil)
		So(f.Mode().IsRegular(), ShouldEqual, true)
	})
	Convey("reinitialize repo", t, func() {
		err = repo.Init()
		So(err,ShouldEqual,nil)
	})

}
