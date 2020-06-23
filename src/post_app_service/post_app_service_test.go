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
	CreatePostForTest()

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

	InitPostTable()
}

func TestGetPostSuccess(t *testing.T) {
	CreatePostForTest()
	post_id := GetPostID()
	CreateLikeForTest(post_id)

	post, err := GetPost(post_id)

	if err != nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
	assert.Equal(t, Title, post.Title, "The two words should be the same.")
	assert.Equal(t, Content, post.Content, "The two words should be the same.")
	assert.Equal(t, PhotoUrl, post.PhotoUrl, "The two words should be the same.")
	assert.Equal(t, UserId, post.UserId, "The two words should be the same.")
	assert.Equal(t, int32(1), post.Likes, "The two words should be the same.")

	InitPostTable()
}

func CreatePostForTest() {
	post := model.Post{Title: Title, Content: Content,
		PhotoUrl: PhotoUrl, UserId: UserId}

	db := db.Connection()
	defer db.Close()

	db.Create(&post)
}

func InitPostTable() {
	db := db.Connection()
	defer db.Close()

	db.Exec("DELETE FROM posts")
	db.Exec("DELETE FROM likes")
	db.Exec("DELETE FROM post_likes")
}

func GetPostID() int32 {
	var id int32

	db := db.Connection()
	defer db.Close()

	db.Table("posts").Count(&id)

	return id
}

func CreateLikeForTest(id int32) {
	var post model.Post

	db := db.Connection()
	defer db.Close()

	like := model.Like{UserId: int32(1)}
	db.Create(&like)
	db.Model(&post).Association("Likes").Append([]model.Like{like})
}
