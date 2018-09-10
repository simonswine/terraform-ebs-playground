package plan

import (
	"github.com/hashicorp/terraform/terraform"
	"strings"
)

const (
	DiffInvalid = iota
	DiffNone
	DiffCreate
	DiffUpdate
	DiffDestroy
	DiffDestroyCreate
)

func IsDestroyingEBSVolume(pl *terraform.Plan) (bool, []string) {
	var resourceNames []string
	isDestroyed := false
	//fmt.Printf("%s: %+v %+v\n", key, resource.ChangeType(), resource.Attributes)

	for _, module := range pl.Diff.Modules {
		for key, resource := range module.Resources {
			//	fmt.Printf("%s: %+v %+v\n", key, resource.ChangeType(), resource.Attributes)
			switch resource.ChangeType() {
			case DiffDestroy, DiffDestroyCreate:
				if strings.Contains(strings.SplitAfter(key, ":")[0], "aws_ebs") {
					resourceNames = append(resourceNames, strings.SplitAfter(key, ":")[0])
					isDestroyed = true
				}
			}
		}
	}
	return isDestroyed, resourceNames
}
