package workswap

import (
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"io/ioutil"
)

const ID = "/ipfs/ping/1.0.0"

type Service struct {
}

func NewService(h host.Host) *Service {
	s := &Service{}
	h.SetStreamHandler(ID, s.Handler)
	return s
}

func (t *Service) Handler(s network.Stream) {
	data, err := ioutil.ReadAll(s)
	if err != nil {
		panic(err)
	}
	println(string(data))
}
