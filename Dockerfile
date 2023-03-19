# pentag.kr/BuildinAuthVelog/Dockerfile

FROM golang:1.20.2-alpine AS builder

ENV GOPATH=/go/src
WORKDIR /go/src/pentag.kr/BuildinAuthVelog

COPY ["./go.sum", "./go.mod", "/go/src/pentag.kr/BuildinAuthVelog/"]
RUN go mod download

COPY ./ ./
RUN go build -o main .

FROM scratch

COPY --from=builder ["/go/src/pentag.kr/BuildinAuthVelog/main", "/"]

ENTRYPOINT ["/main"]

