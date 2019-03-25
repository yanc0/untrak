package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/yanc0/kuntrak/outputs"
	"github.com/yanc0/kuntrak/utils"
	yaml "gopkg.in/yaml.v2"

	"github.com/yanc0/kuntrak/kubernetes"

	"github.com/yanc0/kuntrak/config"
)

func main() {
	// Flags, command line parameters
	var cfgPathOpt = flag.String("config", "./kuntrak.yaml", "Kuntrak Config Path")
	var outputOpt = flag.String("o", "text", "Output format")
	flag.Parse()

	var wg sync.WaitGroup
	var resourcesIn []*kubernetes.Resource
	var resourcesOut []*kubernetes.Resource

	// Config Load
	cfg, err := config.Load(*cfgPathOpt)
	if err != nil {
		fmt.Printf("[ERR] Cannot load %s file: %v\n", *cfgPathOpt, err)
		os.Exit(1)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		resourcesIn, err = getKubernetesResources(cfg.In)
		if err != nil {
			fmt.Printf("[ERR] Failed to get Kubernetes resources (in): %v\n", err)
			os.Exit(1)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		resourcesOut, err = getKubernetesResources(cfg.Out)
		if err != nil {
			fmt.Printf("[ERR] Failed to get Kubernetes resources (out): %v\n", err)
			os.Exit(1)
		}
	}()

	wg.Wait()

	untrackedResources := listUntrackedResources(resourcesIn, resourcesOut, cfg.Exclude)
	switch {
	case *outputOpt == "text":
		outputs.Text(untrackedResources)
	case *outputOpt == "yaml":
		outputs.YAML(untrackedResources)
	default:
		outputs.Text(untrackedResources)
	}
}

func getKubernetesResources(cfgs []*config.CommandConfig) ([]*kubernetes.Resource, error) {
	const yamlSeparator = "---\n"
	var resources []*kubernetes.Resource

	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	for _, cfg := range cfgs {
		wg.Add(1)
		go func(cmd string, args ...string) {
			defer wg.Done()
			stdout, stderr, err := utils.Exec(cmd, args...)
			if err != nil {
				fmt.Println(cmd, args)
				fmt.Println(string(stderr))
				fmt.Println(err)
				os.Exit(1)
			}

			for _, yml := range strings.Split(string(stdout), yamlSeparator) {
				resource := kubernetes.Resource{}

				yaml.Unmarshal([]byte(yml), &resource)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				if resource.Empty() {
					// Failed to ummarshal properly
					continue
				}
				mutex.Lock()
				// When 'kubectl get' is used, an object List is returned with
				// all resources in items array
				if resource.Kind == "List" && resource.Items != nil {
					for _, item := range resource.Items {
						resources = append(resources, item)
					}
				} else {
					resources = append(resources, &resource)
				}
				mutex.Unlock()
			}

		}(cfg.Cmd, cfg.Args...)
	}
	wg.Wait()
	return resources, nil
}

func listUntrackedResources(in []*kubernetes.Resource, out []*kubernetes.Resource, kindExclude []string) []*kubernetes.Resource {
	var untrackedResources []*kubernetes.Resource
	for _, resourceOut := range out {
		if utils.StringInListCaseInsensitive(kindExclude, resourceOut.Kind) {
			continue
		}
		found := false
		for _, resourceIn := range in {
			if resourceOut.ID() == resourceIn.ID() {
				found = true
				break
			}
		}
		if !found {
			untrackedResources = append(untrackedResources, resourceOut)
		}
	}

	return untrackedResources
}
