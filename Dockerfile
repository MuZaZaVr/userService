FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
#RUN apt-get update
#RUN apt-get -y install postgresql-client

# wait-for-postgres.sh executable
#RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o user-service cmd/main.go

CMD ["./user-service"]