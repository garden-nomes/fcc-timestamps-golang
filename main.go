package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

var dateFmt = "January 2, 2006"

type output struct {
	Unix *int64  `json:"unix"`
	Date *string `json:"date"`
}

func emptyResponse() output {
	return output{nil, nil}
}

func makeResponse(t time.Time) output {
	unix := t.Unix()
	natural := t.Format(dateFmt)
	return output{&unix, &natural}
}

func convert(input string) (time.Time, error) {
	if input == "" {
		return time.Time{}, errors.New("empty input string")
	}

	var unixTime = regexp.MustCompile(`^[0-9]+$`)
	if unixTime.MatchString(input) {
		unix, _ := strconv.ParseInt(input, 10, 64)
		return time.Unix(unix, 0), nil
	}

	t, err := time.Parse(dateFmt, input)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := convert(r.URL.Query().Get("time"))

		if err != nil {
			json.NewEncoder(w).Encode(emptyResponse())
		} else {
			json.NewEncoder(w).Encode(makeResponse(t))
		}
	})

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if err != nil {
		port = 6000
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
