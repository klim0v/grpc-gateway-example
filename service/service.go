package service

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	"github.com/klim0v/grpc-gateway-example/api_pb"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type BlockchainServer struct {
	eventBus <-chan interface{}
}

func NewBlockchainServer(eventBus <-chan interface{}) *BlockchainServer {
	return &BlockchainServer{eventBus: eventBus}
}

func (*BlockchainServer) DownloadFile(_ context.Context, _ *empty.Empty) (*httpbody.HttpBody, error) {
	return &httpbody.HttpBody{
		ContentType: "application/octet-stream",
		Data:        []byte("Hello World"),
	}, nil
}

func (b *BlockchainServer) Address(_ context.Context, req *api_pb.AddressRequest) (*api_pb.AddressResponse, error) {
	if req.Address != "Mxb9a117e772a965a3fddddf83398fd8d71bf57ff6" {
		return &api_pb.AddressResponse{}, status.Error(codes.FailedPrecondition, "wallet not found")
	}
	return &api_pb.AddressResponse{
		Balance: map[string]string{
			"BIP": "12345678987654321",
		},
		TransactionsCount: "120",
	}, nil
}

func (b *BlockchainServer) Subscribe(req *api_pb.SubscribeRequest, stream api_pb.BlockchainService_SubscribeServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case event := <-b.eventBus:
			byteData, err := json.Marshal(event)
			if err != nil {
				return err
			}
			var bb bytes.Buffer
			bb.Write(byteData)
			data := &_struct.Struct{Fields: make(map[string]*_struct.Value)}
			if err := (&jsonpb.Unmarshaler{}).Unmarshal(&bb, data); err != nil {
				return err
			}

			if err := stream.Send(&api_pb.SubscribeResponse{
				Query: req.Query,
				Data:  data,
				Events: []*api_pb.SubscribeResponse_Event{
					{
						Key:    "tx.hash",
						Events: []string{"01EFD8EEF507A5BFC4A7D57ECA6F61B96B7CDFF559698639A6733D25E2553539"},
					},
				},
			}); err != nil {
				return err
			}
		case <-time.After(5 * time.Second):
			return nil
		}
	}
}
