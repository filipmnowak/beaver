package runner

import (
	"codeberg.org/filipmnowak/beaver/internal/db/sqlite"
	. "codeberg.org/filipmnowak/beaver/internal/tests"
)

func merge(chs []<-chan []*Test) <-chan []*Test {
	out := make(chan []byte)
	var wg sync.WaitGroup

	output := func(ch <-chan []byte) {
		for n := range ch {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go output(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func goTestWrapper(*Test) <-chan *Test {
	out := make(chan *Test)
	go func() {
		out <- output
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
