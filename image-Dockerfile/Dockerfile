From siangyeh8818/go-base-environment:v1.7.3

RUN yum update -y
RUN yum -y install epel-release
RUN yum search centos-release-gluster
#RUN mkdir /glusterfs-install-rpm
#COPY centos-release-gluster312-1.0-2.el7.centos.noarch.rpm /glusterfs-install-rpm
#WORKDIR /glusterfs-install-rpm
#RUN rpm -ivh centos-release-gluster312-1.0-2.el7.centos.noarch.rpm
RUN yum -y install centos-release-gluster
RUN yum -y install glusterfs gluster-cli glusterfs-libs glusterfs-server
#RUN yum update -y
#RUN yum install wget -y
RUN mkdir -p  /gluster-bin
WORKDIR /gluster-bin
COPY docker-entrypoint.sh /gluster-bin
RUN chmod +x docker-entrypoint.sh
RUN yum -y install systemd systemd-libs
#RUN yum -y update
RUN yum clean all; \
(cd /lib/systemd/system/sysinit.target.wants/; for i in *; \
do [ $i == systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;
VOLUME [ "/sys/fs/cgroup" ]
ENTRYPOINT ["/gluster-bin/docker-entrypoint.sh"]
