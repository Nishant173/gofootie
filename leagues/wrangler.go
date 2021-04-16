package leagues

import (
	"math"
	"sort"
)

func SortLeagueMatchesByDate(matches LeagueMatches, ascending bool) LeagueMatches {
	if ascending {
		sort.SliceStable(matches, func(i, j int) bool {
			return matches[i].Date < matches[j].Date
		})
	} else {
		sort.SliceStable(matches, func(i, j int) bool {
			return matches[i].Date > matches[j].Date
		})
	}
	return matches
}

func GetUniqueLeagueNames(matches LeagueMatches) []string {
	var uniqueLeagueNames []string
	for _, match := range matches {
		league := match.League
		if !StringInSlice(league, uniqueLeagueNames) {
			uniqueLeagueNames = append(uniqueLeagueNames, league)
		}
	}
	sort.Strings(uniqueLeagueNames)
	return uniqueLeagueNames
}

func GetUniqueSeasonNames(matches LeagueMatches) []string {
	var uniqueSeasonNames []string
	for _, match := range matches {
		season := match.Season
		if !StringInSlice(season, uniqueSeasonNames) {
			uniqueSeasonNames = append(uniqueSeasonNames, season)
		}
	}
	sort.Strings(uniqueSeasonNames)
	return uniqueSeasonNames
}

func GetUniqueTeamNames(matches LeagueMatches) []string {
	var uniqueTeamNames []string
	for _, match := range matches {
		homeTeam := match.HomeTeam
		awayTeam := match.AwayTeam
		if !StringInSlice(homeTeam, uniqueTeamNames) {
			uniqueTeamNames = append(uniqueTeamNames, homeTeam)
		}
		if !StringInSlice(awayTeam, uniqueTeamNames) {
			uniqueTeamNames = append(uniqueTeamNames, awayTeam)
		}
	}
	sort.Strings(uniqueTeamNames)
	return uniqueTeamNames
}

func GetUniqueTeamNamesByLeague(matches LeagueMatches) map[string][]string {
	mapUniqueTeamNamesByLeague := make(map[string][]string)
	allLeagues := GetUniqueLeagueNames(matches)
	for _, league := range allLeagues {
		mapUniqueTeamNamesByLeague[league] = []string{}
	}
	for _, match := range matches {
		homeTeam := match.HomeTeam
		awayTeam := match.AwayTeam
		league := match.League
		if !StringInSlice(homeTeam, mapUniqueTeamNamesByLeague[league]) {
			mapUniqueTeamNamesByLeague[league] = append(mapUniqueTeamNamesByLeague[league], homeTeam)
		}
		if !StringInSlice(awayTeam, mapUniqueTeamNamesByLeague[league]) {
			mapUniqueTeamNamesByLeague[league] = append(mapUniqueTeamNamesByLeague[league], awayTeam)
		}
	}
	for _, league := range allLeagues {
		teamsByLeague := mapUniqueTeamNamesByLeague[league]
		sort.Strings(teamsByLeague)
		mapUniqueTeamNamesByLeague[league] = teamsByLeague
	}
	return mapUniqueTeamNamesByLeague
}

func IsMatchDrawn(match LeagueMatch) bool {
	if match.HomeGoals == match.AwayGoals {
		return true
	}
	return false
}

func GetAbsoluteGoalDifference(match LeagueMatch) int {
	return int(math.Abs(float64(match.HomeGoals - match.AwayGoals)))
}

// Returns name of winning team (if any). Returns "NA" if match was a draw.
func GetWinningTeamName(match LeagueMatch) string {
	var winningTeamName string
	if match.HomeGoals > match.AwayGoals {
		winningTeamName = match.HomeTeam
	} else if match.HomeGoals < match.AwayGoals {
		winningTeamName = match.AwayTeam
	} else {
		winningTeamName = "NA"
	}
	return winningTeamName
}

// Returns name of losing team (if any). Returns "NA" if match was a draw.
func GetLosingTeamName(match LeagueMatch) string {
	var losingTeamName string
	if match.HomeGoals > match.AwayGoals {
		losingTeamName = match.AwayTeam
	} else if match.HomeGoals < match.AwayGoals {
		losingTeamName = match.HomeTeam
	} else {
		losingTeamName = "NA"
	}
	return losingTeamName
}
