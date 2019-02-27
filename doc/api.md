服务对应的地址及端口:

* ip以  `192.169.0.`开头
* etcd服务中, 2379 为服务接口, 2380为raft接口.

| service       | host       |ip    | port  | 
| ------------- |:-------------:| -----:|----:|
| registry| registry   | 9   |8080 | 
| mail    | mail       |   10|7070 |
| balance | balance    |    8|5050 |
| userlist | userlist    |    11|6060 |
| mysql | mysql    |    18|3306 |
|etcd1|etcd1| 2| 2379 && 3380|
|etcd2|etcd2| 3| 2379 && 3380|
|etcd3|etcd3| 4| 2379 && 3380|


### *`registry` , `mail`, ` balance`, `userlist` 接口都是post请求, 格式为json.*

# registry 接口

请求参数:

| name       | type          | must  | example
| ------------- |:-------------:| -----:| ---:|
| name     | string| 1 | service/mail
| host      | string      |   1 | mail
| port| int     |    1 |    7070

返回参数:

|httpCode| explain|
|---:|----:|
|200|注册服务成功|
|410| 注册失败|

# mail 接口

请求参数:

| name       | type          | must  | example
| ------------- |:-------------:| -----:| ---:|
| address     | string| 1 | papandadj@gmail.com
| message      | string      |   1 | Php是最好的语言|

返回参数:

|httpCode| explain|
|---:|----:|
|200|发送邮件成功|
|411| 发送失败|

# userlist 接口

请求接口:

* method 为put: name, email 必须
* method 为get: name 必须
* method 为del: name 必须

| name       | type          | must  | example
| ------------- |:-------------:| -----:| ---:|
| method     | string| 1 | put|
| name      | string      |   1 | dong|
| email      | string      |   0 | papandadj@gmail.com|

返回参数:

|httpCode| explain|
|---:|----:|
|200|发送邮件成功|
|410| 发送失败|




# balance 接口

---
* url: mail
* 同mail服务
---
* url:userlist
* 同userlist服务
---