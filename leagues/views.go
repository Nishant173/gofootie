package leagues

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func TeamsView() gin.HandlerFunc {
	matches := ReadLeagueMatchesFromCsv(LeagueMatchesCsvFile)
	allTeams := GetUniqueTeamNames(matches)
	return func(context *gin.Context) {
		if allTeams == nil {
			allTeams = []string{}
		}
		context.JSON(200, allTeams)
	}
}

func LeaguesView() gin.HandlerFunc {
	matches := ReadLeagueMatchesFromCsv(LeagueMatchesCsvFile)
	allLeagues := GetUniqueLeagueNames(matches)
	return func(context *gin.Context) {
		if allLeagues == nil {
			allLeagues = []string{}
		}
		context.JSON(200, allLeagues)
	}
}

func SeasonsView() gin.HandlerFunc {
	matches := ReadLeagueMatchesFromCsv(LeagueMatchesCsvFile)
	allSeasons := GetUniqueSeasonNames(matches)
	return func(context *gin.Context) {
		if allSeasons == nil {
			allSeasons = []string{}
		}
		context.JSON(200, allSeasons)
	}
}

func TeamsByLeagueView() gin.HandlerFunc {
	matches := ReadLeagueMatchesFromCsv(LeagueMatchesCsvFile)
	teamsByLeague := GetUniqueTeamNamesByLeague(matches)
	return func(context *gin.Context) {
		if teamsByLeague == nil {
			teamsByLeague = make(map[string][]string)
		}
		context.JSON(200, teamsByLeague)
	}
}

func LeagueMatchesView() gin.HandlerFunc {
	return func(context *gin.Context) {
		league := context.Query("league")
		season := context.Query("season")
		team := context.Query("team")
		winningTeam := context.Query("winningTeam")
		losingTeam := context.Query("losingTeam")
		goalDifference := context.Query("goalDifference")
		minGoalDifference := context.Query("minGoalDifference")
		maxGoalDifference := context.Query("maxGoalDifference")

		matches := ReadLeagueMatchesFromCsv(LeagueMatchesCsvFile)
		matches = SortLeagueMatchesByDate(matches, true)
		if league != "" {
			matches = FilterByLeague(matches, league)
		}
		if season != "" {
			matches = FilterBySeason(matches, season)
		}
		if team != "" {
			matches = FilterByTeam(matches, team)
		}
		if winningTeam != "" {
			matches = FilterByWinningTeam(matches, winningTeam)
		}
		if losingTeam != "" {
			matches = FilterByLosingTeam(matches, losingTeam)
		}
		if goalDifference != "" {
			gd, _ := strconv.Atoi(goalDifference)
			matches = FilterByGoalDifference(matches, gd)
		}
		if minGoalDifference != "" {
			minGd, _ := strconv.Atoi(minGoalDifference)
			matches = FilterByMinGoalDifference(matches, minGd)
		}
		if maxGoalDifference != "" {
			maxGd, _ := strconv.Atoi(maxGoalDifference)
			matches = FilterByMaxGoalDifference(matches, maxGd)
		}
		if matches == nil {
			matches = []LeagueMatch{}
		}
		context.JSON(200, matches)
	}
}

func LeagueStandingsView() gin.HandlerFunc {
	return func(context *gin.Context) {
		league := context.Query("league")
		season := context.Query("season")
		matches := ReadLeagueMatchesFromCsv(LeagueMatchesCsvFile)
		leagueStandings := GetLeagueStandings(matches, league, season)
		if leagueStandings == nil {
			leagueStandings = []LeagueStanding{}
		}
		context.JSON(200, leagueStandings)
	}
}
