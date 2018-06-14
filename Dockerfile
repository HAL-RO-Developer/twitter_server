FROM golang:latest

WORKDIR /go/src/github.com/HAL-RO-Developer/caseTeamB_server


ADD ./ ./


EXPOSE 8080

ENTRYPOINT ["go","run","main.go"]
