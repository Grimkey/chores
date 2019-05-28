package main

import (
	"fmt"
	"time"
)

func main() {
	grp := FromCsv(time.Now().UnixNano(), "samples.csv")
	grp.Shuffle()
	fmt.Println(grp.String())
}
