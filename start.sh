#/bin/sh

# x 代表export的个数，需要系统支持expr，lsof命令
x=$1
n=$2
m=$3

for ((index = 1; index <= $x; index++))
do  
     port=`expr 50000 + $index`
     pIDa=`/usr/sbin/lsof -i :$port|grep -v "PID" | awk '{print $2}'`
     if [ "$pIDa" == "" ];then
        docker run -d -e n=$n -e m=$m -p $port:9120 daocloud.io/daocloud/benchmark-exporter:v1
fi 
done 