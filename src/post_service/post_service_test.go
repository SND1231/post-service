package post_service

import (
	"github.com/SND1231/post-service/db"
	"github.com/SND1231/post-service/model"
	pb "github.com/SND1231/post-service/proto"
	"testing"
)

const (
	Title    = "テスト"
	Content  = "test@test.com"
	PhotoUrl = "https://test"
	UserId   = int32(1)
)

func TestCheckGetPostsRequestSuccess(t *testing.T) {
	request := pb.GetPostsRequest{Limit: 1, Offset: 0, Id: 0, Title: ""}
	err := CheckGetPostsRequest(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}


func TestCheckGetPostsRequestErrorLimit(t *testing.T) {
	request := pb.GetPostsRequest{Limit: 0, Offset: 0, Id: 0, Title: ""}
	err := CheckGetPostsRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckCreatePostRequestSuccess(t *testing.T) {
	request := pb.CreatePostRequest{Title: Title, Content: Content,
								    PhotoUrl: PhotoUrl, UserId: UserId}
	err := CheckCreatePostRequest(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

func TestCheckCreatePostRequestErrorTitle(t *testing.T) {
	request := pb.CreatePostRequest{Title: "", Content: Content,
								    PhotoUrl: PhotoUrl, UserId: UserId}
	err := CheckCreatePostRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckCreatePostRequestErrorContent(t *testing.T) {
	request := pb.CreatePostRequest{Title: Title, Content: "",
								    PhotoUrl: PhotoUrl, UserId: UserId}
	err := CheckCreatePostRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckGetPostsRequestErrorUserId(t *testing.T) {
	request := pb.CreatePostRequest{Title: Title, Content: Content,
								    PhotoUrl: PhotoUrl, UserId: 0}
	err := CheckCreatePostRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func CreateUserForTest() {
	postParam := model.Post{Title: Title, Content: Content,
		PhotoUrl: PhotoUrl, UserId: UserId}
	db := db.Connection()
	defer db.Close()
	db.Create(&postParam)
}

func InitPostTable() {
	db := db.Connection()
	var post model.Post
	db.Delete(&post)
}
