package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Teamresult struct {
	Name    string
	Matches int
	Wins    int
	Draws   int
	Losses  int
	Points  int
}

func Tally(reader io.Reader, writer io.Writer) error {

	league := map[string]Teamresult{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		entries := strings.Split(line, ";")
		if len(entries) == 3 {
			switch entries[2] {
			case "win":
				league[entries[0]] = AddWin(entries[0], league[entries[0]])
				league[entries[1]] = AddLoss(entries[1], league[entries[1]])
			case "draw":
				league[entries[0]] = AddDraw(entries[0], league[entries[0]])
				league[entries[1]] = AddDraw(entries[1], league[entries[1]])
			case "loss":
				league[entries[0]] = AddLoss(entries[0], league[entries[0]])
				league[entries[1]] = AddWin(entries[1], league[entries[1]])
			default:
				return fmt.Errorf("tally for input \"%s\" should have failed but didn't", line)
			}
		} else {
			if len(line) > 0 && !strings.HasPrefix(line, "#") {
				return fmt.Errorf("invalid format for %s", line)
			}
		}
	}
	a := []Teamresult{}
	for _, v := range league {
		a = append(a, v)
	}
	Write(a, writer)
	return nil
}

func Write(result []Teamresult, writer io.Writer) {
	fmt.Fprintf(writer, "%-30s |%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")

	sort.Slice(result, func(i, j int) bool {
		if result[i].Points == result[j].Points {
			return result[i].Name < result[j].Name
		}
		return result[i].Points > result[j].Points
	})
	for _, r := range result {
		fmt.Fprintf(writer, "%-30s |%3d |%3d |%3d |%3d |%3d\n", r.Name, r.Matches, r.Wins, r.Draws, r.Losses, r.Points)

	}
}

func AddWin(n string, r Teamresult) Teamresult {
	return Teamresult{
		Name:    n,
		Matches: r.Matches + 1,
		Wins:    r.Wins + 1,
		Draws:   r.Draws,
		Losses:  r.Losses,
		Points:  r.Points + 3,
	}
}

func AddDraw(n string, r Teamresult) Teamresult {
	return Teamresult{
		Name:    n,
		Matches: r.Matches + 1,
		Wins:    r.Wins,
		Draws:   r.Draws + 1,
		Losses:  r.Losses,
		Points:  r.Points + 1,
	}
}

func AddLoss(n string, r Teamresult) Teamresult {
	return Teamresult{
		Name:    n,
		Matches: r.Matches + 1,
		Wins:    r.Wins,
		Draws:   r.Draws,
		Losses:  r.Losses + 1,
		Points:  r.Points,
	}
}
