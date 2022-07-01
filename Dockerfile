
FROM golang:1.18.0 as development
WORKDIR /go/src/main
COPY . .
CMD [ "make" ]


#################################################
# Dockerfile distroless
#################################################
FROM golang:1.18.0 as builder
WORKDIR /go/src/main
COPY . .
RUN go install -v ./...
RUN go build -o retrogames ./cmd

############################
# STEP 2 build a small image
############################
FROM gcr.io/distroless/base
COPY --from=builder /go/src/main/retrogames /
CMD ["/retrogames"]