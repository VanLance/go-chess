
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type JsonBoard struct {
	Square map[string] GamePiece
}

type Test struct{
	TestMap
  Testing string
}

type TestMap struct{
	TestNest map[string]Position2
}

type Position2 struct{
	X int
	Y int
}

func handler(w http.ResponseWriter, r *http.Request) {
	testMap := TestMap{make(map[string]Position2)}
	testMap.TestNest["test"] = Position2{X:1,Y:2}  
	test := Test{TestMap: testMap, Testing: "abc" }

	// chess := createChess()
	// mapTest := map[string]Position {"test": {3,8}}
	w.Header().Set("Content-Type","application/json")
	res, err := json.Marshal(test)
	fmt.Println()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}