package filters

import (
	"time"

	"github.com/mikevyt/rollout/models"
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

func UpdateUser(discorduser models.DiscordUser) bson.D {
	return bson.D{primitive.E{
		Key: "$set",
		Value: bson.D{
			primitive.E{Key: "discorduser", Value: discorduser},
			primitive.E{Key: "updatedate", Value: time.Now()},
		},
	}}
}
