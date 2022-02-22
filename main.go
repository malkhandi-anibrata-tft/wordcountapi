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

var inputdatas []string

func indexHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello User!") // write data to response

}

func inputHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		getdata := r.Form["name"]
		//===============process data input========
		process(getdata)

	}
}

func outputHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The count is: %s", inputdatas)
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

	fmt.Println(userinput)

	//====================count word and append====================
	i := MaxWordsInSentences(userinput)
	fmt.Println(i)

	for word, occur := range countsimilarword(userinput) {

		//  ​sort​.​Slice​(​WC​, ​func​(​i​, ​j​ ​int​) ​bool​ {
		// 		​                ​return​ ​WC​[​i​].​count​ ​>​ ​WC​[​j​].​count
		// 		​        })

		occurance := strconv.Itoa(occur)

		inputdatas = append(inputdatas, "word:", word, "occurs", occurance, "times,")

	}

}

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

	fmt.Println(inputdatas)
}

// ​func​ ​PostRequest​(​w​ http.​ResponseWriter​, ​r​ ​*​http.​Request​) {
// 	​        ​content​, ​_​ ​:=​ ​ioutil​.​ReadAll​(​r​.​Body​)
// 	​         ​s​ ​:=​ []​string​{​string​(​content​)}
// 	​        ​new_​ ​:=​ ​WordC​(​s​)
// 	​        ​WC​ ​:=​ ​make​([]​CountW​, ​0​, ​len​(​new_​))
// 	​        ​for​ ​key​, ​val​ ​:=​ ​range​ ​new_​ {
// 	​                ​WC​ ​=​ ​append​(​WC​, ​CountW​{​word​: ​key​, ​count​: ​val​})
// 	​        }

// 	​        ​sort​.​Slice​(​WC​, ​func​(​i​, ​j​ ​int​) ​bool​ {
// 	​                ​return​ ​WC​[​i​].​count​ ​>​ ​WC​[​j​].​count
// 	​        })
// 	​        ​for​ ​i​ ​:=​ ​0​; ​i​ ​<​ ​len​(​WC​) ​&&​ ​i​ ​<​ ​10​; ​i​++​ {
// 	​                ​count_​ ​:=​ ​CountWs​{​CountW_​{​WC​[​i​].​word​, ​WC​[​i​].​count​}}
// 	​                ​json​.​Marshal​(​count_​)
// 	​                ​json​.​NewEncoder​(​w​).​Encode​(​count_​)
// 	​        }
// 	​}
func MaxWordsInSentences(S string) (result int) {

	r, _ := regexp.Compile("[.||?||!]")
	count := strings.Count(S, ".") + strings.Count(S, "!") + strings.Count(S, "?") // Total sentaces

	for i := 0; i < count; i++ {
		sentence := r.Split(S, count)[i]
		splitSentence := strings.Split(sentence, " ")

		var R []string
		for _, str := range splitSentence {
			if str != "" {
				R = append(R, str)
			}
		}

		if len(R) > result {
			result = len(R)
		}
	}

	return

}
