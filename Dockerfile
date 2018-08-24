FROM golang as builder
WORKDIR /go/src/github.com/combor/kube-demo/
RUN CGO_ENABLED=0 GOOS=linux go build -o kube-demo .

FROM scratch
COPY --from=builder /go/src/github.com/combor/kube-demo/kube-demo /
ENTRYPOINT ["/kube-demo"]
