package mongodb

import (
	"context"
	"fmt"
	model "github.com/gobuz/publicspam/internal/domain/model/phone"
	port "github.com/gobuz/publicspam/internal/domain/port/repository"
	entity "github.com/gobuz/publicspam/internal/infrastructure/mongodb/model/phone"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type repository struct {
	context context.Context
	client *mongo.Client
	database *mongo.Database
}

// Init mongo adaptor
// uri to mongodb, db default database name
func Init(context context.Context,uri string, db string) port.Repository {
	client, err := mongo.Connect(context, options.Client().ApplyURI(uri))
	if err == nil {
		err = client.Ping(context, readpref.Primary())
	}
	if err != nil {
		panic(err)
	}
	return &repository{
		context: context,
		client: client,
		database: client.Database(db),
	}
}

func (db *repository) Close() error {
	return db.client.Disconnect(db.context)
}

// GetAllSpamPhones get all of spam phones
func (db *repository) GetAllSpamPhones() []model.SpamPhone {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	cur, err := db.database.Collection(TableSpamPhone).Find(ctx, bson.D{})
	var spamPhones []model.SpamPhone

	if err != nil {
		fmt.Println(err)
		return spamPhones
		// panic(err) //TODO: handle error
	}
	defer func() {
		err = cur.Close(db.context)
		if err != nil {
			fmt.Println(err)
		}
	}()

	for cur.Next(db.context) {
		var result entity.SpamPhone
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		spamPhones = append(spamPhones, entity.ExportSpamPhone(result))
	}
	return spamPhones
}

// GetAllSpamPhones get all of spam phones
func (db *repository) GetSpamPhone(phoneNumber string) model.SpamPhone {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	var spamPhone model.SpamPhone
	var dbPhone entity.SpamPhone
	err := db.database.Collection(TableSpamPhone).FindOne(ctx, bson.M{
			entity.FieldPhoneNumber(): phoneNumber,
	}).Decode(&dbPhone)
	if err != nil {
		fmt.Println(err)
		return spamPhone
		// panic(err) //TODO: handle error
	}
	return entity.ExportSpamPhone(dbPhone)
}