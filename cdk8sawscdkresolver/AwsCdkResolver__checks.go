//go:build !no_runtime_type_checking

package cdk8sawscdkresolver

import (
	"fmt"

	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func (a *jsiiProxy_AwsCdkResolver) validateResolveParameters(context cdk8s.ResolutionContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

