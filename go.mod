module github.com/giantswarm/cluster-api-provider-kvm

go 1.16

require (
	github.com/giantswarm/microkit v0.2.2
	github.com/giantswarm/operatorkit/v4 v4.3.1
	github.com/go-logr/logr v1.2.0
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.18.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.1
	github.com/spf13/pflag v1.0.5
	k8s.io/apimachinery v0.24.2
	k8s.io/client-go v0.24.2
	k8s.io/klog v1.0.0
	k8s.io/klog/v2 v2.60.1
	sigs.k8s.io/cluster-api v1.2.0-beta.2
	sigs.k8s.io/controller-runtime v0.12.2
)

replace (
	github.com/coreos/etcd v3.3.10+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
)
