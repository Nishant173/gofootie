package leagues

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/fatih/structs"
)

// Method that returns slice of stringified elements
func (match LeagueMatch) ListStringifiedValues() []string {
	values := []string{}
	values = append(values, match.HomeTeam)
	values = append(values, match.AwayTeam)
	values = append(values, strconv.Itoa(match.HomeGoals))
	values = append(values, strconv.Itoa(match.AwayGoals))
	values = append(values, match.Season)
	values = append(values, match.League)
	values = append(values, match.Country)
	values = append(values, match.Date)
	return values
}

func (matches LeagueMatches) SaveToCsv(filepath string) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		os.Exit(1)
	}
	sliceStringifiedRecords := [][]string{} // Slice of slice of strings, where each sub-slice represents a record
	statFields := structs.Names(&LeagueMatch{})
	sliceStringifiedRecords = append(sliceStringifiedRecords, statFields)
	for _, match := range matches {
		record := match.ListStringifiedValues()
		sliceStringifiedRecords = append(sliceStringifiedRecords, record)
	}
	csvWriter := csv.NewWriter(file)
	csvWriter.WriteAll(sliceStringifiedRecords)
	csvWriter.Flush()
}

// Method that returns slice of stringified elements
func (leagueStanding LeagueStanding) ListStringifiedValues() []string {
	values := []string{}
	values = append(values, strconv.Itoa(leagueStanding.Position))
	values = append(values, leagueStanding.Team)
	values = append(values, strconv.Itoa(leagueStanding.GamesPlayed))
	values = append(values, strconv.Itoa(leagueStanding.Points))
	values = append(values, strconv.Itoa(leagueStanding.GoalDifference))
	values = append(values, strconv.Itoa(leagueStanding.Wins))
	values = append(values, strconv.Itoa(leagueStanding.Losses))
	values = append(values, strconv.Itoa(leagueStanding.Draws))
	values = append(values, strconv.Itoa(leagueStanding.GoalsScored))
	values = append(values, strconv.Itoa(leagueStanding.GoalsAllowed))
	values = append(values, strconv.Itoa(leagueStanding.CleanSheets))
	values = append(values, strconv.Itoa(leagueStanding.CleanSheetsAgainst))
	values = append(values, strconv.Itoa(leagueStanding.BigWins))
	values = append(values, strconv.Itoa(leagueStanding.BigLosses))
	return values
}

func (leagueStandings LeagueStandings) SaveToCsv(filepath string) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		os.Exit(1)
	}
	sliceStringifiedRecords := [][]string{} // Slice of slice of strings, where each sub-slice represents a record
	statFields := structs.Names(&LeagueStanding{})
	sliceStringifiedRecords = append(sliceStringifiedRecords, statFields)
	for _, leagueStanding := range leagueStandings {
		record := leagueStanding.ListStringifiedValues()
		sliceStringifiedRecords = append(sliceStringifiedRecords, record)
	}
	csvWriter := csv.NewWriter(file)
	csvWriter.WriteAll(sliceStringifiedRecords)
	csvWriter.Flush()
}
