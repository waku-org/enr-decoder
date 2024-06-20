# Waku ENR Decoder
Follows https://github.com/waku-org/specs/blob/master/standards/core/enr.md and https://github.com/waku-org/specs/blob/master/standards/core/relay-sharding.md to decode Waku specific fields.

### To build the binary:

go build -o waku-enr-decoder main.go

### To run the binary:
./waku-enr-decoder --enr="<valid ENR>"
