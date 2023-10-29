# 简介
一个将MySQL服务转化为http接口的程序，有时候并不总是会有一个mysql客户端给你用来查询操作暑假，当然前端同学也无法直接通过JS去连接数据库，所以，这个时候如果可以快速构建起一套`增/删/改/查`MySQL的HTTP服务，是不是就很完美了呢，这个程序就是专门干这个的，一个可执行文件生成另外一个可执行文件。

# 如何运行
## 1. 编写config.json

```json
{
    "debug": true,
    "database" : [
        {
            "name": "mysql2http",
            "dsn" : "root:1234567@tcp(127.0.0.1:3306)/mysql2http?charset=utf8&parseTime=True&loc=Local"
        },
    ],
    "port": 8090,
    "code_folder" : "./tmp_code_folder",
    "target" : "./database.exe"
}
```

## 2. 执行程序

```shell
./mysql2http ./config.json
```

## 3. 使用该服务

```shell
./database.exe
```

- query查询

```shell
curl --location --globoff 'localhost:8090/{database}/{table}/query' \
--header 'Content-Type: application/json' \
--data '{
    "id": "6743743",
    "_setting": {
        "where" : {
            "order_by" : "id desc",
            "limit" : 100,
            "fields" : ["id", "name"]
        }
    }
}'
```
- create数据

```shell
curl --location --globoff 'localhost:8090/{database}/{table}/cretate' \
--header 'Content-Type: application/json' \
--data '{
    "id": "6743743",
    "name" : "app"
}'
```

- modify数据

```shell
curl --location --globoff 'localhost:8090/{database}/{table}/modify' \
--header 'Content-Type: application/json' \
--data '{
    "id": "6743743",
    "name" : "app",
    "_setting" : {
        "where" : {
            "id" : "46637543"
        }
    }
}'
```
- delete数据

```shell
curl --location --globoff 'localhost:8090/{database}/{table}/delete' \
--header 'Content-Type: application/json' \
--data '{
    "id": "6743743"
}'
```

> 备注：看出来了吧，我们尽可能地用最直观易懂得方式去传递参数




