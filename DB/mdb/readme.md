# Mongodb

Mongodb是NoSQL数据库的一种，而什么是mongdb数据库呢？简单来说是一种非关系型数据库，是对不同于传统数据库的其它数据库的统称（Not only SQL）。

## 基础构成
数据库（database）——>集合（collection）——>文档（document）

BSON格式键值对存储，常用数据类型：

- String 字符串 utf-8格式
- Integer 整型数值 根据系统可分为32位和64位
- Boolean 布尔值
- Double 双精度浮点值
- Null 空值
- Date 日期格式 例如  ISODate("2018-03-04T14:58:51.233Z")
- ObjectID 文档ID 类似唯一主键

## 优点

- 数据存储不需要固定的格式，非常容易就可以进行横向拓展
- 最终一致性，而非ACID属性

## 常用语法

#### 创建/删除数据库，mongodb中集合只有在内容插入后才会创建
```javascript
use YourDBName
db.dropDatabase()  //删除当前所在数据库
```

#### 创建/删除集合
```javascript
db.createCollection("YourCollName")
//创建固定集合 mycol，整个集合空间大小 6142800 KB, 文档最大个数为 10000 个。
db.createCollection("mycol", { capped : true, autoIndexId : true, size : 6142800, max : 10000 } )
//插入文档的时候会自动创建
db.YourColl.insert({"testdata":true})
//删除
db.getCollection("YourColl").drop()
```
#### 插入文档、更新文档、删除文档
```javascript
document=({
title: 'MongoDB 教程', 
description: 'MongoDB 是一个 Nosql 数据库',
by: '菜鸟教程',
url: 'http://www.runoob.com',
tags: ['mongodb', 'database', 'NoSQL'],
likes: 100
})
db.getCollection("YourColl").insert(document)
```
#### 更新文档主要使用save和update,update修改原有文档，save通过传入的文档来替换已有的文档
```javascript
/*语法格式
db.getCollection("YourColl").update(
       <query>,
       <update>,
       {
         upsert: <boolean>,
         multi: <boolean>,
         writeConcern: <document>
       }
)
*/
//参数说明
/*
query:update查询条件
update:update的对象和一些更新的操作符（如$,$inc...）等
upsert : 可选，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入
multi : 可选，mongodb 默认是false,只更新找到的第一条记录，如果为true,就把按条件查出来多条记录全部更新。
writeConcern :可选，抛出异常的级别。
*/

db.getCollection("YourColl").update({'title':'MongoDB 教程'},{$set:{'title':'MongoDB'}})
//替换了 _id 为 56064f89ade2f21f36b03136 的文档数据
db.getCollection("YourColl").save({
    "_id" : ObjectId("56064f89ade2f21f36b03136"),
    "title" : "MongoDB",
    "description" : "MongoDB 是一个 Nosql 数据库",
    "by" : "Runoob",
    "url" : "http://www.runoob.com",
    "tags" : [
        "mongodb",
        "NoSQL"
    ],
    "likes" : 110
})
```

#### 删除文档
```javascript
/*
语法格式
db.getCollection("YourColl").remove(
   <query>,
   {
     justOne: <boolean>,
     writeConcern: <document>
   }
)
参数说明
query :（可选）删除的文档的条件。
justOne : （可选）如果设为 true 或 1，则只删除一个文档，如果使用默认值 false，则删除所有匹配条件的文档。
writeConcern :（可选）抛出异常的级别。
*/
db.getCollection("YourColl").remove({'title':'MongoDB 教程'})
db.getCollection("YourColl").remove({})//删除所有数据
```
#### 查询文档
##### 普通查询
```javascript
db.getCollection("YourColl").find().pretty();
db.getCollection("YourColl").find({key1:value1,key2:value2});
db.getCollection("YourColl").find({"money":{$gte:50,$lte:20}});//ne !=  eq =
db.getCollection("YourColl").find({$or:[{key1:value1},{key1:value2}]});
db.getCollection("YourColl").find({
    "likes": {$gt:50},
    $or: [{"by": "菜鸟教程"},{"title": "MongoDB 教程"}]
});
db.getCollection("YourColl").find({"name":{$type:'string'}})
db.getCollection("YourColl").find().limit(number).skip(number)
db.getCollection("YourColl").find().sort({key:1}) //-1降序
db.getCollection("YourColl").find({"name":{$regex:"S1.ll"}})
```
### 聚合查询
&emsp;&emsp;db.collection.aggregate()是基于数据处理的聚合管道，每个文档通过一个由多个阶段（stage）组成的管道，可
以对每个阶段的管道进行分组、过滤等功能，然后经过一系列的处理，输出相应的结果。

1、db.collection.aggregate() 可以用多个构件创建一个管道，对于一连串的文档进行处理。这些构件包括：
筛选操作的match、映射操作的project、分组操作的group、排序操作的sort、限制操作的limit、和跳过操作的
skip。

2、db.collection.aggregate()使用了MongoDB内置的原生操作，聚合效率非常高,支持类似于SQL Group 
By操作的功能，而不再需要用户编写自定义的JavaScript例程。

3、 每个阶段管道限制为100MB的内存。如果一个节点管道超过这个极限,MongoDB将产生一个错误。为了能够
在处理大型数据集,可以设置allowDiskUse为true来在聚合管道节点把数据写入临时文件。这样就可以解决1
00MB的内存的限制。

4、db.collection.aggregate()可以作用在分片集合，但结果不能输在分片集合，MapReduce可以 作用在分
片集合，结果也可以输在分片集合。

5、db.collection.aggregate()方法可以返回一个指针（cursor），数据放在内存中，直接操作。跟Mong
o shell 一样指针操作。

6、db.collection.aggregate()输出的结果只能保存在一个文档中，BSON Document大小限制为16M。
可以通过返回指针解决，版本2.6中后面：DB.collect.aggregate()方法返回一个指针，可以返回任何结果集的大小。

###### 注意

&emsp;&emsp;2.6版中的新增功能：仅当将管道指定为数组时才可用。
使用db.collection.aggregate()直接查询会提示错误，但是传一个空数组如db.collection.aggregate([])则
不会报错，且会和find一样返回所有文档。

```javascript
//group 分组
db.getCollection("YourColl").aggregate([{$group:{_id:"$key1",total:{$sum:1}}}])
db.getCollection("YourColl").aggregate([{$group:{_id:"$key1",total:{$sum:"$money"}}}])
db.getCollection("YourColl").aggregate([{$group:{_id:null,avergeMoney:{$avg:"money"}}}])
//$avg $min $max $push  $addToSe $first	$last
//match 过滤数据
db.getCollection("YourColl").aggregate([{$match:{"name":"564"}}])

//project 修改输入的文档
db.getCollection("YourColl").aggregate([{$project:{title:1,name:1}}])

//skip 跳过指定的文档
db.getCollection("YourColl").aggregate([{$skip:5}])

//limit 限制聚合管道返回的文档数量
db.getCollection("YourColl").aggregate([{$limit:5}])



```

##### 特殊操作
```javascript
/*
Group大约需要一下几个参数。

 1.key：用来分组文档的字段。和keyf两者必须有一个
 2.keyf：可以接受一个javascript函数。用来动态的确定分组文档的字段。和key两者必须有一个
 3.initial：reduce中使用变量的初始化
 4.reduce：执行的reduce函数。函数需要返回值。
 5.cond：执行过滤的条件。
 6.finallize：在reduce执行完成，结果集返回之前对结果集最终执行的函数。可选的
 */
 db.getCollection('recharge').group({key:{userid:true},initial:{sum:0,accid:0,state:0},$reduce:function(doc, result){
 		result.sum += doc.cash
 		result.accid = doc.accid
 	},condition:{cash:{$gt:0},state:{$gt:0}}
 })
 
```

## Links
[浅析mongodb中group分组](https://www.jb51.net/article/65934.htm)


