package backend

import (
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Notes   *mongo.Collection
	DB      *mongo.Database
	limiter *RateLimiter
)

type RateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*Visitor
	limit    int
	window   time.Duration
}

type Visitor struct {
	tokens   int
	lastSeen time.Time
}

type Note struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
