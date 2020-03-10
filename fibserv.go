package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"
	"time"
)

func fib (n int) int {
	if (n > 2) {
		return fib(n-1) + fib(n-2)
	} else {
		return 1
	}
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
	var a = 1
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		nstr, ok := r.URL.Query()["n"]
		if ok {
			n, _ := strconv.Atoi(nstr[0])
			start := time.Now()
			fibn := fib(n)
			elapsed := time.Since(start)
			fmt.Fprintf(w, gen_html(n, fmt.Sprintf("Fibonnaci number #%v is %v. (Serving request #%v, took %v)", n, fibn, a, elapsed)))
			a++
		} else {
			fmt.Fprintf(w, gen_html(20, ""))
		}
	})

	http.ListenAndServe(":8080", nil)
}
