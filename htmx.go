package htmx

import (
	"net/http"
	"strings"
	"time"
)

var (
	DefaultSwapDuration = time.Duration(0 * time.Millisecond)
	DefaultSettleDelay  = time.Duration(20 * time.Millisecond)
)

type (
	HTMX struct{}
)

func New() *HTMX {
	return &HTMX{}
}

func (h *HTMX) NewHandler(w http.ResponseWriter, r *http.Request) *Handler {
	return &Handler{
		w:       w,
		r:       r,
		request: h.HxHeader(r.Context()),
		response: &HxResponseHeader{
			headers: w.Header(),
		},
	}
}

func HxStrToBool(str string) bool {
	if strings.EqualFold(str, "true") {
		return true
	}

	return false
}

func HxBoolToStr(b bool) string {
	if b {
		return "true"
	}

	return "false"
}
