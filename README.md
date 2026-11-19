# gcfbackend

## Testing

Run `go test ./...` for the default test suite. MongoDB integration tests are skipped unless explicitly enabled.

To opt in, use a disposable MongoDB instance and a database name beginning with `test_`, never a production connection or database:

```bash
export GCFBACKEND_RUN_INTEGRATION=1
export GCFBACKEND_TEST_MONGO='mongodb://localhost:27017'
export GCFBACKEND_TEST_DB='test_gis'
go test ./...
```

These integration tests insert, modify and delete records in the selected test database. Use an isolated server with test-only credentials.
