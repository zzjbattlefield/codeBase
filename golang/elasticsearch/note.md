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
  1. ### match查询(匹配查询/模糊查询)
      `match`：模糊匹配，需要指定字段名，但是输入会进行分词，比如"hello world"会进行拆分为hello和world，然后匹配，如果字段中包含hello或者world，或者都包含的结果都会被查询出来，也就是说match是一个部分匹配的模糊查询。查询条件相对来说比较宽松。
      ```json
        GET user/_search
        {
          "query":{
            "match":{"查询字段":"查询的值"}
          }
        }
      ```
  2. ### match_phrase(短语查询)
      `match_phrase`：会对输入做分词，但是需要结果中也包含所有的分词，而且顺序要求一样。以"hello world"为例，要求结果中必须包含hello和world，而且还要求他们是连着的，顺序也是固定的，hello that word不满足，world hello也不满足条件。
      ```json
        GET user/_search
        {
          "query":{
            "match_phrase":{
              "字段":"值"
            }
          }
        }
      ```

  3. ### multi_match
      `multi_match`：查询提供了一个简便的方法用来对多个字段执行相同的查询，即对指定的多个字段进行match查询
      ```json
      GET user/_search
      {
        "query":{
          "multi_match":{
            "query":"要查询的值",
            "fields":["title","desc"]
          }
        }
      }
      ```
  4. ### query_string
      `query_string`：和match类似，但是match需要指定字段名，query_string是在所有字段中搜索，范围更广泛。
      ```json
      GET user/_search
      {
        "query":{
          "query_string":{
            "default_field":"查询字段",//默认查询字段可以不加 不加为全字段搜索
            "query":"Madison AND street"//可以搭配AND 或者 OR 使用
          }
        }
      }
      ```
+ ## term级别的查询
  1. ### term查询(不分词 直接拿输入的值查询)
      `term`:  这种查询和match在有些时候是等价的，比如我们查询单个的词hello，那么会和match查询结果一样，但是如果查询"hello world"，结果就相差很大，因为这个输入不会进行分词，就是说查询的时候，是查询字段分词结果中是否有"hello world"的字样，而不是查询字段中包含"hello world"的字样，elasticsearch会对字段内容进行分词，“hello world"会被分成hello和world，不存在"hello world”，因此这里的查询结果会为空。这也是term查询和match的区别。
      ```json
      GET user/_search
      {
        "query":{
          "term":{
            "address":"需要搜索的值"
          }
        }
      }
      ```
  2. ### 范围range查询
      支持的范围格式`gt`, `gte`, `lt`, `lte` 
      ```json
      GET user/_search
      {
        "query":{
          "range":{
            "需要搜索的字段":{
              "gte":"30",
              "lte":"40"
            }
          }
        }
      }
      ```
  3. ### exists查询
      `exists`:查询index中存在此field的记录
      ```json
      GET user/_search{
        "query":{
          "exists":{
            "需要搜索的字段":"需要搜索的值"
          }
        }
      }
      ```
  4. ### fuzzy模糊查询
    使用普通的match也可以开启模糊查询
    ```json
    GET user/_seach
    {
      "query":{
        "match":{
          "搜索的字段":{
            "query":"需要搜索的值",
            "fuzziness":1,
          }
        }
      }
    }
    ```
    普通的fuzzy查询:
    ```json
    GET user/_search
    {
      "query":{
        "fuzzy":{
          "搜索的字段":{
            "value":"需要搜寻的值"
          }
        }
      }
    }
    ```
+ ## 组合bool查询
  组合查询的格式:
  ```json
    {
      "query":{
        "bool":{
          "must":[],//计分
          "shout":[],//计分
          "must_not":[],//不计分
          "filter":[],//不计分
        }
      }
    }
  ```
  示例:
  ```json
  GET user/_search
    {
      "query": {
        "bool": {
          "must": [
            {
              "term": {
                "state": "tn"
              }
            },
            {
              "range": {
                "age": {
                  "gte": 20,
                  "lte": 30
                }
              }
            }
          ],
          "must_not": [
            {
              "term": {
                "gender": "m"
              }
            }
          ],
          "should": [
            {
              "match": {
                "firstname": "Decker"
              }
            }
          ],
          "filter": [
            {
              "range": {
                "age": {
                  "gte": 25,
                  "lte": 30
                }
              }
            }
          ]
        }
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
