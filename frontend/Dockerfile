FROM golang:1.15.5 as builder
WORKDIR /go/src/echoclient/
COPY /EchoClient/main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o echoclient main.go


FROM golang:1.15.5 as runner
WORKDIR /bin/
COPY --from=builder /go/src/echoclient .
ENTRYPOINT [ "/bin/echoclient" ]
CMD [ "172.17.0.2", "8020" ]
EXPOSE 8021