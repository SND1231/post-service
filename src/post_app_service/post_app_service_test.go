package post_app_service

import (
	"github.com/SND1231/post-service/db"
	"github.com/SND1231/post-service/model"
	pb "github.com/SND1231/post-service/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Title    = "テスト"
	Content  = "test@test.com"
	PhotoUrl = "https://test"
	UserId   = int32(1)
)

func TestGetPostsSuccess(t *testing.T) {
	InitPostTable()
	CreateUserForTest()
	request := pb.GetPostsRequest{Limit: 1, Offset: 0, Id: 0, Title: ""}
	posts, count, err := GetPosts(request)

	if err != nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
	assert.Equal(t, Title, posts[0].Title, "The two words should be the same.")
	assert.Equal(t, Content, posts[0].Content, "The two words should be the same.")
	assert.Equal(t, PhotoUrl, posts[0].PhotoUrl, "The two words should be the same.")
	assert.Equal(t, UserId, posts[0].UserId, "The two words should be the same.")
	assert.Equal(t, int32(1), count, "The two words should be the same.")
}

func CreateUserForTest() {
	post_param := model.Post{Title: Title, Content: Content,
		PhotoUrl: PhotoUrl, UserId: UserId}
	db := db.Connection()
	defer db.Close()
	db.Create(&post_param)
}

func InitPostTable() {
	db := db.Connection()
	var post model.Post
	db.Delete(&post)
}
