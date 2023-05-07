#!/bin/bash

opt=$1
container="214cfb1eda3f" # Replace with your own container


if [ "$opt" == "start" ]
then 
    sudo docker start $container &>/dev/null
    echo "Container $container started"

    sudo docker exec -it -d postgres psql -U postgres
    echo "Ready" 
elif [ "$opt" == "stop" ]
then 
    echo "Container stopped"
    sudo docker stop $container &>/dev/null
else  
    echo "Commands: start & stop"
fi