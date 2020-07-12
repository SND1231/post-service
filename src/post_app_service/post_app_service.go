package post_app_service

import (
	"github.com/SND1231/post-service/db"
	"github.com/SND1231/post-service/model"
	"github.com/SND1231/post-service/post_service"
	pb "github.com/SND1231/post-service/proto"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetPosts(request pb.GetPostsRequest) ([]*pb.Post, int32, error) {
	var posts []model.Post
	var postList []*pb.Post
	var count int32

	err := post_service.CheckGetPostsRequest(request)
	if err != nil {
		return postList, 0, err
	}

	limit := request.Limit
	offset := limit * (request.Offset - 1)

	db := db.Connection()
	defer db.Close()
	postDb := db
	countDb := db.Table("posts")

	if request.Id != 0 {
		postDb = postDb.Where("user_id = ?", request.Id)
		countDb = countDb.Where("user_id = ?", request.Id)
	}
	if request.Title != "" {
		postDb = postDb.Where("title LIKE ?", "%"+request.Title+"%")
		countDb = countDb.Where("title LIKE ?", "%"+request.Title+"%")
	}
	postDb.Limit(limit).Offset(offset).Order("id desc").Find(&posts).Scan(&postList)
	countDb.Table("posts").Find(&posts).Count(&count)

	for i := 0; i < len(postList); i++ {
		postList[i].Likes = post_service.CountLikes(postList[i].Id)
	}

	return postList, count, nil
}

func GetPost(id int32) (pb.Post, error) {
	var post model.Post
	var postParam pb.Post

	db := db.Connection()
	defer db.Close()
	db.Find(&post, id).Scan(&postParam)

	postParam.Likes = post_service.CountLikes(postParam.Id)

	return postParam, nil
}

func CreatePost(request pb.CreatePostRequest) (int32, error) {
	err := post_service.CheckCreatePostRequest(request)
	if err != nil {
		return -1, err
	}

	postParam := model.Post{Title: request.Title, Content: request.Content,
		PhotoUrl: request.PhotoUrl, UserId: request.UserId, StoreInfo: request.StoreInfo}

	db := db.Connection()
	defer db.Close()
	db.Create(&postParam)
	if db.NewRecord(postParam) == false {
		return postParam.ID, err
	}
	return -1, status.New(codes.Unknown, "作成失敗").Err()
}

func UpdatePost(request pb.UpdatePostRequest) (int32, error) {
	err := post_service.CheckUpdatePostRequest(request)
	if err != nil {
		return -1, err
	}

	id := request.Id

	postParam := model.Post{Title: request.Title, Content: request.Content,
		PhotoUrl: request.PhotoUrl, StoreInfo: request.StoreInfo}

	db := db.Connection()
	defer db.Close()
	post := model.Post{}
	db.Find(&post, id)

	db.Model(&post).UpdateColumns(postParam)
	return id, nil

}

func DeletePost(request pb.DeletePostRequest) (int32, error) {
	err := post_service.CheckDeletePostRequest(request)
	if err != nil {
		return -1, err
	}

	id := request.Id
	userId := request.UserId

	db := db.Connection()
	defer db.Close()
	db.Where("id = ? AND user_id = ?", id, userId).Delete(model.Post{})
	return id, nil
}

func CreateLike(request pb.CreateLikeRequest) (int32, int32, error) {
	err := post_service.CheckLikeExists(request)
	if err != nil {
		return -1, 0, err
	}
	var post model.Post
	db := db.Connection()
	defer db.Close()
	db.First(&post, request.PostId)

	like := model.Like{UserId: request.UserId}
	db.Create(&like)
	db.Model(&post).Association("Likes").Append([]model.Like{like})
	if db.NewRecord(like) == false {
		return like.ID, post_service.CountLikes(request.PostId), nil
	}
	return -1, 0, status.New(codes.Unknown, "作成失敗").Err()
}

func CheckLiked(request pb.CheckLikedRequest) (bool, int32) {
	return post_service.LikeExists(request.PostId, request.UserId)
}

func DeleteLike(request pb.DeleteLikeRequest) (int32, int32, error) {
	var postId int32

	db := db.Connection()
	defer db.Close()
	row := db.Table("post_likes").Where("like_id = ?", request.Id).Select("post_id").Row()
	row.Scan(&postId)

	db.Where("id = ?", request.Id).Delete(model.Like{})
	db.Exec("DELETE FROM post_likes WHERE like_id = ?", request.Id)
	return request.Id, post_service.CountLikes(postId), nil
	row.Scan(&postId)

	db.Where("id = ?", request.Id).Delete(model.Like{})
	db.Exec("DELETE FROM post_likes WHERE like_id = ?", request.Id)
	return request.Id, post_service.CountLikes(postId), nil
}
