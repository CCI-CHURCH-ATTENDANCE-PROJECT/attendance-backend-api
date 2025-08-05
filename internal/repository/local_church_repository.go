package repository

import (
	"context"
	"errors"

	"cci-api/internal/database"
	"cci-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LocalChurchRepository struct {
	db         *database.Database
	collection *mongo.Collection
}

func NewLocalChurchRepository(db *database.Database) *LocalChurchRepository {
	return &LocalChurchRepository{
		db:         db,
		collection: db.Collection("church_info"),
	}
}

func (r *LocalChurchRepository) Create(ctx context.Context, church *models.LocalChurch) error {
	result, err := r.collection.InsertOne(ctx, church)
	if err != nil {
		return err
	}

	church.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *LocalChurchRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.LocalChurch, error) {
	var church models.LocalChurch
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&church)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &church, nil
}

func (r *LocalChurchRepository) GetFirst(ctx context.Context) (*models.LocalChurch, error) {
	var church models.LocalChurch
	err := r.collection.FindOne(ctx, bson.M{}).Decode(&church)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &church, nil
}

// func (r *LocalChurchRepository) GetAll(ctx context.Context) (*models.LocalChurch, error) {
// 	var church models.LocalChurch
// 	err := r.collection.GetAll(ctx, bson.M{}).Decode(&church)
// 	if err != nil {
// 		if  errors.Is(err, mongo.ErrNoDocuments){
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &church, nil
// }

func (r *LocalChurchRepository) GetAll(ctx context.Context, page, limit int) ([]*models.LocalChurch, int, error) {
	skip := (page - 1) * limit

	// Get total count
	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	// Get churches with pagination
	cursor, err := r.collection.Find(ctx, bson.M{}, options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var churches []*models.LocalChurch
	if err := cursor.All(ctx, &churches); err != nil {
		return nil, 0, err
	}

	return churches, int(total), nil
}

func (r *LocalChurchRepository) Update(ctx context.Context, church *models.LocalChurch) error {
	filter := bson.M{"_id": church.ID}
	update := bson.M{"$set": church}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *LocalChurchRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
