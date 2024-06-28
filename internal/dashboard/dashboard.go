package dashboard

import (
	"context"
	"github.com/jfyne/live"
	"html/template"
	"log"
	"math/rand/v2"
	"net/http"
	"time"
)

const tick = "tick"

var Matrix3x3 = newMatrix3x3()

func colors() []string {
	return []string{"green", "red"}
}

type matrix3x3 struct {
	Data [][]string
}

func newMatrix3x3() *matrix3x3 {
	return &matrix3x3{
		Data: [][]string{
			matrixRow(),
			matrixRow(),
			matrixRow(),
		},
	}
}

func matrixValue() string {
	return colors()[rand.IntN(2)]
}

func matrixRow() []string {
	return []string{matrixValue(), matrixValue(), matrixValue()}
}

type CronEngine struct {
	*live.HttpEngine
}

func NewCronEngine(h live.Handler) *CronEngine {
	e := &CronEngine{
		live.NewHttpHandler(live.NewCookieStore("session-name", []byte("weak-secret")), h),
	}
	return e
}

func (e *CronEngine) Start() {
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for {
			<-ticker.C
			Matrix3x3 = newMatrix3x3()
			e.Broadcast(tick, nil)

		}
	}()
}

func main() {
	t, err := template.ParseFiles("view.html")
	if err != nil {
		log.Fatal(err)
	}

	h := live.NewHandler(live.WithTemplateRenderer(t))

	// Set the mount function for this handler.
	h.HandleMount(func(ctx context.Context, _ live.Socket) (interface{}, error) {
		// This will initialise the counter if needed.
		return Matrix3x3, nil
	})

	// Client side events.

	h.HandleSelf(tick, func(ctx context.Context, _ live.Socket, _ any) (interface{}, error) {
		return Matrix3x3, nil
	})

	ce := NewCronEngine(h)
	ce.Start()

	// Run the server.
	http.Handle("/", ce)
	http.Handle("/live.js", live.Javascript{})
	http.Handle("/auto.js.map", live.JavascriptMap{})
	http.ListenAndServe(":8080", nil)
}
