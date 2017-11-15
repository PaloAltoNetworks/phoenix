package phoenix

import (
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/rufusmodels/v1/golang"
)

// HookFunc is the type of a function to that can be used as a Hook.
type HookFunc func(string, elemental.Operation, rufusmodels.RemoteProcessorModeValue, elemental.Identifiable, []string) error

// HooksRegistry represents a list of HookFunc.
type HooksRegistry []HookFunc

// NewHooksRegistry returns a HooksRegistry with the given list of HookFunc.
func NewHooksRegistry(plugins ...HookFunc) HooksRegistry {
	return append(HooksRegistry{}, plugins...)
}
