package main

import (
	"context"
	"log"
	"net"

	post_app_service "github.com/SND1231/post-service/post_app_service"
	pb "github.com/SND1231/post-service/proto"
	"google.golang.org/grpc"
)

const (
	port = ":9002"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedPostServiceServer
}

// GET Post
func (s *server) GetPost(ctx context.Context, in *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	id := in.Id
	post, err := post_app_service.GetPost(id)
	return &pb.GetPostResponse{Post: &post}, err
}

// GET Posts
func (s *server) GetPosts(ctx context.Context, in *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	posts, count, err := post_app_service.GetPosts(*in)
	return &pb.GetPostsResponse{Posts: posts, Count: count}, err
}

// Create Post
func (s *server) CreatePost(ctx context.Context, in *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	request := *in
	id, err := post_app_service.CreatePost(request)
	if err == nil {
		return &pb.CreatePostResponse{Id: id}, nil
	} else {
		return &pb.CreatePostResponse{}, err
	}
}

// Update Post
func (s *server) UpdatePost(ctx context.Context, in *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	request := *in
	id, err := post_app_service.UpdatePost(request)
	if err == nil {
		return &pb.UpdatePostResponse{Id: id}, nil
	} else {
		return &pb.UpdatePostResponse{}, err
	}
}

// Delete Post
func (s *server) DeletePost(ctx context.Context, in *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	request := *in
	id, err := post_app_service.DeletePost(request)
	if err == nil {
		return &pb.DeletePostResponse{Id: id}, nil
	} else {
		return &pb.DeletePostResponse{}, err
	}
}

// Create Like
func (s *server) CreateLike(ctx context.Context, in *pb.CreateLikeRequest) (*pb.CreateLikeResponse, error) {
	request := *in
	id, count, err := post_app_service.CreateLike(request)
	if err == nil {
		return &pb.CreateLikeResponse{Id: id, Count: count}, nil
	} else {
		return &pb.CreateLikeResponse{}, err
	}
}

// Delete Like
func (s *server) DeleteLike(ctx context.Context, in *pb.DeleteLikeRequest) (*pb.DeleteLikeResponse, error) {
	request := *in
	id, count, err := post_app_service.DeleteLike(request)
	if err == nil {
		return &pb.DeleteLikeResponse{Id: id, Count: count}, nil
	} else {
		return &pb.DeleteLikeResponse{}, err
	}
}

// Check Liked
func (s *server) CheckLiked(ctx context.Context, in *pb.CheckLikedRequest) (*pb.CheckLikedResponse, error) {
	request := *in
	is_liked, id := post_app_service.CheckLiked(request)
	log.Println(id)
	return &pb.CheckLikedResponse{Liked: is_liked, Id: id}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
