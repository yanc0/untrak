package outputs

import (
	"fmt"

	"github.com/yanc0/kuntrak/kubernetes"
)

// Text output resources as text
func Text(resources []*kubernetes.Resource) {
	var output string
	for _, r := range resources {
		out := r.ID()
		output += fmt.Sprintf("- %s\n", out)
	}
	fmt.Println(output)
}
