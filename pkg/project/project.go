package project

var (
	description string = ""
	gitSHA             = "n/a"
	name        string = "cluster-api-core-app"
	source      string = "https://github.com/giantswarm/cluster-api-core-app"
	version            = "0.0.0"
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
