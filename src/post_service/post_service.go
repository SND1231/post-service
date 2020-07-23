package post_service

import (
	"fmt"
	"github.com/SND1231/post-service/db"
	"github.com/SND1231/post-service/model"
	pb "github.com/SND1231/post-service/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//errorList [], err_msg string
func CreateError(code codes.Code, errorList []*errdetails.BadRequest_FieldViolation) error {
	st := status.New(codes.InvalidArgument, "エラー発生")
	// add error message detail
	st, err := st.WithDetails(
		&errdetails.BadRequest{
			FieldViolations: errorList,
		},
	)
	// unexpected error
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %+v", err))
	}

	// return error
	return st.Err()
}

func CreateBadRequestFieldViolation(feild string, desc string) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       feild,
		Description: desc,
	}
}

func CheckGetPostsRequest(request pb.GetPostsRequest) error {
	var errorList []*errdetails.BadRequest_FieldViolation
	if request.Limit == 0 {
		errorList = append(errorList, CreateBadRequestFieldViolation("Limit", "値が設定されていません"))
	}

	if len(errorList) > 0 {
		return CreateError(codes.InvalidArgument, errorList)
	} else {
		return nil
	}
}

func CheckCreatePostRequest(request pb.CreatePostRequest) error {
	var errorList []*errdetails.BadRequest_FieldViolation
	if request.Title == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("タイトル", "必須です"))
	}
	if request.Content == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("内容", "必須です"))
	}
	if request.UserId == 0 {
		errorList = append(errorList, CreateBadRequestFieldViolation("ユーザID", "必須です"))
	}

	if len(errorList) > 0 {
		return CreateError(codes.InvalidArgument, errorList)
	} else {
		return nil
	}
}

func CheckUpdatePostRequest(request pb.UpdatePostRequest) error {
	var errorList []*errdetails.BadRequest_FieldViolation
	if request.Id == 0 {
		errorList = append(errorList, CreateBadRequestFieldViolation("ID", "必須です"))
	}
	if request.Title == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("タイトル", "必須です"))
	}
	if request.Content == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("内容", "必須です"))
	}

	if len(errorList) > 0 {
		return CreateError(codes.InvalidArgument, errorList)
	} else {
		return nil
	}
	return nil
}

func CheckDeletePostRequest(request pb.DeletePostRequest) error {
	var post model.Post
	db := db.Connection()
	defer db.Close()

	id := request.Id
	userId := request.UserId

	db.Where("id = ? AND user_id = ?", id, userId).First(&post)
	fmt.Println(post.ID)
	if post.ID == 0 {
		return status.New(codes.NotFound, "ユーザが違うか、投稿がすでに存在しません").Err()
	}
	return nil
}

func CheckLikeExists(request pb.CreateLikeRequest) error {
	result, _ := LikeExists(request.PostId, request.UserId)
	if result {
		return status.New(codes.AlreadyExists, "すでに、いいね済みです").Err()
	}
	return nil
}


func LikeExists(postId int32, userId int32) (bool, int32) {
	db := db.Connection()
	defer db.Close()

	var post model.Post
	db.First(&post, postId)

	var likes []model.Like
	db.Model(&post).Where("user_id = ?", userId).Related(&likes, "Likes")
	if len(likes) == 0 {
		return false, 0
	}
	return true, likes[0].ID
}

func CountLikes(postId int32) int32 {
	var count int32
	db := db.Connection()
	defer db.Close()

	db.Table("post_likes").Where("post_id = ?", postId).Count(&count)
	return count
}
