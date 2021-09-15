package graph

import (
	"gosocial/feed"
	"gosocial/graph/generated"
	"gosocial/post"
	"gosocial/profile"
	"gosocial/relation"

	"github.com/99designs/gqlgen/graphql"
)

type Server struct {
	profileClient *profile.Client
	postClient *post.Client
	relationClient *relation.Client
	feedClient *feed.Client
}

func NewGraphQLServer(profileURL,relationURL,postURL ,feedURL string) (*Server,error){
	profileClient,err:=profile.NewClient(profileURL)
	if err != nil {
		return nil,err
	}
	relationClient,err:=relation.NewClient(relationURL)
	if err != nil {
		profileClient.Close()
		return nil,err
	}
	postClient,err:=post.NewClient(postURL)
	if err != nil {
		profileClient.Close()
		relationClient.Close()
		return nil,err
	}
	feedClient,err:=feed.NewClient(feedURL)
	if err != nil {
		profileClient.Close()
		relationClient.Close()
		postClient.Close()
		return nil,err
	}
	return &Server{
		profileClient,
		postClient,
		relationClient,		
		feedClient,
	},nil
}

func (s *Server) Mutation() generated.MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() generated.QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: s,
	})
}