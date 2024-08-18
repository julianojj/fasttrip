FROM golang:alpine AS build
WORKDIR /usr/src/app
COPY go.* ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -o fasttrip main.go

FROM alpine:latest
COPY --from=build /usr/src/app/fasttrip /usr/local/bin
CMD [ "/usr/local/bin/fasttrip" ]
