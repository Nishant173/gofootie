package leagues

type LeagueMatch struct {
	HomeTeam  string
	AwayTeam  string
	HomeGoals int
	AwayGoals int
	Season    string
	League    string
	Country   string
	Date      string
}

type LeagueMatches []LeagueMatch

type LeagueStanding struct {
	Position           int
	Team               string
	GamesPlayed        int
	Points             int
	GoalDifference     int
	Wins               int
	Losses             int
	Draws              int
	GoalsScored        int
	GoalsAllowed       int
	CleanSheets        int
	CleanSheetsAgainst int
	BigWins            int
	BigLosses          int
}

type LeagueStandings []LeagueStanding

type LeagueStandingNormalized struct {
	Position    int
	Team        string
	GamesPlayed int
	Ppg         float64
	Gdpg        float64
	WinPct      float64
	LossPct     float64
	DrawPct     float64
	Gspg        float64
	Gapg        float64
	CsPct       float64
	CsaPct      float64
	BigWinPct   float64
	BigLossPct  float64
}

type LeagueStandingsNormalized []LeagueStandingNormalized
