package repository

import (
	"c1pherten/yet-webapp2/config"
	"c1pherten/yet-webapp2/appctx"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Repository interface { }

type repository struct {
	c appctx.Container
	db *mongo.Database
	mongoClient *mongo.Client
	redisClient *redis.Client
}

func connectMongo(c config.Config) (*mongo.Database, *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	// uri := fmt.Sprintf("mongo://%s:%s", c.Database.Host, c.Database.Port)
	uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// todo
	// db := client.Database(c.Database.DBName)
	db := client.Database("mydb")
	return db, client
}

func connectRedis(c config.Config) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := cli.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	return cli
}

// func NewRepository(c appctx.Container) * {
func NewRepository(c appctx.Container) repository {
	db, mCli := connectMongo(*c.Config())
	redisCli := connectRedis(*c.Config())
	return repository{
		c:           c,
		db:          db,
		mongoClient: mCli,
		redisClient: redisCli,
	}
}
