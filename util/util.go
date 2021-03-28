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

func SortMatches(matches []model.Match, limit int) []model.Match {
	if limit < 0 {
		return nil
	}
	var sortedTopMatches []model.Match // Len not set, since there can be less matches than the number of the limit
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
