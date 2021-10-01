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

func (m *Base) FindOneAndDelete(ctx context.Context, filter interface{}, result interface{}, opts ...*options.FindOneAndDeleteOptions) error {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.FindOneAndDelete(ctx, filter, opts...).Decode(result)
}

func (m *Base) FindOneAndReplace(ctx context.Context,filter, replacement, result interface{},opts ...*options.FindOneAndReplaceOptions) error {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.FindOneAndReplace(ctx, filter, replacement, opts...).Decode(result)
}

func (m *Base) FindOneAndUpdate(ctx context.Context, filter, update, result interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.FindOneAndUpdate(ctx, filter, update, opts...).Decode(result)
}

func (m *Base) FindOneAndUpsert(ctx context.Context, filter, update, result interface{}, opts ...*options.FindOneAndUpdateOptions) error {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	rd := options.After
	optUpsert := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(rd)
	opts = append(opts, optUpsert)
	return collection.FindOneAndUpdate(ctx, filter, update, opts...).Decode(result)
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

func (m *Base) Upsert(ctx context.Context, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	optUpsert := options.Update().SetUpsert(true)
	opts = append(opts, optUpsert)
	return collection.UpdateOne(ctx, filter, replacement, opts...)
}

//删除文档
func (m *Base) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.DeleteOne(ctx, filter, opts...)

}

//删除多个文档
func (m *Base) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.DeleteMany(ctx, filter, opts...)
}

func (m *Base) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)  {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.CountDocuments(ctx, filter, opts...)
}

func (m *Base) ReplaceOne(ctx context.Context,filter, replacement interface{},opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	collection := m.Mongo.Database(m.database).Collection(m.collection)
	return collection.ReplaceOne(ctx, filter, replacement, opts...)
}