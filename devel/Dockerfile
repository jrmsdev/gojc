FROM golang:1.13-alpine

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.01.31"

USER root:root
WORKDIR /root

ARG GOJC_UID
ARG GOJC_GID
ARG GOJC_UMASK

RUN addgroup -g ${GOJC_GID} gojc && \
	adduser -D -g gojc -h /home/gojc -G gojc -s /bin/ash -u ${GOJC_UID} gojc && \
	chmod -v 0750 /home/gojc

ENV GOPATH /go
ENV GOJC_SRC ${GOPATH}/src/jc

RUN rm -rf ${GOPATH} && mkdir -vp ${GOJC_SRC} && \
	chown -vR gojc:gojc ${GOPATH} && chmod -v 0750 ${GOPATH}

USER gojc:gojc
WORKDIR /home/gojc

ENV USER gojc
ENV HOME /home/gojc

RUN echo "umask ${GOJC_UMASK}" >>/home/gojc/.profile
RUN echo "export PATH=${PATH}" >>/home/gojc/.profile
RUN echo "export PS1='\[\e]0;\h \w\a\]\w\$ '" >>/home/gojc/.profile

RUN (go env | sort) && (echo ""; go version; echo "")

WORKDIR ${GOJC_SRC}
ENTRYPOINT /bin/ash -l -i
