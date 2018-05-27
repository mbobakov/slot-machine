package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jessevdk/go-flags"
	"github.com/mbobakov/slot-machine/account"
	"github.com/mbobakov/slot-machine/slot"
)

type spinner interface {
	Spin(int) (int, bool, [5]uint8)
}
type accountRepo interface {
	InfoFromContext(ctx context.Context) (*account.Info, error)
	SignInfo(i *account.Info) (string, error)
}

func main() {
	var opts struct {
		Verbose     bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
		Listen      string `short:"l" long:"listen" default:"127.0.0.1:8080" description:"Listen on this interface"`
		DebugListen string `long:"debug.listen" default:"127.0.0.1:6060" description:"Listen pprof on this interface"`
		Key         string `long:"jwt.key" default:"dummySecret" description:"A signing key for JWT operations"`
	}

	_, err := flags.Parse(&opts)

	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}
	acc := &account.Repo{SigningKey: []byte(opts.Key)}
	srv := &Server{
		Account: acc,
		Slots:   map[string]spinner{"atkins-diet": slot.NewAtkinsDiet()},
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	if opts.Verbose {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.RequestID)
	// Spin the slot-machine
	r.With(acc.Middleware, middleware.NoCache).
		Post("/api/machines/{slotID}/spins", srv.postSlotSpins)
	go func() {
		errd := http.ListenAndServe(opts.DebugListen, nil)
		if errd != nil {
			log.Printf("[ERROR] Couldn't start listen pprof interface. E: '%v'", err)
		}
	}()
	err = http.ListenAndServe(opts.Listen, r)
	if err != nil {
		log.Fatal(err)
	}
}
