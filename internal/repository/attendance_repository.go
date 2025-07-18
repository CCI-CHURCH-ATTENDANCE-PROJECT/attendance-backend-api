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

type AttendanceRepository struct {
	db         *database.Database
	collection *mongo.Collection
}

func NewAttendanceRepository(db *database.Database) *AttendanceRepository {
	return &AttendanceRepository{
		db:         db,
		collection: db.Collection("attendance"),
	}
}

func (r *AttendanceRepository) Create(ctx context.Context, attendance *models.Attendance) error {
	attendance.DateTimeOfAttendance = time.Now()

	result, err := r.collection.InsertOne(ctx, attendance)
	if err != nil {
		return err
	}

	attendance.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *AttendanceRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Attendance, error) {
	var attendance models.Attendance
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&attendance)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &attendance, nil
}

func (r *AttendanceRepository) GetByUserAndDate(ctx context.Context, userID primitive.ObjectID, date time.Time) (*models.Attendance, error) {
	// Create date range for the day
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	var attendance models.Attendance
	filter := bson.M{
		"user": userID,
		"date_time_of_attendance": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}

	err := r.collection.FindOne(ctx, filter).Decode(&attendance)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &attendance, nil
}

func (r *AttendanceRepository) GetHistory(ctx context.Context, startDate, endDate *time.Time, page, limit int) ([]*models.Attendance, int, error) {
	offset := (page - 1) * limit

	filter := bson.M{}
	if startDate != nil && endDate != nil {
		filter["date_time_of_attendance"] = bson.M{
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
		SetSort(bson.D{{Key: "date_time_of_attendance", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var attendances []*models.Attendance
	if err = cursor.All(ctx, &attendances); err != nil {
		return nil, 0, err
	}

	return attendances, int(total), nil
}

func (r *AttendanceRepository) CountTotalForDate(ctx context.Context, date time.Time) (int, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	filter := bson.M{
		"date_time_of_attendance": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	return int(total), err
}

func (r *AttendanceRepository) CountMembersForDate(ctx context.Context, date time.Time) (int, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "date_time_of_attendance", Value: bson.D{
				{Key: "$gte", Value: startOfDay},
				{Key: "$lt", Value: endOfDay},
			}},
		}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "user"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "user_info"},
		}}},
		{{Key: "$match", Value: bson.D{
			{Key: "user_info.member", Value: true},
		}}},
		{{Key: "$count", Value: "total"}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		return 0, err
	}

	if len(result) == 0 {
		return 0, nil
	}

	return int(result[0]["total"].(int32)), nil
}

func (r *AttendanceRepository) CountVisitorsForDate(ctx context.Context, date time.Time) (int, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "date_time_of_attendance", Value: bson.D{
				{Key: "$gte", Value: startOfDay},
				{Key: "$lt", Value: endOfDay},
			}},
		}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "user"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "user_info"},
		}}},
		{{Key: "$match", Value: bson.D{
			{Key: "user_info.visitor", Value: true},
		}}},
		{{Key: "$count", Value: "total"}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		return 0, err
	}

	if len(result) == 0 {
		return 0, nil
	}

	return int(result[0]["total"].(int32)), nil
}

func (r *AttendanceRepository) GetAttendanceByDateRange(ctx context.Context, startDate, endDate time.Time) ([]map[string]interface{}, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "date_time_of_attendance", Value: bson.D{
				{Key: "$gte", Value: startDate},
				{Key: "$lte", Value: endDate},
			}},
		}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "user"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "user_info"},
		}}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{
				{Key: "year", Value: bson.D{{Key: "$year", Value: "$date_time_of_attendance"}}},
				{Key: "month", Value: bson.D{{Key: "$month", Value: "$date_time_of_attendance"}}},
				{Key: "day", Value: bson.D{{Key: "$dayOfMonth", Value: "$date_time_of_attendance"}}},
			}},
			{Key: "total_attendance", Value: bson.D{{Key: "$sum", Value: 1}}},
			{Key: "members", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$cond", Value: bson.A{
				bson.D{{Key: "$eq", Value: bson.A{"$user_info.member", true}}},
				1,
				0,
			}}}}}},
			{Key: "visitors", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$cond", Value: bson.A{
				bson.D{{Key: "$eq", Value: bson.A{"$user_info.visitor", true}}},
				1,
				0,
			}}}}}},
		}}},
		{{Key: "$sort", Value: bson.D{
			{Key: "_id.year", Value: -1},
			{Key: "_id.month", Value: -1},
			{Key: "_id.day", Value: -1},
		}}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []map[string]interface{}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
