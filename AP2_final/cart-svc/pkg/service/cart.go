package service

import (
	"context"
	"net/http"

	"github.com/anv01208/online-flower-shop/cart-svc/pkg/client"
	"github.com/anv01208/online-flower-shop/cart-svc/pkg/db"
	"github.com/anv01208/online-flower-shop/cart-svc/pkg/models"
	"github.com/anv01208/online-flower-shop/cart-svc/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedCartServiceServer
	H          db.Handler
	ProductSvc client.ProductServiceClient
}

func (s *Server) GetCartItems(req *pb.ViewCartRequest, stream pb.CartService_GetCartItemsServer) error {
	var cartItems []models.Cart

	if result := s.H.DB.Where("user_id = ?", req.UserId).Find(&cartItems); result.Error != nil {
		return status.Errorf(codes.Internal, "Failed to retrieve cart items: %v", result.Error)
	}

	response := &pb.ViewCartResponse{}

	// Iterate over the cart items and send each as a response
	for _, cartItem := range cartItems {
		cartItemData := &pb.CartItem{
			CartItemId: cartItem.Id,
			ProductId:  cartItem.ProductId,
			Quantity:   int64(cartItem.Quantity),
		}
		response.CartItem = append(response.CartItem, cartItemData)
	}

	if err := stream.Send(response); err != nil {
		return status.Errorf(codes.Internal, "Failed to send response: %v", err)
	}

	return nil
}

func (s *Server) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	var cartItem models.Cart

	cartItem.ProductId = req.ProductId
	cartItem.Quantity = int32(req.Quantity)
	cartItem.UserId = req.UserId

	if result := s.H.DB.Create(&cartItem); result.Error != nil {
		return &pb.AddToCartResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.AddToCartResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) RemoveFromCart(ctx context.Context, req *pb.RemoveFromCartRequest) (*pb.RemoveFromCartResponse, error) {
	if result := s.H.DB.Delete(&models.Cart{}, req.CartItemId).Where("user_id = ?", req.UserId); result.Error != nil {
		return &pb.RemoveFromCartResponse{
			Status: http.StatusInternalServerError,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.RemoveFromCartResponse{
		Status: http.StatusOK,
	}, nil
}
