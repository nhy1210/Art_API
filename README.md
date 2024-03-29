# Art_API
This project is about buildinga RESTful JSON API using Golang. 
It uses the gorrilla mux package for HTTP request multiplexer. Like the standard http.ServeMux, mux. Router matches incoming requests to a list of registered routes and calls a handler for the route that matches the URL or other conditions. 
We have 3 main endpoints to handle the request. 

I. Create an article with the POST method, store it within the service

II. Get an article using the articles id with GET method

III. Get all the tag name on the given date and some additional data using GET method.

All the articles and the data requested in endpoint number (III) are written in JSON format.
Below are step by step guideline for running this project. 

System requirement and installation
Article format and tag name information return format
Intergration test - Postman tool 


1. System requirement: 
Any operating system with Golang and test tool Postinstalled. If your OS is macOS, refer to installation guideline in repository.
2. Article JSON format

  a)An article has some attibutes below:
  
	{
		"id":        "1",
		"title":     "latest science shows that potato chips are better for you than sugar",
		"date": 	 "2016-09-22",
		"body": 	 "some text, potentially containing simple markup about how potato chips are great",
		"tags":       ["health", "fitness", "science"],
	} 
	
  b) A tagEvent has below JSON format
  
	{
		"tag":       		"health",
		"count":     		"17",
		"articles": 		["1", "7"],   	//list of article id for last 10 articles for that day 
		"related_tags":     ["fitness", "science"], //contains a list of tags that are on the articles that the current tag is on       for the same day _ No duplicate
	}
	
 3. Handling the endpoint requests

a) Create an article with the POST method. Below is example of creating article.


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
}
    
 b) Get an article using the articles id with GET method


	func getOneEvent(w http.ResponseWriter, r *http.Request) {
	articleID := mux.Vars(r)["id"]
	for _, singleArticle := range articlesData {
		if singleArticle.ID == articleID {
			json.NewEncoder(w).Encode(singleArticle)
		}
	}
}

c) Get all the tag name on the given date and some additional data using GET method.



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

4) Test the solution with POSTMAN tool. Installation guideline refer to this link https://www.code2bits.com/how-to-install-postman-on-macos-using-homebrew/


