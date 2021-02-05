# person-svc-go


## Local test results:


Run that connects to Postgres and queries for persons and returns the results.

```
wrk -c 100 -d 60 -t 8 http://localhost:8080/persons
Running 1m test @ http://localhost:8080/persons
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.48ms    3.90ms  45.51ms   63.03%
    Req/Sec     3.02k   157.93     4.21k    68.25%
  1442493 requests in 1.00m, 537.89MB read
Requests/sec:  24031.57
Transfer/sec:      8.96MB
```


Run that calles a status enpoint with a static JSON response.

```
wrk -c 100 -d 60 -t 8 http://localhost:8080/status
Running 1m test @ http://localhost:8080/status
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.17ms    3.11ms  57.00ms   87.35%
    Req/Sec     9.80k     2.23k   23.82k    68.21%
  4682922 requests in 1.00m, 553.78MB read
Requests/sec:  77947.51
Transfer/sec:      9.22MB
```
