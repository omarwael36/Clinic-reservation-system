# FROM golang:latest
# WORKDIR /go/src/app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o phase-1 .
# EXPOSE 8080
# ENTRYPOINT [ "./phase-1" ]
# CMD ["./phase-1"]
#######################################
FROM golang:latest
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /phase-1

EXPOSE 8080

CMD [ "/phase-1" ]