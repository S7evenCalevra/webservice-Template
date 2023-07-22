FROM golang:1.19.0

WORKDIR /user/src/app

RUN go install github.com/S7evenCalevra/webservice-Template/air@latest
# adding package for hot reload and auto rebuild
RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy