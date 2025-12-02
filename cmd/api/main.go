package main

import (
	"GoPracZ13-BorisovDA/internal/work"
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"
)

func main() {

	// Эндпоинт, вызывающий “тяжёлую” работу.
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		defer work.TimeIt("Fib(38)")()
		res := work.Fib(38)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte((fmtInt(res))))
	})

	log.Println("Server on :8080; pprof on /debug/pprof/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fmtInt(v int) string { return fmt.Sprintf("%d\n", v) }
