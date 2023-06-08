package client

import (
	"Projects/Car24/car24_api_gateway/config"
	"Projects/Car24/car24_api_gateway/genproto/client_service"
	"Projects/Car24/car24_api_gateway/genproto/order_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	UserService() client_service.ClientServiceClient
	//order
	OrderService() order_service.OrderServiceClient
	CarService() order_service.CarServiceClient
	DiscountService() order_service.DiscountServiceClient
	MechanicService() order_service.MechanicServiceClient
	ModelService() order_service.ModelServiceClient
	TarifService() order_service.TarifServiceClient
}

type grpcClients struct {
	userService client_service.ClientServiceClient
	//order
	orderService    order_service.OrderServiceClient
	carService      order_service.CarServiceClient
	discountService order_service.DiscountServiceClient
	mechanicService order_service.MechanicServiceClient
	modelService    order_service.ModelServiceClient
	tarifService    order_service.TarifServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connUserService, err := grpc.Dial(
		cfg.UserServiceHost+cfg.UserServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connOrderService, err := grpc.Dial(
		cfg.OrderServiceHost+cfg.OrderServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connCarService, err := grpc.Dial(
		cfg.OrderServiceHost+cfg.OrderServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connDisService, err := grpc.Dial(
		cfg.OrderServiceHost+cfg.OrderServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connMechService, err := grpc.Dial(
		cfg.OrderServiceHost+cfg.OrderServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connModelService, err := grpc.Dial(
		cfg.OrderServiceHost+cfg.OrderServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connTarifService, err := grpc.Dial(
		cfg.OrderServiceHost+cfg.OrderServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		userService:     client_service.NewClientServiceClient(connUserService),
		orderService:    order_service.NewOrderServiceClient(connOrderService),
		carService:      order_service.NewCarServiceClient(connCarService),
		discountService: order_service.NewDiscountServiceClient(connDisService),
		mechanicService: order_service.NewMechanicServiceClient(connMechService),
		modelService:    order_service.NewModelServiceClient(connModelService),
		tarifService:    order_service.NewTarifServiceClient(connTarifService),
	}, nil
}

func (g *grpcClients) UserService() client_service.ClientServiceClient {
	return g.userService
}

func (g *grpcClients) OrderService() order_service.OrderServiceClient {
	return g.orderService
}

func (g *grpcClients) CarService() order_service.CarServiceClient {
	return g.carService
}

func (g *grpcClients) DiscountService() order_service.DiscountServiceClient {
	return g.discountService
}

func (g *grpcClients) MechanicService() order_service.MechanicServiceClient {
	return g.mechanicService
}

func (g *grpcClients) ModelService() order_service.ModelServiceClient {
	return g.modelService
}

func (g *grpcClients) TarifService() order_service.TarifServiceClient {
	return g.tarifService
}
