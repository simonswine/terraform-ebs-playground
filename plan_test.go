package plan

import (
	"github.com/hashicorp/terraform/terraform"
	"os"
	"reflect"
	"testing"
)

var expResources = map[string]bool{
	"module.etcd.aws_ebs_volume.volume.0": false,
	"module.etcd.aws_ebs_volume.volume.1": false,
	"module.etcd.aws_ebs_volume.volume.2": false,
}

func openReadPlan(t *testing.T, testCase string) *terraform.Plan {
	file, err := os.Open(testCase)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	plan, err := terraform.ReadPlan(file)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	return plan
}

func TestIsDestroyedCreate(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "create.plan"))

	if exp, act := false, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) != 0 {
		t.Errorf("unexpected resourceName returned %+v", resourceName)
	}

}

func TestIsDestroyedTainted(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "tainted.plan"))

	if exp, act := true, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) != 1 {
		t.Errorf("unexpected resourceName returned %+v", resourceName)
	}

	if exp, act := []string{"module.etcd.aws_ebs_volume.volume.0"}, resourceName; !reflect.DeepEqual(exp, act) {
		t.Errorf("unexpected slice exp=%+v act=%+v", exp, act)
	}

}

func TestIsDestroyedModify(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "modify.plan"))

	if exp, act := false, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) != 0 {
		t.Errorf("unexpected resourceName returned %+v", resourceName)
	}

}

func TestIsDestroyedDestroy(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "destroy.plan"))

	if exp, act := true, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) != 3 {
		t.Errorf("unexpected resourceName returned %+v", resourceName)

	}

	for _, resource := range resourceName {
		if _, ok := expResources[resource]; ok {
			expResources[resource] = true
		} else {
			t.Errorf("unexpected slice act=%+v", resource)
		}
	}
}

func TestIsDestroyedRecreate(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "recreate.plan"))

	if exp, act := true, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) != 3 {
		t.Errorf("expected resourceName returned %+v", resourceName)
	}

	for _, resource := range resourceName {
		if _, ok := expResources[resource]; ok {
			expResources[resource] = true
		} else {
			t.Errorf("unexpected slice  act=%+v", resource)
		}
	}

}

func TestIsDestroyedNochanges(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "nochanges.plan"))

	if exp, act := false, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) != 0 {
		t.Errorf("expected resourceName returned %+v", resourceName)
	}
}

func TestIsDestroyedNonEbs(t *testing.T) {
	isDestroyed, resourceName := IsDestroyingEBSVolume(openReadPlan(t, "destroy_non_ebs.plan"))

	if exp, act := false, isDestroyed; exp != act {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", exp, act)
	}

	if len(resourceName) != 0 {
		t.Errorf("expected resourceName returned %+v", resourceName)
	}

}
