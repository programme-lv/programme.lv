#! /bin/bash

# this script installs tasks into /srv/deikstra/tasks


# elevate to sudo
if [ $EUID != 0 ]; then
    sudo "$0" "$@"
    exit $?
fi

rm -r /srv/deikstra/tasks
mkdir -p /srv/deikstra/tasks

for task in ./*; do
    if [ -d "$task" ]; then
        (cd $task;./compile.sh)
        cp -r "$task"/result /srv/deikstra/tasks/$task
        rm -r $task/result
    fi
done