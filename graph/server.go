package main

import (
	"gosocial/graph"
	
	"log"
	"net/http"	
	"github.com/99designs/gqlgen/handler"		
)

func main() {
	profileURL:="localhost:8000"
	postURL:="localhost:8001"
	relationURL:="localhost:8002"
	feedURL:="localhost:8003"
	srv,err := graph.NewGraphQLServer(profileURL,relationURL,postURL,feedURL)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/graphql", handler.GraphQL(srv.ToExecutableSchema()))
	http.Handle("/playground", handler.Playground("gosocial", "/graphql"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":"+"8080", nil))
}
