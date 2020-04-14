package filters

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetUserByDiscordID returns a filter to retreive documents by discordID
func GetUserByDiscordID(discordID string) bson.D {
	return bson.D{
		primitive.E{
			Key:   "discorduser.id",
			Value: discordID,
		},
	}
}
