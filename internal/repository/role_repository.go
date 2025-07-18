package repository

import (
	"context"
	"errors"
	"time"

	"church-attendance-api/internal/database"
	"church-attendance-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoleRepository struct {
	db         *database.Database
	collection *mongo.Collection
}

func NewRoleRepository(db *database.Database) *RoleRepository {
	return &RoleRepository{
		db:         db,
		collection: db.Collection("roles"),
	}
}

func (r *RoleRepository) Create(ctx context.Context, role *models.Role) error {
	role.DateAdded = time.Now()
	role.DateUpdated = time.Now()

	result, err := r.collection.InsertOne(ctx, role)
	if err != nil {
		return err
	}

	role.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *RoleRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Role, error) {
	var role models.Role
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&role)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) GetAll(ctx context.Context, page, limit int) ([]*models.Role, int, error) {
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
		SetSort(bson.D{{Key: "date_added", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var roles []*models.Role
	if err = cursor.All(ctx, &roles); err != nil {
		return nil, 0, err
	}

	return roles, int(total), nil
}

func (r *RoleRepository) GetByName(ctx context.Context, name string) (*models.Role, error) {
	var role models.Role
	err := r.collection.FindOne(ctx, bson.M{"role_name": name}).Decode(&role)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) Update(ctx context.Context, role *models.Role) error {
	role.DateUpdated = time.Now()

	filter := bson.M{"_id": role.ID}
	update := bson.M{"$set": role}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *RoleRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *RoleRepository) UpdateMemberCount(ctx context.Context, roleID primitive.ObjectID, count int) error {
	filter := bson.M{"_id": roleID}
	update := bson.M{"$set": bson.M{"total_members": count, "date_updated": time.Now()}}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}
