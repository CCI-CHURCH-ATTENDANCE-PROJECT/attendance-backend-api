package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"cci-api/internal/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewConnection(cfg *config.Config) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var uri string
	// Use DB_URI to connect to the db if it is in the .env, else build connection URI
	if cfg.DB_URI != "" {
		uri = cfg.DB_URI
	} else {
		if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBName == "" {
			return nil, fmt.Errorf("db_host, db_port, and db_name must be set in the configuration for they are required")
		}
		// Build connection URI
		if cfg.DBUser != "" && cfg.DBPassword != "" {
			uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
				cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
		} else {
			uri = fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
		}
	}




	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	db := client.Database(cfg.DBName)

	log.Println("Successfully connected to MongoDB!")

	return &Database{
		Client: client,
		DB:     db,
	}, nil
}

func (d *Database) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return d.Client.Disconnect(ctx)
}

// Collection returns a MongoDB collection
func (d *Database) Collection(name string) *mongo.Collection {
	return d.DB.Collection(name)
}

// CreateIndexes creates necessary indexes for the collections
func (d *Database) CreateIndexes() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Users collection indexes
	usersCollection := d.Collection("users")
	_, err := usersCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    map[string]interface{}{"email": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    map[string]interface{}{"user_id": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: map[string]interface{}{"qr_code_token": 1},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create users indexes: %w", err)
	}

	// Attendance collection indexes
	attendanceCollection := d.Collection("attendance")
	_, err = attendanceCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: map[string]interface{}{"user": 1},
		},
		{
			Keys: map[string]interface{}{"date_time_of_attendance": -1},
		},
		{
			Keys: bson.D{
				{Key: "user", Value: 1},
				{Key: "date_time_of_attendance", Value: -1},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create attendance indexes: %w", err)
	}

	// Family members collection indexes
	familyCollection := d.Collection("family_members")
	_, err = familyCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: map[string]interface{}{"family_head": 1},
		},
		{
			Keys: map[string]interface{}{"family_members": 1},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create family_members indexes: %w", err)
	}

	// Sermons collection indexes
	sermonsCollection := d.Collection("sermons")
	_, err = sermonsCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: map[string]interface{}{"date_of_meeting": -1},
		},
		{
			Keys: map[string]interface{}{"entry_made_by": 1},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create sermons indexes: %w", err)
	}

	// Announcements collection indexes
	announcementsCollection := d.Collection("announcements")
	_, err = announcementsCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: map[string]interface{}{"announcement_date": -1},
		},
		{
			Keys: map[string]interface{}{"status": 1},
		},
		{
			Keys: map[string]interface{}{"announcement_entry_made_by": 1},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create announcements indexes: %w", err)
	}

	// Refresh tokens collection indexes
	refreshTokensCollection := d.Collection("refresh_tokens")
	_, err = refreshTokensCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    map[string]interface{}{"token": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: map[string]interface{}{"user_id": 1},
		},
		{
			Keys:    map[string]interface{}{"expires_at": 1},
			Options: options.Index().SetExpireAfterSeconds(0),
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create refresh_tokens indexes: %w", err)
	}

	log.Println("Database indexes created successfully!")
	return nil
}
