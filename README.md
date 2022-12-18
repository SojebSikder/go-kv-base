# go-kv-base

Simple key value in-memory database created with golang just for fun.

**Note: data stores in memory, so restarting the server will loose data**

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

Example:

```
set name sojeb
```

```
get name
```

```
delete name
```

```
flush
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

  - Get - read value by key
  - Set - set key value
  - Delete - delete value by key
  - Flush - delete all key value data
