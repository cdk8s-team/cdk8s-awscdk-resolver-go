//go:build no_runtime_type_checking

package cdk8sawscdkresolver

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsCdkResolver) validateResolveParameters(context cdk8s.ResolutionContext) error {
	return nil
}

