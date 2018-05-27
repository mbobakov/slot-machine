package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Server service request to slots
type Server struct {
	Slots   map[string]spinner
	Account accountRepo
}
type resp struct {
	Total int
	Spins []spinResult
	JWT   string
}
type spinResult struct {
	Total int
	Type  string
	Stops [5]uint8
}

func (s *Server) postSlotSpins(w http.ResponseWriter, r *http.Request) {

	slotID := chi.URLParam(r, "slotID")
	if slotID == "" {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("[ERROR] reqID='%s' no slotID for request", middleware.GetReqID(r.Context()))
		w.Write([]byte("internal error. Request ID = " + middleware.GetReqID(r.Context()))) // nolint: errcheck
		return
	}
	slot, ok := s.Slots[slotID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("[ERROR] reqID='%s' Slot ('%s') not found ", middleware.GetReqID(r.Context()), slotID)
		w.Write([]byte("Slot '" + slotID + "' not found. Request ID = " + middleware.GetReqID(r.Context()))) // nolint: errcheck
		return
	}
	info, err := s.Account.InfoFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("[ERROR] reqID='%s' Account info error", middleware.GetReqID(r.Context()))
		w.Write([]byte("Account error. Request ID = " + middleware.GetReqID(r.Context()))) // nolint: errcheck
		return
	}
	if info.Chips < info.Bet {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("[ERROR] reqID='%s' Insufficient сhips.", middleware.GetReqID(r.Context()))
		w.Write([]byte("Insufficient сhips. Request ID = " + middleware.GetReqID(r.Context()))) // nolint: errcheck
		return
	}
	var (
		total int
		ss    []spinResult
	)
	bonus, freeSpin, stops := slot.Spin(1)
	total += bonus
	ss = append(ss, spinResult{Stops: stops, Total: bonus, Type: "main"})
	for freeSpin {
		freeSpin = false
		for i := 0; i < 10; i++ {
			b, f, st := slot.Spin(3)
			total += b
			ss = append(ss, spinResult{Stops: st, Total: b, Type: "free spin"})
			if f {
				freeSpin = true
			}
		}
	}
	info.Chips = info.Chips - info.Bet + total
	tok, err := s.Account.SignInfo(info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("[ERROR] reqID='%s'. Sign account error", middleware.GetReqID(r.Context()))
		w.Write([]byte("Sign account error. Request ID = " + middleware.GetReqID(r.Context()))) //nolint:errcheck
		return
	}
	render.JSON(w, r, resp{Spins: ss, Total: total, JWT: tok})
}
