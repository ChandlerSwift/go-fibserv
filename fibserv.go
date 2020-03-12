package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

var requestCount = 1

func fibHelper(n int) int {
	if n > 2 {
		return fibHelper(n-1) + fibHelper(n-2)
	} else {
		return 1
	}
}

type FibResponse struct {
	Index           int
	Value           int
	CalculationTime time.Duration
}

func fib(n int) FibResponse {
	start := time.Now()
	fibn := fibHelper(n)
	elapsed := time.Since(start)
	return FibResponse{n, fibn, elapsed}
}

func gen_html(current int, content string) string {
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
	nstr, has_n := r.URL.Query()["n"]
	if !has_n {
		fmt.Fprintf(w, gen_html(20, ""))
		return
	}

	n, err := strconv.Atoi(nstr[0])
	if err != nil {
		fmt.Fprintf(w, gen_html(20, "ERROR: %v.\n"))
		return
	}

	f := fib(n)
	fmt.Fprintf(w, gen_html(n, fmt.Sprintf("Fibonnaci number #%v is %v. (Serving request #%v, took %v)", f.Index, f.Value, requestCount, f.CalculationTime)))
	requestCount++
}

func serveAPI(w http.ResponseWriter, r *http.Request) {
	nstr, has_n := r.URL.Query()["n"]
	if !has_n {
		http.Error(w, "parameter n (fib index) is required", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(nstr[0])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	f, err := json.Marshal(fib(n))

	fmt.Fprintf(w, string(f))
	requestCount++
}

func main() {

	http.HandleFunc("/", servePage)
	http.HandleFunc("/api", serveAPI)

	fmt.Printf("Serving on port 8080...\n")
	http.ListenAndServe(":8080", nil)
}
