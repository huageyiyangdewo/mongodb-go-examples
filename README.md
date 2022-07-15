# mongodb-go-examples
mongodb go examples

启动步骤，首先进入 mongodb-go-examples

```
docker-compose up -d
```

需要根据实际情况，修改docker-compose 里面的信息，主要是
修改
```
ME_CONFIG_MONGODB_URL: mongodb://root:123456@192.168.170.66:27017/

把 192.168.170.66 改成自己的。
```

[官方例子](https://github.com/mongodb/mongo-go-driver/blob/master/examples/documentation_examples/examples.go)

其实看官方的例子就够了，又多又清晰

bson.D{}
```
bson.D{} 是一个有序的文档对象，当插入的文档对顺序由要求时，
应该使用这个类型。假如对文档顺序没有要求，则使用 bson.M{}
```

bson.M{}
``` 
对文档顺序没有要求，则使用 bson.M{}
```

bson.A{}
``` 
A是BSON阵列的有序表示。
```

bson.E{}
``` 
通常使用在 bson.D{} 里面使用
```


``` 
error decoding key comment_id: cannot decode 32-bit integer into a string type

https://cloud.tencent.com/developer/ask/sof/1465278

```