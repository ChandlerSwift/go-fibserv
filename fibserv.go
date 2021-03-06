package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/chandlerswift/fib"
)

var requestCount = 1
var port int

type fibResponse struct {
	Index           int
	Value           int
	CalculationTime time.Duration
}

func timeFib(n int) fibResponse {
	start := time.Now()
	fibn := fib.Fib(n)
	elapsed := time.Since(start)
	return fibResponse{n, fibn, elapsed}
}

func genHTML(current int, content string) string {
	hostname, _ := os.Hostname()
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
	<head>
		<title></title>
	</head>
	<body>
		<form>
			<input type='number' value='%v' name='n'>
			<input type='submit'>
		</form>
		<br>
		%v
		<br><br><br><br>
		Served from %v
	</body>
</html>`, current, content, hostname)
}

func servePage(w http.ResponseWriter, r *http.Request) {
	nstr, hasN := r.URL.Query()["n"]
	if !hasN {
		fmt.Fprintf(w, genHTML(20, ""))
		return
	}

	n, err := strconv.Atoi(nstr[0])
	if err != nil {
		fmt.Fprintf(w, genHTML(20, "ERROR: %v.\n"))
		return
	}

	f := timeFib(n)
	fmt.Fprintf(w, genHTML(n, fmt.Sprintf("Fibonnaci number #%v is %v. (Serving request #%v, took %v)", f.Index, f.Value, requestCount, f.CalculationTime)))
	requestCount++
}

func serveAPI(w http.ResponseWriter, r *http.Request) {
	nstr, hasN := r.URL.Query()["n"]
	if !hasN {
		http.Error(w, "parameter n (fib index) is required", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(nstr[0])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	f, err := json.Marshal(timeFib(n))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, string(f))
	requestCount++
}

func main() {
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.Parse()

	if port < 1 || port > 65535 {
		log.Fatalf("Invalid port selected")
	}

	http.HandleFunc("/", servePage)
	http.HandleFunc("/api", serveAPI)

	fmt.Printf("Serving on port %v...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
