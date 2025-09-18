package main

import (
	"net/http"

	_ "github.com/eyanshu1997/tunnlrx/common/proto"
	"github.com/eyanshu1997/tunnlrx/common/serviceutils"
)

func main() {
	serviceutils.InitServiceUtils()
	serviceutils.Log.Info("🚀 TunnlrX server starting on :8080")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.ListenAndServe(":8080", nil)
}
