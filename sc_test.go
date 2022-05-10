package hwmm_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nobishino/hwmm"
)

func TestSCMachine(t *testing.T) {
	m := hwmm.NewMachine()

	m.AddThread(
		hwmm.Write("x", 1), hwmm.Write("y", 2),
	)
	result := m.Run()
	want := hwmm.Result{
		hwmm.SharedMemory{"x": 1, "y": 2},
	}
	if diff := cmp.Diff(result, want); diff != "" {
		t.Error(diff)
	}
}
