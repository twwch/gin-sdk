package mongoBase

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Base struct {
	database   string
	collection string
	Mongo *mongo.Client
}

func NewBase(host, database, collection string) *Base {
	return &Base{
		database,
		collection,
		setConnect(host),
	}
}

//连接设置
func setConnect(mongoUrl string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 连接池
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl).SetMaxPoolSize(20))
	if err != nil {
		panic(err)
	}
	return client
}

//插入单个文档
func (m *Base) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.InsertOne(ctx, document, opts...)
}

//插入多个文档
func (m *Base) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.InsertMany(ctx, documents, opts...)
}

// 查询单个文档
func (m *Base) FindOne(ctx context.Context, filter, result interface{}, opts ...*options.FindOneOptions) error {
	collection, err := m.Mongo.Database(m.database).Collection(m.collection).Clone()
	if err != nil {
		return err
	}
	return collection.FindOne(ctx, filter, opts...).Decode(result)
}

//查询多个文档
func (m *Base) Find(ctx context.Context, filter, results interface{}, opts ...*options.FindOptions) error {
	collection, err := m.Mongo.Database(m.database).Collection(m.collection).Clone()
	if err != nil {
		return err
	}
	cur, err := collection.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	return cur.All(ctx, results)
}

//更新单个文档
func (m *Base) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.UpdateOne(ctx, filter, update, opts...)
}

//更新多个文档
func (m *Base) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.UpdateMany(ctx, filter, update, opts...)
}

//删除文档
func (m *Base) Delete(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.DeleteOne(ctx, filter, opts...)

}

//删除多个文档
func (m *Base) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.DeleteMany(ctx, filter, opts...)
}