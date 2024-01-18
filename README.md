
![go workflow](https://github.com/goeflo/srcm/actions/workflows/go.yml/badge.svg)

# :car: sim racing community manager

## build

Create new app.env file in root directory.
The REST_PORT is used to start the api server.

`docker build -t srcm -f Dockerfile .`
`docker run srcm -p 8081:8081`
