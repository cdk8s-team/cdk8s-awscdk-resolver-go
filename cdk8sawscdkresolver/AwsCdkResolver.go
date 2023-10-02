package cdk8sawscdkresolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-awscdk-resolver-go/cdk8sawscdkresolver/jsii"

	"github.com/cdk8s-team/cdk8s-awscdk-resolver-go/cdk8sawscdkresolver/internal"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type AwsCdkResolver interface {
	cdk8s.IResolver
	// This function is invoked on every property during cdk8s synthesis.
	//
	// To replace a value, implementations must invoke `context.replaceValue`.
	Resolve(context cdk8s.ResolutionContext)
}

// The jsii proxy struct for AwsCdkResolver
type jsiiProxy_AwsCdkResolver struct {
	internal.Type__cdk8sIResolver
}

func NewAwsCdkResolver() AwsCdkResolver {
	_init_.Initialize()

	j := jsiiProxy_AwsCdkResolver{}

	_jsii_.Create(
		"@cdk8s/awscdk-resolver.AwsCdkResolver",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAwsCdkResolver_Override(a AwsCdkResolver) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdk8s/awscdk-resolver.AwsCdkResolver",
		nil, // no parameters
		a,
	)
}

func (a *jsiiProxy_AwsCdkResolver) Resolve(context cdk8s.ResolutionContext) {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"resolve",
		[]interface{}{context},
	)
}

