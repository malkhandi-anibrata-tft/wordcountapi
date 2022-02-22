package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

var inputdata string

func indexHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello User!") // write data to response

}

func inputHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("input:", r.Form["name"])
		getdata := r.Form["name"]
		//===============process data input========
		process(getdata)

	}
}

func outputHandler(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
	//id := "xyzabc"
	///import user input here
	if userinput == "" {
		http.Error(w, "The id query parameter is missing", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "<h1>The user id is: %s</h1>", userinput)
}

func process(getdata []string) {

	str, err := json.Marshal(getdata)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	userinput := string(str[:])

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}

	userinput = re.ReplaceAllString(userinput, " ")
	userinput = strings.ToLower(userinput)

	//====================count word and append====================

	for word, occur := range countsimilarword(userinput) {

		occurance := strconv.Itoa(occur)

		var userinput []string

		userinput = append(userinput, word, occurance)

		fmt.Println(userinput) //need to print

		//inputdata = userinput
		//hii(userinput)

	}

}

// func hii(userinput []string) {

// 	if userinput == "" {
// 		http.Error(w, "The id query parameter is missing", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Fprintf(w, "<h1>The user id is: %s</h1>", userinput)

// }

// ======================count similar word===========
func countsimilarword(st string) map[string]int {

	words := strings.Fields(st)

	m := make(map[string]int)
	for _, word := range words {
		_, ok := m[word]
		if !ok {
			m[word] = 1
		} else {
			m[word]++
		}
	}

	return m
}

func main() {
	http.HandleFunc("/index", indexHandler) // welcome page

	http.HandleFunc("/input", inputHandler)
	http.HandleFunc("/output", outputHandler)

	err := http.ListenAndServe(":8000", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// fmt.Println(userinput)
}
