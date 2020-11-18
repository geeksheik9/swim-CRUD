package db

import "go.mongodb.org/mongo-driver/mongo"

//WorkoutDB is the data access object for accessing swim workouts
type WorkoutDB struct {
	client        *mongo.Client
	datbaseName   string
	setCollection string
}
