package db

import (
	"context"
	"fmt"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Environments -
type Environments struct {
	Cluster  string `json:"cluster"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
	Mode     string `json:"mode"`
	Host     string `json:"host"`
}

//Connect - new connection
func Connect(ctx context.Context, envs *Environments) (*mongo.Database, error) {
	var defaultURI string

	switch envs.Mode {
	case "dev":
		defaultURI = fmt.Sprintf("mongodb%s://%s:%s@%s/%s", "", envs.User, url.QueryEscape(envs.Password), envs.Host, envs.Database)
	default:
		defaultURI = fmt.Sprintf("mongodb%s://%s:%s@%s", "+srv", envs.User, url.QueryEscape(defaultURI), envs.Cluster)
	}
	fmt.Println(defaultURI)
	options := options.Client().ApplyURI(defaultURI)
	client, err := mongo.NewClient(options)

	if err != nil {
		fmt.Println("DATABASE mongo: Error try to start a new client")
		return nil, err
	}

	err = client.Connect(ctx)

	if err != nil {
		fmt.Println("DATABASE mongo: Error try to start a new connection")
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("DATABASE mongo: Error try pinging")
		return nil, err
	}

	check := client.Database(envs.Database).RunCommand(ctx, bson.M{"serverStatus": 1})
	if check.Err() != nil {
		fmt.Println("DATABASE mongo: Error try to check status")
		return nil, err
	}

	fmt.Println("DATABAE mongo: connection with database: ", envs.Database)
	return client.Database(envs.Database), nil
}
