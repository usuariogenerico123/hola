package fbok

import (
	"encoding/json"
	"fmt"
)




type Comment struct{

	Entry []Info
	Object string
}

type Info struct{

	Id string 
	Time int
	Changes []Values
}

type Values struct {
	Value Value
	Field string

}

type Value struct {
	From FromUser
	Post Post
	Message string `json:"message"`
    PostId string `json:"post_id"`
	CommentId string `json:"comment_id"`
    CreatedTime int `json:"created_time"`
    Item string `json:"item"`
	ParentId string `json:"parent_id"`
    Verb string `json:"verb"`
}

type FromUser struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type Post struct{
	StatusType string `json:"status_type"`
	IsPublished bool `json:"is_published"`
	UpdatedTime string `json:"updated_time"`
    PermalinkUrl string `json:"permalink_url"`
    PromotionStatus string `json:"promotion_status"`
	Id string `json:"id"`
}




func Unm(data string)Comment{
	var comment Comment

	var resp error = json.Unmarshal([]byte(data), &comment)
	if(resp != nil){
		fmt.Println(resp.Error())
	}
	return comment
	
}