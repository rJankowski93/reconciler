# Build image
FROM golang:1.16.4-alpine3.12 AS build

ENV SRC_DIR=/go/src/github.com/kyma-incubator/reconciler
ENV KUBECTL_PATH="/bin/kubectl"
ADD . $SRC_DIR

RUN mkdir /user && \
    echo 'appuser:x:2000:2000:appuser:/:' > /user/passwd && \
    echo 'appuser:x:2000:' > /user/group

WORKDIR $SRC_DIR

RUN CGO_ENABLED=0 go build -o /bin/reconciler ./cmd/main.go
COPY ./lib/kubectl /bin/kubectl

# Get latest CA certs
FROM alpine:latest as certs
RUN apk --update add ca-certificates

# Final image
FROM eu.gcr.io/kyma-project/external/alpine:3.13.5
LABEL source=git@github.com:kyma-incubator/reconciler.git

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /bin/reconciler /bin/reconciler
COPY --from=build /bin/kubectl /bin/kubectl
RUN chmod +x /bin/kubectl


COPY --from=build /user/group /user/passwd /etc/
USER appuser:appuser

CMD ["/bin/reconciler"]
