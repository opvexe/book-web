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

#### 1.4 mysql 主从搭配

```shell

```

