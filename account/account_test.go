package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepo_SignInfo(t *testing.T) {
	tests := []struct {
		name   string
		expect string
		input  *Info
	}{
		{"simple", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVSUQiOiJtYm9iYWtvdiIsIkNoaXBzIjoxMDAwLCJCZXQiOjEwMH0.m8IAZCAXL2A51YN-8nTUb0RHnOVlb3rWluZJ0O236jw", &Info{Bet: 100, Chips: 1000, UID: "mbobakov"}},
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{SigningKey: []byte("dummySecret")}
			tk, err := r.SignInfo(tt.input)
			if err != nil {
				t.Error(err)
				t.Fail()
			}
			assert.Equal(t, tt.expect, tk)
		})
	}
}
