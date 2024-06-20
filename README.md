# Waku ENR Decoder
Follows https://github.com/waku-org/specs/blob/master/standards/core/enr.md and https://github.com/waku-org/specs/blob/master/standards/core/relay-sharding.md to decode Waku specific fields.

### To build the binary:

`go build -o waku-enr-decoder main.go`

### To run the binary:
`./waku-enr-decoder --enr="<valid ENR>"`

Sample Usage 

`./waku-enr-decoder--enr="enr:-QEGuEAfJdcVdiAFr_e79ilATxUUGDouAqrvvEWEINUioIzTYRgsfYAYfxUlSLCs95w1NrFe7y2Wp54WlN6V6V_j31q3hgGQM9Ki8IJpZIJ2NIJpcITAqAEJim11bHRpYWRkcnO4VgBUNiVib290LTAxLmRvLWFtczMuc2hhcmRzLnRlc3Quc3RhdHVzLmltBnZfpQMnACUIAhIhAt60bRUEoHNuLlnsM12sU2PIQwBwfLIJ8a_ZPEY2-RnkgnJzhwAQAgBAACCJc2VjcDI1NmsxoQJ5lIphhEl668B0TOuIzNRyTUJCA_s4xvhAvQVqNa7CYYN0Y3CC2zKDdWRwguqLhXdha3UyDQ"`

```
Decoded ENR:
seq: 1718856360688
signature: 0x1f25d715762005aff7bbf629404f1514183a2e02aaefbc458420d522a08cd361182c7d80187f152548b0acf79c3536b15eef2d96a79e1694de95e95fe3df5ab7
peer-id:  16Uiu2HAm3cGgUdycLBx3zeKiJnVaLWVgPwcJYM8LSgiJh7fyiPYU
ipv4: 192.168.1.9
tcp: 56114
cluster-id:  16
shards:  [64 32]
Wakuv2 Protocols Supported:
/vac/waku/relay/2.0.0
/vac/waku/filter-subscribe/2.0.0-beta1
/vac/waku/lightpush/2.0.0-beta1
multiaddresses:
/ip4/192.168.1.9/tcp/56114/p2p/16Uiu2HAm3cGgUdycLBx3zeKiJnVaLWVgPwcJYM8LSgiJh7fyiPYU
/dns4/boot-01.do-ams3.shards.test.status.im/tcp/30303/p2p/16Uiu2HAmAR24Mbb6VuzoyUiGx42UenDkshENVDj4qnmmbabLvo31/p2p-circuit/p2p/16Uiu2HAm3cGgUdycLBx3zeKiJnVaLWVgPwcJYM8LSgiJh7fyiPYU```

