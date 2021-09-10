package mongoData

import (
	"crashPri/config"
	"context"
	"crashPri/logger"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var MongoDB *MongoDatabase

func init() {
	MongoDB = &MongoDatabase{
		MongoClient: SetConnect(),
	}
}

type MongoDatabase struct {
	MongoClient *mongo.Client
}

func (m *MongoDatabase) GetCollection(database string, collection string) *mongo.Collection {
	return m.MongoClient.Database(database).Collection(collection)
}

func (m *MongoDatabase) Refresh() {

}
func (m *MongoDatabase) FindOne(database string, collection_ string, filter bson.D, object interface{}, retryTime int) error {
	time := 0
	collection := m.GetCollection(database, collection_)
	err := collection.FindOne(context.Background(), filter).Decode(object)
	for err != nil && err != mongo.ErrNoDocuments && time < retryTime {
		err = collection.FindOne(context.Background(), filter).Decode(object)
		time++
	}
	if err != nil && err != mongo.ErrNoDocuments {
		return errors.Wrap(err, "mongo FindOne error")
	} else if err == mongo.ErrNoDocuments {
		return err
	}
	return nil
}

func (m *MongoDatabase) InsertOne(database string, collection_ string, data bson.M, retryTime int) error {
	time := 0
	collection := m.GetCollection(database, collection_)
	_, err := collection.InsertOne(context.Background(), data)
	for err != nil && time < retryTime {
		_, err = collection.InsertOne(context.Background(), data)
		time++
	}
	if err != nil {
		return errors.Wrap(err, "mongo InsertOne error")
	}
	return nil
}

//	通过指定更新操作符，来实现refresh操作
func (m *MongoDatabase) UpdateOne(database string, collection_ string, filter bson.D, data bson.D, retryTime int) error {
	time := 0
	collection := m.GetCollection(database, collection_)
	_, err := collection.UpdateOne(context.Background(), filter, data)
	for err != nil && time < retryTime {
		_, err = collection.UpdateOne(context.Background(), filter, data)
		time++
	}
	if err != nil {
		return errors.Wrap(err, "mongo UpdateOne error")
	}
	return nil
}

// 连接设置
func SetConnect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MONGO_URL).SetMaxPoolSize(20)) // 连接池
	if err != nil {
		panic(err)
	}
	logger.Logger.Info().Str("mongodb link", config.MONGO_URL).Msg("连接上mongodb")
	return client
}
