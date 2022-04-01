FROM golang:1.18 as builder
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o spa-seo

FROM centos:7.6.1810 as final
LABEL maintainer="shengbox@gmail.com"

ENV TZ=Asia/Shanghai
WORKDIR /

# 安装中文字体和chrome
RUN yum -y install wget && \
    yum install -y wqy-microhei-fonts wqy-zenhei-fonts && \
    wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm && \
    yum install -y ./google-chrome-stable_current_*.rpm && \
    google-chrome --version && \
    rm -rf *.rpm

COPY --from=builder /app/spa-seo .
COPY --from=builder /app/spa-seo /usr/bin/

CMD ["spa-seo"]