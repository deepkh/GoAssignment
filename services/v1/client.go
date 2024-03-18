package v1

import (
	"context"
	"log"

	"go-recommendation-system/protos"

	"google.golang.org/grpc"
)

var defaultAddress string = "localhost:5001"
var connn *grpc.ClientConn = nil

func dial(addr string) (*grpc.ClientConn, error) {
	if connn == nil {
		// Set up a connection to the server.
		c, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("grpc did not connect: %v", err)
			return nil, err
		}
		connn = c
	}

	return connn, nil
}

func Reg(in *protos.RegReq) (*protos.RegRep, error) {
	c, err := dial(defaultAddress)
	if err != nil {
		return nil, err
	}

	r, err2 := protos.NewRegServClient(c).Reg(context.Background(), in)
	if err != nil {
		return nil, err2
	}

	return r, nil
}

func Confirm(in *protos.ConfirmReq) (*protos.ConfirmRep, error) {
	c, err := dial(defaultAddress)
	if err != nil {
		return nil, err
	}

	r, err2 := protos.NewRegServClient(c).Confirm(context.Background(), in)
	if err != nil {
		return nil, err2
	}

	return r, nil
}

func Auth(in *protos.AuthReq) (*protos.AuthRep, error) {
	c, err := dial(defaultAddress)
	if err != nil {
		return nil, err
	}

	r, err2 := protos.NewAuthServClient(c).Auth(context.Background(), in)
	if err != nil {
		return nil, err2
	}

	return r, nil
}

func CheckToken(in *protos.CheckTokenReq) (*protos.CheckTokenRep, error) {
	c, err := dial(defaultAddress)
	if err != nil {
		return nil, err
	}

	r, err2 := protos.NewAuthServClient(c).CheckToken(
		context.Background(), in)
	if err != nil {
		return nil, err2
	}

	return r, nil
}

func GetRecommendation(in *protos.GetRecommendationReq) (*protos.GetRecommendationRep, error) {
	c, err := dial(defaultAddress)
	if err != nil {
		return nil, err
	}

	r, err2 := protos.NewRecommendationServClient(c).GetRecommendation(
		context.Background(), in)
	if err != nil {
		return nil, err2
	}

	return r, nil
}
