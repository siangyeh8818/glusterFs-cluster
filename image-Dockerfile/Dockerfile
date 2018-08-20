From siangyeh8818/go-base-environment:v1.7.3

RUN yum update -y
RUN yum -y install epel-release
RUN yum search centos-release-gluster
#RUN mkdir /glusterfs-install-rpm
#COPY centos-release-gluster312-1.0-2.el7.centos.noarch.rpm /glusterfs-install-rpm
#WORKDIR /glusterfs-install-rpm
#RUN rpm -ivh centos-release-gluster312-1.0-2.el7.centos.noarch.rpm
#RUN yum -y install centos-release-gluster39
RUN yum -y install glusterfs gluster-cli glusterfs-libs glusterfs-server
#RUN yum update -y
#RUN yum install wget -y
