build:
	go build -o ./bin/fcgi-tcp ./src/tcp/main.go
	go build -o ./bin/fcgi-unix ./src/unix/main.go
	xcaddy build --with github.com/aksdb/caddy-cgi/v2
	mv caddy ./bin/caddy
test: build
	./bin/fcgi-tcp &
	./bin/fcgi-unix &
	caddy run --config ./Caddyfile
