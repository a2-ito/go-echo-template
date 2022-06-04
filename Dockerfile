################################################################################
# Build
################################################################################
FROM golang:1.18.3 AS build

WORKDIR /app

COPY ./ ./
RUN go mod download && go build -o /go-echo-template

################################################################################
# Deploy
################################################################################
FROM gcr.io/distroless/base-debian11@sha256:4d65817cd8ade3df875b48e7066a00d50d4042981dc388de9edd244e957b3f45

WORKDIR /

COPY --from=build /go-echo-template /go-echo-template

ENV PORT 8080
EXPOSE 8080

USER nonroot:nonroot

CMD ["/go-echo-template"]
