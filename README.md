# dns
A Golang DNS implementation developed for learning purpose.

## Building

The command below will build the program and create the executable file `dns`.

```bash
mkdir bin
go build -o bin/dns
```

## Avaiable Commands

### decode

The decode command prints the content of a dns message. The message may be provided
from the stdin or from a file.

#### reading from a file

```bash
dns --path ./message
```

#### reading from stdin

```bash
cat ./message | dns --stdin
```

#### retrieving a dns message

There are many ways you can get a DNS message, but a quick way to do so is using
dig and nc.

In a terminal, run the nc command:

```bash
nc -u -l -p 5553 > ./message
```

This will make nc to listen for a udp datagram on port 5553 and to write the datagram
content to a file called message.

In another terminal, run the dig command:

```bash
dig +retry=0 -p 5553 @127.0.0.1 +noedns www.google.com
```

This instructs dig to send a DNS query to a server in localhost port 5553 (or nc command).
