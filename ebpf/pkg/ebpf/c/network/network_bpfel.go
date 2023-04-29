// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package network

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type networkEventData struct {
	Ipv4 struct {
		Ip   uint32
		Desc [10]uint8
		_    [2]byte
	}
}

// loadNetwork returns the embedded CollectionSpec for network.
func loadNetwork() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_NetworkBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load network: %w", err)
	}

	return spec, err
}

// loadNetworkObjects loads network and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*networkObjects
//	*networkPrograms
//	*networkMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadNetworkObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadNetwork()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// networkSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type networkSpecs struct {
	networkProgramSpecs
	networkMapSpecs
}

// networkSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type networkProgramSpecs struct {
	XdpProgFunc *ebpf.ProgramSpec `ebpf:"xdp_prog_func"`
}

// networkMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type networkMapSpecs struct {
	Event *ebpf.MapSpec `ebpf:"event"`
}

// networkObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadNetworkObjects or ebpf.CollectionSpec.LoadAndAssign.
type networkObjects struct {
	networkPrograms
	networkMaps
}

func (o *networkObjects) Close() error {
	return _NetworkClose(
		&o.networkPrograms,
		&o.networkMaps,
	)
}

// networkMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadNetworkObjects or ebpf.CollectionSpec.LoadAndAssign.
type networkMaps struct {
	Event *ebpf.Map `ebpf:"event"`
}

func (m *networkMaps) Close() error {
	return _NetworkClose(
		m.Event,
	)
}

// networkPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadNetworkObjects or ebpf.CollectionSpec.LoadAndAssign.
type networkPrograms struct {
	XdpProgFunc *ebpf.Program `ebpf:"xdp_prog_func"`
}

func (p *networkPrograms) Close() error {
	return _NetworkClose(
		p.XdpProgFunc,
	)
}

func _NetworkClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed network_bpfel.o
var _NetworkBytes []byte
