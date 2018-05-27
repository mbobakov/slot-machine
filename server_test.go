package main

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mbobakov/slot-machine/account"
	"github.com/stretchr/testify/assert"
)

//nolint
func TestServer_postSlotSpins(t *testing.T) {
	tests := []struct {
		name           string
		expectCode     int
		expectBody     string
		inputSlotIDCtx string
		inputMethod    string
		inputURL       string
		inputBody      string
	}{

		{"OK", 200, "{\"Total\":100,\"Spins\":[{\"Total\":100,\"Type\":\"main\",\"Stops\":[0,1,2,3,4]}],\"JWT\":\"signed\"}\n",
			"test",
			"POST",
			"127.0.0.1/api/machines/test/spins",
			`{"jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiJtYm9iYWtvdiIsIkNoaXBzIjoxMDAwLCJCZXQiOjEwMH0.m8IAZCAXL2A51YN-8nTUb0RHnOVlb3rWluZJ0O236jw"}`,
		},

		{"NoSlot", 404, "Slot 'noslot' not found. Request ID = treqID",
			"noslot",
			"POST",
			"127.0.0.1/api/machines/noslot/spins",
			`{"jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiJtYm9iYWtvdiIsIkNoaXBzIjoxMDAwLCJCZXQiOjEwMH0.m8IAZCAXL2A51YN-8nTUb0RHnOVlb3rWluZJ0O236jw"}`,
		},

		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{Account: &mockAccount{}, Slots: map[string]spinner{"test": &mockSlot{}}}
			req, err := http.NewRequest(tt.inputMethod, tt.inputURL, bytes.NewBufferString(tt.inputBody))
			if err != nil {
				t.Error(err)
				t.Fail()
			}

			cctx := chi.NewRouteContext()
			cctx.RoutePatterns = []string{"/api/machines/{slotID}/spins"}
			cctx.URLParams = chi.RouteParams{Keys: []string{"slotID"}, Values: []string{tt.inputSlotIDCtx}}
			ctx := context.WithValue(context.Background(), chi.RouteCtxKey, cctx)
			ctx = context.WithValue(ctx, middleware.RequestIDKey, "treqID")
			req = req.WithContext(ctx)
			w := httptest.NewRecorder()
			s.postSlotSpins(w, req)
			assert.Equal(t, tt.expectCode, w.Code)
			assert.Equal(t, tt.expectBody, w.Body.String())
		})
	}
}

type mockAccount struct{}

func (m *mockAccount) InfoFromContext(ctx context.Context) (*account.Info, error) {
	return &account.Info{UID: "test", Bet: 100, Chips: 1000}, nil
}
func (m *mockAccount) SignInfo(i *account.Info) (string, error) {
	return "signed", nil
}

type mockSlot struct{}

func (m *mockSlot) Spin(int) (int, bool, [5]uint8) {
	return 100, false, [5]uint8{0, 1, 2, 3, 4}
}
