package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

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

func main() {

	var requestCount = 1
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	fmt.Printf("Serving on port 8080...\n")
	http.ListenAndServe(":8080", nil)
}
