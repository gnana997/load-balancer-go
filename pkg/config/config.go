package config

// Config is a representation of the configuration given to load balancer from a config source
type Config struct {
	Services []Service `yaml:"services"`

	Strategy string `yaml:"defaultStrategy"` // Name of the strategy for load balancing between the replicas
}

type Replica struct {
	Host     string            `yaml:"host"`
	Metadata map[string]string `yaml:"metadata"`
}

type Service struct {
	Name string `yaml:"name"` // name of the serive
	// can be subdomain or regex
	Matcher  string    `yaml:"matcher"`  // prefix of the url to match the service
	Replicas []Replica `yaml:"replicas"` // replicas of the service like all the ips of the service
	Strategy string    `yaml:"strategy"` // name of the strategy for load balancing between the replicas
}
