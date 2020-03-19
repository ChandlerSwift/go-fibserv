package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"testing"
)

func TestAPI(t *testing.T) {

	cases := []struct{ index, value int }{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{20, 6765},
		{30, 832040},
		{40, 102334155},
	}

	for _, c := range cases {
		t.Run(strconv.Itoa(c.index), func(t *testing.T) {

			// Create HTTP Request
			req, err := http.NewRequest("GET", "/api", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Add URL Params
			q := req.URL.Query()
			q.Add("n", strconv.Itoa(c.index))
			req.URL.RawQuery = q.Encode()

			// Fetch response into rr
			rr := httptest.NewRecorder()
			serveAPI(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			// Create regex to check response body against
			re, err := regexp.Compile(fmt.Sprintf(`^{"Index":%v,"Value":%v,"CalculationTime":\d+}$`, c.index, c.value))
			if err != nil {
				t.Fatalf("Could not compile regular expression: %v", err)
			}

			// Check response body against regex
			if !re.Match(rr.Body.Bytes()) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), fmt.Sprintf(`{"Index":%v,"Value":%v,"CalculationTime":###}`, c.index, c.value))
			}
		})
	}
}
