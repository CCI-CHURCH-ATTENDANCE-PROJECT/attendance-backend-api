package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"church-attendance-api/internal/database"
	"church-attendance-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	db         *database.Database
	collection *mongo.Collection
}

func NewUserRepository(db *database.Database) *UserRepository {
	return &UserRepository{
		db:         db,
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	user.DateJoined = time.Now()
	user.DateUpdated = time.Now()

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByUserID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByQRToken(ctx context.Context, token string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"qr_code_token": token}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	user.DateUpdated = time.Now()

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *UserRepository) Search(ctx context.Context, query string, page, limit int) ([]*models.User, int, error) {
	offset := (page - 1) * limit

	// Create search filter
	searchFilter := bson.M{
		"$or": []bson.M{
			{"fname": bson.M{"$regex": query, "$options": "i"}},
			{"lname": bson.M{"$regex": query, "$options": "i"}},
			{"email": bson.M{"$regex": query, "$options": "i"}},
			{"user_id": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	// Count total documents
	total, err := r.collection.CountDocuments(ctx, searchFilter)
	if err != nil {
		return nil, 0, err
	}

	// Find documents
	findOptions := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "date_joined", Value: -1}})

	cursor, err := r.collection.Find(ctx, searchFilter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func (r *UserRepository) GetAll(ctx context.Context, page, limit int) ([]*models.User, int, error) {
	offset := (page - 1) * limit

	// Count total documents
	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	// Find documents
	findOptions := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "date_joined", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func (r *UserRepository) Filter(ctx context.Context, field, value string, page, limit int) ([]*models.User, int, error) {
	offset := (page - 1) * limit

	// Create filter
	filter := bson.M{}

	// Handle different field types
	switch field {
	case "member", "visitor", "usher", "admin", "family_head":
		if value == "true" {
			filter[field] = true
		} else if value == "false" {
			filter[field] = false
		} else {
			return nil, 0, fmt.Errorf("invalid boolean value for field %s", field)
		}
	case "gender", "campus_state", "campus_country", "profession":
		filter[field] = bson.M{"$regex": value, "$options": "i"}
	default:
		return nil, 0, fmt.Errorf("filtering not supported for field: %s", field)
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
		SetSort(bson.D{{Key: "date_joined", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func (r *UserRepository) UpdateQRToken(ctx context.Context, userID, token string) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$set": bson.M{
			"qr_code_token": token,
			"date_updated":  time.Now(),
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *UserRepository) UpdateQRCodeImage(ctx context.Context, userID, qrImage string) error {
	filter := bson.M{"user_id": userID}
	update:=bson.M{
		"$set":bson.M{
			"qr_code_image":qrImage,
			"date_updated": time.Now(), 
		},
	}
	_, err:=r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, token string) error {
	filter := bson.M{"qr_code_token": token}

	// Check if user exists
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no user with QR code token %s found", token)
	}

	// Delete user qr code token
	_, err = r.collection.UpdateOne(ctx, filter, bson.M{"$unset": bson.M{"qr_code_token": ""}})
	if err != nil {
		return fmt.Errorf("failed to unset QR code token: %w", err)
	}
	// Delete user document
	_, err = r.collection.DeleteOne(ctx, filter)
	return err
}

func (r *UserRepository) CountTotal(ctx context.Context) (int, error) {
	total, err := r.collection.CountDocuments(ctx, bson.M{})
	return int(total), err
}

func (r *UserRepository) CountMembers(ctx context.Context) (int, error) {
	total, err := r.collection.CountDocuments(ctx, bson.M{"member": true})
	return int(total), err
}

func (r *UserRepository) CountVisitors(ctx context.Context) (int, error) {
	total, err := r.collection.CountDocuments(ctx, bson.M{"visitor": true})
	return int(total), err
}
