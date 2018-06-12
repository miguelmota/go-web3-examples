# Go Web3 Examples

> Example of how to use "Web3" in golang.

---

## UPDATE

Check out my little book on [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) as a better guide.

---

## Synopsis

Shows how to

- compile a smart contract
- generate a golang package from the ABI
- connect to rpc or websocket provider
- load a contract from it's address
- load a private key
- call a contract method
- subscribe to contract events
- decode event log

## Development

Solidity contract to ABI

```bash
solc --abi contracts/Greeter.sol
```

ABI to Go package

```bash
abigen --abi contracts/Greeter.abi --pkg greeter --out greeter.go
```

## Resources

- [Native DApps: Go bindings to Ethereum contracts](https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts)

## License

MIT
