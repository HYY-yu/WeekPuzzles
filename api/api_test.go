package api

import (
	"testing"
	"fmt"
	"WeekPuzzles/saturday"
)

func TestApiToServer(t *testing.T) {
	sas := make(saturday.Saturdays, 0, 2)

	sa1 := saturday.Saturday{
		Date: "20180428",
	}

	sa2 := saturday.Saturday{
		Date: "20180505",
	}

	sas = append(sas, sa1, sa2)

	ApiToServer(sas)
	fmt.Println(sas)
}
