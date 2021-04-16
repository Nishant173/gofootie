package leagues

import "github.com/gin-gonic/gin"

func RouteHandler() {
	router := gin.Default()
	router.GET("/teams", TeamsView())
	router.GET("/leagues", LeaguesView())
	router.GET("/seasons", SeasonsView())
	router.GET("/teams-by-league", TeamsByLeagueView())
	router.GET("/league-matches", LeagueMatchesView())
	router.GET("/league-standings", LeagueStandingsView())
	router.Run()
}
