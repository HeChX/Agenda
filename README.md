# Agend
## Agenda是使用Go语言开发的一个简单的会议管理系统          
### 文件结构
+ cmd文件夹内存储了各个子指令的处理文件       
+ entity文件内存储了用户、会议、日期三个基础对象的定义类，Fileio.go是对于用户和会议数据的文件读写以提供数据持久化支持，使用了json的序列和反序列化，Storage.go是对y用户和会议数据的增删查改    
+ service文件夹内的文件service.go是各命令的逻辑操作，agenda.log为项目的活动日志
+ User.txt和Meeting.txt为用户信息和会议信息的存储文件

### 业务需求及命令列表
+ 用户：每个用户的账号名称具有唯一性，且用户在同一时间只能参加一场会议
+ 会议：每场会议的题目具有唯一性

|需求|命令|描述|
|----|-----|----|
|用户注册 |register|用户通过设置一个唯一的用户名和密码以及邮箱和电话注册账号|
|用户登录 |login|用户提供正确的用户名和密码以登录系统|
|用户登出 |logout|登录用户登出当前账号|
|注销账号 |logoff|已登录用户可以注销删除自己的账号|
|查询指定用户|queryuser|已登录用户可以通过一个指定的用户名查询该用户的账号邮箱和手机号|
|查看用户列表|queryalluser|已登录用户可以产看已注册的所有用户的账号邮箱和手机号|
|创建会议|createmeeting|已登录用户以会议发起者身份创建一个会议|
|取消会议|cancelmeeting|已登录用户可以取消一场自己参与的会议|
|增加会议成员|addparticipator|已登录用户可在自己发起的会议中添加一名复合要求的成员|
|删除会议成员|deleteparticipator|已登录用户可在自己发起的会议中删除一个参与者|
|清空会议|clearmeeting|已登录的用户可以清空自己发起的所有会议安排|
|退出会议|exitmeeting|已登录的用户可以退出自己参与的某一会议安排|
|查询会议|querymeeting|已登录用户可以查询某一时间段内的所有会议|

### 命令使用
#### agenda
```
agenda -h
```
可以查看所有该系统的介绍和所有子指令
#### 用户注册 register
```
agenda register -n [name] -p [password] -e [email] -t [phone]
```
+ `-u`参数为用户名，具有唯一性
+ `-p`参数为账号密码
+ `-e`参数为账号邮箱
+ `-t`参数为账号手机号码
#### 用户登录 login
```
agenda login -n [name] -p [password]
```
+ `-u`参数为用户名，具有唯一性
+ `-p`参数为账号密码                  
需要提供用户名和正确的密码
#### 用户登出 logout
```
agenda logout
```
无参数，登出当前账号
#### 注销账号 logoff
```
agenda logoff
```
无参数，注销当前账号
#### 查询指定用户 queryuser
```
agenda queryuser -n [name] 
```
+ `-n`参数为要查询的用户的用户名
#### 查询所有用户 queryalluser
```
agenda queryuser -a 
```
可以列出已注册的所有账号
#### 创建会议 createmeeting
```
agenda createmeeting -t [title] -s [starttime] -e [endtime] -p [participator] 
```
+ `-t`参数为会议标题，具有唯一性
+ `-s`参数为会议开始时间
+ `-e`参数为会议结束时间
+ `-p`参数为参加会议的成员
#### 取消会议 cancelmeeting
```
agenda cancelmeeting -t [title]
```
+ `-t`参数为会议标题，具有唯一性
#### 添加会议成员 addparticipator
```
agenda addpaticipator -n [name] -t [title]
```
+ `-n`参数为成员用户名，具有唯一性
+ `-t`参数为会议标题，具有唯一性
#### 删除会议成员 deleteparticipator
```
agenda deleteparticipator -n [name] -t [meeting title]
```
+ `-n`参数为成员用户名，具有唯一性
+ `-t`参数为会议标题
#### 清空会议 clearmeeting
```
agenda clearmeeting 
```
已登录用户清空所有与自己有关的会议，无参数
#### 退出会议 exitmeeting
```
agenda exitmeeting -t [title] 
```
+ `-t`参数为会议标题，具有唯一性
#### 查询会议 querymeeting
```
agenda querymeeting -s [start time] -e [end time]
```
+ `-s`参数为查询时间段的开始时间
+ `-t`参数为查询时间段的结束时间
