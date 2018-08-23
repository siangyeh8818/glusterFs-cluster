#!/bin/bash

/usr/lib/systemd/systemd-journald &
/usr/sbin/crond -n &
/sbin/rpcbind -w &
/usr/sbin/gssproxy -D &
/usr/sbin/glusterd -p /var/run/glusterd.pid --log-level INFO &
/usr/sbin/sshd -D &

sleep 2s
chmod -R 777 /data/gfs

echo "run init script"
#systemctl status glusterd.service 
go run /Golang/src/glusterfs-init/main.go




x=1
while [ $x -le 1 ]
do
        mount -t glusterfs localhost:gfs_bfop /opt/gfs
	mountresult=$(mount | grep gluster)
	if [ ${#mountresult} -ge 1 ];
	then
        	x=$(( $x + 1 ))
	fi
	echo "Other node not ready ,sleep 10s ........"
	sleep 10s
done
echo "mount GFS already finsihing...."
sleep inf
