package outputs

import (
	"fmt"

	"github.com/yanc0/kuntrak/kubernetes"
	yaml "gopkg.in/yaml.v2"
)

// YAML outputs resources as YAML
func YAML(resources []*kubernetes.Resource) {
	var output string
	for _, r := range resources {
		out, err := yaml.Marshal(r)
		if err != nil {
			panic(err)
		}
		output += fmt.Sprintf("---\n%s\n", string(out))
	}
	fmt.Println(output)
}
