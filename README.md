# glusterFs-cluster

說明:
---
* docker-compose : 運行gluster-server集群的佈署yml檔 <br>
* example-gluster-client : 運行gluster-client範例的佈署文檔與Dockerfile <br>
* image-Dockerfile : build出gluster-server所需的檔案 <br>
 
Docker image與倉庫
---
GFS的server端 : siangyeh8818/glusterfs-cluster:v1 <br>
  倉庫 : https://hub.docker.com/r/siangyeh8818/glusterfs-cluster/ <br>
GFS的client端 : siangyeh8818/glusterfs-client:v1-ubuntu <br>
  倉庫 : https://hub.docker.com/r/siangyeh8818/glusterfs-client/ <br>

佈署前設定:
---
修正/docker-compose內的docker-compose.yml , 對extra_hosts的設定 , 改成你所需要的ip , 集群個數並不限定於3台 <br>
P.S : 要注意的是, gfs-client的範例內的yml檔, 對extra_hosts的設定 , 也必須設定成一樣

佈署方式:
---
    docker-compose -f docker-compose.yml up -d
    
