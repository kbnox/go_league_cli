package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var once sync.Once
var regexPattern *regexp.Regexp
var subexpNames []string

func main() {
	fileName := flag.String("file", "", "The file containing the scores for each game.")
	flag.Parse()

	//ensure our file has been parsed
	if *fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	//We open our file.
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("Error Opening file. Error: %v", err)
	}
	//Close our file when we're done with it.
	defer file.Close()

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	// keep a variable to help us tell the user what line we couldn't process
	l := 1

	// we'll put our results in a map for each team
	leagueTableMap := make(map[string]int64)

	// Read each line of the file line by line
	for scanner.Scan() {
		/*
			Each line is the result of a game.
			The results of the game are split by a comma.
			[Team 1] [score], [Team 2] [score]
		*/
		teamsStrs := strings.Split(scanner.Text(), ",")

		/*
			If:
			- the line is empty
			- we have less than 2 teams for each game
			- or more than 2 teams each game
			We fail
		*/
		if len(teamsStrs) != 2 {
			log.Fatalf("Incorrect number of teams on line %d", l)
		}

		// setup our regex
		setupRegex()

		// Get the game so we can process it
		game := GetGame(teamsStrs[0], teamsStrs[1])

		/*
			Game Rules:
			- Team with highest score earns 3 points
			- A game that's a draw results in both teams earning 1 point
			- Loosing team receives 0 points
		*/

		if game.Team1.Score == game.Team2.Score {
			leagueTableMap[game.Team1.Name]++
			leagueTableMap[game.Team2.Name]++
		} else if game.Team1.Score > game.Team2.Score {
			leagueTableMap[game.Team1.Name] += 3
			leagueTableMap[game.Team2.Name] += 0
		} else {
			leagueTableMap[game.Team1.Name] += 0
			leagueTableMap[game.Team2.Name] += 3
		}

		l++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error occured while reading from file. Error: %v", err)
	}

	// make a lice of team structs for our league results
	leagueResults := make([]*Team, 0)

	// populate our slice with the league results for each team
	for name, points := range leagueTableMap {
		leagueResults = append(leagueResults, &Team{
			Name:   name,
			Points: points,
		})
	}

	/*
		Sort:
		- Sort by highest points first
		- Where points are equal, we sort alphabetically
	*/
	sort.Slice(leagueResults, func(i, j int) bool {
		if leagueResults[i].Points == leagueResults[j].Points {
			return leagueResults[i].Name < leagueResults[j].Name
		}
		return leagueResults[i].Points > leagueResults[j].Points
	})

	//print out the league standings
	for i, team := range leagueResults {
		fmt.Printf("%d. %s, %d pts\n", i+1, team.Name, team.Points)
	}
}

/*
	setupRegex allows us to ensure that we'll only ever compile our regex once
	- using MustCompile will ensure our regex comiles. If it fails it will panic
*/
func setupRegex() {
	once.Do(func() {
		regexPattern = regexp.MustCompile(`(?P<Alpha>[a-z A-Z]*)(?P<Numeric>[0-9])`)
		subexpNames = regexPattern.SubexpNames()
	})
}

type Team struct {
	Name   string
	Score  int64
	Points int64
}

type Game struct {
	Team1 *Team
	Team2 *Team
}

// GetGame returns a game struct
func GetGame(s1, s2 string) *Game {

	return &Game{
		Team1: GetTeam(s1),
		Team2: GetTeam(s2),
	}
}

/*
	GetTeam return a pointer of struct Team
	- Matches the team string to regex to find our groups
	- Populates our team struct
*/
func GetTeam(s string) *Team {
	// run our string through our regex
	matches := regexPattern.FindAllStringSubmatch(s, -1)

	// if we can't match it then it must be funky
	if len(matches) < 1 {
		log.Fatalf("Failed to process team: %s", s)
	}

	// create a map of match data
	md := make(map[string]string)

	// we put our match subexpressions in the map so we can look them up easier
	for i, n := range matches[0] {
		md[subexpNames[i]] = n
	}

	// convert our name from the alpha subexpression and trim spaces
	name := strings.TrimSpace(md["Alpha"])

	// convert our score to int64
	score, err := strconv.ParseInt(md["Numeric"], 10, 64)
	if err != nil {
		log.Fatalf("Error converting string to int64: %s", md["Numeric"])
	}

	// return our team struct
	return &Team{
		Name:  name,
		Score: score,
	}
}
