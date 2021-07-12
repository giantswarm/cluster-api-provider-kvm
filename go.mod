module github.com/giantswarm/cluster-api-provider-kvm

go 1.16

require (
	github.com/giantswarm/exporterkit v0.2.1
	github.com/giantswarm/microendpoint v0.2.0
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/microkit v0.2.2
	github.com/giantswarm/micrologger v0.5.0
	github.com/giantswarm/operatorkit/v4 v4.3.1
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/viper v1.8.1
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
	sigs.k8s.io/cluster-api v0.4.0-beta.1
	sigs.k8s.io/controller-runtime v0.9.2
)

replace (
	github.com/coreos/etcd v3.3.10+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
)
