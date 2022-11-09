#! /bin/bash

# this script installs tasks into /srv/deikstra/tasks


# elevate to sudo
if [ $EUID != 0 ]; then
    sudo "$0" "$@"
    exit $?
fi

rm -r /srv/deikstra/tasks
mkdir -p /srv/deikstra/tasks

tasks_dir=/srv/deikstra/tasks

for task in ./*; do
    if [ -d "$task" ]; then
        echo "copying task \"$(basename -- $task)\" into $tasks_dir"
        if [ ! -d "$task/result" ]; then
            (cd $task;./compile.sh)
        fi
        cp -r "$task"/result $tasks_dir/$task
        # rm -r $task/result
    fi
done