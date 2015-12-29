package gitcli

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestVersion(t *testing.T) {
	var r Repo

	Convey("Returns version", t, func() {
		z, err := r.Version()

		So(z, ShouldContainSubstring, "git version")
		So(err, ShouldEqual, nil)
	})
}
