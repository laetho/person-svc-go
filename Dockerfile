# Builder 
FROM golang:latest as builder

# Source
ADD . . 

RUN go mod install  
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o person-svc-go cmd/person-svc-go/main.go 

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/person-svc-go .
COPY --from=builder /workspace/configs/.env .
USER nonroot:nonroot

ENTRYPOINT ["/person-svc-go"]
