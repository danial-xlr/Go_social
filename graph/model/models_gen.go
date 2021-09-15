// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Feed struct {
	ID        int         `json:"id"`
	ProfileID int         `json:"profile_id"`
	Posts     []*FeedPost `json:"posts"`
}

type FeedInput struct {
	ProfileID int `json:"profile_id"`
}

type FeedPost struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	ProfileID int       `json:"profile_id"`
	CreatAt   time.Time `json:"creat_at"`
}

type Post struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	ProfileID int       `json:"profile_id"`
	CreatAt   time.Time `json:"creat_at"`
}

type PostInput struct {
	Body      string `json:"body"`
	ProfileID int    `json:"profile_id"`
}

type Profile struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
}

type ProfileInput struct {
	UserName string `json:"user_name"`
}

type Relation struct {
	ID          int `json:"id"`
	ProfileID   int `json:"profile_id"`
	FollowingID int `json:"following_id"`
}

type RelationInput struct {
	ProfileID   int `json:"profile_id"`
	FollowingID int `json:"following_id"`
}
