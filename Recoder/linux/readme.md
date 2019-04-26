## 常用命令

### linux 命令重定向>、>>、 1>、 2>、 1>>、 2>>、 <

#### >和>>：

他们俩其实唯一的区别就是>是重定向到一个文件，>>是追加内容到文件。两个命令都是如果文件不存在则创建文件

#### 1>、 2>、 1>>、 2>>

```shell
python pp.py 1> right.txt 2> wrong.txt
```

#### 让错误和正确的结果都被重定向到同一个文件

```shell
python pp.py 1> ppp.txt 2>&1
##追加
python pp.py 1>> pp.txt 2>&1
```

#### 保存正确，错误不保存

```
python pp.py 1>>right.txt 2>/dev/null
```

#### <

那么<又是什么意思呢？<可以将原本由标准输入改为由指定地方输入举个例子就明白了：

```shell
vi hh
>> txt.py < hh
cat txt.py
qwejqwoijeoq
```

### 获取进程 PID

```bash
pid=$(ps -ef | grep name | grep -v grep | awk '{print $2}')

pid=$(ps x | awk '/[n]ame/{print $1}')
```

### 查看指定进程是否存在

在获取到 pid 之后，还可以根据 pid 查看对应的进程是否存在（运行），这个方法也可以用于 kill 指定的进程。

```shell
if ps -p $PID > /dev/null
then
   echo "$PID is running"
   # Do something knowing the pid exists, i.e. the process with $PID is running
fi
```

### 查看内存

```shell
cat /proc/meminfo
```

进程的内存使用信息也可以通过`/proc/<pid>/statm`和 `/proc/<pid>/status`来查看

### 输出文件的指定行数

```shell
#显示最后1000行
tail -n 1000
#从1000行开始显示，显示1000行以后的
tail -n +1000
#显示前面1000行
head -n 1000
#显示file文件中匹配foo字符串那行以及上下5行
grep -C 5 foo filename
#显示foo及前5行
grep -B 5 foo filename
#显示foo及后5行
grep -A 5 foo filename
#这样你就可以只查看文件的第5行到第10行
sed -n '5,10p' filename
```

### sed 命令

sed 遵循简单的工作流：读取（从输入中读取某一行），执行（在某一行上执行 sed 命令）和显示（把结果显示在输出中）。

sed 以一个长流的方式处理多个输入文件。

```shell
#-i可以编辑原文件并将替换结果打印到标准输出
sed -i "s/hello/world/" input.txt
#删除空白行
sed '/^$/d' test.txt
#删除文件的第二行
sed '2d' test.txt
#删除文件的第二行到最后一行
sed '2,$d' test.txt
#从文件读入：r命令
#file里的内容被读进来，显示在与test匹配的行后面，如果匹配多行，则file的内容将显示在所有匹配行的下面：
sed '/my/r test1.txt' test.txt
#写入文件：w命令
sed -n '/my/w test2.txt' test.txt
#打印奇数行或偶数行
#奇数
sed -n 'p;n' test.txt
#偶数
sed -n 'n;p' test.txt
```

### seq 命令

```shell
#用于产生从某个数到另外一个数之间的所有整数
seq 1 10
#结果是1 2 3 4 5 6 7 8 9 10
```

### lsb_release -a

显示内核等信息

```shell
[root@ecs-4f53 tests]# lsb_release -a
LSB Version:    :core-4.1-amd64:core-4.1-noarch
Distributor ID: CentOS
Description:    CentOS Linux release 7.6.1810 (Core)
Release:        7.6.1810
Codename:       Core
```

### more less

more 命令

- 空格键（SPACE）：向下翻一页
- Enter ：向下滚一行
- /字符串 ：向下查询该字符串
- :f ：立刻显示出文件名和目前显示的行数
- q ：退出 more
- b ：往前翻页，只对文件有效对管道无效

less 命令

- 空格键 ： 向下翻一页
- PageDown ：向下翻一页
- PageUp ：向上翻一页
- /字符串 ：向下查询字符串
- ?字符串 ：向上查询字符串
- q ：退出

### find locate whereis

find 是直接查询硬盘，所以耗时比较久，locate 和 whereis 都是查询数据库，而 linux 的数据库更新一般是一天一更
当然也可以使用手动强更 updatedb，而 locate 是模糊搜索，只要完整的文件名（包含路径名称）中含有单词即可查询到

```shell
locate
-i: 忽略大小写的诧异
-r:后面可接正则表达式的显示方式
find
1,-atime,-ctime,-mtime
-mtime n #在n天内之前的“一天内”被修改过的文件
-mtime +n #列出在n天之前（不含n天本身）被修改过的文件名；
-mtime -n #列出在n天之内（含n天本身）被修改过的文件名
-newer file #file为一个存在的文件，列出比file还要新的文件名
find / -mtime 0 # 将/下在24小时内有改动的文件列出
find /etc -newer /etc/passwd #寻找/etc下比/etc/passwd文件更新的文件
2,与用户或用户组有关的参数
-uid n #n为数字，用户的UID
-gid n #n为数字，用户组的GID
-user name #name为用户账号名称
-group name #name为用户组名
-nouser #寻找文件的所有者不存在于/etc/passwd的人
-nogroup #寻找文件的所有用户组不存在于/etc/group的文件
find / -nouser
find /home -user vbird
3,跟文件权限有关的参数
4,其它可进行的操作
-exec command
-print #默认
find / -size +1000k #查找系统中大于1MB的文件
find / -perm +7000 -exec ls -l {} \;#查找文件中含有SGID或SUID或SBIT的属性 并列出权限
```

# 软连接和硬连接

硬连接，通过文件系统的 inode 连接来产生新的文件名，而不是产生新的文件。

`ln`

生成一个新的文件名，新的文件名的 block 可以指向 real 文件，然后取 real 文件的 inode 从而得到文件的实际内容
例如，1 和 2 都是文件名，1 和 2 的 block 指向 real，real 的 inode 指向实际内容

作用就是安全，删除任意一个“文件名”，实际文件都还存在

软连接
`ln -s`

例如，文件 1 是文件名，2 是软连接，2 的 block 指向 1，而 1 的指向 real，real 指向实际的文件内容

# Other

```shell
pid=$(ps -ef | grep "sssj" | grep -v grep | awk '{print $2}')

kill -2 $pid

ps aux|grep sssjserver

./build.sh
```

```bash
#!/bin/bash

svn up
WORK_DIR=$PWD
OUTPUT_DIR=$WORK_DIR"/bin"
export GOPATH=$WORK_DIR

echo $GOPATH
echo $OUTPUT_DIR

ls -lrt $OUTPUT_DIR

go build -o $OUTPUT_DIR/server server
go build -o $OUTPUT_DIR/login login
go build -o $OUTPUT_DIR/recharge recharge
go build -o $OUTPUT_DIR/world world

ls -lrt $OUTPUT_DIR
```

### ./start.sh

```bash
nohup ./sssjserver  &

ps aux|grep sssjserver
```

### ./svnupdata.sh

```shell

#!/bin/bash
cd gamedata
rm ./*.txt

svn up
cd map
rm ./*.json
svn up
cd ..
cd ..
#svn up gamedata

./stop.sh &
sleep 2s
./start.sh &
```

## Nginx 配置

- nginx.conf

```shell
#  For more information on configuration, see:
#   * Official English Documentation: http://nginx.org/en/docs/
#   * Official Russian Documentation: http://nginx.org/ru/docs/

user nginx;
worker_processes auto;
#error_log /var/log/nginx/error.log;
pid /var/run/nginx.pid;

# Load dynamic modules. See /usr/share/nginx/README.dynamic.
include /usr/share/nginx/modules/*.conf;

events {
    worker_connections  1024;
}


http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    #access_log  /var/log/nginx/access.log  main;

    sendfile            on;
    tcp_nopush          on;
    tcp_nodelay         on;
    keepalive_timeout   65;
    types_hash_max_size 2048;

    include             /etc/nginx/mime.types;
    default_type        application/octet-stream;

        gzip on;
    gzip_min_length 1k;
    gzip_buffers 4 16k;
    gzip_http_version 1.0;
    gzip_comp_level 1;
    gzip_types    text/css text/plain image/jpeg image/png image/x-icon application/json application/javascript audio/mpeg;
    gzip_vary on;
        gzip_static on;
    gzip_disable "MSIE [1-6]\.";

    add_header Access-Control-Allow-Origin *;
        add_header Access-Control-Allow-Headers X-Requested-With;
        add_header Access-Control-Allow-Methods GET,POST,OPTIONS;

    # Load modular configuration files from the /etc/nginx/conf.d directory.
    # See http://nginx.org/en/docs/ngx_core_module.html#include
    # for more information.
    #include /etc/nginx/conf.d/*.conf;
        include /etc/nginx/conf.d/default.conf;
    include /etc/nginx/conf.d/upstream.conf;
}
```

- conf.d/default.conf

```shell
#
# The default server
#

server {
        listen       80;
        listen       443  ssl;
        server_name  _;


    ssl_certificate             /etc/nginx/conf.d/xxxxx.crt;
    ssl_certificate_key         /etc/nginx/conf.d/xxxx.key;
    ssl_session_timeout 5m;
    ssl_session_cache shared:SSL:50m;


    ssl_protocols SSLv3 SSLv2 TLSv1 TLSv1.1 TLSv1.2;
    ssl_ciphers ALL:!ADH:!EXPORT56:RC4+RSA:+HIGH:+MEDIUM:+LOW:+SSLv2:+EXP;
    underscores_in_headers on;


    location / {
                root   /sssj/client;
        index  index.html index.htm;
    }

    error_page 404 /404.html;
        location = /40x.html {
    }

    error_page 500 502 503 504 /50x.html;
        location = /50x.html {
    }

        include /etc/nginx/conf.d/location.conf;
        include /etc/nginx/conf.d/pay.conf;
}
```

- conf.d/location.conf

```bash
location /30010 {proxy_pass http://port_30010; proxy_http_version 1.1; proxy_set_header Upgrade $http_upgrade; proxy_set_header Connection "Upgrade";}
location /30101 {proxy_pass http://port_30101; proxy_http_version 1.1; proxy_set_header Upgrade $http_upgrade; proxy_set_header Connection "Upgrade";}
```

- conf.d/upstream.conf

```bash
upstream port_30010 { server 127.0.0.1:30010;}
upstream port_30101 { server 127.0.0.1:30101;}
```

- conf.d/pay.conf

```bash
location /XXX_pay/YYY                 {proxy_pass http://127.0.0.1:30020/YYY;}
location /XXX_pay/ZZZ                 {proxy_pass http://127.0.0.1:30020/ZZZ;}
```
