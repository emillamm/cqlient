# cqlient
Terminal client for sending CQL commands to Cassandra.

## Installation
First make sure that GOPATH is set. Then run `go install github.com/emillamm/cqlient@latest`.

## Usage
```
Flags:
  -command string
    	CQL command to execute
  -host string
    	Cassandra host. Defaults to localhost. (default "localhost")
  -keyspace string
    	Cassandra Keyspace
  -pass string
    	Password - will not use authentication if not set
  -port int
    	Cassandra host port. Defaults to 9042. (default 9042)
  -user string
    	Username - will not use authentication if not set
```

