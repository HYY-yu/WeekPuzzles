package api

import (
	"calcuate/saturday"
	"testing"
	"fmt"
)

func TestApiToServer(t *testing.T) {
	sa := &saturday.Saturday{
		Date:"20180428",
	}

		ApiToServer(sa)
		fmt.Println(sa)
}
