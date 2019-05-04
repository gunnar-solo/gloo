// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type TranslatorSnapshot struct {
	Secrets   gloo_solo_io.SecretsByNamespace
	Upstreams gloo_solo_io.UpstreamsByNamespace
	Ingresses IngressesByNamespace
}

func (s TranslatorSnapshot) Clone() TranslatorSnapshot {
	return TranslatorSnapshot{
		Secrets:   s.Secrets.Clone(),
		Upstreams: s.Upstreams.Clone(),
		Ingresses: s.Ingresses.Clone(),
	}
}

func (s TranslatorSnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashSecrets(),
		s.hashUpstreams(),
		s.hashIngresses(),
	)
}

func (s TranslatorSnapshot) hashSecrets() uint64 {
	return hashutils.HashAll(s.Secrets.List().AsInterfaces()...)
}

func (s TranslatorSnapshot) hashUpstreams() uint64 {
	return hashutils.HashAll(s.Upstreams.List().AsInterfaces()...)
}

func (s TranslatorSnapshot) hashIngresses() uint64 {
	return hashutils.HashAll(s.Ingresses.List().AsInterfaces()...)
}

func (s TranslatorSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("secrets", s.hashSecrets()))
	fields = append(fields, zap.Uint64("upstreams", s.hashUpstreams()))
	fields = append(fields, zap.Uint64("ingresses", s.hashIngresses()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type TranslatorSnapshotStringer struct {
	Version   uint64
	Secrets   []string
	Upstreams []string
	Ingresses []string
}

func (ss TranslatorSnapshotStringer) String() string {
	s := fmt.Sprintf("TranslatorSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Secrets %v\n", len(ss.Secrets))
	for _, name := range ss.Secrets {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Ingresses %v\n", len(ss.Ingresses))
	for _, name := range ss.Ingresses {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s TranslatorSnapshot) Stringer() TranslatorSnapshotStringer {
	return TranslatorSnapshotStringer{
		Version:   s.Hash(),
		Secrets:   s.Secrets.List().NamespacesDotNames(),
		Upstreams: s.Upstreams.List().NamespacesDotNames(),
		Ingresses: s.Ingresses.List().NamespacesDotNames(),
	}
}
