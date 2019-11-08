# Art_API
This project is about buildinga RESTful JSON API using Golang
It uses the gorrilla mux package for HTTP request multiplexer. Like the standard http.ServeMux, mux.Router matches incoming requests to a list of registered routes and calls a handler for the route that matches the URL or other conditions. We have 3 main endpoints to handle the request. 
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
    a) Create an article with the POST method
    
