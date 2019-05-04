// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"

	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type ApiSnapshot struct {
	Gateways        GatewaysByNamespace
	VirtualServices VirtualServicesByNamespace
}

func (s ApiSnapshot) Clone() ApiSnapshot {
	return ApiSnapshot{
		Gateways:        s.Gateways.Clone(),
		VirtualServices: s.VirtualServices.Clone(),
	}
}

func (s ApiSnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashGateways(),
		s.hashVirtualServices(),
	)
}

func (s ApiSnapshot) hashGateways() uint64 {
	return hashutils.HashAll(s.Gateways.List().AsInterfaces()...)
}

func (s ApiSnapshot) hashVirtualServices() uint64 {
	return hashutils.HashAll(s.VirtualServices.List().AsInterfaces()...)
}

func (s ApiSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("gateways", s.hashGateways()))
	fields = append(fields, zap.Uint64("virtualServices", s.hashVirtualServices()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type ApiSnapshotStringer struct {
	Version         uint64
	Gateways        []string
	VirtualServices []string
}

func (ss ApiSnapshotStringer) String() string {
	s := fmt.Sprintf("ApiSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Gateways %v\n", len(ss.Gateways))
	for _, name := range ss.Gateways {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  VirtualServices %v\n", len(ss.VirtualServices))
	for _, name := range ss.VirtualServices {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s ApiSnapshot) Stringer() ApiSnapshotStringer {
	return ApiSnapshotStringer{
		Version:         s.Hash(),
		Gateways:        s.Gateways.List().NamespacesDotNames(),
		VirtualServices: s.VirtualServices.List().NamespacesDotNames(),
	}
}
