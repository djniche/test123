package getchu

import (
	"testing"

	"github.com/djniche/test123/provider/internal/testkit"
)

func TestGetchu_GetMovieInfoByID(t *testing.T) {
	testkit.Test(t, New, []string{
		"4018339",
		"4042392",
		"4041955",
		"4042404",
		"4042423",
	})
}
