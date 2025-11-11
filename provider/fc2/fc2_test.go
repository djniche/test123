package fc2

import (
	"testing"

	"github.com/djniche/test123/provider/internal/testkit"
)

func TestFC2_GetMovieInfoByID(t *testing.T) {
	testkit.Test(t, New, []string{
		"2812904",
		"4669533",
	})
}
