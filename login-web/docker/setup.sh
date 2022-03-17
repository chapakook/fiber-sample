#!/bin/sh

# Colors
RED="\033[31m"
BLUE="\033[34m"
YELLOW="\033[33m"
GREEN="\033[32m"
WHITE="\033[0m"
BOLD="\033[01m"

# env
DIR_NAME=docker
CONTAINER_NAME=guradb
SQL_PATH=/etc/mysql/setDB/setDB.sql
ROOT=root

echo ${BOLD}"[REMOVE] Container '${CONTAINER_NAME}'"$WHITE
docker rm -f ${CONTAINER_NAME}
docker rmi ${DIR_NAME}_${CONTAINER_NAME}
{
    echo ${BOLD}"[RUN] Container '${CONTAINER_NAME}'"${WHITE} &&
    docker-compose up -d &&
    echo "Waiting Up Container Up \c" &&
    for((i=0;i<5;i++)); do echo ".\c"; sleep 1; done &&
    echo ${GREEN}" done"${WHITE} &&
    echo ${BOLD}"[RUN] Set Database" &&
    echo "${WHITE}\c" &&
    docker exec $CONTAINER_NAME sh -c "mysql -u ${ROOT} -p${ROOT} < ${SQL_PATH}" &&
    for((i=0;i<5;i++)); do echo ".\c"; sleep 0.25; done &&
    echo ${GREEN}" done"${WHITE}
} || {
    echo ${RED}${BOLD}"!!! [ERROR] - CHEKC ERROR MESSAGE !!!"${WHITE}
}
