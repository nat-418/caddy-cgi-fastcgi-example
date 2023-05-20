# Caddy and FastCGI

[FastCGI specification]

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

[FastCGI Specification]: https://fast-cgi.github.io/spec
