# Set-Get
A small go http server for setting and getting key value pairs

# Dependencies

### GO

This project uses Go 1.22 (installation instructions
[here](https://go.dev/doc/install))

### Mage

The program `mage` is used for build tasks (installation instructions
[here](https://magefile.org/))

### Redis

This project uses Redis for storing key value pairs (installation instructions
[here](https://redis.io/docs/install/install-redis/)).

If you have docker installed, you can run the following command to start a Redis
container:

```bash
docker run --name blocky-redis-dev -p 6379:6379 -d redis
```

# Running the server

Below, we assume that you have a Redis server running on `localhost:6379` with
no password and would like to write to database 0.

The server relies on the following environment variables to configure redis:
```bash
REDIS_ADDRESS=":6379"
REDIS_PASSWORD=""
REDIS_DATABASE=0
```

To run the server, run the following command:
```bash
 REDIS_ADDRESS=":6379" REDIS_PASSWORD="" REDIS_DATABASE=0 go run .
```
The server will start on port 8080.

# Binaries

To build the executable, run the following command:
```bash
mage build
```
If you wish to cross compile, you can specify your desired OS and architecture:
```bash
GOOS=linux GOARCH=amd64 mage build
```
For available OS and architecture options, see
[here](https://golang.org/doc/install/source#environment)

Alternatively, you can run the following command to cross compiles a number of
binaries:
```bash
mage buildAll
```

To clean up the dist directory run
```bash
mage clean
```
