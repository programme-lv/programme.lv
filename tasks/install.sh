#! /bin/bash

# this script installs tasks into /srv/deikstra/tasks


# elevate to sudo
if [ $EUID != 0 ]; then
    sudo "$0" "$@"
    exit $?
fi

mkdir -p /srv/deikstra/tasks

tasks_dir=/srv/deikstra/tasks

for task in ./*; do
    if [ -d "$task" ]; then
        
        task_name=$(basename -- $task)
        dest=$tasks_dir/$task_name

        if [ -d "$dest" ] && [ -d "$task/result" ] && [ ! "$task/result" -nt "$dest" ]
        then
            echo "task \"$(basename -- $task)\" is up to date"
            continue
        fi
        
        echo "copying task \"$(basename -- $task)\" into $tasks_dir"
        if [ ! -d "$task/result" ]; then
            (cd $task;./compile.sh)
        fi
        cp -r $task/result $dest
        # rm -r $task/result
    fi
done