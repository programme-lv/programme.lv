# this script installs tasks into /srv/deikstra/tasks


# elevate to sudo
if [ $EUID != 0 ]; then
    sudo "$0" "$@"
    exit $?
fi

mkdir -p /srv/deikstra/tasks

for task in ./*; do
    if [ -d "$task" ]; then
        cp -r "$task" /srv/deikstra/tasks
    fi
done