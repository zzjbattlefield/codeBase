<font size="4">

# es中的index type mapping
mysql|es
--|:--:|
database
table | index(7之后type固定为_doc)
row | document
column | field
schema | mapping
sql | DSL

# 索引(index)
索引有两个含义一个是动词(插入)名词(表)

# 如何新建数据
+ ## 通过put+id新建数据
```
PUT /[index名称]/_doc/id
```
使用put必须要携带id(如果id不存在则会新建 存在就会更新)
```json
PUT /testmd/_doc/1
{
  "name": "go_es",
  "age": "18",
  "company": [
    {
      "name": "google"
    },
    {
      "name": "baidu"
    }
  ]
}
//返回
{
  "_index": "testmd",
  "_type": "_doc",
  "_id": "1",
  "_version": 1,
  "result": "created",
  "_shards": {
    "total": 2,
    "successful": 1,
    "failed": 0
  },
  "_seq_no": 0,
  "_primary_term": 1
}
```
+ ## 发送POST新增数据
```json
POST user/_doc/
{
  "name":"go",
  "company":"google"
}
POST+id = PUT方法
```

+ ## POST+_creat 没有就创建 已存在就报错
```json
POST user/_create/1
{
  "name":"go",
  "company":"google"
}
//返回
[1]: version conflict, document already exists (current version [1])
```

# 获取索引数据
+ ## 获取具体的数据
```json
  GET [indexname]/_doc/id
  GET user/_doc/1
```
+ ## 通过ReqeustBody进行查询
   ### 查询全部
  ```json
    GET user/_search
    {
      "query":{
        "match_all":{}
      }
    }
  ```
+ ## 全文查询(分词)
  1. ### match查询(匹配查询)
      `match`：模糊匹配，需要指定字段名，但是输入会进行分词，比如"hello world"会进行拆分为hello和world，然后匹配，如果字段中包含hello或者world，或者都包含的结果都会被查询出来，也就是说match是一个部分匹配的模糊查询。查询条件相对来说比较宽松。
      ```json
        GET user/_search
        {
          "query":{
            "match":{"查询字段":"查询的值"}
          }
        }
      ```

# 更新数据
## 使用POST(PUT) [index]/_doc/id{}的方式更新数据会覆盖原来的数据
+ ## 更新数据的方法
```json
POST [index]/_update/id
{
  "doc":{"updateField":"value"}
}
//如果更新的值和原本的值一样 则不会执行语句
```

# 删除数据
+ ## 删除指定id的数据
```json
  DELETE [index]/_doc/id{}
```

+ ## 删除整个索引
```json
  DELETE index
```

# 批量操作(_bulk)
+ ## 批量增删改
```json
POST _bulk
{ "index" : { "_index" : "test", "_id" : "1" } }
{ "field1" : "value1" }
{ "delete" : { "_index" : "test", "_id" : "2" } }
{ "create" : { "_index" : "test", "_id" : "3" } }
{ "field1" : "value3" }
{ "update" : {"_id" : "1", "_index" : "test"} }
{ "doc" : {"field2" : "value2"} }

//新增 插入 更新操作有两行 删除操作有一行
{ "行为" : { "索引名称" : "test", "_id" : "1" } }
{ "field1" : "value1" }
or
{"doc":{ "field1" : "value1" }}
```

+ ## 批量查询(_mget)
```json
GET /_mget
{
  "docs": [
    {
      "_index": "user",
      "_id": "1"
    },
    {
      "_index": "account",
      "_id": "2"
    }
  ]
}
```
