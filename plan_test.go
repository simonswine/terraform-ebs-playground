package plan

import (
	"fmt"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

func OpenReadPlan(testCase string) *terraform.Plan {
	file, err := os.Open(testCase)
	if err != nil {
		fmt.Errorf("unexpected error %v", err)
	}

	plan, err := terraform.ReadPlan(file)
	if err != nil {
		fmt.Errorf("unexpected error %v", err)
	}

	return plan
}

func TestIsDestroyedCreate(t *testing.T) {
	isDestroyed, _ := IsDestroyingEBSVolume(OpenReadPlan("create.plan"))
	if isDestroyed != false {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", false, true)
	}
}

func TestIsDestroyedTainted(t *testing.T) {
	isDestroyed, _ := IsDestroyingEBSVolume(OpenReadPlan("tainted.plan"))

	if isDestroyed != true {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", false, true)
	}

}

func TestIsDestroyedModify(t *testing.T) {
	isDestroyed, _ := IsDestroyingEBSVolume(OpenReadPlan("modify.plan"))

	if isDestroyed != false {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", false, true)
	}

}

func TestIsDestroyedDestroy(t *testing.T) {
	isDestroyed, _ := IsDestroyingEBSVolume(OpenReadPlan("destroy.plan"))
	if isDestroyed != true {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", false, true)
	}
}

func TestIsDestroyedRecreate(t *testing.T) {
	isDestroyed, _ := IsDestroyingEBSVolume(OpenReadPlan("recreate.plan"))

	if isDestroyed != true {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", false, true)
	}

}
func TestIsDestroyedNochanges(t *testing.T) {
	isDestroyed, _ := IsDestroyingEBSVolume(OpenReadPlan("nochanges.plan"))

	if isDestroyed != false {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", false, true)
	}

}
