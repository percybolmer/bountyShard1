# Deploy smart contracts
To deploy needed smart contracts first use the Deploy command

```bash
go build -o rpctester
./rpctester deploy
```


# Tests
Run these tests using 

Make sure you take the address from the deployed smart contract into the `.env` file  
in the `SMART_CONTRACT_ADDRESS` variabel

```
cd harmony && go test -v
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

Some Methods does not exist
hmyv2_getPendingCrossLinks

TransactionHistory is always empty on Shard1

# Questions to harmony
How to perform a SendRawStakingTransaction

Diffrent endpoitns for getTransactionHistory in docs, but not true

