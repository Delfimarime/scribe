package config

type ProjectDefinition struct {
	Id                string            `yaml:"id,omitempty"`
	Solution          *Solution         `yaml:"solution"`
	ServiceDefinition ServiceDefinition `yaml:"service"`
	Name              string            `yaml:"name,omitempty"`
	Version           string            `yaml:"version,omitempty"`
	ShortDescription  string            `yaml:"short_description,omitempty"`
}

type Solution struct {
	Id string `yaml:"id"`
}

type ServiceDefinition struct {
	Type          string                  `yaml:"type,omitempty"`
	Language      string                  `yaml:"language,omitempty"`
	Runtime       string                  `yaml:"runtime,omitempty"`
	DependsOn     []DependencyDefinition  `yaml:"depends_on,omitempty"`
	Environment   []EnvironmentDefinition `yaml:"environment,omitempty"`
	Specification []byte                  `yaml:"specification,omitempty"`
}

type DependencyDefinition struct {
	Taggable
	Id            string                 `yaml:"id,omitempty"`
	Type          string                 `yaml:"type,omitempty"`
	Name          string                 `yaml:"name,omitempty"`
	DependsOn     []DependencyDefinition `yaml:"depends_on,omitempty"`
	ShortInfo     string                 `yaml:"short_info,omitempty"`
	Specification []byte                 `yaml:"specification,omitempty"`
}

type Taggable struct {
	Tags []TagDefinition `yaml:"tags,omitempty"`
}

type TagDefinition struct {
	Name  string `yaml:"name,omitempty"`
	Value string `yaml:"value,omitempty"`
}

type EnvironmentDefinition struct {
	Name       string `yaml:"name,omitempty"`
	DefaultsTo string `yaml:"defaults_to,omitempty"`
	Type       string `yaml:"type,omitempty"`
}
