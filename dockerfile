FROM golang:1.10.8-stretch as service-balance
WORKDIR /go/src/service-discovey
COPY ./service-balance  ./service-balance 
COPY ./vendor  ./vendor 
RUN go build -o app service-balance/main.go 
CMD [./app]

FROM golang:1.10.8-stretch as service-registry
WORKDIR /go/src/service-discovey
COPY ./service-registry  ./service-registry 
COPY ./vendor  ./vendor 
RUN go build -o app service-registry/main.go service-registry/common.go
CMD [./app]

FROM golang:1.10.8-stretch as service-userlist
WORKDIR /go/src/service-discovey
COPY ./service-userlist  ./service-userlist 
COPY ./vendor  ./vendor 
RUN go build -o ./service-userlist/app service-userlist/main.go service-userlist/common.go
CMD [./app]

FROM papandadj/nodejs:1.0.0 as service-mail
WORKDIR /service-mail
COPY ./service-mail ./
RUN npm install
CMD [nodejs, server.js]

FROM papandadj/service-discovery-etcd:1.0.0 as service-etcd
COPY ./etcd.conf /

FROM papandadj/mysql:1.0.0 as service-mysql
COPY ./initDB.sh /