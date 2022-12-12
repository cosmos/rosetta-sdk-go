# cosmos/rosetta-sdk-go

This repo is a fork of [coinbase/rosetta-sdk-go](https://github.com/coinbase/rosetta-sdk-go), limited to the package required by the Cosmos SDK.

- asserter (unchanged)
- errors (unchanged)
- server (unchanged)

The Cosmos SDK Team forked this repo to avoid importing the entire `rosetta-sdk-go` package, which includes a lot of dependencies that are not required by the SDK.
