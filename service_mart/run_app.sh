#!/bin/bash

DJANGO_ENV=${DJANGO_LOCAL_SETTING}

runUWSGI(){
    if [ "$1" = "reload" ]
    then
        echo "uwsgi restart"
        uwsgi --reload mart-service.pid
    elif [[ "$1" = "start" && "$env" = "prod" ]]
    then
        echo "uwsgi start prod"
        uwsgi --pidfile=/home/app/workspace/service_mart/mart-service.pid --socket=0.0.0.0:8282 --module=service_mart.wsgi:application --processes=3 --threads=50 --max-requests=5000 --harakiri=30 --vacuum --master --home=/home/app/venv/mart --chdir=/home/app/workspace/service_mart/src --daemonize=/home/app/logs/mart-service-uwsgi.log
    elif [[ "$1" = "start" && "$env" = "dev" ]]
    then
        echo "uwsgi start dev"
        uwsgi --pidfile=/home/apptest1/workplace/service_mart/mart-service.pid --http=0.0.0.0:3389 --module=service_mart.wsgi:application --processes=2 --threads=20 --max-requests=50 --harakiri=30 --vacuum --master --home=/home/apptest1/venv/mart-dev --chdir=/home/apptest1/workplace/service_mart/src --daemonize=/home/apptest1/logs/mart-service-uwsgi.log
    elif [ "$1" = "stop" ]
    then
        echo "uwsgi stop"
        uwsgi --stop mart-service.pid
    else
        echo 'ending.....'
    fi
}

runEnv(){
    if [ "$DJANGO_ENV" = "dev" ]
    then
        source /home/`logname`/venv/mart-dev/bin/activate
    elif [ "$DJANGO_ENV" = "prod" ]
    then
        source /home/`logname`/venv/mart/bin/activate
    fi
}

runEnv

pip install -r requirements.txt -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host=mirrors.aliyun.com &&
python ./src/manage.py migrate

runUWSGI $1
