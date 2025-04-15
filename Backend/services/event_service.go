package services

import (
	pb "AI-Powered-Automated-Loan-Underwriting-System/created_proto/event"
	"AI-Powered-Automated-Loan-Underwriting-System/models"
	"AI-Powered-Automated-Loan-Underwriting-System/repositories"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventServiceServer struct {
	pb.UnimplementedEventServiceServer
	Repo *repositories.EventRepo
}

func NewEventServiceServer(repo *repositories.EventRepo) *EventServiceServer {
	return &EventServiceServer{
		Repo: repo,
	}
}

// CreateEvent implements the CreateEvent RPC method
func (s *EventServiceServer) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	event := models.Event{
		EventType: req.EventType,
		Payload:   req.Payload,
		Timestamp: req.Timestamp.AsTime(),
	}

	// Call repository to save event
	if err := s.Repo.CreateEvent(ctx, event); err != nil {
		return nil, err
	}

	return &pb.CreateEventResponse{
		EventId: uint64(event.ID),
		Status:  "Event created successfully",
	}, nil
}

// GetEvent implements the GetEvent RPC method
func (s *EventServiceServer) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	var event models.Event

	// Fetch event using the repository
	if err := s.Repo.GetEventByID(ctx, event.ID, &event); err != nil {
		return nil, err
	}

	// Return the event details as response
	return &pb.GetEventResponse{
		EventId:   uint64(event.ID),
		EventType: event.EventType,
		Payload:   event.Payload,
		Timestamp: timestamppb.New(event.Timestamp),
	}, nil
}

// GetAllEvents implements the GetAllEvents RPC method
func (s *EventServiceServer) GetAllEvents(ctx context.Context, req *pb.Empty) (*pb.EventList, error) {
	var events []models.Event

	// Fetch all events using the repository
	if err := s.Repo.GetAllEvents(ctx, &events); err != nil {
		return nil, err
	}

	var eventResponses []*pb.EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, &pb.EventResponse{
			EventId:   uint64(event.ID),
			EventType: event.EventType,
			Payload:   event.Payload,
			Timestamp: timestamppb.New(event.Timestamp),
		})
	}

	return &pb.EventList{
		Events: eventResponses,
	}, nil
}
