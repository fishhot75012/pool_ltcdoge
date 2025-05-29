# pool_ltcdoge
Test Pool Merged Mining Litecoin Dogecoin
Schema of file structure
stratum/
├── main.go
├── connection.go
├── handlers.go
├── job_generator.go
├── mining_authorize.go
├── mining_notify.go
├── mining_submit.go
├── utils/
│   ├── block.go
│   └── rpc.go
├── db/
│   ├── db.go
│   └── schema.sql
├── tests/
│   └── stratum_flow_test.go
Part 1: the first functional brick of the Stratum server, integrated into a pool structure adapted to Litecoin and Dogecoin 

pool_ltcdoge/
├── go.mod
├── main.go
├── stratum/
│   ├── server.go
│   ├── handler.go
│   ├── notify.go
│   ├── submit.go
│   └── jobs.go
├── rpc/
│   └── client.go
├── utils/
│   └── block.go
├── database/
│   └── postgres.go
├── sql/
│   └── init.sql
└── certs/
    ├── cert.pem
    └── key.pem
