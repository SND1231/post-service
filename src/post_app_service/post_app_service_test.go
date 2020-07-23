package post_app_service

import (
	"github.com/SND1231/post-service/db"
	"github.com/SND1231/post-service/model"
	pb "github.com/SND1231/post-service/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Title     = "テスト"
	Content   = "test@test.com"
	PhotoUrl  = "https://test"
	StoreInfo = "https://store1"
	UserId    = int32(1)

	Title2     = "タイトル検索"
	Content2   = "title@test.com"
	PhotoUrl2  = "https://title"
	StoreInfo2 = "https://store1"
	UserId2    = int32(2)
)

// 投稿一覧の取得 複数取得できるか
func TestGetPostsSuccess(t *testing.T) {
	InitPostTable()
	_ = CreatePostForTest1()
	_ = CreatePostForTest2()

	request := pb.GetPostsRequest{Limit: 2, Offset: 0, Id: 0, Title: ""}
	posts, count, err := GetPosts(request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
	assert.Equal(t, int32(2), count, "The two words should be the same.")

	assert.Equal(t, Title2, posts[0].Title, "The two words should be the same.")
	assert.Equal(t, Content2, posts[0].Content, "The two words should be the same.")
	assert.Equal(t, PhotoUrl2, posts[0].PhotoUrl, "The two words should be the same.")
	assert.Equal(t, StoreInfo2, posts[0].StoreInfo, "The two words should be the same.")
	assert.Equal(t, UserId2, posts[0].UserId, "The two words should be the same.")
	assert.Equal(t, int32(1), posts[0].Likes, "The two words should be the same.")

	assert.Equal(t, Title, posts[1].Title, "The two words should be the same.")
	assert.Equal(t, Content, posts[1].Content, "The two words should be the same.")
	assert.Equal(t, PhotoUrl, posts[1].PhotoUrl, "The two words should be the same.")
	assert.Equal(t, StoreInfo, posts[0].StoreInfo, "The two words should be the same.")
	assert.Equal(t, UserId, posts[1].UserId, "The two words should be the same.")
	assert.Equal(t, int32(1), posts[1].Likes, "The two words should be the same.")
}

// 投稿一覧の取得 オフセットできてるか
func TestGetPostsSuccessOffset(t *testing.T) {
	InitPostTable()
	_ = CreatePostForTest1()
	_ = CreatePostForTest2()

	request := pb.GetPostsRequest{Limit: 1, Offset: 2, Id: 0, Title: ""}
	posts, count, err := GetPosts(request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
	assert.Equal(t, int32(2), count, "The two words should be the same.")

	assert.Equal(t, Title, posts[0].Title, "The two words should be the same.")
	assert.Equal(t, Content, posts[0].Content, "The two words should be the same.")
	assert.Equal(t, PhotoUrl, posts[0].PhotoUrl, "The two words should be the same.")
}

// 投稿一覧の取得 ユーザIDで絞り込みできてるか
func TestGetPostsSuccessById(t *testing.T) {
	InitPostTable()
	_ = CreatePostForTest1()
	_ = CreatePostForTest2()

	request := pb.GetPostsRequest{Limit: 1, Offset: 0, Id: 1, Title: ""}
	posts, count, err := GetPosts(request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}

	assert.Equal(t, int32(1), count, "The two words should be the same.")
	assert.Equal(t, Title, posts[0].Title, "The two words should be the same.")
	assert.Equal(t, Content, posts[0].Content, "The two words should be the same.")
	assert.Equal(t, PhotoUrl, posts[0].PhotoUrl, "The two words should be the same.")
}

//  投稿一覧の取得 タイトルで絞り込みできてるか
func TestGetPostsSuccessByTitle(t *testing.T) {
	InitPostTable()
	_ = CreatePostForTest1()
	_ = CreatePostForTest2()

	request := pb.GetPostsRequest{Limit: 1, Offset: 0, Id: 0, Title: "タイトル"}
	posts, count, err := GetPosts(request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}

	assert.Equal(t, int32(1), count, "The two words should be the same.")
	assert.Equal(t, Title2, posts[0].Title, "The two words should be the same.")
	assert.Equal(t, Content2, posts[0].Content, "The two words should be the same.")
	assert.Equal(t, PhotoUrl2, posts[0].PhotoUrl, "The two words should be the same.")
}

// 投稿一覧の取得エラー
func TestGetPostsError(t *testing.T) {
	InitPostTable()
	
	request := pb.GetPostsRequest{Limit: 0, Offset: 0, Id: 0, Title: ""}
	_, _, err := GetPosts(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// 投稿を取得できるか
func TestGetPostSuccess(t *testing.T) {
	InitPostTable()
	postId := CreatePostForTest1()

	post, err := GetPost(postId)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
	assert.Equal(t, Title, post.Title, "The two words should be the same.")
	assert.Equal(t, Content, post.Content, "The two words should be the same.")
	assert.Equal(t, PhotoUrl, post.PhotoUrl, "The two words should be the same.")
	assert.Equal(t, UserId, post.UserId, "The two words should be the same.")
	assert.Equal(t, int32(1), post.Likes, "The two words should be the same.")
}

// 投稿作成できるか
func TestCreatePostSuccess(t *testing.T) {
	InitPostTable()
	request := pb.CreatePostRequest{Title: Title, Content: Content,
		PhotoUrl: PhotoUrl, UserId: UserId, StoreInfo: StoreInfo}

	postId, err := CreatePost(request)
	post := GetPostForTest(postId)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
	assert.Equal(t, Title, post.Title, "The two words should be the same.")
	assert.Equal(t, Content, post.Content, "The two words should be the same.")
	assert.Equal(t, PhotoUrl, post.PhotoUrl, "The two words should be the same.")
	assert.Equal(t, StoreInfo, post.StoreInfo, "The two words should be the same.")
	assert.Equal(t, UserId, post.UserId, "The two words should be the same.")

}

//　投稿作成エラー
func TestCreatePostError(t *testing.T) {
	InitPostTable()
	request := pb.CreatePostRequest{Title: "", Content: Content,
		PhotoUrl: PhotoUrl, UserId: UserId, StoreInfo: StoreInfo}

	_, err := CreatePost(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// 投稿更新
func TestUpdatePostSuccess(t *testing.T) {
	InitPostTable()
	postId := CreatePostForTest1()
	request := pb.UpdatePostRequest{Title: Title2, Content: Content2,
		PhotoUrl: PhotoUrl2, StoreInfo: StoreInfo2, Id: postId}

	postId, err := UpdatePost(request)
	post := GetPostForTest(postId)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
	assert.Equal(t, Title2, post.Title, "The two words should be the same.")
	assert.Equal(t, Content2, post.Content, "The two words should be the same.")
	assert.Equal(t, PhotoUrl2, post.PhotoUrl, "The two words should be the same.")
	assert.Equal(t, StoreInfo2, post.StoreInfo, "The two words should be the same.")
	assert.Equal(t, UserId, post.UserId, "The two words should be the same.")
}

// 投稿更新エラー
func TestUpdatePostError(t *testing.T) {
	InitPostTable()
	_ = CreatePostForTest1()
	request := pb.UpdatePostRequest{Title: Title2, Content: Content2,
		PhotoUrl: PhotoUrl2, StoreInfo: StoreInfo2, Id: 0}

	_, err := UpdatePost(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// 投稿削除
func TestDeletePostSuccess(t *testing.T) {
	InitPostTable()
	postId := CreatePostForTest1()
	request := pb.DeletePostRequest{Id: postId, UserId: UserId}

	postId, err := DeletePost(request)
	count := CountPostNum()

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
	assert.Equal(t, 0, count, "The two words should be the same.")
}

// 投稿削除エラー
func TestDeletePostError(t *testing.T) {
	InitPostTable()
	postId := CreatePostForTest1()
	request := pb.DeletePostRequest{Id: postId, UserId: 0}

	_, err := DeletePost(request)
	
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// いいね作成
func TestCreateLikeSuccess(t *testing.T) {
	InitPostTable()
	postId := CreatePostForTest1()
	request := pb.CreateLikeRequest{PostId: postId, UserId: 5}

	_, likeCount, err := CreateLike(request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
	assert.Equal(t, int32(2), likeCount, "The two words should be the same.")
}

// いいね作成エラー
func TestCreateLikeError(t *testing.T) {
	InitPostTable()
	postId := CreatePostForTest1()
	request := pb.CreateLikeRequest{PostId: postId, UserId: 0}

	_, _ ,err := CreateLike(request)

	if err != nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// いいね済みのチェック いいね済み
func TestCheckLikedLiked(t *testing.T) {
	InitPostTable()
	postId := CreatePostForTest1()
	request := pb.CheckLikedRequest{PostId: postId, UserId: UserId}

	isLiked, _ := CheckLiked(request)

	assert.Equal(t, true, isLiked, "The two words should be the same.")
}

// いいね済みのチェック いいねしてない
func TestCheckLikedNotLiked(t *testing.T) {
	InitPostTable()
	postId := CreatePostForTest1()
	request := pb.CheckLikedRequest{PostId: postId, UserId: 999}

	isLiked, _ := CheckLiked(request)

	assert.Equal(t, false, isLiked, "The two words should be the same.")
}

// いいね取り消し
func TestDeleteLikeSuccess(t *testing.T) {
	InitPostTable()
	postId := CreatePostForTest1()
	likeId := CreateLikeForTest(postId, 5)
	request := pb.DeleteLikeRequest{Id: likeId}

	_, likeCount, _ := DeleteLike(request)

	isDeleteLike := LikeExistsForTest(likeId)
	assert.Equal(t, true, isDeleteLike, "The two words should be the same.")
	assert.Equal(t, int32(1), likeCount, "The two words should be the same.")
}

func CreatePostForTest1() int32 {
	post := model.Post{Title: Title, Content: Content,
		PhotoUrl: PhotoUrl, UserId: UserId, StoreInfo: StoreInfo}
	
	postId := CreatePostForTest(post)
	_ = CreateLikeForTest(postId, UserId)

	return postId
}

func CreatePostForTest2() int32 {
	post := model.Post{Title: Title2, Content: Content2,
		PhotoUrl: PhotoUrl2, UserId: UserId2, StoreInfo: StoreInfo2}
	
	postId := CreatePostForTest(post)
	_ = CreateLikeForTest(postId, UserId)

	return postId
}

func CreatePostForTest(post model.Post) int32 {
	db := db.Connection()
	defer db.Close()
	db.Create(&post)

	return post.ID
}

func InitPostTable() {
	db := db.Connection()
	var post model.Post
	db.Delete(&post)
	defer db.Close()

	db.Exec("DELETE FROM posts")
	db.Exec("DELETE FROM likes")
	db.Exec("DELETE FROM post_likes")
}

func CreateLikeForTest(postId int32, userId int32) int32 {
	var post model.Post

	db := db.Connection()
	defer db.Close()

	db.First(&post, postId)
	like := model.Like{UserId: userId}
	db.Create(&like)
	db.Model(&post).Association("Likes").Append([]model.Like{like})

	return like.ID
}

func GetPostForTest(id int32) model.Post {
	var post model.Post

	db := db.Connection()
	defer db.Close()
	db.Find(&post, id)

	return post
}

func CountPostNum() int{
	var count int
	var posts []model.Post
	db := db.Connection()
	defer db.Close()

	db.Table("posts").Find(&posts).Count(&count)

	return count

}

func LikeExistsForTest(likeId int32) bool {
	var like model.Like
	db := db.Connection()
	defer db.Close()

	db.Table("likes").Find(like, likeId)

	if likeId != 0 {
		return true
	} else {
		return false
	}
}