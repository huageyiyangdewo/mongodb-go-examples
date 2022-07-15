package curd

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func Update()  {
	//updateOne()

	updateMany()
}

// db.collection.update({}, {$set:{stat: NumberInt(143)}})  局部修改
func updateOne()  {
	commentCol := MongoDBClient.Database(MongoDBName).Collection(CollectionName)
	// db.collection.update({comment_id: 1tg}, {start: NumberInt(15678)})

	result, err := commentCol.UpdateOne(context.TODO(),
		bson.D{
		{"comment_id", "1tg"},
		},
		bson.D{
		{"$set", bson.D{{"star", 15678}}},
		},
		)
	if err != nil {
		fmt.Printf("updateOne failed, err:%s \n", err)
	}


	fmt.Printf("result --> %+v \n", result)
}

// db.collection.update({}, {$set:{stat: NumberInt(143)}}, {multi: true})  批量修改
func updateMany()  {
	commentCol := MongoDBClient.Database(MongoDBName).Collection(CollectionName)
	// db.collection.update({comment_id: 1tg}, {start: NumberInt(15678)}, {multi: true})

	result, err := commentCol.UpdateMany(context.TODO(),
		bson.D{
			{"comment_id", "1"},
		},
		bson.D{
			{"$set", bson.D{{"star", 15678}, {"create_at", time.Now()}}},
		},
	)
	if err != nil {
		fmt.Printf("updateMany failed, err:%s \n", err)
	}


	fmt.Printf("result --> %+v \n", result)
}

