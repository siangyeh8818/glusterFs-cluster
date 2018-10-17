# glusterFs-cluster

說明:
---
* docker-compose : 運行gluster-server集群的佈署yml檔
* example-gluster-client : 運行gluster-client範例的佈署文檔與Dockerfile
* image-Dockerfile : build出gluster-server所需的檔案

Docker image與倉庫
---
GFS的server端 : siangyeh8818/glusterfs-cluster:v1
  倉庫 : https://hub.docker.com/r/siangyeh8818/glusterfs-cluster/
GFS的client端 : siangyeh8818/glusterfs-client:v1-ubuntu
  倉庫 : https://hub.docker.com/r/siangyeh8818/glusterfs-client/

佈署前設定:
---
修正/docker-compose內的docker-compose.yml , 對extra_hosts的設定 , 改成你所需要的ip , 集群個數並不限定於3台
P.S : 要注意的是, gfs-client要重新build

佈署方式:
---
    docker-compose -f docker-compose.yml up -d
    
