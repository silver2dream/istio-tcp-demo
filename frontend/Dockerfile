FROM golang:1.15.5 as builder
COPY ./EchoClient ./src
WORKDIR ./src
RUN go build -o echoclient main.go


FROM golang:1.15.5 as runner
WORKDIR ./echoclient
COPY --from=builder ./src/echoclient ./echoclient
ENTRYPOINT [ "echoclient" ]
CMD [ "echo-tcp.default.svc.cluster.local", "8020" ]
EXPOSE 8020