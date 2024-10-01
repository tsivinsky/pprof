# pprof

## slow version

```
real    0m15.695s
user    0m12.612s
sys     0m11.064s
```

```bash
go run . foo
go tool pprof -top mem.pprof
```

## fast(er) version

```
real    0m0.833s
user    0m0.015s
sys     0m0.110s
```

```bash
go run . bar
go tool pprof -top mem.pprof
```
