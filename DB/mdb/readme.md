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

//unwind 拆分文档中的某个数组类型，每条包含数组中的一个值
db.getCollection("YourColl").aggregate([{$unwind:{path:"$sizes"}}])

//sort 将输入文档排序后输出
db.getCollection("YourColl").aggregate([{$sort:{age:-1,posts:1}}])

//out 将处理后的文档输入到某个集合
db.getCollection("YourColl").aggregate([{$group: {_id:"$author",books:{$push: "$title"}}},{ $out : "authors" }])
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

## 索引
&emsp;&emsp;索引通常能够极大的提高查询的效率，如果没有索引，MongoDB在读取数据时必须扫描集合中的每个文件并选取那些符合查询条件的记录。
这种扫描全集合的查询效率是非常低的，特别在处理大量的数据时，查询可以要花费几十秒甚至几分钟，这对网站的性能是非常致命的。
索引是特殊的数据结构，索引存储在一个易于遍历读取的数据集合中，索引是对数据库表中一列或多列的值进行排序的一种结构
```javascript
//语法
db.getCollection('YourColl').createIndex(keys, options)
//1升序 -1降序 复合索引
db.getCollection('YourColl').createIndex({"title":1,"description":-1})

/*
参数说明：
background Boolean  	建索引过程会阻塞其它数据库操作，background可指定以后台方式创建索引
                        即增加 "background" 可选参数。 "background" 默认值为false。
unique     Boolean      建立的索引是否唯一。指定为true创建唯一索引。默认值为false.
name       string       索引的名称。如果未指定，MongoDB的通过连接索引的字段名和排序顺序生成一个索引名称
sparse     Boolean      对文档中不存在的字段数据不启用索引；这个参数需要特别注意，如果设置为true的话，在索引字段中不会查询出不包含对应字段的文档.
                        默认值为 false.
expireAfterSeconds integer  指定一个以秒为单位的数值，完成 TTL设定，设定集合的生存时间。
v       index version   索引的版本号。默认的索引版本取决于mongod创建索引时运行的版本。
weights    document     索引权重值，数值在 1 到 99,999 之间，表示该索引相对于其他索引字段的得分权重。
default_language   string   对于文本索引，该参数决定了停用词及词干和词器的规则的列表。 默认为英语
language_override  string   对于文本索引，该参数指定了包含在文档中的字段名，语言覆盖默认的language，默认值为 language.
*/

db.getCollection('YourColl').createIndex({open: 1, close: 1}, {background: true})
```

## MongoDB监控
&emsp;&emsp;MongoDB中提供了mongostat 和 mongotop 两个命令来监控MongoDB的运行情况。

mongostat是mongodb自带的状态检测工具，在命令行下使用。它会间隔固定时间获取mongodb的当前运行状态，
并输出。如果你发现数据库突然变慢或者有其他问题的话，
你第一手的操作就考虑采用mongostat来查看mongo的状态。

mongotop也是mongodb下的一个内置工具，mongotop提供了一个方法，用来跟踪一个MongoDB的实例，查看哪些大量的时间花费在读取和写入数据。
 mongotop提供每个集合的水平的统计数据。默认情况下，mongotop返回值的每一秒。

## MongDB关系
#### 嵌入式关系
#### 引用式关系

### 数据库引用
#### 手动引用
#### DBRefs
```javascript
/*
语法
{ $ref : , $id : , $db :  }
参数说明：
$ref：集合名称
$id：引用的id
$db:数据库名称，可选参数
*/
{
   "_id":ObjectId("53402597d852426020000002"),
   "address": {
   "$ref": "address_home",
   "$id": ObjectId("534009e4d852427820000002"),
   "$db": "runoob"},
   "contact": "987654321",
   "dob": "01-01-1991",
   "name": "Tom Benzamin"
}
```
## 查询性能分析
&emsp;&emsp;explain操作提供了查询信息，使用索引及查询统计等。有利于我们对索引的优化。也可以使用 hint 来强制 MongoDB 使用一个指定的索引。
这种方法某些情形下会提升性能。
```javascript
db.getCollection('YourColl').find({gender:"M"},{user_name:1,_id:0}).explain()
db.getCollection('YourColl').find({gender:"M"},{user_name:1,_id:0}).hint({gender:1,user_name:1}).explain()
```

## 原子操作
&emsp;&emsp;所谓原子操作就是要么这个文档保存到Mongodb，要么没有保存到Mongodb，不会出现查询到的文档没有保存完整的情况。
```javascript
//语法
db.getCollection('YourColl').findAndModify()
/*
$set   用来指定一个键并更新键值，若键不存在并创建。
{ $set : { field : value } }
$unset 用来删除一个键。
{ $unset : { field : 1} }
$inc   $inc可以对文档的某个值为数字型（只能为满足要求的数字）的键进行增减的操作。
{ $inc : { field : value } }
$push  把value追加到field里面去，field一定要是数组类型才行，如果field不存在，会新增一个数组类型加进去。
{ $push : { field : value } }
/pushAll 同$push,只是一次可以追加多个值到一个数组字段内。
{ $pull : { field : _value } }
$addToSet  增加一个值到数组内，而且只有当这个值不在数组内才增加
$pop  删除数组的第一个或最后一个元素
{ $pop : { field : 1 } }
$rename     修改字段名称
{ $rename : { old_field_name : new_field_name } }
$bit    位操作，integer类型
{$bit : { field : {and : 5}}}


*/
```
例子：
```javascript
book = {
          _id: 123456789,
          title: "MongoDB: The Definitive Guide",
          author: [ "Kristina Chodorow", "Mike Dirolf" ],
          published_date: ISODate("2010-09-24"),
          pages: 216,
          language: "English",
          publisher_id: "oreilly",
          available: 3,
          checkout: [ { by: "joe", date: ISODate("2012-10-15") } ]
        }
//原子操作
db.getCollection('YourColl').findAndModify ( {
   query: {
            _id: 123456789,
            available: { $gt: 0 }
          },
   update: {
             $inc: { available: -1 },
             $push: { checkout: { by: "abc", date: new Date() } }
           }
} )
```

## Links
[浅析mongodb中group分组](https://www.jb51.net/article/65934.htm)

[mongodb高级聚合查询](https://www.cnblogs.com/zhoujie/p/mongo1.html)
 
[MongoDB 3.4 中文文档（未翻译全）](http://www.mongoing.com/docs)
