package plan

import (
	"fmt"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

func TestIsDestroyedCreate(t *testing.T) {
	// testing if the resources are deleted when i create plan is applied.
	file, err := os.Open("create.plan")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	plan, err := terraform.ReadPlan(file)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	isDestroyed, resourceName := IsDestroyingEBSVolume(plan)

	if isDestroyed != false {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", false, true)
	}

	fmt.Println(resourceName)
}

func TestIsDestroyedTainted(t *testing.T) {
	// testing if the resources are deleted when i create plan is applied.
	file, err := os.Open("tainted.plan")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	plan, err := terraform.ReadPlan(file)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	isDestroyed, resourceName := IsDestroyingEBSVolume(plan)

	if isDestroyed != false {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", false, true)
	}

	fmt.Println(resourceName)
}

func TestIsDestroyedModify(t *testing.T) {
	// testing if the resources are deleted when i create plan is applied.
	file, err := os.Open("modify.plan")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	plan, err := terraform.ReadPlan(file)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	isDestroyed, resourceName := IsDestroyingEBSVolume(plan)

	if isDestroyed != false {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", false, true)
	}

	fmt.Println(resourceName)
}

func TestIsDestroyedDestroy(t *testing.T) {
	// testing if the resources are deleted when i create plan is applied.
	file, err := os.Open("destroy.plan")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	plan, err := terraform.ReadPlan(file)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	isDestroyed, resourceName := IsDestroyingEBSVolume(plan)

	if isDestroyed != false {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", false, true)
	}

	fmt.Println(resourceName)
}

func TestIsDestroyedRecreate(t *testing.T) {
	// testing if the resources are deleted when i create plan is applied.
	file, err := os.Open("recreate.plan")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	plan, err := terraform.ReadPlan(file)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	isDestroyed, resourceName := IsDestroyingEBSVolume(plan)

	if isDestroyed != false {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", false, true)
	}

	fmt.Println(resourceName)
}
func TestIsDestroyedNochanges(t *testing.T) {
	// testing if the resources are deleted when i create plan is applied.
	file, err := os.Open("nochanges.plan")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	plan, err := terraform.ReadPlan(file)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	isDestroyed, resourceName := IsDestroyingEBSVolume(plan)

	if isDestroyed != false {
		t.Errorf("unexpected value exp=%+v\n act=%+v\n", false, true)
	}

	fmt.Println(resourceName)
}
