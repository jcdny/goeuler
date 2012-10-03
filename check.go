package euler

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
)

const solutions = "/Dropbox/Euler_Answers.txt"

var (
	once        sync.Once
	solutionset map[int]string
)

func initSolutionSet() {
	solutionset = make(map[int]string, 1000)
	s, err := ioutil.ReadFile(os.Getenv("HOME") + solutions)
	if err != nil {
		solutionset[0] = "BADPARSE"
	} else {
		f := strings.Fields(string(s))
		for i := 0; i < len(f); i += 2 {
			np, err := strconv.Atoi(f[i])
			if err != nil {
				log.Print("Bad parse ", f[i], err)
			} else {
				solutionset[np] = f[i+1]
			}
		}
	}
}

func Check(p int, answer interface{}) (bool, bool) {
	once.Do(initSolutionSet)

	if correct, ok := solutionset[p]; ok {
		return correct == fmt.Sprint(answer), true
	}
	// no answer found return false for correct and solved.
	return false, false
}

func Validate(t *testing.T, p int, answer interface{}) {
	correct, solved := Check(p, answer)
	if !solved {
		t.Log(p, " answer not checked:", answer)
	} else if !correct {
		t.Error(p, " answer incorrect:", answer)
	}
}
