package grpc

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/application/mocks"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serverSuite struct {
	app *mocks.App
	server *grpc.Server
	client consumerpb.ConsumerServiceClient
	suite.Suite
}

func TestServer(t *testing.T) {
	suite.Run(t, &serverSuite{})
}

func (s *serverSuite) SetupSuite()    {}
func (s *serverSuite) TearDownSuite() {}

func (s *serverSuite) SetupTest() {
	const grpcTestPort = ":10912"

	var err error
	// create server
	s.server = grpc.NewServer()
	var listener net.Listener
	listener, err = net.Listen("tcp", grpcTestPort)
	if err != nil {
		s.T().Fatal(err)
	}

	// create app mock
	s.app = mocks.NewApp(s.T())

	// register app with the server
	if err = RegisterServer(s.app, s.server); err != nil {
		s.T().Fatal(err)
	}
	go func(listener net.Listener) {
		err := s.server.Serve(listener)
		if err != nil {
			s.T().Fatal(err)
		}
	}(listener)

	// create client
	var conn *grpc.ClientConn
	conn, err = grpc.NewClient(grpcTestPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.T().Fatal(err)
	}
	s.client = consumerpb.NewConsumerServiceClient(conn)
}

func(s *serverSuite)TearDownTest(){
	s.server.GracefulStop()
}

func(s *serverSuite)TestConsumerService_RegisterConsumer(){
	s.app.On("RegisterConsumer", mock.Anything, mock.AnythingOfType("application.RegisterConsumer")).Return(nil)
	_, err := s.client.RegisterConsumer(context.Background(),&consumerpb.RegisterConsumerRequest{
		Name: "name",
	})
	s.NoError(err)
}
func(s *serverSuite)TestConsumerService_RegisterConsumerFailed(){
	s.app.On("RegisterConsumer", mock.Anything, mock.AnythingOfType("application.RegisterConsumer")).
	Return(errors.New("RegisterConsumer failed"))
	_, err := s.client.RegisterConsumer(context.Background(),&consumerpb.RegisterConsumerRequest{
		Name: "name",
	})
	s.Error(err)
}

func(s *serverSuite)TestConsumerService_GetAddress(){
	s.app.On("GetConsumerAddress", mock.Anything, application.GetConsumerAddress{ConsumerID: "id",AddressID: "address-id"}).
	Return(domain.Address{Street1: "street"}, nil)
	_, err := s.client.GetAddress(context.Background(),&consumerpb.GetAddressRequest{
		ConsumerID: "id",
		AddressID: "address-id",
	})
	s.NoError(err)
}
