package runner

import (
	"codeberg.org/filipmnowak/beaver/internal/db/sqlite"
	. "codeberg.org/filipmnowak/beaver/internal/tests"
)

func runTestFamilies(tfs []*TestFamily) <-chan []*Test {
	return make(<-chan []*Test)
}

func runTestGroups(tgs []*TestGroup) <-chan []*Test {
	return make(<-chan []*Test)
}

func runTests(ts []*Test) <-chan []*Test {
	out := make(<-chan []*Test)
	for i, t := range ts {
		go func() {
		}
	}
}

func persistResults() {
}

func Run() {
}
