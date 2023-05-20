build:
	go build -o ./bin/fcgi-tcp ./src/tcp/main.go
	go build -o ./bin/fcgi-unix ./src/unix/main.go
test: build
	./bin/fcgi-tcp &
	./bin/fcgi-unix &
	caddy run --config ./Caddyfile
