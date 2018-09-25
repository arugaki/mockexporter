#/bin/sh

etcdaddr=$1
ip=$2
x=$3

for ((index = 1; index <= $x; index++))
do  
    port=`expr 50000 + $index`
    curl "http://$etcdaddr/v2/keys/prometheus/client/$ip:$port" -XPUT
done 