package leagues

import (
	"math"
	"sort"
)

func GetGamesPlayedCount(matches LeagueMatches, team string) int {
	count := 0
	for _, match := range matches {
		if match.HomeTeam == team {
			count++
		} else if match.AwayTeam == team {
			count++
		}
	}
	return count
}

func GetWinCount(matches LeagueMatches, team string) int {
	count := 0
	for _, match := range matches {
		if match.HomeTeam == team && match.HomeGoals > match.AwayGoals {
			count++
		} else if match.AwayTeam == team && match.AwayGoals > match.HomeGoals {
			count++
		}
	}
	return count
}

func GetLossCount(matches LeagueMatches, team string) int {
	count := 0
	for _, match := range matches {
		if match.HomeTeam == team && match.HomeGoals < match.AwayGoals {
			count++
		} else if match.AwayTeam == team && match.AwayGoals < match.HomeGoals {
			count++
		}
	}
	return count
}

func GetDrawCount(matches LeagueMatches, team string) int {
	count := 0
	for _, match := range matches {
		if match.HomeTeam == team && match.HomeGoals == match.AwayGoals {
			count++
		} else if match.AwayTeam == team && match.AwayGoals == match.HomeGoals {
			count++
		}
	}
	return count
}

func GetGoalsScored(matches LeagueMatches, team string) int {
	goalsScored := 0
	for _, match := range matches {
		if match.HomeTeam == team {
			goalsScored += match.HomeGoals
		} else if match.AwayTeam == team {
			goalsScored += match.AwayGoals
		}
	}
	return goalsScored
}

func GetGoalsAllowed(matches LeagueMatches, team string) int {
	goalsAllowed := 0
	for _, match := range matches {
		if match.HomeTeam == team {
			goalsAllowed += match.AwayGoals
		} else if match.AwayTeam == team {
			goalsAllowed += match.HomeGoals
		}
	}
	return goalsAllowed
}

func GetCleanSheets(matches LeagueMatches, team string) int {
	cleanSheetCount := 0
	for _, match := range matches {
		if match.HomeTeam == team && match.AwayGoals == 0 {
			cleanSheetCount++
		} else if match.AwayTeam == team && match.HomeGoals == 0 {
			cleanSheetCount++
		}
	}
	return cleanSheetCount
}

func GetCleanSheetsAgainst(matches LeagueMatches, team string) int {
	cleanSheetAgainstCount := 0
	for _, match := range matches {
		if match.HomeTeam == team && match.HomeGoals == 0 {
			cleanSheetAgainstCount++
		} else if match.AwayTeam == team && match.AwayGoals == 0 {
			cleanSheetAgainstCount++
		}
	}
	return cleanSheetAgainstCount
}

func GetBigWinCount(matches LeagueMatches, team string, margin int) int {
	bigWinCount := 0
	for _, match := range matches {
		hg := match.HomeGoals
		ag := match.AwayGoals
		goalMargin := int(math.Abs(float64(hg - ag)))
		if match.HomeTeam == team && hg > ag && goalMargin >= margin {
			bigWinCount++
		} else if match.AwayTeam == team && ag > hg && goalMargin >= margin {
			bigWinCount++
		}
	}
	return bigWinCount
}

func GetBigLossCount(matches LeagueMatches, team string, margin int) int {
	bigLossCount := 0
	for _, match := range matches {
		hg := match.HomeGoals
		ag := match.AwayGoals
		goalMargin := int(math.Abs(float64(hg - ag)))
		if match.HomeTeam == team && hg < ag && goalMargin >= margin {
			bigLossCount++
		} else if match.AwayTeam == team && ag < hg && goalMargin >= margin {
			bigLossCount++
		}
	}
	return bigLossCount
}

func sortLeagueStandingsByRankingMetrics(leagueStandings LeagueStandings) LeagueStandings {
	sort.SliceStable(leagueStandings, func(i, j int) bool {
		pointsOfI := 3*leagueStandings[i].Wins + leagueStandings[i].Draws
		pointsOfJ := 3*leagueStandings[j].Wins + leagueStandings[j].Draws
		ppgOfI := float64(pointsOfI) / float64(leagueStandings[i].GamesPlayed)
		ppgOfJ := float64(pointsOfJ) / float64(leagueStandings[j].GamesPlayed)
		return ppgOfI > ppgOfJ
	})
	return leagueStandings
}

func attachRankingToSortedLeagueStandings(leagueStandings LeagueStandings) LeagueStandings {
	var leagueStandingsRanked LeagueStandings
	for idx, leagueStanding := range leagueStandings {
		leagueStanding.Position = idx + 1
		leagueStandingsRanked = append(leagueStandingsRanked, leagueStanding)
	}
	return leagueStandingsRanked
}

// Gets league standings for the given league and season
func GetLeagueStandings(matches LeagueMatches, league string, season string) LeagueStandings {
	matches = FilterByLeague(matches, league)
	matches = FilterBySeason(matches, season)
	teams := GetUniqueTeamNames(matches)
	bigResultGoalMargin := 3 // Will be considered as big result if GoalDifference >= this number
	var leagueStandings LeagueStandings
	for _, team := range teams {
		wins := GetWinCount(matches, team)
		draws := GetDrawCount(matches, team)
		gs := GetGoalsScored(matches, team)
		ga := GetGoalsAllowed(matches, team)
		gd := gs - ga
		points := 3*wins + draws
		leagueStanding := LeagueStanding{
			Team:               team,
			GamesPlayed:        GetGamesPlayedCount(matches, team),
			Points:             points,
			GoalDifference:     gd,
			Wins:               wins,
			Losses:             GetLossCount(matches, team),
			Draws:              draws,
			GoalsScored:        gs,
			GoalsAllowed:       ga,
			CleanSheets:        GetCleanSheets(matches, team),
			CleanSheetsAgainst: GetCleanSheetsAgainst(matches, team),
			BigWins:            GetBigWinCount(matches, team, bigResultGoalMargin),
			BigLosses:          GetBigLossCount(matches, team, bigResultGoalMargin),
		}
		leagueStandings = append(leagueStandings, leagueStanding)
	}
	leagueStandings = sortLeagueStandingsByRankingMetrics(leagueStandings)
	leagueStandings = attachRankingToSortedLeagueStandings(leagueStandings)
	return leagueStandings
}
