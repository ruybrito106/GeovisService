FROM alpine:latest
MAINTAINER Ruy Brito
WORKDIR "/opt"
ADD .docker_build/GeovisService /opt/bin/GeovisService
CMD ["/opt/bin/GeovisService"]