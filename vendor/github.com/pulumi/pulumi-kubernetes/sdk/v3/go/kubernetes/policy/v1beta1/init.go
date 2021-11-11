// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubernetes:policy/v1beta1:PodDisruptionBudget":
		r = &PodDisruptionBudget{}
	case "kubernetes:policy/v1beta1:PodDisruptionBudgetList":
		r = &PodDisruptionBudgetList{}
	case "kubernetes:policy/v1beta1:PodSecurityPolicy":
		r = &PodSecurityPolicy{}
	case "kubernetes:policy/v1beta1:PodSecurityPolicyList":
		r = &PodSecurityPolicyList{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := kubernetes.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"kubernetes",
		"policy/v1beta1",
		&module{version},
	)
}
