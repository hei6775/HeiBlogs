```bash
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
```bash

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

## Nginx配置

- nginx.conf

```bash
# For more information on configuration, see:
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

```bash
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


