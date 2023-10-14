package context

import (
	"fmt"
	"github.com/delfimarime/scribe/pkg/config"
)

func DoResolve(definition config.ProjectDefinition) (config.ProjectDefinition, error) {
	seenIDs := make(map[string]bool)
	for index, dep := range definition.ServiceDefinition.DependsOn {
		if dep.Id == "" {
			dep.Id = dep.Name
		}
		if _, hasValue := seenIDs[dep.Id]; hasValue {
			return config.ProjectDefinition{}, fmt.Errorf(`service.depends_on[%d].id="%s" cannot be used twice`, index, dep.Id)
		}
		seenIDs[dep.Id] = true
		definition.ServiceDefinition.DependsOn[index] = dep
	}
	findById := func(i, j int, id string) (*config.DependencyDefinition, error) {
		for _, dep := range definition.ServiceDefinition.DependsOn {
			if dep.Id == id {
				return &dep, nil
			}
		}
		return nil, fmt.Errorf(`service.depends_on[%d].depends_on[%d].id="%s" not found`, i, j, id)
	}
	for index, dep := range definition.ServiceDefinition.DependsOn {
		for dIndex, dOn := range dep.DependsOn {
			resolvedDep, err := findById(index, dIndex, dOn.Id)
			if err != nil {
				return config.ProjectDefinition{}, err
			}
			definition.ServiceDefinition.DependsOn[index].DependsOn[dIndex] = *resolvedDep
		}
	}
	return definition, nil
}
