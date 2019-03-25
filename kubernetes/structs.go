package kubernetes

import (
	"fmt"
)

type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type Resource struct {
	APIVersion string      `yaml:"apiVersion"`
	Kind       string      `yaml:"kind"`
	Metadata   *Metadata   `yaml:"metadata"`
	Items      []*Resource `yaml:"items,omitempty"`
}

// ID of the resource
func (r *Resource) ID() string {
	return fmt.Sprintf("%s/%s/%s",
		r.Metadata.Namespace,
		r.Kind,
		r.Metadata.Name,
	)
}

//Empty return true if resource was not correctly loaded
func (r *Resource) Empty() bool {
	return r.APIVersion == "" ||
		r.Kind == "" ||
		r.Metadata == nil
}