package phone

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SpamPhone struct {
	ID primitive.ObjectID `bson:"_id"`
	PhoneNumber string `bson:"phoneNumber"`
	Note string `bson:"note"`
	Type string `bson:type`
	ReportedTimes int `bson:"reportedTimes"`
	Reports []Report `bson:"reports"`
}


type Report struct {
	UserId primitive.ObjectID  `bson:"user_id"`
	Note string `bson:"note"`
	Type string `bson:type`
	CreatedDate time.Time `bson:"createdDate"`
}
