package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"strconv"
	"github.com/gorilla/mux" //using gorilla web toolkit
)
//For convert time type
const (
    layoutISO = "2006-01-02"
    layoutUS  = "20060102"
)
/*
	// =========================================================================
	An article has some attibutes like below example:
	{
		"id":        "1",
		"title":     "latest science shows that potato chips are better for you than sugar",
		"date": 	 "2016-09-22",
		"body": 	 "some text, potentially containing simple markup about how potato chips are great",
		"tags":       ["health", "fitness", "science"],
	}
	// =========================================================================
*/
type article struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Date		string `json:"date"`
	Body 		string `json:"body"`
	Tags		[]string `json:"tags"`
}

/*
	// =========================================================================
	An tagEvent has JSON format like below example:
	{
		"tag":       		"health",
		"count":     		"17",
		"articles": 		["1", "7"],   	//list of article id for last 10 articles for that day 
		"related_tags":     ["fitness", "science"], //contains a list of tags that are on the articles that the current tag is on for the same day _ No duplicate
	}
	// ==========================================================================
*/
type tagEvent struct {
	Tag          	string `json:"tag"`
	Count       	string `json:"count"`
	Articles		[]string `json:"articles"`
	Related_tags	[]string `json:"realated_tags"`
}

// ==============================================================================
type articles []article
var articlesData articles
// ==============================================================================



// ==============================================================================
// POST /articles handles the receipt of article data in json format and store it  
func createArticle(w http.ResponseWriter, r *http.Request) {
	var newArticle article
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	err1 := json.Unmarshal(reqBody, &newArticle)
	if err1 !=nil{
		fmt.Println(err1)
	}
	articlesData = append(articlesData, newArticle)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newArticle)
}

// ==============================================================================
//GET /articles/{id} returns the JSON representation of the article
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	articleID := mux.Vars(r)["id"]

	for _, singleArticle := range articlesData {
		if singleArticle.ID == articleID {
			json.NewEncoder(w).Encode(singleArticle)
		}
	}
}


// ==============================================================================
/* 
	GET /tags/{tagName}/{date} will return the list of articles that have 
	that tag name on the given date and some summary data about that tag 
	for that day
*/
// ==============================================================================
func getTagNameOnDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagName := vars["tagName"]
	tempdate := vars["date"]
	//Convert date from "20160922" to "2016-09-22" to match with article.Date 
	t,_ := time.Parse(layoutUS, tempdate)
	tagDate := t.Format(layoutISO)

	//Init count value and declare a tagMap for storing related_tags, preventing duplicate articles
	var tagOnDate tagEvent
	count:= 0	
	tagMap := make(map[string]int)
	for _, singleArticle := range articlesData {
		if singleArticle.Date == tagDate {
			mymap:= make(map[string]int)
			for index, singleTag:= range singleArticle.Tags{
				mymap[singleTag]=index
			}
			if _,found := mymap[tagName]; found{
				count++
				if len(tagOnDate.Articles) <10 {
						tagOnDate.Articles = append(tagOnDate.Articles, singleArticle.ID)
					} else {
						tagOnDate.Articles = tagOnDate.Articles[1:]
						tagOnDate.Articles = append(tagOnDate.Articles, singleArticle.ID)
					}

				delete(mymap, tagName)
				for otherTag:= range mymap{
					if _,found := tagMap[otherTag]; found == false {
						tagMap[otherTag]++
					}
				}
			}
		}
	}
	tagOnDate.Tag = tagName
	tagOnDate.Count = strconv.Itoa(count)
	for reTag:= range tagMap{
		tagOnDate.Related_tags = append(tagOnDate.Related_tags, reTag)
	} 
	json.NewEncoder(w).Encode(tagOnDate)
}


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/articles", createArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/tags/{tagName}/{date}", getTagNameOnDate).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}