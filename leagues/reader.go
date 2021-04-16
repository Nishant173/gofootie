package leagues

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// Read CSV file having columns [HomeTeam, AwayTeam, HomeGoals, AwayGoals, Season, Date, League, Country] in that order
func ReadLeagueMatchesFromCsv(filepath string) LeagueMatches {
	csvfile, err := os.Open(filepath)
	if err != nil {
		log.Fatalln("Couldn't open the CSV file", err)
	}
	r := csv.NewReader(csvfile)
	lineCount := 0
	var records LeagueMatches
	for {
		lineCount++
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if lineCount != 1 {
			homeGoals, strConvErrHome := strconv.Atoi(record[2])
			awayGoals, strConvErrAway := strconv.Atoi(record[3])
			if strConvErrHome != nil {
				log.Fatalln("Error while converting HomeGoals to int", strConvErrHome)
			}
			if strConvErrAway != nil {
				log.Fatalln("Error while converting AwayGoals to int", strConvErrAway)
			}
			record := LeagueMatch{
				HomeTeam:  record[0],
				AwayTeam:  record[1],
				HomeGoals: homeGoals,
				AwayGoals: awayGoals,
				Season:    record[4],
				Date:      record[5],
				League:    record[6],
				Country:   record[7],
			}
			records = append(records, record)
		}
	}
	return records
}
