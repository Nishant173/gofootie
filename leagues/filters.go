package leagues

func FilterByTeam(matches LeagueMatches, team string) LeagueMatches {
	var matchesByTeam LeagueMatches
	for _, match := range matches {
		if match.HomeTeam == team || match.AwayTeam == team {
			matchesByTeam = append(matchesByTeam, match)
		}
	}
	return matchesByTeam
}

func FilterByLeague(matches LeagueMatches, league string) LeagueMatches {
	var matchesByLeague LeagueMatches
	for _, match := range matches {
		if match.League == league {
			matchesByLeague = append(matchesByLeague, match)
		}
	}
	return matchesByLeague
}

func FilterBySeason(matches LeagueMatches, season string) LeagueMatches {
	var matchesBySeason LeagueMatches
	for _, match := range matches {
		if match.Season == season {
			matchesBySeason = append(matchesBySeason, match)
		}
	}
	return matchesBySeason
}

func FilterByWinningTeam(matches LeagueMatches, winningTeam string) LeagueMatches {
	var matchesByWinningTeam LeagueMatches
	for _, match := range matches {
		winningTeamName := GetWinningTeamName(match)
		if winningTeam == winningTeamName {
			matchesByWinningTeam = append(matchesByWinningTeam, match)
		}
	}
	return matchesByWinningTeam
}

func FilterByLosingTeam(matches LeagueMatches, losingTeam string) LeagueMatches {
	var matchesByLosingTeam LeagueMatches
	for _, match := range matches {
		losingTeamName := GetLosingTeamName(match)
		if losingTeam == losingTeamName {
			matchesByLosingTeam = append(matchesByLosingTeam, match)
		}
	}
	return matchesByLosingTeam
}

func FilterByGoalDifference(matches LeagueMatches, goalDifference int) LeagueMatches {
	var matchesByGoalDifference LeagueMatches
	for _, match := range matches {
		absoluteGoalDifference := GetAbsoluteGoalDifference(match)
		if goalDifference == absoluteGoalDifference {
			matchesByGoalDifference = append(matchesByGoalDifference, match)
		}
	}
	return matchesByGoalDifference
}

func FilterByMinGoalDifference(matches LeagueMatches, minGoalDifference int) LeagueMatches {
	var matchesByMinGoalDifference LeagueMatches
	for _, match := range matches {
		absoluteGoalDifference := GetAbsoluteGoalDifference(match)
		if absoluteGoalDifference >= minGoalDifference {
			matchesByMinGoalDifference = append(matchesByMinGoalDifference, match)
		}
	}
	return matchesByMinGoalDifference
}

func FilterByMaxGoalDifference(matches LeagueMatches, maxGoalDifference int) LeagueMatches {
	var matchesByMaxGoalDifference LeagueMatches
	for _, match := range matches {
		absoluteGoalDifference := GetAbsoluteGoalDifference(match)
		if absoluteGoalDifference <= maxGoalDifference {
			matchesByMaxGoalDifference = append(matchesByMaxGoalDifference, match)
		}
	}
	return matchesByMaxGoalDifference
}
