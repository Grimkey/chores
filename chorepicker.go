package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// Group is a set of people and chores from the CSV file
type Group struct {
	seed  int64
	Name  []string
	Chore []string
}

// FromCsv loads a person slice from a csv file
func FromCsv(seed int64, filename string) *Group {
	grp := Group{seed: seed}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(file))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range records {
		grp.Name = append(grp.Name, i[0])
		grp.Chore = append(grp.Chore, i[1])
	}

	return &grp
}

// Shuffle randomized chores
func (g *Group) Shuffle() {
	vals := g.Chore
	r := rand.New(rand.NewSource(g.seed))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func (g *Group) String() string {
	bldr := strings.Builder{}

	for idx := 0; idx < len(g.Name); idx++ {
		bldr.WriteString(fmt.Sprintf("\"%s\",\"%s\"\n", g.Name[idx], g.Chore[idx]))
	}

	return bldr.String()
}
