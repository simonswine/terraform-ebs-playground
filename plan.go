package plan

import (
	"github.com/hashicorp/terraform/terraform"
	"strings"
)

func IsDestroyingEBSVolume(pl *terraform.Plan) (bool, []string) {
	var resourceNames []string
	isDestroyed := false
	//fmt.Printf("%s: %+v %+v\n", key, resource.ChangeType(), resource.Attributes)

	for _, module := range pl.Diff.Modules {
		for key, resource := range module.Resources {
			switch resource.ChangeType() {
			case terraform.DiffDestroy, terraform.DiffDestroyCreate:
				if strings.Split(key, ".")[0] == "aws_ebs_volume" {
					resourceNames = append(resourceNames, key)
					isDestroyed = true
				}
			}
		}
	}
	return isDestroyed, resourceNames
}
