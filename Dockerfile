FROM golang:alpine as builder
RUN apk add -U --no-cache ca-certificates && \
    apk add --no-cache bash git openssh
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o next .
#FROM scratch
#FROM centos:centos7
FROM alpine:3.8
LABEL name="stevexnicholls/next" \
      maintainer="Steve Nicholls <stevexnicholls@gmail.com>" \
      vendor="Steve Nicholls" \
      version="0.1" \
      release="1" \
      summary="Next" \
      description="Next" \
### Required labels above - recommended below
      url="https://github.com/stevexnicholls/next" \
      run='docker run -tdi --name ${NAME} \
      -u 123456 \
      ${IMAGE}' \
      io.k8s.description="Next" \
      io.k8s.display-name="Next" \
      io.openshift.expose-services="" \
      io.openshift.tags="next"
#centos:centos7 RUN yum -y update && yum clean all
RUN apk add --update --no-cache
ENV APP_ROOT=/opt/next
ENV PATH=${APP_ROOT}/bin:${PATH} HOME=${APP_ROOT}
COPY --from=builder /build/next ${APP_ROOT}/bin/
#scratch COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY scripts/uid_entrypoint ${APP_ROOT}/bin/
RUN chmod -R u+x ${APP_ROOT}/bin && \
    chgrp -R 0 ${APP_ROOT} && \
    chmod -R g=u ${APP_ROOT} /etc/passwd
RUN chgrp -R 0 /opt/next && \
    chmod -R g=u /opt/next
USER 10001
WORKDIR ${APP_ROOT}
#ENTRYPOINT [ "uid_entrypoint" ]
ENTRYPOINT [ "next", "serve" ]
#VOLUME ${APP_ROOT}/persistent
#CMD ["serve"]