package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type consumerSuite struct {
	container testcontainers.Container
	db        *sql.DB
	repo      ConsumerReopsitory
	tableName string
	suite.Suite
}

func TestConsumerRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("short mode: skipping")
	}
	suite.Run(t, &consumerSuite{tableName: "consumer.consumers"})
}

func (c *consumerSuite) SetupSuite() {
	var err error

	ctx := context.Background()
	initDir, err := filepath.Abs("./../../../docker/database")
	if err != nil {
		c.T().Fatal(err)
	}
	const dbUrl = "postgres://ftgogo_user:ftgogo_pass@%s:%s/ftgogo?sslmode=disable"
	c.container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:15-alpine",
			Hostname:     "postgres",
			ExposedPorts: []string{"5432/tcp"},
			Env: map[string]string{
				"POSTGRES_PASSWORD": "itsasecret",
			},
			Mounts: []testcontainers.ContainerMount{
				testcontainers.BindMount(initDir, "/docker-entrypoint-initdb.d"),
			},
			WaitingFor: wait.ForSQL("5432/tcp", "pgx", func(host string, port nat.Port) string {
				return fmt.Sprintf(dbUrl, host, port.Port())
			}).WithStartupTimeout(10 * time.Second),
		},
		Started: true,
	})
	if err != nil {
		c.T().Fatal(err)
	}

	endpoint, err := c.container.Endpoint(ctx, "")
	if err != nil {
		c.T().Fatal(err)
	}

	c.db, err = sql.Open("pgx", fmt.Sprintf("postgres://ftgogo_user:ftgogo_pass@%s/ftgogo?sslmode=disable", endpoint))
	if err != nil {
		c.T().Fatal(err)
	}

}

func (c *consumerSuite) TearDownSuite() {
	err := c.db.Close()
	if err != nil {
		c.T().Fatal(err)
	}
	if err := c.container.Terminate(context.Background()); err != nil {
		c.T().Fatal(err)
	}
}

func (c *consumerSuite) TearDownTest() {
	query := fmt.Sprintf("DELETE from %s", c.tableName)
	_, err := c.db.ExecContext(context.Background(), query)
	if err != nil {
		c.T().Fatal(err)
	}
}

func (c *consumerSuite) SetupTest() {
	c.repo = NewConsumerReopsitory(c.tableName, c.db)
}
func (c *consumerSuite) TestConsumerReopsitory_Save() {
	consumer := domain.NewConsumer("id")
	consumer.Name = "name"
	consumer.Addresses =  map[string]domain.Address{"address-id": domain.Address{Street1: "street"}}

	err := c.repo.Save(context.Background(), consumer)
	c.NoError(err)
	query := fmt.Sprintf("SELECT name FROM %s WHERE id = $1", c.tableName)
	row := c.db.QueryRow(query, "id")
	c.NoError(row.Err())
	var name string
	c.NoError(row.Scan(&name))
	c.Equal(name, "name")
}

func (c *consumerSuite) TestConsumerReopsitory_Find() {
	consumer := domain.NewConsumer("id")
	consumer.Name = "name"
	consumer.Addresses =  map[string]domain.Address{"address-id": domain.Address{Street1: "street"}}
	err := c.repo.Save(context.Background(), consumer)
	c.NoError(err)
	consumer, err = c.repo.Find(context.Background(), "id")
	c.NoError(err)
	c.Equal(consumer.Name, "name")
	address, ok := consumer.Addresses["address-id"]
	c.True(ok)
	c.Equal(address.Street1, "street")
}

func (c *consumerSuite) TestConsumerReopsitory_Update() {
	consumer := domain.NewConsumer("id")
	consumer.Name = "name"
	consumer.Addresses =  map[string]domain.Address{"address-id": domain.Address{Street1: "street"}}
	err := c.repo.Save(context.Background(), consumer)
	c.NoError(err)
	consumer.Name = "changed"
	consumer.Addresses=map[string]domain.Address{}
	err = c.repo.Update(context.Background(),consumer)
	c.NoError(err)
	query := fmt.Sprintf("SELECT name FROM %s WHERE id = $1", c.tableName)
	row := c.db.QueryRow(query, "id")
	c.NoError(row.Err())
	var name string
	c.NoError(row.Scan(&name))
	c.Equal(name, "changed")

}
