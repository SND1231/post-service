package post_service

import (
	"github.com/SND1231/post-service/db"
	"github.com/SND1231/post-service/model"
	pb "github.com/SND1231/post-service/proto"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//error_list [], err_msg string
func CreateError(code codes.Code, error_list []*errdetails.BadRequest_FieldViolation) error {
	st := status.New(codes.InvalidArgument, "エラー発生")
	// add error message detail
	st, err := st.WithDetails(
		&errdetails.BadRequest{
			FieldViolations: error_list,
		},
	)
	// unexpected error
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %+v", err))
	}

	// return error
	return st.Err()
}

func CreateBadRequest_FieldViolation(feild string, desc string) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       feild,
		Description: desc,
	}
}

func CheckGetPostsRequest(request pb.GetPostsRequest) error {
	var error_list []*errdetails.BadRequest_FieldViolation
	if request.Limit == 0 {
		error_list = append(error_list, CreateBadRequest_FieldViolation("Limit", "値が設定されていません"))
	}
	if request.Limit == 0 {
		error_list = append(error_list, CreateBadRequest_FieldViolation("id", "値が設定されていません"))
	}

	if len(error_list) > 0 {
		return CreateError(codes.InvalidArgument, error_list)
	} else {
		return nil
	}
}

func CheckCreatePostRequest(request pb.CreatePostRequest) error {
	var error_list []*errdetails.BadRequest_FieldViolation
	if request.Title == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("タイトル", "必須です"))
	}
	if request.Content == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("内容", "必須です"))
	}
	if request.UserId == 0 {
		error_list = append(error_list, CreateBadRequest_FieldViolation("ユーザID", "必須です"))
	}

	if len(error_list) > 0 {
		return CreateError(codes.InvalidArgument, error_list)
	} else {
		return nil
	}
}

func CheckUpdatePostRequest(request pb.UpdatePostRequest) error {
	var error_list []*errdetails.BadRequest_FieldViolation
	if request.Title == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("タイトル", "必須です"))
	}
	if request.Content == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("内容", "必須です"))
	}

	if len(error_list) > 0 {
		return CreateError(codes.InvalidArgument, error_list)
	} else {
		return nil
	}
}

func CheckDeletePostRequest(request pb.DeletePostRequest) error {
	var post model.Post
	db := db.Connection()
	defer db.Close()

	id := request.Id
	user_id := request.UserId

	db.Where("id = ? AND user_id = ?", id, user_id).First(&post)
	fmt.Println(post.ID)
	if post.ID == 0 {
		return status.New(codes.AlreadyExists, "ユーザが違うか、投稿がすでに存在しません").Err()
	}
	return nil
}

func LikeExists(post_id int32, user_id int32) (bool, int32) {
	db := db.Connection()
	defer db.Close()

	var post model.Post
	db.First(&post, post_id)

	var likes []model.Like
	db.Model(&post).Where("user_id = ?", user_id).Related(&likes, "Likes")
	if len(likes) == 0 {
		return false, 0
	}
	return true, likes[0].ID
}

func CheckLikeExists(request pb.CreateLikeRequest) error {
	result, _ := LikeExists(request.PostId, request.UserId)
	if result {
		return status.New(codes.AlreadyExists, "すでに、いいね済みです").Err()
	}
	return nil
}

func CountLikes(post_id int32) int32 {
	var count int32
	db := db.Connection()
	defer db.Close()

	db.Table("post_likes").Where("post_id = ?", post_id).Count(&count)
	return count
}
