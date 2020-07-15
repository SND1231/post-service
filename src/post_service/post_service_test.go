package post_service

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
	PostId    = int32(1)
)

// 投稿取得のリクエスト確認　成功
func TestCheckGetPostsRequestSuccess(t *testing.T) {
	request := pb.GetPostsRequest{Limit: 1, Offset: 0, Id: 0, Title: ""}
	err := CheckGetPostsRequest(request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

// 投稿取得のリクエスト確認　失敗
func TestCheckGetPostsRequestErrorLimit(t *testing.T) {
	request := pb.GetPostsRequest{Limit: 0, Offset: 0, Id: 0, Title: ""}
	err := CheckGetPostsRequest(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// 投稿作成のリクエスト確認　成功
func TestCheckCreatePostRequestSuccess(t *testing.T) {
	request := pb.CreatePostRequest{Title: Title, Content: Content,
									PhotoUrl: PhotoUrl, UserId: UserId, 
									StoreInfo: StoreInfo}
	err := CheckCreatePostRequest(request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

// 投稿作成のリクエスト確認　失敗　タイトルが空
func TestCheckCreatePostRequestErrorTitle(t *testing.T) {
	request := pb.CreatePostRequest{Title: "", Content: Content,
									PhotoUrl: PhotoUrl, UserId: UserId,
									StoreInfo: StoreInfo}
	err := CheckCreatePostRequest(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// 投稿作成のリクエスト確認　失敗　内容が空
func TestCheckCreatePostRequestErrorContent(t *testing.T) {
	request := pb.CreatePostRequest{Title: Title, Content: "",
									PhotoUrl: PhotoUrl, UserId: UserId,
									StoreInfo:StoreInfo}
	err := CheckCreatePostRequest(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// 投稿作成のリクエスト確認　失敗　ユーザIDが空
func TestCheckCreatePostsRequestErrorUserId(t *testing.T) {
	request := pb.CreatePostRequest{Title: Title, Content: Content,
								    PhotoUrl: PhotoUrl, UserId: 0}
	err := CheckCreatePostRequest(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}


//-----------------update
// 投稿更新のリクエスト確認　成功
func TestCheckUpdatePostRequestSuccess(t *testing.T) {
	request := pb.UpdatePostRequest{Title: Title, Content: Content,
									Id: PostId, StoreInfo: StoreInfo}
	err := CheckUpdatePostRequest(request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

// 投稿作成のリクエスト確認　失敗　投稿IDが空
func TestCheckUpdatePostRequestErrorId(t *testing.T) {
	request := pb.UpdatePostRequest{Title: Title, Content: Content,
									Id: 0, StoreInfo: StoreInfo}
	err := CheckUpdatePostRequest(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// 投稿作成のリクエスト確認　失敗　タイトルが空
func TestCheckUpdatePostRequestErrorTitle(t *testing.T) {
	request := pb.UpdatePostRequest{Title: "", Content: Content,
									Id: PostId, StoreInfo: StoreInfo}
	err := CheckUpdatePostRequest(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// 投稿作成のリクエスト確認　失敗　内容が空
func TestCheckUpdatePostRequestErrorContent(t *testing.T) {
	request := pb.UpdatePostRequest{Title: Title, Content: "",
									Id: PostId, StoreInfo: StoreInfo}
	err := CheckUpdatePostRequest(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// 投稿削除のリクエスト確認　成功
func TestCheckDeletePostRequestSuccess(t *testing.T){
	InitPostTable()
	postId := CreatePostForTest1()
	request := pb.DeletePostRequest{Id: postId, UserId: UserId}
	err := CheckDeletePostRequest(request)

	if err != nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// 投稿削除のリクエスト確認　エラー　すでに削除済み
func TestCheckDeletePostRequestError(t *testing.T){
	InitPostTable()
	postId := CreatePostForTest1()
	request := pb.DeletePostRequest{Id: postId, UserId: 999}
	err := CheckDeletePostRequest(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// いいね作成前の確認　成功
func TestCheckLikeExistsSuccess(t *testing.T){
	InitPostTable()
	postId := CreatePostForTest1()
	request := pb.CreateLikeRequest{PostId: postId, UserId: 5}
	err := CheckLikeExists(request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

// いいね作成前の確認　エラー
func TestCheckLikeExistsLikedError(t *testing.T){
	InitPostTable()
	postId := CreatePostForTest1()
	request := pb.CreateLikeRequest{PostId: postId, UserId: UserId}
	err := CheckLikeExists(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// いいね済みの確認　いいね済み
func TestLikeExistsLiked(t *testing.T){
	InitPostTable()
	postId := CreatePostForTest1()
	isLiked, _ := LikeExists(postId, UserId)

	assert.Equal(t, true, isLiked, "The two words should be the same.")
}

// いいね済みの確認　いいね未済み
func TestLikeExistsNotLiked(t *testing.T){
	InitPostTable()
	postId := CreatePostForTest1()
	isLiked, _ := LikeExists(postId, 999)

	assert.Equal(t, false, isLiked, "The two words should be the same.")
}

// いいね数のカウント
func TestCountLikes(t *testing.T){
	InitPostTable()
	postId := CreatePostForTest1()
	CreateLikeForTest(postId, 2)
	CreateLikeForTest(postId, 3)

	count := CountLikes(postId)

	assert.Equal(t, int32(3), count, "The two words should be the same.")
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

func CreatePostForTest1() int32 {
	post := model.Post{Title: Title, Content: Content,
		PhotoUrl: PhotoUrl, UserId: UserId, StoreInfo: StoreInfo}
	
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

