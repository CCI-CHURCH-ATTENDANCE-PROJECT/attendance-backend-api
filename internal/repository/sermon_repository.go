package repository

import (
	"context"
	"errors"
	"time"

	"cci-api/internal/database"
	"cci-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SermonRepository struct {
	db         *database.Database
	collection *mongo.Collection
}

func NewSermonRepository(db *database.Database) *SermonRepository {
	return &SermonRepository{
		db:         db,
		collection: db.Collection("sermons"),
	}
}

func (r *SermonRepository) Create(ctx context.Context, sermon *models.Sermon) error {
	// sermon.DateOfMeeting = time.Now()

	result, err := r.collection.InsertOne(ctx, sermon)
	if err != nil {
		return err
	}

	sermon.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *SermonRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Sermon, error) {
	var sermon models.Sermon
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&sermon)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &sermon, nil
}

func (r *SermonRepository) GetAll(ctx context.Context, page, limit int, startDate, endDate *time.Time) ([]*models.Sermon, int, error) {
	offset := (page - 1) * limit

	filter := bson.M{}
	if startDate != nil && endDate != nil {
		filter["date_of_meeting"] = bson.M{
			"$gte": *startDate,
			"$lte": *endDate,
		}
	}

	// Count total documents
	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	// Find documents
	findOptions := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "date_of_meeting", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var sermons []*models.Sermon
	if err = cursor.All(ctx, &sermons); err != nil {
		return nil, 0, err
	}

	return sermons, int(total), nil
}

func (r *SermonRepository) Update(ctx context.Context, sermon *models.Sermon) error {
	filter := bson.M{"_id": sermon.ID}
	update := bson.M{"$set": sermon}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *SermonRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
