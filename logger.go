package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

type rlog struct {
	*logrus.Logger
}

func (r *rlog) NewLogEntry(req *http.Request) middleware.LogEntry {
	reqID := middleware.GetReqID(req.Context())
	nr := &rlog{r.WithField("reqID", reqID).Logger}
	return nr
}

func (r *rlog) Write(status, bytes int, elapsed time.Duration) {
	r.Infof("Request done in %s with '%d' status and %d bytes", elapsed, status, bytes)
}
func (r *rlog) Panic(v interface{}, stack []byte) {
	r.Errorf("Panic: %v. Stack: %s", v, stack)
}
