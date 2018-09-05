package main

import (
	"fmt"
	"github.com/hashicorp/terraform/terraform"
	"os"
)

func main() {
	file, err := os.Open("enablekms.plan")
	if err != nil {
		panic(err)
	}

	plan, err := terraform.ReadPlan(file)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("plan: %s", plan.String())

	for _, module := range plan.Diff.Modules {
		for key, resource := range module.Resources {
			fmt.Printf("%s: %+v %+v\n", key, resource.ChangeType(), resource.Attributes)
		}
	}

}
