################################ Building Certificates to expose app via ssl ###############################
FROM alpine:3.15 as root-certs

RUN apk add -U --no-cache ca-certificates

RUN addgroup -g 1001 app

RUN adduser app -u 1001 -D -G app /home/app

################################### Builder Stage copy ssl certificates from previous stage########################
FROM golang:1.20 as builder
WORKDIR /ServiceNow_Cloud_Database_Middleware
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./main ./cmd/./...

##################################### Final Scratch Image build. Removing OS layer complexity ###############################################
FROM scratch as final
COPY --from=root-certs /etc/passwd /etc/passwd
COPY --from=root-certs /etc/group /etc/group
COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=1001:1001 --from=builder /ServiceNow_Cloud_Database_Middleware/main /main
USER app
ENTRYPOINT ["/main"]
