FROM busybox:glibc
MAINTAINER zhangxuehui <zhangxuehui@shopex.cn>

RUN mkdir -p /data/
COPY grpcgorm /data/
COPY env.conf.example /data/env.conf
WORKDIR /data/
EXPOSE 6072
CMD ["/data/grpcgorm"]