package project

var (
	description = "The cluster-api-provider-kvm does something."
	gitSHA      = "n/a"
	name        = "cluster-api-provider-kvm"
	source      = "https://github.com/giantswarm/cluster-api-provider-kvm"
	version     = "0.1.0-dev"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
