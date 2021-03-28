package util_test

import (
	"fmt"
	"middleware/model"
	"middleware/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortMatches(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		Name      string
		Limit     int
		Matches   []model.Match
		SortedIDs []int
	}{
		{
			Name:      "empty array",
			Limit:     0,
			Matches:   make([]model.Match, 0),
			SortedIDs: nil,
		},
		{
			Name:  "negative uts",
			Limit: 0,
			Matches: []model.Match{
				{
					PlayTime: model.PlayTime{
						UTS: -2531935833,
					},
				},
			},
			SortedIDs: nil,
		},
		{
			Name:      "negative limit",
			Limit:     -1,
			Matches:   getTestMatches(),
			SortedIDs: nil,
		},
		{
			Name:      "limit smaller than the number of mathes",
			Limit:     1,
			Matches:   getTestMatches(),
			SortedIDs: []int{5},
		},
		{
			Name:      "5 matches, 2 not yet played, should return 3 sorted",
			Limit:     5,
			Matches:   getTestMatches(),
			SortedIDs: []int{5, 4, 2},
		},
	}
	for _, tc := range testCases {
		sortedMatches := util.SortMatches(tc.Matches, tc.Limit)
		assert.Equal(len(tc.SortedIDs), len(sortedMatches))
		for i, match := range sortedMatches {
			assert.Equal(tc.SortedIDs[i], match.ID, fmt.Sprintf("Failed test %s", tc.Name))
		}
	}
}

func TestGetYears(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		Name           string
		TournamentYear string
		Years          []int64
	}{
		{
			Name:           "empty string",
			TournamentYear: "",
			Years:          nil,
		},
		{
			Name:           "invalid year string",
			TournamentYear: "invalid year",
			Years:          nil,
		},
		{
			Name:           "unsupported delimiter",
			TournamentYear: "20,21",
			Years:          nil,
		},
		{
			Name:           "1 full year",
			TournamentYear: "2021",
			Years:          []int64{2021},
		},
		{
			Name:           "1 shortened year",
			TournamentYear: "21",
			Years:          []int64{21},
		},
		{
			Name:           "negative year",
			TournamentYear: "-21",
			Years:          nil,
		},
		{
			Name:           "2 years",
			TournamentYear: "20/21",
			Years:          []int64{20, 21},
		},
	}
	for _, tc := range testCases {
		years := util.GetYears(tc.TournamentYear)
		assert.Equal(tc.Years, years)
	}
}

func getTestMatches() []model.Match {
	return []model.Match{
		{
			ID: 1,
			PlayTime: model.PlayTime{
				Date: "26/03/50", //2050
				UTS:  2531935833,
			},
		},
		{
			ID: 2,
			PlayTime: model.PlayTime{
				Time: "19:25",
				Date: "19/03/21",
				UTS:  1616181900,
			},
		},
		{
			ID: 3,
			PlayTime: model.PlayTime{
				Date: "26/03/42", //2042
				UTS:  2279475033,
			},
		},
		{
			ID: 4,
			PlayTime: model.PlayTime{
				Time: "11:30",
				Date: "20/03/21",
				UTS:  1616239800,
			},
		},
		{
			ID: 5,
			PlayTime: model.PlayTime{
				Time: "14:30",
				Date: "20/03/21",
				UTS:  1616250600,
			},
		},
	}
}
