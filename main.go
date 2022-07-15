package main

import (
	"context"
	"mongodb-go-examples/curd"
)



func main()  {
	curd.ConnectMongo()
	defer func() {
		_ = curd.MongoDBClient.Disconnect(context.TODO())
	}()
	//curd.Create()

	curd.Delete()

	//curd.Query()

	//curd.Update()
}

