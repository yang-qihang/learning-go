package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	name        string
	playerNames []string
}

type League struct {
	Teams []Team         // maps to each team
	Wins  map[string]int // maps a team's name to its number of wins
}

func (l *League) MatchResult(name1 string, score1 int, name2 string, score2 int) {
	// a tie is not a win
	if score1 > score2 {
		l.Wins[name1]++
	} else if score1 < score2 {
		l.Wins[name2]++
	}
}

func (l League) Ranking() []string {
	ranking := make([]string, 0, len(l.Wins))

	// we need to create entries for all teams from the map
	for name := range l.Wins {
		ranking = append(ranking, name)
	}

	// the comparison function returns `true`
	// when the element at index i should precede the element at index j
	// Here, we sort in descending order of wins
	sort.Slice(ranking, func(i, j int) bool {
		return l.Wins[ranking[i]] > l.Wins[ranking[j]]
	})

	// Note:
	// the current comparator doesn't specify an ordering
	// for teams with equal wins.
	// That's perfectly valid for sorting, but their relative order
	// isn't guaranteed.
	// If you want deterministic results, add a secondary alphabetical comparison.

	// wi := l.Wins[ranking[i]]
	// wj := l.Wins[ranking[j]]

	// if wi != wj {
	//     return wi > wj
	// }

	// return ranking[i] < ranking[j]

	return ranking
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(rker Ranker, w io.Writer) {
	for _, str := range rker.Ranking() {
		io.WriteString(w, str+"\n")
	}
}

func main() {
	league := League{
		Teams: []Team{
			{name: "Eagles"},
			{name: "Lions"},
			{name: "Tigers"},
		},
		Wins: map[string]int{
			"Eagles": 0,
			"Lions":  0,
			"Tigers": 0,
		},
	}

	league.MatchResult("Eagles", 3, "Lions", 1)
	league.MatchResult("Tigers", 2, "Eagles", 1)
	league.MatchResult("Lions", 3, "Tigers", 0)
	league.MatchResult("Eagles", 2, "Lions", 1)

	fmt.Println(league.Ranking())

	// os.Stdout can be passed directly as an io.Writer.
	RankPrinter(league, os.Stdout)
}
