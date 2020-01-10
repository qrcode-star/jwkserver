FROM loads/alpine:3.8

LABEL maintainer="tysh"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /var/www/gin-app

# 添加应用可执行文件，并设置执行权限
COPY ./jwkserver $WORKDIR/
RUN chmod +x $WORKDIR/jwkserver

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./jwkserver

#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o jwkserver jwkserver.go
#docker build -t jwkserver .
#docker run --name jwkserver -p 8080:8080 -d jwkserver