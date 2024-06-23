package runner

import (
	//"codeberg.org/filipmnowak/beaver/internal/db/sqlite"
	. "codeberg.org/filipmnowak/beaver/internal/tests"
	"sync"
)

func Merge(chs []chan *Test) <-chan *Test {
	out := make(chan *Test)
	var wg sync.WaitGroup
	wg.Add(len(chs))

	output := func(ch <-chan *Test) {
		for n := range ch {
			out <- n
		}
		wg.Done()
	}
	for _, ch := range chs {
		go output(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func runTestFamilies(tfs []*TestFamily) <-chan []*Test {
	return make(<-chan []*Test)
}

func runTestGroups(tgs []*TestGroup) <-chan []*Test {
	return make(<-chan []*Test)
}

func RunTests(ts []*Test) []chan *Test {
	var out []chan *Test
	for i := 0; i < len(ts); i++ {
		out = append(out, make(chan *Test))
	}
	for i, t := range ts {
		go func() {
			t.Run()
			out[i] <- t
			close(out[i])
		}()
	}
	return out
}

func persistResults() {
}

func Run() {
}
