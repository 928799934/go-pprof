package pprof

import (
	"net/http"
	"runtime/pprof"
)

func mem(wr http.ResponseWriter, r *http.Request) {
	wr.Header().Set("X-Content-Type-Options", "nosniff")
	wr.Header().Set("Content-Type", "application/octet-stream")
	wr.Header().Set("Content-Disposition", `attachment; filename="memory"`)
	_ = pprof.WriteHeapProfile(wr)
	//http.Error(wr, http.StatusText(http.StatusOK), http.StatusOK)
}
