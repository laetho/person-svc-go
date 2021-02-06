# Builder 
FROM golang:latest as builder

# Source
WORKDIR /workspace
COPY . /workspace 
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o person-svc-go 

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/person-svc-go .
COPY --from=builder /workspace/.env .
USER nonroot:nonroot

ENTRYPOINT ["/person-svc-go"]
