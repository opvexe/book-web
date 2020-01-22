#### 1.1 iBook项目

创建项目:

```shell
$ bee new ibook
```

#### 1.2 数据库表设计

```shell
不同数据放在不同表
一对多：把对应关系放在”多“表中
多对多：把对应关系放在单独表中
相同数据的对应表：
一对多：形成自联表 [一级分类中有二级分类]
多对多：对应关系单独表 [一个用户多个粉丝，一个粉丝有多个用户]
```

#### 1.3 elasticsearch安装与使用

```shell
$ brew install elasticsearch@5.6
$ brew services start elasticsearch@5.6
# 访问http://localhost:9200
```

```go
//put请求
func (this *ElasticSearchClient) put(api string) (req *httplib.BeegoHTTPRequest) {
	return httplib.Put(api).Header("Content-Type", "application/json").SetTimeout(this.Timeout, this.Timeout)
}

//post请求
func (this *ElasticSearchClient) post(api string) (req *httplib.BeegoHTTPRequest) {
	return httplib.Post(api).Header("Content-Type", "application/json").SetTimeout(this.Timeout, this.Timeout)
}

//delete请求
func (this *ElasticSearchClient) delete(api string) (req *httplib.BeegoHTTPRequest) {
	return httplib.Delete(api).Header("Content-Type", "application/json").SetTimeout(this.Timeout, this.Timeout)
}

//get请求
func (this *ElasticSearchClient) get(api string) (req *httplib.BeegoHTTPRequest) {
	return httplib.Get(api).Header("Content-Type", "application/json").SetTimeout(this.Timeout, this.Timeout)
}

```

#### 1.4 elasticsearch基本使用

添加:

```json
请求地址：http://localhost:9200/testbookindex/books/testbookindex/book/1
请求头：Content-Type:application/json;charset=UTF-8
请求方式：PUT
请求体:
{
 "book_id":3,
 "book_name":"Java",
 "description":"一门垃圾语言"
}
```

返回结果:

```json
{
  "_index": "testbookindex",
  "_type": "books",
  "_id": "3",
  "_version": 1,
  "result": "created",
  "_shards": {
    "total": 2,
    "successful": 1,
    "failed": 0
  },
  "created": true
}
```

查询:

```json
请求地址：http://localhost:9200/testbookindex/books/_search
请求头：Content-Type:application/json;charset=UTF-8
请求方式：POST
请求体:
{
 "query":{
 "match":{"description":"Go"} 
	}
}
```

返回结果:

```json
{
  "took": 1,
  "timed_out": false,
  "_shards": {
    "total": 5,
    "successful": 5,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": 1,
    "max_score": 0.38080662,
    "hits": [{
      "_index": "testbookindex",
      "_type": "books",
      "_id": "1",
      "_score": 0.38080662,
      "_source": {
        "book_id": 1,
        "book_name": "Go语言课本",
        "description": "Go性能并发，支持Go程"
      }
    }]
  }
}
```

分类查询:

```json
请求地址：http://localhost:9200/testbookindex/books/_search
请求头：Content-Type:application/json;charset=UTF-8
请求方式：POST
请求体:
{
"query":{
"bool":{
"should":{
{"match":{"book_name":"Go"}},
{"match":{"description":"ios"}}
}
}
}
}
```

#### 1.4 mysql 主从搭配

```shell

```

