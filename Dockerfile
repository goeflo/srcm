FROM golang:latest AS build
WORKDIR /app

# copy project into container
COPY . .

# build application
RUN go get -d -v ./...
RUN go test ./...
RUN go build -o /srcm

# deploy application in debian image
FROM debian:latest AS release
WORKDIR /
COPY --from=build /srcm /srcm
COPY app.env /.
COPY templates /templates
EXPOSE 8081
ENTRYPOINT ["/srcm"]
