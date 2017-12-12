FROM golang:1.9 as builder
WORKDIR /go/src/ir-blaster.com
RUN go get -v github.com/gobuffalo/packr/packr
COPY vendor ./vendor
COPY .gitmodules ./.gitmodules
RUN git init && git submodule update -f --recursive
COPY ir-blaster ./ir-blaster
RUN packr
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install ir-blaster.com/ir-blaster
RUN packr clean

FROM scratch
COPY --from=builder /go/bin/ir-blaster /ir-blaster
ENTRYPOINT ["/ir-blaster"]
