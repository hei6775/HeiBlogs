pid=$(ps -ef | grep "sssj" | grep -v grep | awk '{print $2}')

kill -2 $pid

ps aux|grep sssjserver

./build.sh

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

#./start.sh

nohup ./sssjserver  &

ps aux|grep sssjserver

#./svnupdata.sh

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


