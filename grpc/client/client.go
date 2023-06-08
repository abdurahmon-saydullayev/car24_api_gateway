package client

import (
	"Projects/Car24/car24_api_gateway/config"
	"Projects/Car24/car24_api_gateway/genproto/client_service"
	"Projects/Car24/car24_api_gateway/genproto/order_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	//userservice
	UserService() client_service.ClientServiceClient
	//orderservice
	OrderService() order_service.OrderServiceClient
}

type grpcClients struct {
	userService  client_service.ClientServiceClient
	orderService order_service.OrderServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	//connecting to user service
	connUserService, err := grpc.Dial(
		cfg.UserServiceHost+cfg.UserServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	// connecting to order service
	connOrderService, err := grpc.Dial(
		cfg.OrderServiceHost+cfg.OrderServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		userService:  client_service.NewClientServiceClient(connUserService),
		orderService: order_service.NewOrderServiceClient(connOrderService),
	}, nil
}

func (g *grpcClients) UserService() client_service.ClientServiceClient {
	return g.userService
}

func (g *grpcClients) OrderService() order_service.OrderServiceClient {
	return g.orderService
}
