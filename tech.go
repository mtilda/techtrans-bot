package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
)

type Body struct {
	Results [][]string `json:"results"`
	Count   int        `json:"count"`
	Total   int        `json:"total"`
	Perpage int        `json:"perpage"`
	Page    int        `json:"page"`
}

func FetchTech() string {

	res, err := http.Get("https://api.nasa.gov/techtransfer/patent/?wing&api_key=DEMO_KEY")
	if err != nil {
		Error(err)
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		if nil != err {
			Error(err)
		}

		// Inform(string(data))

		var body Body

		err = json.Unmarshal(data, &body)
		if nil != err {
			// always throws the following error:
			// json: cannot unmarshal number into Go struct field Body.results of type string
			// I think because it si unmarshalling an array of mixed data types.
			// We do not need the one that is erroring out.
			// Error(err)
		}

		index := rand.Intn(body.Count)

		return removeTags(body.Results[index][2]) + "\n" + body.Results[index][10]
	}
	return ""
}

func removeTags(input string) string {
	re := regexp.MustCompile("<[^>]*>")
	return strings.Join(re.Split(input, -1), "")
}
