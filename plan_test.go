package plan

import (
	"github.com/hashicorp/terraform/terraform"
	"os"
	"strings"
	"testing"
)

func openReadPlan(t *testing.T, testCase string) *terraform.Plan {
	file, err := os.Open(testCase)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	plan, err := terraform.ReadPlan(file)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	return plan
}

func TestIsDestroyedCreate(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "create.plan"))

	if exp, act := false, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) != 0 {
		t.Errorf("unexpected value")
	}

}

func TestIsDestroyedTainted(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "tainted.plan"))

	if exp, act := true, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) == 0 {
		t.Errorf("unexpected value no resource to add")
	}

	for _, resource := range resourceName {
		if strings.Split(resource, ".")[0] != "aws_ebs_volume" {
			t.Errorf("expected ebs resource but recieved=%s", resource)
		}
	}

}

func TestIsDestroyedModify(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "modify.plan"))

	if exp, act := false, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) != 0 {
		t.Errorf("unexpected value")
	}

}

func TestIsDestroyedDestroy(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "destroy.plan"))

	if exp, act := true, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) == 0 {
		t.Errorf("unexpected value")
	}

	for _, resource := range resourceName {
		if strings.Split(resource, ".")[0] != "aws_ebs_volume" {
			t.Errorf("expected ebs resource but received=%s", resource)
		}
	}

}

func TestIsDestroyedRecreate(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "recreate.plan"))

	if exp, act := true, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) == 0 {
		t.Errorf("expected value")
	}

	for _, resource := range resourceName {
		if strings.Split(resource, ".")[0] != "aws_ebs_volume" {
			t.Errorf("expected ebs resource but received=%s", resource)
		}
	}

}

func TestIsDestroyedNochanges(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "nochanges.plan"))

	if exp, act := false, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) != 0 {
		t.Errorf("unexpected value")
	}
}

func TestIsDestroyedNonEbs(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "destroy_non_ebs.plan"))

	if exp, act := false, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) != 0 {
		t.Errorf("unexpected value")
	}

}
