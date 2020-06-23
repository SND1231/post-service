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
