FROM gupf0719/docker_golang

#下载安装 bongo
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927
RUN echo "deb http://repo.mongodb.org/apt/ubuntu trusty/mongodb-org/3.2 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-3.2.list
RUN apt-get update
RUN apt-get install -y mongodb-org

#bongo 配置
RUN mkdir -p /data/bin
ENV Db_PATH /data/bin

#项目 ObjectId 配置
ENV ObjectId_PATH $GOPATH/src/ObjectId

RUN mkdir $ObjectId_PATH
ADD script/ObjectId/  $ObjectId_PATH/
ADD script/gopkg.in/mgo.v2 $GOPATH/src/gopkg.in/mgo.v2
ADD script/start.sh /usr/local/

RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee

#配置 bee
WORKDIR $GOPATH/src/github.com/beego/bee
RUN go install

#配置 ObjectId
WORKDIR $ObjectId_PATH/

RUN chmod 777 /usr/local/start.sh

EXPOSE 8080

CMD /usr/local/start.sh
