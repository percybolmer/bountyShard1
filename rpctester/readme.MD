# Deploy smart contracts
To deploy needed smart contracts first use the Deploy command

```bash
go build -o rpctester
./rpctester deploy
```


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

## Known Issues
Filter Logs are always empty no matter what
Created filter sometimes does not exist, even after a 2 second delay