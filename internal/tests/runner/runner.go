package runner

import (
	"codeberg.org/filipmnowak/beaver/internal/db/sqlite"
	. "codeberg.org/filipmnowak/beaver/internal/tests"
	"fmt"
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

func PersistResults(ch <-chan *Test, dbPath string) error {
	for t := range ch {
		// TODO: update in batches, timeout or closed channel
		db := sqlite.NewDB(nil, dbPath, "")
		input := []map[string]string{
			{"family": t.FQN[0], "_group": t.FQN[1], "test": t.FQN[2], "variant": t.FQN[3], "key": "/success", "value": fmt.Sprintf("%s", t.Variants[0].Success())},
			{"family": t.FQN[0], "_group": t.FQN[1], "test": t.FQN[2], "variant": t.FQN[3], "key": "/log", "value": string(t.Variants[0].Result.Log)},
			{"family": t.FQN[0], "_group": t.FQN[1], "test": t.FQN[2], "variant": t.FQN[3], "key": "/error", "value": fmt.Sprintf("%s", t.Variants[0].Result.Err)},
		}
		out, err := db.TransactUpserts(input, "test_results", "family, _group, test, variant, key")
		fmt.Printf("out: %s\n", out)
		fmt.Printf("err: %s\n", err)
	}
	return nil
}

func Run() {
}
