package curd

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Query()  {
	//find()

	//findAndSortSkipLimit()

	//findByRegex()

	findByAndOr()
}

// db.collection.find()
func find()  {
	commentCol := MongoDBClient.Database(MongoDBName).Collection(CollectionName)
	// db.collection.find({comment_id: 2}, {comment_id:1, _id:0, author:1})

	// 这里做用 原先定义好的结构体接受时，不需要的字段会被赋予空值，暂时没有找到如何除去，
	// 结构体中添加 omitempty 无效
	projection := bson.D{
		{"comment_id", 1},
		{"_id", 0},
		{"author", 1},
	}
	commentCursor, err := commentCol.Find(context.TODO(),
		bson.D{
		{"comment_id", 2},
		},
		options.Find().SetProjection(projection),
		)
	if err != nil {
		fmt.Printf("commentCol.Find failed, err:%s \n", err)
	}
	// 注意这里使用了自定义的 string 解析器，具体看 conf.ConnectMongo,否则会报错
	var comment []Comment

	err = commentCursor.All(context.TODO(), &comment)

	if err != nil {
		fmt.Printf("commentCursor.All failed, err:%s \n", err)
	}

	fmt.Printf("comment --> %+v \n", comment)

}

// db.collection.find().sort().skip().limit()
func findAndSortSkipLimit()  {
	commentCol := MongoDBClient.Database(MongoDBName).Collection(CollectionName)
	// db.comment.find({star: {$gt: 3}}).sort({"star":1}).skip(1).limit(1)


	commentCursor, err := commentCol.Find(context.TODO(),
		bson.D{
			{"star", bson.D{{"$gt", 3}}},
		},
		options.Find().SetSort(bson.D{{"star", 1}}),
		options.Find().SetLimit(1),
		options.Find().SetSkip(1),
	)
	if err != nil {
		fmt.Printf("commentCol.Find failed, err:%s \n", err)
	}
	// 注意这里使用了自定义的 string 解析器，具体看 conf.ConnectMongo,否则会报错
	var comment []Comment

	err = commentCursor.All(context.TODO(), &comment)

	if err != nil {
		fmt.Printf("commentCursor.All failed, err:%s \n", err)
	}

	fmt.Printf("comment --> %+v \n", comment)
}

// db.collection.find({content: /周杰伦/})
func findByRegex()  {
	commentCol := MongoDBClient.Database(MongoDBName).Collection(CollectionName)
	// db.comment.find({content: /周杰伦/})


	commentCursor, err := commentCol.Find(context.TODO(),
		bson.D{
		{"content", primitive.Regex{Pattern: "周杰伦"}},
		},
	)
	if err != nil {
		fmt.Printf("commentCol.Find failed, err:%s \n", err)
	}
	// 注意这里使用了自定义的 string 解析器，具体看 conf.ConnectMongo,否则会报错
	var comment []Comment

	err = commentCursor.All(context.TODO(), &comment)

	if err != nil {
		fmt.Printf("commentCursor.All failed, err:%s \n", err)
	}

	fmt.Printf("comment --> %+v \n", comment)
}


// db.collection.find({$or: [{content: /周杰伦/}, {star: {$gte: 15}}})
func findByAndOr()  {
	commentCol := MongoDBClient.Database(MongoDBName).Collection(CollectionName)
	// db.collection.find({$or: [{content: /周杰伦/}, {star: {$gte: 15}}})

	commentCursor, err := commentCol.Find(context.TODO(),
		bson.D{
			{"$or", bson.A{
				bson.D{{"content", primitive.Regex{Pattern: "周杰伦"}}},
				bson.D{{"star", bson.D{{"$gte", 15}}},
			}},
		}},
	)
	if err != nil {
		fmt.Printf("findByAndOr failed, err:%s \n", err)
	}
	// 注意这里使用了自定义的 string 解析器，具体看 conf.ConnectMongo,否则会报错
	var comment []Comment

	err = commentCursor.All(context.TODO(), &comment)

	if err != nil {
		fmt.Printf("findByAndOr failed, err:%s \n", err)
	}

	fmt.Printf("comment --> %+v \n", comment)
}