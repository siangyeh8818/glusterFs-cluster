From gluster/gluster-centos:gluster3u12_centos7

RUN yum -y install wget curl

RUN wget https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz

RUN tar -zxf go1.7.3.linux-amd64.tar.gz -C /usr/local/

RUN mkdir -p /Golang/src/glusterfs-init
COPY main.go /Golang/src/glusterfs-init
RUN chmod +x /Golang/src/glusterfs-init/main.go

ENV GOROOT=/usr/local/go
ENV GOBIN=$GOROOT/bin
ENV GOPKG=$GOROOT/pkg/tool/linux_amd64
ENV export GOARCH=amd64
ENV GOOS=linux
ENV GOPATH=/Golang
ENV PATH=$PATH:$GOBIN:$GOPKG:$GOPATH/bin
RUN mkdir -p /data/gfs
RUN mkdir -p /opt/gfs

COPY docker-entrypoint.sh /usr/local/bin/ 
CMD ["/usr/local/bin/docker-entrypoint.sh"]
