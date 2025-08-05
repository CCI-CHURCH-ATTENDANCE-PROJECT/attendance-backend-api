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

type FamilyMemberRepository struct {
	db         *database.Database
	collection *mongo.Collection
}

func NewFamilyMemberRepository(db *database.Database) *FamilyMemberRepository {
	return &FamilyMemberRepository{
		db:         db,
		collection: db.Collection("family_members"),
	}
}

func (r *FamilyMemberRepository) Create(ctx context.Context, familyMember *models.FamilyMember) error {
	familyMember.DateJoined = time.Now()

	result, err := r.collection.InsertOne(ctx, familyMember)
	if err != nil {
		return err
	}

	familyMember.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *FamilyMemberRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.FamilyMember, error) {
	var familyMember models.FamilyMember
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&familyMember)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &familyMember, nil
}

func (r *FamilyMemberRepository) GetByFamilyHead(ctx context.Context, familyHeadID primitive.ObjectID, page, limit int) ([]*models.FamilyMember, int, error) {
	offset := (page - 1) * limit

	filter := bson.M{"family_head": familyHeadID}

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

	var familyMembers []*models.FamilyMember
	if err = cursor.All(ctx, &familyMembers); err != nil {
		return nil, 0, err
	}

	return familyMembers, int(total), nil
}

func (r *FamilyMemberRepository) GetAll(ctx context.Context, page, limit int) ([]*models.FamilyMember, int, error) {
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

	var familyMembers []*models.FamilyMember
	if err = cursor.All(ctx, &familyMembers); err != nil {
		return nil, 0, err
	}

	return familyMembers, int(total), nil
}

// UpdateFamilyMember updates a family member's details by ID
func (r *FamilyMemberRepository) Update(ctx context.Context, familyMember *models.FamilyMember) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": familyMember.ID}, bson.M{"$set": familyMember})
	return err
}

func (r *FamilyMemberRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *FamilyMemberRepository) DeleteByFamilyHead(ctx context.Context, familyHeadID primitive.ObjectID) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{"family_head": familyHeadID})
	return err
}

func (r *FamilyMemberRepository) DeleteByMember(ctx context.Context, memberID primitive.ObjectID) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{"family_members": memberID})
	return err
}
