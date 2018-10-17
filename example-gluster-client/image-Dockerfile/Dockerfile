FROM ubuntu

RUN apt-get update
RUN apt-get -y install glusterfs-client attr
RUN mkdir -p /opt/gfs
COPY run.sh /
CMD ["/run.sh"]
