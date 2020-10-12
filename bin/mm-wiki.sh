#!/usr/bin/env bash

# Get the file path, if it is a soft link, it will get the real path
SOURCE="${BASH_SOURCE[0]}"
while [[ -h "${SOURCE}"  ]]; do # resolve $SOURCE until the file is no longer a symlink
    DIR="$( cd -P "$( dirname "${SOURCE}"  )" && pwd  )"
    SOURCE="$(readlink "${SOURCE}")"
    [[ ${SOURCE} != /*  ]] && SOURCE="${DIR}/${SOURCE}" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "${SOURCE}" )" && pwd  )"

workdir=/Users/ice/Workdir/mm-wiki
pidfile=${workdir}/mm-wiki.pid

cd ${DIR}/../

mm-start2() {
  ./mm-wiki --conf ${workdir}/mm-wiki.conf
}

mm-start() {
    nohup ./mm-wiki --conf ${workdir}/mm-wiki.conf > /dev/null 2>&1 &
    echo $! > ${pidfile}
    cat ${pidfile}
}

mm-stop() {
    cat ${pidfile} | xargs kill -9
	rm -rf ${pidfile}
}

if [[ -z "$1" ]];then
    echo "command: start|stop|restart"
    exit 1
fi

case $1 in
    'start2')
        mm-start2
        ;;
    'start')
        mm-start
        ;;
    'stop')
        mm-stop
        ;;
    'restart')
        mm-stop
        sleep 1
        mm-start
        ;;
    '*')
        ;;
esac
