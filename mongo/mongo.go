package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Database defining basic interface
type Database interface {
	Collection(string) Collection
	Client() Client
}

// mongoDatabase struct defining mongo database
type mongoDatabase struct {
	db *mongo.Database
}

// Collection interface defining database operation
type Collection interface {
	FindOne(context.Context, interface{}) SingleResult
	Find(context.Context, interface{}, ...*options.FindOptions) (Cursor, error)
	InsertOne(context.Context, interface{}) (interface{}, error)
	InsertMany(context.Context, []interface{}) ([]interface{}, error)
	DeleteOne(context.Context, interface{}) (int64, error)
	CountDocuments(context.Context, interface{}, ...*options.CountOptions) (int64, error)
	Aggregate(context.Context, interface{}) (Cursor, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

// mongoCollection struct defining mongo collection
type mongoCollection struct {
	coll *mongo.Collection
}

// SingleResult interface defining basic database operation
type SingleResult interface {
	Decode(interface{}) error
}

// mongoSingleResult struct for fetching single document
type mongoSingleResult struct {
	sr *mongo.SingleResult
}

// Cursor interface defining properties to implement pagination
type Cursor interface {
	Next(context.Context) bool
	Close(context.Context) error
	Decode(interface{}) error
	All(context.Context, interface{}) error
}

// mongoCursor struct for fetching multiple documents
type mongoCursor struct {
	mc *mongo.Cursor
}

// Client interface defining basic database operation
type Client interface {
	Database(string) Database
	Disconnect(context.Context) error
	StartSession() (mongo.Session, error)
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
	Ping(context.Context) error
}

// mongoClient struct defining mongo client
type mongoClient struct {
	cl *mongo.Client
}

// mongoSession struct for defining mongo session
type mongoSession struct {
	mongo.Session
}

// NewClient function creating new mongo client
func NewClient(ctx context.Context, connection string) (Client, error) {
	time.Local = time.UTC
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))

	if err != nil {
		return nil, err
	}

	return &mongoClient{cl: c}, nil
}

// Client methods

// Database method returning database handler
func (mc *mongoClient) Database(dbName string) Database {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

// Disconnect method to disconnect the database connection
func (mc *mongoClient) Disconnect(ctx context.Context) error {
	return mc.cl.Disconnect(ctx)
}

// StartSession method to start the session
func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.cl.StartSession()
	return &mongoSession{session}, err
}

// UseSession method to create new session and use it as SessionContext
func (mc *mongoClient) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return mc.cl.UseSession(ctx, fn)
}

// Ping method to check database connection
func (mc *mongoClient) Ping(ctx context.Context) error {
	return mc.cl.Ping(ctx, readpref.Primary())
}

// Database methods

// Collection method to get the handle for a particular collection
func (md *mongoDatabase) Collection(colName string) Collection {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

// Client method to return the client through which the database was performed
func (md *mongoDatabase) Client() Client {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

// Collection methods

// FindOne method to find a particular document using filters
func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

// Find method to return cursor over the matching documents
func (mc *mongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	findResult, err := mc.coll.Find(ctx, filter, opts...)
	return &mongoCursor{mc: findResult}, err
}

// InsertOne method to insert a particular document into specified collection
func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)
	return id.InsertedID, err
}

// InsertMany method to insert multiple documents into specified collection
func (mc *mongoCollection) InsertMany(ctx context.Context, document []interface{}) ([]interface{}, error) {
	res, err := mc.coll.InsertMany(ctx, document)
	return res.InsertedIDs, err
}

// DeleteOne method to delete a particular document using filter
func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := mc.coll.DeleteOne(ctx, filter)
	return count.DeletedCount, err
}

// CountDocuments method to return counts of documents based on filters
func (mc *mongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return mc.coll.CountDocuments(ctx, filter, opts...)
}

// Aggregate method returns cursor over resulting documents after executing aggregate cmd
func (mc *mongoCollection) Aggregate(ctx context.Context, pipeline interface{}) (Cursor, error) {
	aggregateResult, err := mc.coll.Aggregate(ctx, pipeline)
	return &mongoCursor{mc: aggregateResult}, err
}

// UpdateOne method to update a paritcular document using filters
func (mc *mongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.coll.UpdateOne(ctx, filter, update, opts[:]...)
}

// UpdateMany method updates multiple documents based on filters
func (mc *mongoCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.coll.UpdateMany(ctx, filter, update, opts[:]...)
}

// Cursor methods

// Next method returns the next document by the cursor
func (mr *mongoCursor) Next(ctx context.Context) bool {
	return mr.mc.Next(ctx)
}

// Close method closes the cursor
func (mr *mongoCursor) Close(ctx context.Context) error {
	return mr.mc.Close(ctx)
}

// Decode method to unmarshall the document
func (mr *mongoCursor) Decode(v interface{}) error {
	return mr.mc.Decode(v)
}

// All method iterates the cursor and decode each document
func (mr *mongoCursor) All(ctx context.Context, result interface{}) error {
	return mr.mc.All(ctx, result)
}

// SingleResult methods

// Decode method to unmarshall the document
func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}
