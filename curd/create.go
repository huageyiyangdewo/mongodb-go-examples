package curd

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func Create()  {
	//createCollection()

	collectionInsert()

	//collectionInertMany()
}


// db.createCollection(name)
func createCollection()  {
	// 显式创建 数据库和 集合
	err := MongoDBClient.Database(MongoDBName).CreateCollection(context.TODO(), CollectionName)
	if err != nil {
		fmt.Printf("createCollection failed, err:%s \n", err)
	}
}

// db.collection.insert()
func collectionInsert()  {
	commentColl := MongoDBClient.Database(MongoDBName).Collection(CollectionName)

	document := Comment{
		CommentId: "1tg",
		Author: "张三",
		Content: "周杰伦6首新歌5首由方文山作词",
		Star: 10,
		CreateAt: time.Now(),
	}

	// 插入一条数据,如何集合不存在，会隐式的创建集合
	result, err := commentColl.InsertOne(context.TODO(), document)
	if err != nil {
		fmt.Printf("collectionInsert failed, err:%s \n", err)
	}
	fmt.Printf("collectionInsert --> result:%+v \n", result)
	// collectionInsert --> result:&{InsertedID:ObjectID("62d0e5a42c4d4cecae4f92ad")}


	result1, err := commentColl.InsertOne(context.TODO(), bson.D{
		{"comment_id", 2},
		{"author", 2},
		{"content", "上海二季度GDP同比下降13.7%新"},
		{"star", 2},
		{"create_at", time.Now()},
	})
	if err != nil {
		fmt.Printf("collectionInsert failed, err:%s \n", err)
	}
	fmt.Printf("collectionInsert --> result1:%+v \n", result1)
	// collectionInsert --> result1:&{InsertedID:ObjectID("62d0e902d66e799268036ee7")}

}


// db.collection.insertMany()
func collectionInertMany()  {
	commentColl := MongoDBClient.Database(MongoDBName).Collection(CollectionName)

	document1 := Comment{
		CommentId: "11",
		Author: "张三1",
		Content: "分析师称中国黄牛都看好iPhone14",
		Star: 11,
		CreateAt: time.Now(),
	}
	document2 := Comment{
		CommentId: "12",
		Author: "李四",
		Content: "活虾从菜场拎到家被“热熟”",
		Star: 11,
		CreateAt: time.Now(),
	}
	document := []interface{}{document1,document2}

	// 插入多条数据,如何集合不存在，会隐式的创建集合
	result, err := commentColl.InsertMany(context.TODO(), document)
	if err != nil {
		fmt.Printf("collectionInertMany failed, err:%s \n", err)
	}
	fmt.Printf("collectionInertMany --> result:%+v \n", result)
	// collectionInertMany --> result:&{InsertedIDs:[ObjectID("62d0ec9beda9c3e20a745f1d") ObjectID("62d0ec9beda9c3e20a745f1e")]}


	doc := []interface{}{
		bson.M{
			"comment_id": 13,
			"author": "李四",
			"content": "定格山东舰霸气瞬间热",
			"star": 109,
			"create_at": time.Now(),
		},
		bson.M{
			"comment_id": 13,
			"author": "李四",
			"content": "广西新增本土无症状感染者165例热",
			"star": 34,
			"create_at": time.Now(),
		},
	}
	result1, err := commentColl.InsertMany(context.TODO(), doc)
	if err != nil {
		fmt.Printf("collectionInertMany failed, err:%s \n", err)
	}
	fmt.Printf("collectionInertMany --> result1:%+v \n", result1)
	// collectionInertMany --> result1:&{InsertedIDs:[ObjectID("62d0ec9beda9c3e20a745f1f") ObjectID("62d0ec9beda9c3e20a745f20")]}
}