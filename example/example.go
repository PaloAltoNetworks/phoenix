package main

import (
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/rufusmodels/v1/golang"
	"github.com/aporeto-inc/gaia/squallmodels/v1/golang"
)

func exampleHookFunc(requestIdentifier string, op elemental.Operation, mode rufusmodels.RemoteProcessorModeValue, obj elemental.Identifiable, claims []string) error {

	if o, ok := obj.(*squallmodels.Namespace); ok && o.Description == "" {
		o.Description = "Y U NO put description?"
	}

	return nil
}
