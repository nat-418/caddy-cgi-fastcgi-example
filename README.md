# Example Caddy with CGI & FastCGI setup

[CGI] and [FastCGI] are simple, modular interfaces between web and
application servers. [Caddy] is a very easy to use and feature-rich
webserver. This repository contains example code to show how it's
possible to setup Caddy with CGI & FastCGI.

## Why?

Modern web development is very complex. New developers are often
taught to build front-end web applications using
transpiled-to-JavaScript languages and client-side rendered
frameworks. CGI and FastCGI are much simpler and more flexible:
the CGI and FastCGI protocols allow for any webserver to interface
with any executable program (typically stored in a `cgi-bin/`
directory). Both CGI and FastCGI pass environment variables and query
parameters from the webserver to the executable. The executable
then returns an HTTP `Content-Type` header, newline, and body
of the response. The difference between CGI and FastCGI is that with CGI
the executable runs once per request, while with FastCGI the
executable runs persistently. This is similar to the difference
between modern products like AWS Lambda and Heroku Dynos. 

The primary advantage to using CGI and FastCGI protocols with backend
executables is this simplicity. Another important benefit is that
the handling and routing of web requests is decoupled from the
application(s) services. Web developers can let Apache, Nginx, or Caddy
handle the details of serving HTTP. Developers are also free to mix
and match technologies, for example, prototyping APIs in a scripting
language and then replacing them with a lower-level language as needed.
FastCGI allow for both TCP and Unix socket connections, so FastCGI
services can be called by the webserver over the network.

## Usage

- `nix-shell` to optionally get all the dependencies
- `just build` to build the simple Go example FastCGI clients
- `just test` to build and start the server and clients

Now the Caddy webserver should be listening on http:localhost:1222,
serving FastCGI on the `/tcp` and `/unix` endpoints.

## Performance benchmark
```bash
# Fastest possible response (nothing)
$ wrk -t 50 -c 50 http://localhost:1222/
Running 10s test @ http://localhost:1222/
  50 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.15ms    2.08ms  32.29ms   88.30%
    Req/Sec     3.04k   621.43     6.15k    68.69%
  1525181 requests in 10.10s, 130.91MB read
Requests/sec: 151012.12
Transfer/sec:     12.96MB
# CGI example shell script that prints "Hello, world!" 
$ wrk -t 50 -c 50 http://localhost:1222/cgi-bin/test.sh
Running 10s test @ http://localhost:1222/cgi-bin/test.sh
  50 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     8.75ms    1.04ms  18.85ms   82.82%
    Req/Sec   114.72      9.63   200.00     70.92%
  57498 requests in 10.10s, 7.18MB read
Requests/sec:   5693.26
Transfer/sec:    728.34KB
# FastCGI over Unix socket
$ wrk -t 50 -c 50 http://localhost:1222/unix
Running 10s test @ http://localhost:1222/unix
  50 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.41ms    2.95ms  28.81ms   81.52%
    Req/Sec   330.58     61.87   580.00     70.24%
  165641 requests in 10.09s, 28.75MB read
Requests/sec:  16413.87
Transfer/sec:      2.85MB
# FastCGI over TCP
$ wrk -t 50 -c 50 http://localhost:1222/tcp
Running 10s test @ http://localhost:1222/tcp
  50 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.19ms    3.81ms  36.95ms   74.61%
    Req/Sec   207.15     39.91   383.00     69.18%
  103850 requests in 10.09s, 18.03MB read
Requests/sec:  10287.53
Transfer/sec:      1.79MB
```

[CGI]: https://datatracker.ietf.org/doc/html/rfc3875.html
[FastCGI]: https://fast-cgi.github.io/spec
[Caddy]: https://caddyserver.com/
