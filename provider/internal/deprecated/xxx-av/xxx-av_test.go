//go:build deprecated

package xxx_av

import (
	"testing"

	"github.com/djniche/test123/provider/internal/testkit"
)

func TestXXXAV_GetMovieInfoByID(t *testing.T) {
	testkit.Test(t, New, []string{
		"24719",
		"23395",
		"19337",
	})
}
