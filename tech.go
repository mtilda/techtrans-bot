// Fetch and manipulate data from NASA's Technology Transfer API

package main

import (
	"encoding/json"
	"errors"
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
	PerPage int        `json:"perpage"`
	Page    int        `json:"page"`
}

/* Fetch a random technology
   Return the title and image URL
*/
func FetchTech() (string, error) {

	res, err := http.Get("https://api.nasa.gov/techtransfer/patent/?engine&api_key=DEMO_KEY")
	if err != nil {
		Error(err)
	} else {
		data, err := ioutil.ReadAll(res.Body)
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

		return removeTags(body.Results[index][2]), nil // + " " + body.Results[index][10], nil
	}

	return "", errors.New("Unable to fetch tech")
}

/* Remove all substrings that look like HTML tags
 */
func removeTags(input string) string {
	re := regexp.MustCompile("<[^>]*>")
	return strings.Join(re.Split(input, -1), "")
}
