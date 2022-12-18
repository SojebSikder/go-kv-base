# go-kv-base

Simple key value database created with golang just for fun.

# Usage

## Server

Start server:

```
go run main.go start-server <port>
```

## Client

### Start cli:

```
go run main.go cli <server_url>
```

### Api:

POST http://host:port

Set command:

```bash
curl \
'http://host:port' \
-d '{"key":"name","value":"sojeb","command":"set"}'
```

Get command:

```bash
curl \
'http://host:port' \
-d '{"key":"name","command":"get"}'
```

Delete command:

```bash
curl \
'http://host:port' \
-d '{"key":"name","command":"delete"}'
```

Flush command:

```bash
curl \
'http://host:port' \
-d '{"command":"flush"}'
```

## Supported commands:

- Database oparations

  - Get
  - Set
  - Delete
  - Flush
