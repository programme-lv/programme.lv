package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Execution struct {
	LangId  string
	SrcCode string
	StdIn   string
}

func enqueueExecution(w http.ResponseWriter, r *http.Request) {
	// TODO read https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
	var e Execution
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Execution: %+v", e)
}

func listExecutions(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func getExecution(w http.ResponseWriter, r *http.Request) {
	// TODO
}
