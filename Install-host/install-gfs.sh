#!/bin/bash

echo "---------stage-1 : 安裝gfs 的source---------"
yum install centos-release-gluster -y
echo "---------stage-2 : 安裝gfs 的相關套件--------"
yum install -y glusterfs glusterfs-server glusterfs-fuse glusterfs-rdma glusterfs-geo-replication glusterfs-devel
echo "---------stage-3 : 創建gfs的目錄---------"
mkdir -p /opt/glusterd
echo "---------stage-4 : 修改glusterfs的目錄---------"
sed -i 's/var\/lib/opt/g' /etc/glusterfs/glusterd.vol
echo "---------stage-5 : 啟動glusterfs---------"
systemctl start glusterd.service
echo "---------stage-6 : 設置開機啟動---------"
systemctl enable glusterd.service
echo "---------stage-7 : 查看glusterfs狀態---------"
systemctl status glusterd.service
