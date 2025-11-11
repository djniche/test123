//go:build deprecated

package prestige

import (
	"testing"

	"github.com/djniche/test123/provider/internal/testkit"
)

func TestPRESTIGE_GetMovieInfoByID(t *testing.T) {
	testkit.Test(t, New, []string{
		"hrv-014",
		"edd-010",
	})
}

func TestPRESTIGE_SearchMovie(t *testing.T) {
	testkit.Test(t, New, []string{
		"edd-013",
	})
}
