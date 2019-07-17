package node

import "github.com/romarq/go-tezos-tezaria/block"

type TezosNodeService interface {
	MonitorHeads(chain string, heads chan block.Header, errc chan error, done chan bool)
	Bootstrapped() (Bootstrap, error)
	CommitHash() (string, error)
}
