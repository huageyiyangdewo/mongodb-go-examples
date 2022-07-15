package curd

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func Delete()  {
	//collectionDrop()

	//removeOne()

	removeMany()
}

// db.collection.drop()
func collectionDrop()  {
	// 删除 集合
	err := MongoDBClient.Database(MongoDBName).Collection(CollectionName).Drop(context.TODO())
	if err != nil {
		fmt.Printf("collectionDrop failed, err:%s \n", err)
	}
}


// db.collection.remove({comment_id: "1tg"})
func removeOne()  {
	commentCol := MongoDBClient.Database(MongoDBName).Collection(CollectionName)
	// db.collection.remove({comment_id: "1tg"})

	result, err := commentCol.DeleteOne(context.TODO(),
		bson.D{
			{"comment_id", "1tg"},
		},
		)

	if err != nil {
		fmt.Printf("removeOne failed, err:%s \n", err)
	}


	fmt.Printf("result --> %+v \n", result)
}



// db.collection.remove()
func removeMany()  {
	commentCol := MongoDBClient.Database(MongoDBName).Collection(CollectionName)
	// db.collection.remove({star: {$lt: 5}})


	result, err := commentCol.DeleteMany(context.TODO(),
		bson.D{
			{"star", bson.D{{"$lt", 5}}},
		},
	)

	if err != nil {
		fmt.Printf("removeMany failed, err:%s \n", err)
	}


	fmt.Printf("result --> %+v \n", result)
}
