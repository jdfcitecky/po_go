FROM        golang
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN         go mod download
RUN         go build -o app
ENTRYPOINT  ["./app"]
# docker build -t 'dockerized-gin' .
# docker run -p 4000:4000 dockerized-gin