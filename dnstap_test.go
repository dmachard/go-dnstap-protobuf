package dnstap

import (
	"testing"

	"google.golang.org/protobuf/proto"
)

var dnstap_ref = []byte{10, 8, 100, 110, 115, 116, 97, 112, 112, 98, 18, 14, 100, 110, 115, 116, 97, 112,
	112, 98, 32, 48, 46, 48, 46, 48, 114, 90, 8, 5, 16, 1, 24, 1, 34, 4, 10, 0, 0, 210, 42, 4, 10, 0, 0,
	210, 48, 230, 174, 3, 56, 53, 64, 142, 143, 198, 255, 5, 77, 28, 102, 44, 21, 82, 53, 218, 186, 1, 32,
	0, 1, 0, 0, 0, 0, 0, 1, 3, 119, 119, 119, 6, 103, 111, 111, 103, 108, 101, 2, 102, 114, 0, 0, 1, 0, 1,
	0, 0, 41, 16, 0, 0, 0, 0, 0, 0, 12, 0, 10, 0, 8, 230, 56, 227, 142, 1, 222, 120, 120, 1}

func TestMarshal(t *testing.T) {
	// init
	dt := &Dnstap{}

	dt.Reset()

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
		t.Errorf("error on encode dnstap message %s", err)
	}
	if len(wiremessage) != len(dnstap_ref) {
		t.Errorf("size of the encoded message is different from reference")
	}
}

func TestUnmarshal(t *testing.T) {
	// init
	dt := &Dnstap{}

	// Unmarshal parses a wire-format message and places the decoded results in dt.
	err := proto.Unmarshal(dnstap_ref, dt)
	if err != nil {
		t.Errorf("error on decode dnstap message %s", err)
	}

	if string(dt.GetIdentity()) != "dnstappb" {
		t.Errorf("mismatch identity %s", string(dt.GetIdentity()))
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	// init
	dt := &Dnstap{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(dnstap_ref, dt)
		if err != nil {
			break
		}
	}
}

func BenchmarkMarshal(b *testing.B) {
	// init
	dt := &Dnstap{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dt.Reset()

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

		_, err := proto.Marshal(dt)
		if err != nil {
			break
		}
	}
}
