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
)

type RefreshTokenRepository struct {
	db         *database.Database
	collection *mongo.Collection
}

func NewRefreshTokenRepository(db *database.Database) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db:         db,
		collection: db.Collection("refresh_tokens"),
	}
}

func (r *RefreshTokenRepository) Create(ctx context.Context, token *models.RefreshToken) error {
	token.CreatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, token)
	if err != nil {
		return err
	}

	token.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *RefreshTokenRepository) GetByToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	err := r.collection.FindOne(ctx, bson.M{"token": token}).Decode(&refreshToken)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &refreshToken, nil
}

func (r *RefreshTokenRepository) DeleteByToken(ctx context.Context, token string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"token": token})
	return err
}

func (r *RefreshTokenRepository) DeleteByUserID(ctx context.Context, userID primitive.ObjectID) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{"user_id": userID})
	return err
}

func (r *RefreshTokenRepository) DeleteExpired(ctx context.Context) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{
		"expires_at": bson.M{"$lt": time.Now()},
	})
	return err
}
