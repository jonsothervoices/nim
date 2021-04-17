package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	fmt.Fprintf(w, "%v", rules)
}

func playPostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("unable to read request, %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var daGame game
	err = json.Unmarshal(body, &daGame)
	if err != nil {
		fmt.Printf("unable to read request body, %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = daGame.play()
	if err != nil {
		fmt.Printf("%s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(daGame); err != nil {
		fmt.Printf("could not write response, %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func playGetHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var daGame game
	coins := query["coins"]
	var err error
	if len(coins) > 0 {
		daGame.Coins, err = strconv.Atoi(coins[0])
		if err != nil {
			fmt.Printf("Coins must be an integer, %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	move := query["move"]
	if len(move) > 0 {
		daGame.Move, err = strconv.Atoi(move[0])
		if err != nil {
			fmt.Printf("Move must be an integer, %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	err = daGame.play()
	if err != nil {
		fmt.Printf("%s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(daGame); err != nil {
		fmt.Printf("could not write response, %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func varPlayGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var daGame game
	coins := vars["coins"]
	var err error
	if len(coins) > 0 {
		daGame.Coins, err = strconv.Atoi(coins)
		if err != nil {
			fmt.Printf("Coins must be an integer, %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	move := vars["move"]
	if len(move) > 0 {
		daGame.Move, err = strconv.Atoi(move)
		if err != nil {
			fmt.Printf("Move must be an integer, %s", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	err = daGame.play()
	if err != nil {
		fmt.Printf("%s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(daGame); err != nil {
		fmt.Printf("could not write response, %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
