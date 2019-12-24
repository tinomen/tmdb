package rename

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanFilename(t *testing.T) {
	var movies = []struct {
		name     string
		expected string
	}{
		{"Joker.2019.720p.BluRay.x264-[YTS.LT].avi", "Joker 2019 720p BluRay x264 YTS LT"},
		{"Avengers Endgame.2019.720p.BluRay.x264-[YTS.LT].avi", "Avengers Endgame 2019 720p BluRay x264 YTS LT"},
	}

	for _, test := range movies {
		assert.Equal(t, cleanFilename(test.name), test.expected)
	}
}
