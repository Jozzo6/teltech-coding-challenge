package main

import (
	"coding_challenge/middleware"
	"coding_challenge/models"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	addr := os.Getenv("ADDR")

	mux := http.NewServeMux()

	mux.HandleFunc("/add", AddHandler)
	mux.HandleFunc("/subtract", SubtractHandler)
	mux.HandleFunc("/divide", DivideHandler)
	mux.HandleFunc("/multiply", MultiplyHandler)

	wrappedMux := middleware.NewLogger(middleware.NewCache(mux))

	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := GetParams(w, r)

	if err {
		return
	}

	result := models.Result{Action: "add", X: x, Y: y, Answer: x + y, Cached: false}
	middleware.SaveRequest(r.URL.String(), result)
	SendResponse(w, result)
}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := GetParams(w, r)

	if err {
		return
	}

	result := models.Result{Action: "subtract", X: x, Y: y, Answer: (x - y), Cached: false}
	middleware.SaveRequest(r.URL.String(), result)
	SendResponse(w, result)
}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := GetParams(w, r)

	if err {
		return
	}

	if y == 0 {
		SendErrorResponse(w, models.ErrorResponse{Message: "Cannot divide with 0"})
		return
	}

	result := models.Result{Action: "divide", X: x, Y: y, Answer: x / y, Cached: false}
	middleware.SaveRequest(r.URL.String(), result)
	SendResponse(w, result)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	x, y, err := GetParams(w, r)

	if err {
		return
	}

	result := models.Result{Action: "multiply", X: x, Y: y, Answer: x * y, Cached: false}
	middleware.SaveRequest(r.URL.String(), result)
	SendResponse(w, result)
}

func SendResponse(w http.ResponseWriter, result models.Result) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		SendErrorResponse(w, models.ErrorResponse{Message: "Failed to encode"})
	}
}

func SendErrorResponse(w http.ResponseWriter, err models.ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
}

func GetParams(w http.ResponseWriter, r *http.Request) (float64, float64, bool) {
	x, err := strconv.ParseFloat(r.URL.Query().Get("x"), 64)

	if err != nil {
		SendErrorResponse(w, models.ErrorResponse{Message: "X is missing"})
		return 0, 0, true
	}

	y, e := strconv.ParseFloat(r.URL.Query().Get("y"), 64)
	if e != nil {
		SendErrorResponse(w, models.ErrorResponse{Message: "Y is missing"})
		return 0, 0, true
	}

	return x, y, false
}
