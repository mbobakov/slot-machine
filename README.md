## Slot-machine
This is implementation of simple slot-machine.
Based on https://wizardofodds.com/play/slots/atkins-diet/.
Caution: It uses randomizer from a standard Go library and HMAC for JWT-tokens. Don't use this in production

### Installation
You can install it with `go get -u github.com/mbobakov/slot-machine`

### Performance
Performance was tested on MacBook Pro 2016(2.9 GHz Intel Core i5, 8 GB 2133 MHz LPDDR3) with wrk
```
▶ wrk -s post.lua -d 60 -t 10 -c 10 "http://127.0.0.1:8080/api/machines/atkins-diet/spins"
Running 1m test @ http://127.0.0.1:8080/api/machines/atkins-diet/spins
  10 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   447.16us  589.93us  18.75ms   94.26%
    Req/Sec     2.76k   226.57     5.55k    73.73%
  1650655 requests in 1.00m, 748.18MB read
Requests/sec:  27464.60
Transfer/sec:     12.45MB
```

### Usage
```
Usage:
  slot-machine [OPTIONS]

Application Options:
  -v, --verbose       Show verbose debug information
  -l, --listen=       Listen on this interface (default: 127.0.0.1:8080)
      --debug.listen= Listen pprof on this interface (default: 127.0.0.1:6060)
      --jwt.key=      A signing key for JWT operations (default: dummySecret)

Help Options:
  -h, --help          Show this help message
```
