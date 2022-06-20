# Tests
Run these tests using 

```
go test -v
```

It will run tests and generate a report in a file named `results.json`

## Results
Results are printed into a `results.json` file

It looks like the following

```json
[
    {
        "method": "hmyv2_getBalanceByBlockNumber",
        "test": "genesis_addr_initial_balance",
        "pass": false,
        "duration": "885.094µs",
        "error": "0 != 10000000000000000000000000000 "
    },
    {
        "method": "hmyv2_getBalanceByBlockNumber",
        "test": "missing block",
        "pass": true,
        "duration": "327.28µs"
    }
]

```