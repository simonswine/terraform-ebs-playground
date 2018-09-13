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
					if module.Path == nil || len(module.Path) == 1 {
						resourceNames = append(resourceNames, key)
					} else {
						modulePath := module.Path[1:len(module.Path)]
						resourceNames = append(resourceNames, "module."+strings.Join(modulePath, ".")+"."+key)
					}
					isDestroyed = true
				}
			}
		}
	}
	return isDestroyed, resourceNames
}
