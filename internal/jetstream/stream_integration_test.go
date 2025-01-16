package jetstream

// import (
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/nats-io/nats.go"
// 	"github.com/rezaAmiri123/ftgogoV3/internal/am"
// 	"github.com/stretchr/testify/suite"
// 	"github.com/testcontainers/testcontainers-go"
// 	"github.com/testcontainers/testcontainers-go/wait"
// )

// type JSSuite struct {
// 	container  testcontainers.Container
// 	nc         *nats.Conn
// 	js         nats.JetStreamContext
// 	streamName string
// 	stream     *Stream
// 	suite.Suite
// }

// func TestJetStreamSuite(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("short mode: skipping")
// 	}
// 	suite.Run(t, &JSSuite{streamName: "ftgogo"})
// }

// func (c *JSSuite) SetupSuite() {
// 	var err error

// 	ctx := context.Background()
// 	c.container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
// 		ContainerRequest: testcontainers.ContainerRequest{
// 			Image:        "nats:2-alpine",
// 			Hostname:     "nats",
// 			ExposedPorts: []string{"4222/tcp"},
// 			Cmd:          []string{"-js", "-sd", "/var/lib/nats/data"},
// 			WaitingFor: wait.ForLog("Server is ready").WithStartupTimeout(10*time.Second),
// 		},
// 		Started: true,
// 	})

// 	if err != nil {
// 		c.T().Fatal(err)
// 	}

// 	endpoint, err := c.container.Endpoint(ctx, "")
// 	if err != nil {
// 		c.T().Fatal(err)
// 	}
// 	c.nc, err = nats.Connect(endpoint)
// 	if err != nil {
// 		c.T().Fatal(err)
// 	}

// 	c.js, err = c.nc.JetStream()
// 	if err != nil {
// 		c.T().Fatal(err)
// 	}

// 	_, err = c.js.AddStream(&nats.StreamConfig{
// 		Name:     c.streamName,
// 		Subjects: []string{fmt.Sprintf("%s.>", c.streamName)},
// 	})
// 	if err != nil {
// 		c.T().Fatal(err)
// 	}

// }

// func (c *JSSuite) TearDownSuite() {
// 	c.nc.Close()
// 	if err := c.container.Terminate(context.Background()); err != nil {
// 		c.T().Fatal(err)
// 	}
// }

// func (c *JSSuite) TearDownTest() {

// }

// func (c *JSSuite) SetupTest() {
// 	c.stream = NewStream(c.streamName, c.js)
// }
// func (c *JSSuite) TestStream_PublishSubscribe() {
// 	topic := fmt.Sprintf("%s.%s", c.streamName, "topic")
// 	err := c.stream.Publish(context.Background(), topic, &rawMessage{
// 		id:   "id",
// 		name: "name",
// 		data: []byte{},
// 	})
// 	if err != nil {
// 		c.T().Fatal(err)
// 	}
// 	fn := am.MessageHandlerFunc[am.RawMessage](func(ctx context.Context, msg am.RawMessage) error {
// 		c.Equal(msg.ID(),"id")
// 		c.Equal(msg.MessageName(),"name")
// 		return nil
// 	})
// 	c.stream.Subscribe(topic,fn)

// }
