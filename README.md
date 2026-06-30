# RealEstateToken

A proof-of-concept platform for **tokenizing physical real estate on Ethereum**:
each property on the map is minted as its own ERC-20 token, and those tokens can
be held and transferred between user wallets. Built in Go with a `go-ethereum`
backend, a Solidity token contract, and a Bing Maps property browser.

> **Context / status.** This is a 2018–2019 learning project — my deep dive into
> Ethereum, smart contracts, and wallet mechanics. It runs, but it targets the
> Solidity `0.4.x` era and is not production-hardened. It's published as a record
> of the architecture and what I learned, not as a maintained library. The direct
> precursor was [`simple-blockchain`](https://github.com/MichaelPGifford/simple-blockchain),
> where I built a blockchain from scratch to decide whether to roll my own chain
> or issue tokens on an existing one — this repo is the "issue tokens on Ethereum"
> conclusion of that experiment.

## What it does

- **One token per property.** Adding a property deploys a fresh ERC-20 contract
  (its own symbol/name) representing fractional ownership of that listing.
- **Map-first browsing.** Properties render as pins on a Bing Maps view; an admin
  flow adds new properties/pins.
- **Wallets & transfers.** Users sign up, get an Ethereum wallet, check their ETH
  balance, and send ETH — all brokered server-side through `go-ethereum`.
- **Auth + identity.** Sign-up / sign-in with a basic ID-verification step,
  persisted in MongoDB.

## Architecture

```
main.go                     httprouter server: pages + REST API
├── handlers/               server-rendered pages (homepage, balance, admin)
├── api/                    JSON endpoints
│   ├── sign_up / sign_in   auth
│   ├── generate_address    wallet creation
│   ├── get_ether_balance   chain reads
│   ├── send_ether          signed transactions
│   ├── map_pins            property pins for the map
│   └── verify_id
├── tokens/                 contract lifecycle (deploy / fetch / update)
│   └── contracts/          Solidity sources + generated Go bindings
├── user/ , user/wallet/    user model + Ethereum wallet/transfer logic
└── mongodb/                persistence
```

**Stack:** Go (`julienschmidt/httprouter`), `go-ethereum`, Solidity `^0.4.24`,
MongoDB, Bing Maps, Materialize CSS.

## Attribution

The ERC-20 base contract (`tokens/contracts/property.sol`) is adapted from
BokkyPooBah's widely-used fixed-supply token template (MIT, credited in the file
header). The original work in this repo is the Go platform around it: per-property
contract deployment, wallet handling, ETH transfers, the map UI, and auth.

## Running it (historical)

This depends on a synced Ethereum node / endpoint and a funded account, plus a
MongoDB instance, configured for the Solidity `0.4.x` toolchain of the era. It is
not expected to run unmodified against today's networks.

```bash
go run main.go   # serves on :8080
```

## What I took away from it

- How ERC-20 transfers, allowances, and `approve`/`transferFrom` actually work at
  the contract level — and where the footguns are (e.g. the approval double-spend
  the EIP itself calls out).
- Driving contract deployment and signed transactions from a backend with
  `go-ethereum` rather than from the browser.
- Why "one contract per asset" is elegant conceptually but expensive in gas and
  operational complexity at scale — a tradeoff I'd weigh very differently today.

## License

MIT
