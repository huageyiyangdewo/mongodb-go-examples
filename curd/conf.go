package curd

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDBClient *mongo.Client

	MongoDBAddr = "192.168.170.66"
	MongoDBPort = "27017"
	MongoDBUser = "root"
	MongoDBPwd = "123456"
	MongoDBName = "test"
	CollectionName = "comment"
)

type Comment struct {
	CommentId string `bson:"comment_id"` // 评论id
	Author string  `bson:"author"` // 评论者
	Content string `bson:"content"` // 评论内容
	Star int  `bson:"star"`  // 点赞数
	CreateAt time.Time `bson:"create_at"` // 创建时间
}

// stringDecoderValue 为 string 类型编写自定义解析器, 在使用 find.ALL() 方法时，能够将
// 数据库中的 BSON boolean, int32, int64, double, or null 转换为 字符串。
// BSON boolean, int32, int64, double, or null 为 零值 时，转换为 "" (空字符串)
func stringDecoderValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {

	// All decoder implementations should check that val is valid, settable,
	// and is of the correct kind before proceeding.
	if !val.IsValid() || !val.CanSet() || val.Kind() != reflect.String {
		return bsoncodec.ValueDecoderError{
			Name:     "lenientStringDecodeValue",
			Kinds:    []reflect.Kind{reflect.String},
			Received: val,
		}

	}

	var result string
	switch vr.Type() {
	case bsontype.Boolean:
		b, err := vr.ReadBoolean()
		if err != nil {
			return err
		}
		// b -> true -> "true"
		if b {
			result = "true"
		} else {
			result = ""
		}
	case bsontype.Int32:
		i32, err := vr.ReadInt32()
		if err != nil {
			return err
		}
		if i32 != 0 {
			result = fmt.Sprintf("%d", i32)
		} else {
			result = "0"
		}
	case bsontype.Int64:
		i64, err := vr.ReadInt64()
		if err != nil {
			return err
		}
		if i64 != 0 {
			result = fmt.Sprintf("%d", i64)
		} else {
			result = "0"
		}
	case bsontype.Double:
		f64, err := vr.ReadDouble()
		if err != nil {
			return err
		}
		if f64 != 0.0 {
			result = fmt.Sprintf("%f", f64)
		} else {
			result = "0.0"
		}
	case bsontype.Null:
		if err := vr.ReadNull(); err != nil {
			return err
		}
		result = ""
	case bsontype.String:
		s, err := vr.ReadString()
		if err != nil {
			return err
		}
		result = s
	default:
		return fmt.Errorf(
			"received invalid BSON type to decode into lenientString: %s",
			vr.Type())
	}

	val.SetString(result)
	return nil

}

func ConnectMongo()  {

	decoder := bsoncodec.ValueDecoderFunc(stringDecoderValue)

	registry := bson.NewRegistryBuilder().
		RegisterDefaultDecoder(reflect.String, decoder).Build()

	path := fmt.Sprintf("mongodb://%s:%s@%s:%s", MongoDBUser, MongoDBPwd, MongoDBAddr, MongoDBPort)
	clientOptions := options.Client().SetRegistry(registry).ApplyURI(path)

	var err error
	MongoDBClient, err = mongo.Connect(context.TODO(), clientOptions)
	// 这样赋值有问题，没明白？
	//MongoDBClient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Printf("mongo connect failed, err:%s \n", err)
		panic(err)
	}

	err = MongoDBClient.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Printf("mongo ping failed, err:%s \n", err)
		panic(err)
	}


	fmt.Println("mongo connect success")
}
