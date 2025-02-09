<img src="https://img.shields.io/badge/go%20version-min%201.21-green" alt="Go version"/>

# go-dnstap-protobuf

Dnstap Protocol Buffers encoder and decoder implementation in Golang.

This library is based on the following
- [protocol buffers schema](https://raw.githubusercontent.com/dnstap/dnstap.pb/master/dnstap.proto)
- https://github.com/dnstap/dnstap.pb


## Installation

```go
go get -u github.com/dmachard/go-dnstap-protobuf
```

## Usage example

Example to use the Dnstap decoder

```go
var dnstap_wiremessage = []byte{10, 8, 100, 110, 115, 116, 97, 112, 112, 98, 18, 14, 100, 110, 115, 116, 97, 112,
                                112, 98, 32, 48, 46, 48, 46, 48, 114, 90, 8, 5, 16, 1, 24, 1, 34, 4, 10, 0, 0, 210, 42, 4, 10, 0, 0,
                                210, 48, 230, 174, 3, 56, 53, 64, 142, 143, 198, 255, 5, 77, 28, 102, 44, 21, 82, 53, 218, 186, 1, 32,
                                0, 1, 0, 0, 0, 0, 0, 1, 3, 119, 119, 119, 6, 103, 111, 111, 103, 108, 101, 2, 102, 114, 0, 0, 1, 0, 1,
                                0, 0, 41, 16, 0, 0, 0, 0, 0, 0, 12, 0, 10, 0, 8, 230, 56, 227, 142, 1, 222, 120, 120, 1}

dt := &dnstap.Dnstap{}
err := proto.Unmarshal(dnstap_ref, dt)
if err != nil {
    t.Errorf("error to decode dnstap message %s", err)
}

identity := dt.GetIdentity()
```

Example to use the Dnstap encoder

```go
dt := &dnstap.Dnstap{}

dt.Type = Dnstap_Type.Enum(1)
dt.Version = []byte("dnstappb 0.0.0")
dt.Identity = []byte("dnstappb")

dt.Message = &Message{}
dt.Message.Type = Message_Type.Enum(5)
dt.Message.SocketProtocol = SocketProtocol.Enum(1)
dt.Message.SocketFamily = SocketFamily.Enum(1)
dt.Message.QueryAddress = []byte("\n\x00\x00\xd2")
dt.Message.QueryPort = proto.Uint32(55142)
dt.Message.ResponseAddress = []byte("\n\x00\x00\xd2")
dt.Message.ResponsePort = proto.Uint32(53)
dt.Message.QueryTimeSec = proto.Uint64(1609664398)
dt.Message.QueryTimeNsec = proto.Uint32(355231260)
dt.Message.QueryMessage = []byte("ں\x01 \x00\x01\x00\x00\x00\x00\x00\x01\x03www\x06google\x02fr\x00\x00\x01\x00\x01\x00\x00)\x10\x00\x00\x00\x00\x00\x00\x0c\x00\n\x00\x08\xe68\xe3\x8e\x01\xdex")

wiremessage, err := proto.Marshal(dt)
if err != nil {
    t.Errorf("error to encode dnstap message %s", err)
}
```

## Testing

```bash
$ go test -v
=== RUN   TestMarshal
--- PASS: TestMarshal (0.00s)
=== RUN   TestUnmarshal
--- PASS: TestUnmarshal (0.00s)
PASS
ok      github.com/dmachard/go-dnstap-protobuf  0.003s
```

## Benchmark

```bash
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/dmachard/go-dnstap-protobuf
cpu: AMD EPYC 7763 64-Core Processor                
BenchmarkUnmarshal-4   	 2134670	       552.7 ns/op
BenchmarkMarshal-4     	 1889336	       601.4 ns/op
PASS
ok      github.com/dmachard/go-dnstap-protobuf  4.596s
```

# Development

Download some tools

```bash
wget https://raw.githubusercontent.com/dnstap/dnstap.pb/master/dnstap.proto
```

Edit and past the following line in the dnsmessage.proto

```bash
option go_package = "github.com/dmachard/go-dnstap-protobuf;dnstap";
```

Install protoc

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go mod tidy

export PROTOC_VER=29.3
export GITHUB_URL=https://github.com/protocolbuffers
wget $GITHUB_URL/protobuf/releases/download/v$PROTOC_VER/protoc-$PROTOC_VER-linux-x86_64.zip
unzip protoc-$PROTOC_VER-linux-x86_64.zip
```

Generate protobuf file

```bash
bin/protoc --proto_path=. --go_out=. --go_opt=paths=source_relative --plugin protoc-gen-go=$(go env GOPATH)/bin/protoc-gen-go dnstap.proto
```

Rename the `dnstap.pb.go` file to `dnstap.go`
