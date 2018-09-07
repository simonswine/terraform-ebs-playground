package plan

import (
	"fmt"
	"github.com/hashicorp/terraform/terraform"
)

func IsDestroyingEBSVolume(pl *terraform.Plan) (bool, []string) {
	var resourceNames []string

	for _, module := range pl.Diff.Modules {
		for key, resource := range module.Resources {
			fmt.Printf("%s: %+v %+v\n", key, resource.ChangeType(), resource.Attributes)
			switch resource.ChangeType() {
			case 5:
				return true, resourceNames
			case 4:
				return true, resourceNames
			case 3:
				return true, resourceNames
			case 2:
				return false, resourceNames
			}
		}
	}
	return false, []string{}
}
