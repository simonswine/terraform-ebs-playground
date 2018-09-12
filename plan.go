package plan

import (
	"github.com/hashicorp/terraform/terraform"
	"strings"
)

func IsDestroyingEBSVolume(pl *terraform.Plan) (bool, []string) {
	var resourceNames []string
	isDestroyed := false

	for _, module := range pl.Diff.Modules {
		for key, resource := range module.Resources {
			switch resource.ChangeType() {
			case terraform.DiffDestroy, terraform.DiffDestroyCreate:
				if strings.Split(key, ".")[0] == "aws_ebs_volume" {

					modulePath := "module." + module.Path[1] + "."
					resourceNames = append(resourceNames, modulePath+key)
					isDestroyed = true
				}
			}
		}
	}
	return isDestroyed, resourceNames
}
