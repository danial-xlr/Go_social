package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"	
	"gosocial/graph/model"
	"log"
	"time"
)

func (r *mutationResolver) CreatProfile(ctx context.Context, profile model.ProfileInput) (*model.Profile, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	p, err := r.server.profileClient.PostProfile(ctx, profile.UserName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.Profile{
		ID:       p.ID,
		UserName: p.UserName,
	}, nil
}

func (r *mutationResolver) CreatPost(ctx context.Context, post model.PostInput) (*model.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	p, err := r.server.postClient.CreatPost(ctx, post.Body, post.ProfileID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.Post{
		ID:        p.ID,
		Body:      p.Body,
		ProfileID: p.ProfileId,
		CreatAt:   p.CreatAt,
	}, nil
}

func (r *mutationResolver) CreatRlation(ctx context.Context, relation model.RelationInput) (*model.Relation, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	p, err := r.server.relationClient.CreatRelation(ctx, relation.ProfileID, relation.FollowingID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.Relation{
		ID:          p.ID,
		ProfileID:   p.ProfileId,
		FollowingID: p.FollowingId,
	}, nil
}

func (r *mutationResolver) CreatFeed(ctx context.Context, feed model.FeedInput) (*model.Feed, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	p, err := r.server.feedClient.CreatFeed(ctx, feed.ProfileID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	posts := []*model.FeedPost{}
	for _, p := range p.Posts {
		posts = append(posts, &model.FeedPost{
			ID:        p.ID,
			ProfileID: p.ProfileId,
			Body:p.Body,
			CreatAt:p.CreatAt,
		})
	}
	return &model.Feed{
		ID:        p.ID,
		ProfileID: p.ProfileId,
		Posts:     posts,
	}, nil
}

func (r *queryResolver) Profiles(ctx context.Context, id *int) ([]*model.Profile, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if id != nil {
		p, err := r.server.profileClient.GetProfile(ctx, *id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return []*model.Profile{&model.Profile{
			ID:       p.ID,
			UserName: p.UserName,
		}}, nil
	}
	p, err := r.server.profileClient.GetAllProfiles(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	profiles := []*model.Profile{}
	for _, a := range p {
		profiles = append(profiles, &model.Profile{
			ID:       a.ID,
			UserName: a.UserName,
		})
	}
	return profiles, nil
}

func (r *queryResolver) Posts(ctx context.Context, id *int, profileID *int) ([]*model.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if id == nil {
		p, err := r.server.postClient.GetPostsByProfileId(ctx, *profileID)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		posts := []*model.Post{}
		for _, a := range p {
			posts = append(posts, &model.Post{
				ID:        a.ID,
				ProfileID: a.ProfileId,
				Body:      a.Body,
				CreatAt:   a.CreatAt,
			})
		}
	}
	p, err := r.server.postClient.GetPostById(ctx, *id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return []*model.Post{&model.Post{
		ID:        p.ID,
		ProfileID: p.ProfileId,
		Body:      p.Body,
		CreatAt:   p.CreatAt,
	}}, nil
}

func (r *queryResolver) Relations(ctx context.Context, profileID *int) ([]int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rel, err := r.server.relationClient.GetProfileFollower(ctx, *profileID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	follower := []int{}
	for _, f := range rel {
		follower = append(follower, int(f))
	}
	return follower, nil
}

func (r *queryResolver) Feeds(ctx context.Context, profileID *int) (*model.Feed, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	f, err := r.server.feedClient.GetProfileFeed(ctx, *profileID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	posts := []*model.FeedPost{}
	for _, p := range f.Posts {
		posts = append(posts, &model.FeedPost{
			ID:        p.ID,
			ProfileID: p.ProfileId,
			Body:p.Body,
			CreatAt:p.CreatAt,
		})
	}
	return &model.Feed {
		ID:        f.ID,
		ProfileID: f.ProfileId,
		Posts:     posts,
	}, nil
}

type mutationResolver struct{ 
	*Resolver
	server *Server
 }

type queryResolver struct{ 
	*Resolver
	server *Server
}

var (
	ErrInvalidParameter = errors.New("invalid parameter")
)
