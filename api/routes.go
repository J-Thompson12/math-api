package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justin/sojern-code-challenge/equations"
)

type params struct {
	List       []int
	Quantifier int
}

func routes(router *mux.Router) {
	router.Path("/min").Methods(http.MethodGet).HandlerFunc(min)
	router.Path("/max").Methods(http.MethodGet).HandlerFunc(max)
	router.Path("/avg").Methods(http.MethodGet).HandlerFunc(average)
	router.Path("/median").Methods(http.MethodGet).HandlerFunc(median)
	router.Path("/percentile").Methods(http.MethodGet).HandlerFunc(percentile)
}

func httpErrorHandler(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func min(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/text")

	var req params

	err := json.NewDecoder(r.Body).Decode(&req)
	httpErrorHandler(w, err)

	result, err := equations.Min(req.List, req.Quantifier)
	httpErrorHandler(w, err)

	body, err := json.Marshal(result)
	httpErrorHandler(w, err)
	w.Write(body)
}

func max(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/text")

	var req params

	err := json.NewDecoder(r.Body).Decode(&req)
	httpErrorHandler(w, err)

	result, err := equations.Max(req.List, req.Quantifier)
	httpErrorHandler(w, err)

	body, err := json.Marshal(result)
	httpErrorHandler(w, err)
	w.Write(body)
}

func average(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/text")

	var req params

	err := json.NewDecoder(r.Body).Decode(&req)
	httpErrorHandler(w, err)

	result, err := equations.Average(req.List)
	httpErrorHandler(w, err)

	body, err := json.Marshal(result)
	httpErrorHandler(w, err)
	w.Write(body)

}

func median(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/text")

	var req params

	err := json.NewDecoder(r.Body).Decode(&req)
	httpErrorHandler(w, err)

	result, err := equations.Median(req.List)
	httpErrorHandler(w, err)

	body, err := json.Marshal(result)
	httpErrorHandler(w, err)
	w.Write(body)

}

func percentile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/text")

	var req params

	err := json.NewDecoder(r.Body).Decode(&req)
	httpErrorHandler(w, err)

	result, err := equations.Percentile(req.List, req.Quantifier)
	httpErrorHandler(w, err)

	body, err := json.Marshal(result)
	httpErrorHandler(w, err)
	w.Write(body)

}
