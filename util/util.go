package util

import (
	"log"
	"middleware/model"
	"sort"
	"strconv"
	"strings"
	"time"
)

const _dateLayout = "02/01/06"

// SortMatches sorts limit number of matches by date descending.
func SortMatches(matches []model.Match, limit int) []model.Match {
	if limit < 0 {
		return nil
	}
	var sortedTopMatches []model.Match
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].PlayTime.UTS > matches[j].PlayTime.UTS
	})
	for _, match := range matches {
		t, err := time.Parse(_dateLayout, match.PlayTime.Date)
		if err != nil {
			log.Printf("[SortMatches] Could not parse date for match id %d: %s", match.ID, err.Error())
			continue
		}
		if t.After(time.Now()) {
			//Skip this match since it wasn't yet held.
			continue
		}
		sortedTopMatches = append(sortedTopMatches, match)
		if len(sortedTopMatches) == limit {
			break
		}
	}
	return sortedTopMatches
}

// GetYears converts string tournamentYear, which can e.g be like 2020 or 20/21 and returns an array of
// int values. Negative numbers are discarded.
func GetYears(tournamentYear string) []int64 {
	if tournamentYear == "" {
		return nil
	}
	splittedYears := strings.Split(tournamentYear, "/")
	years := make([]int64, len(splittedYears))
	for i, str := range splittedYears {
		year, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			log.Printf("[GetTournamentYears] Failed converting string to int: %s", err.Error())
			return nil
		}
		if year < 0 {
			return nil
		}
		years[i] = year
	}
	return years
}
