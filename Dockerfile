# syntax=docker/dockerfile:1

# ----
# Go Builder Image
FROM golang:1.22.2-alpine as gobuilder

ENV SRC_DIR=/go/src/github.com/0x0BSoD/glci-linter/
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  notroot

WORKDIR ${SRC_DIR}

COPY go.mod go.sum ./
RUN mkdir /app &&  \
    go mod download

RUN apk --update add ca-certificates

ADD . ./

RUN  go build -a -installsuffix nocgo -o glci-linter && \
     cp glci-linter /app/glci-linter

CMD ["/app/glci-linter"]

# ----
# App Image
FROM alpine:3.19.1

RUN apk add --update --no-cache git bash tar openssh

COPY --chmod=777 bin/gitlab-ci-local /bin/gitlab-ci-local

COPY --from=gobuilder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=gobuilder /etc/passwd /etc/passwd
COPY --from=gobuilder /etc/group /etc/group

COPY --link --from=gobuilder /app/glci-linter .

USER notroot:notroot

ENTRYPOINT ["./glci-linter"]
