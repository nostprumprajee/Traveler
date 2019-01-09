package main

import (
	"log"
	"net"

	pb "github.com/mongmx/grpc-tutorial-topic-service/proto/topic"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type Repository interface {
	Create(*pb.Topic) (*pb.Topic, error)
	GetAll() []*pb.Topic
}

type TopicRepository struct {
	topics []*pb.Topic
}

func (repo *TopicRepository) Create(topic *pb.Topic) (*pb.Topic, error) {
	updated := append(repo.topics, topic)
	repo.topics = updated
	return topic, nil
}

func (repo *TopicRepository) GetAll() []*pb.Topic {
	return repo.topics
}

type service struct {
	repo Repository
}

func (s *service) CreateTopic(ctx context.Context, req *pb.Topic) (*pb.Response, error) {
	topic, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	log.Println("Client call: CreateTopic")
	return &pb.Response{Created: true, Topic: topic}, nil
}

func (s *service) GetTopics(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	topics := s.repo.GetAll()
	log.Println("Client call: GetTopics")
	return &pb.Response{Topics: topics}, nil
}

func main() {
	repo := &TopicRepository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterTopicServiceServer(s, &service{repo})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
