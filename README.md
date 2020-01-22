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

修改my.cnf:

```shell
docker pull mysql:5.7
Master(主)：
docker run -p 3339:3306 --name mysql-master -e MYSQL_ROOT_PASSWORD=135246 -d mysql:5.7
Slave(从)：
docker run -p 3340:3306 --name mysql-slave -e MYSQL_ROOT_PASSWORD=135246 -d mysql:5.7
# 进入主容器
docker exec -it mysql-master /bin/bash
mysql -uroot -p135246

# mysql user_w，密码是 135246
CREATE USER 'user_w'@'％' IDENTIFIED BY '135246';
# 数据库的所有操作权限都授权给了用户user_w
GRANT ALL PRIVILEGES ON * .* TO 'user_w'@'％';
flush privileges;
容器内部设置解决报错问题
ALTER USER 'user_w'@'%' IDENTIFIED WITH mysql_native_password BY '135246';

# 容器中安装vim
$ apt-get update
$ apt-get install vim
# mysql master容器
$ cd /etc/mysql
# 编辑my.cnf
[mysqld]
## 同一局域网内注意要唯一
server-id=100  
## 开启二进制日志功能，可以随便取（关键）
log-bin=mysql-bin

# 重启master容器
mysql -uroot -p135246
# 创建slave并授权
CREATE USER 'slave'@'%' IDENTIFIED BY '135246';
GRANT REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'slave'@'%';

# 和配置Master(主)一样，在Slave配置文件my.cnf中添加如下配置：
[mysqld]
## 设置server_id,注意要唯一
server-id=101  
## 开启二进制日志功能，以备Slave作为其它Slave的Master时使用
log-bin=mysql-slave-bin   
## relay_log配置中继日志
relay_log=edu-mysql-relay-bin 
```

获取容器IP地址

```shell
docker inspect -f '{{.Name}} - {{.NetworkSettings.IPAddress }}' $(docker ps -aq)
```
设置主从同步:

```shell
在Master进入mysql，执行

show master status;

+------------------+----------+--------------+------------------+-------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
+------------------+----------+--------------+------------------+-------------------+
| mysql-bin.000001 |      155 |              |                  |                   |
+------------------+----------+--------------+------------------+-------------------+
1 row in set (0.00 sec)

在Slave进入mysql，设置其主数据库

change master to master_host='172.17.0.2', master_user='slave', master_password='135246', master_port=3306, master_log_file='mysql-bin.000001', master_log_pos= 155, master_connect_retry=30;

增加从数据库连接帐号

mysql -uroot -p135246
CREATE USER 'user_r'@'％' IDENTIFIED BY '135246';
GRANT SElECT ON * .* TO 'user_r'@'％';
ALTER USER 'user_r'@'%' IDENTIFIED WITH mysql_native_password BY '135246';

在Slave 中的mysql终端执行show slave status \G;用于查看主从同步状态。
当
Slave_IO_Running: Yes
Slave_SQL_Running: Yes
都为Yes时，表示正常同步了

show slave status G;

*************************** 1. row ***************************
               Slave_IO_State: Waiting for master to send event
                  Master_Host: 172.17.0.2
                  Master_User: slave
                  Master_Port: 3306
                Connect_Retry: 30
              Master_Log_File: mysql-bin.000001
          Read_Master_Log_Pos: 23093
               Relay_Log_File: edu-mysql-relay-bin.000002
                Relay_Log_Pos: 322
        Relay_Master_Log_File: mysql-bin.000001
             Slave_IO_Running: Yes
            Slave_SQL_Running: Yes
              Replicate_Do_DB: 
          Replicate_Ignore_DB: 
           Replicate_Do_Table: 
       Replicate_Ignore_Table: 
      Replicate_Wild_Do_Table: 
  Replicate_Wild_Ignore_Table: 
                   Last_Errno: 0
                   Last_Error: 
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 23093
              Relay_Log_Space: 534
              Until_Condition: None
               Until_Log_File: 
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File: 
           Master_SSL_CA_Path: 
              Master_SSL_Cert: 
            Master_SSL_Cipher: 
               Master_SSL_Key: 
        Seconds_Behind_Master: 0
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error: 
               Last_SQL_Errno: 0
               Last_SQL_Error: 
  Replicate_Ignore_Server_Ids: 
             Master_Server_Id: 100
                  Master_UUID: 12b2e6dd-bf46-11e9-ba82-0242ac110002
             Master_Info_File: mysql.slave_master_info
                    SQL_Delay: 0
          SQL_Remaining_Delay: NULL
      Slave_SQL_Running_State: Slave has read all relay log; waiting for more updates
           Master_Retry_Count: 86400
                  Master_Bind: 
      Last_IO_Error_Timestamp: 
     Last_SQL_Error_Timestamp: 
               Master_SSL_Crl: 
           Master_SSL_Crlpath: 
           Retrieved_Gtid_Set: 
            Executed_Gtid_Set: 
                Auto_Position: 0
         Replicate_Rewrite_DB: 
                 Channel_Name: 
           Master_TLS_Version: 
       Master_public_key_path: 
        Get_master_public_key: 0
            Network_Namespace: 
1 row in set (0.00 sec)

ERROR: 
No query specified


从库停止主从复制

STOP SLAVE IO_THREAD FOR CHANNEL ''

```

