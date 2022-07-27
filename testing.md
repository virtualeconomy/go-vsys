Timeout needed to avoid default 10min limit for running tests. With -run flag you can supply regexp for tests to run.
```
go test ./vsys/ -v -timeout 0 -run AsWhole
```