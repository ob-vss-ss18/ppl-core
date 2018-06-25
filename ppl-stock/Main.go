package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"io/ioutil"
	"net/http"
	"github.com/ob-vss-ss18/ppl-stock/schema"
)

func main() {

	http.HandleFunc("/graphql", func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
		// httpRequest.URL.Query().Get("query"):
		// returns "message" (without the qoutes) for the example below
		bodyBytes, error := ioutil.ReadAll(httpRequest.Body)
		if error == nil {
			bodyString := string(bodyBytes[:])
			result := executeQuery(bodyString)
			json.NewEncoder(httpResponseWriter).Encode(result)
		} else {
			fmt.Println(error)
		}
		httpRequest.Body.Close()
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Curl test: curl --data 'query{Ski(id:10){id gender category length price_new}}' http://localhost:8080/graphql")
	fmt.Println("Curl test: curl --data 'query{Stick(id:5){id gender price_new availability}}' http://localhost:8080/graphql")
	http.ListenAndServe(":8080", nil)

}

/*
	A small function to which will call graphql.Do(), which will handle resolving the
	request and returns the results.
*/
func executeQuery(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema: schema.PPLStockSchema,
		// this variable is set to "message" for the example below
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
