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

type AnnouncementRepository struct {
	db         *database.Database
	collection *mongo.Collection
}

func NewAnnouncementRepository(db *database.Database) *AnnouncementRepository {
	return &AnnouncementRepository{
		db:         db,
		collection: db.Collection("announcements"),
	}
}

func (r *AnnouncementRepository) Create(ctx context.Context, announcement *models.Announcement) error {
	announcement.DateAdded = time.Now()
	announcement.Status = "Pending"

	result, err := r.collection.InsertOne(ctx, announcement)
	if err != nil {
		return err
	}

	announcement.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *AnnouncementRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Announcement, error) {
	var announcement models.Announcement
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&announcement)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &announcement, nil
}

func (r *AnnouncementRepository) GetAll(ctx context.Context, page, limit int, status string) ([]*models.Announcement, int, error) {
	offset := (page - 1) * limit

	filter := bson.M{}
	if status != "" {
		filter["status"] = status
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
		SetSort(bson.D{{Key: "announcement_date", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var announcements []*models.Announcement
	if err = cursor.All(ctx, &announcements); err != nil {
		return nil, 0, err
	}

	return announcements, int(total), nil
}

func (r *AnnouncementRepository) Update(ctx context.Context, announcement *models.Announcement) error {
	filter := bson.M{"_id": announcement.ID}
	update := bson.M{"$set": announcement}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *AnnouncementRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *AnnouncementRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status}}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}
