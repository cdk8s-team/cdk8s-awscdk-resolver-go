// @cdk8s/awscdk-resolver
package cdk8sawscdkresolver

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdk8s/awscdk-resolver.AwsCdkResolver",
		reflect.TypeOf((*AwsCdkResolver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
		},
		func() interface{} {
			j := jsiiProxy_AwsCdkResolver{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sIResolver)
			return &j
		},
	)
}
