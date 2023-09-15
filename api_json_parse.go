package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

// Given this API - https://dummyjson.com/comments
// Parse the response of this API and print the following results:
// group the comments by user
// sort the user based on the comments count in descending order

type CommentResponse struct { // the field of struct for which we want to parse json should always be exported
	Comments []Comment `json:"comments"`
	Total int `json:"total"`
	Skip int `json:"skip"`
	Limit int `json:"limit"`
}

type User struct { // why exported fields? Read here - https://stackoverflow.com/questions/11126793/json-and-dealing-with-unexported-fields
	Id int `json:"id"`
	Username string `json:"username"`
}

type Comment struct {
	Id int `json:"id"`
	Body string `json:"body"`
	PostId int `json:"postId"`
	User User `json:"user"`
}

func getResponse(url string) (CommentResponse, error) {
	var commentResponse CommentResponse
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return CommentResponse{}, err
	}
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return CommentResponse{}, err
	}

	if err = json.Unmarshal(responseData, &commentResponse); err != nil {
		fmt.Println("Error parsing json",err.Error())
	}
	//fmt.Printf("Comment Response Struct: %+v", commentResponse)
	return commentResponse, nil
}


func main() {

	url := "https://dummyjson.com/comments"
	commentResponse, err := getResponse(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	userComments := make(map[int][]Comment)

	for _, comment := range commentResponse.Comments {
		if _, ok := userComments[comment.User.Id]; ok {
			userComments[comment.User.Id] = append(userComments[comment.User.Id], comment)
		} else {
			userComments[comment.User.Id] = []Comment{comment}
		}
	}

	var userList []int

	for user, cm := range userComments {
		fmt.Printf("User Id: %d has following comments:\n", user)
		userList = append(userList, user)
		for _, comment := range cm {
			fmt.Printf("%+v\n", comment)
		}
		fmt.Println()
	}

	// sort the user based on comment count in descending order
	// this is how you sort the keys of the map based on the values in decreasing order
	sort.SliceStable(userList, func(i, j int) bool {
		return len(userComments[userList[i]]) > len(userComments[userList[j]])
	})

	fmt.Println(userList)
}