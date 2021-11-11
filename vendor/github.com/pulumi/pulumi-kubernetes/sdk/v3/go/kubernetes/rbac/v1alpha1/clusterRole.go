// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRole, and will no longer be served in v1.22.
type ClusterRole struct {
	pulumi.CustomResourceState

	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
	AggregationRule AggregationRulePtrOutput `pulumi:"aggregationRule"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Rules holds all the PolicyRules for this ClusterRole
	Rules PolicyRuleArrayOutput `pulumi:"rules"`
}

// NewClusterRole registers a new resource with the given unique name, arguments, and options.
func NewClusterRole(ctx *pulumi.Context,
	name string, args *ClusterRoleArgs, opts ...pulumi.ResourceOption) (*ClusterRole, error) {
	if args == nil {
		args = &ClusterRoleArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("ClusterRole")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1:ClusterRole"),
		},
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRole"),
		},
	})
	opts = append(opts, aliases)
	var resource ClusterRole
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterRole gets an existing ClusterRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterRoleState, opts ...pulumi.ResourceOption) (*ClusterRole, error) {
	var resource ClusterRole
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterRole resources.
type clusterRoleState struct {
}

type ClusterRoleState struct {
}

func (ClusterRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleState)(nil)).Elem()
}

type clusterRoleArgs struct {
	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
	AggregationRule *AggregationRule `pulumi:"aggregationRule"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Rules holds all the PolicyRules for this ClusterRole
	Rules []PolicyRule `pulumi:"rules"`
}

// The set of arguments for constructing a ClusterRole resource.
type ClusterRoleArgs struct {
	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
	AggregationRule AggregationRulePtrInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPtrInput
	// Rules holds all the PolicyRules for this ClusterRole
	Rules PolicyRuleArrayInput
}

func (ClusterRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterRoleArgs)(nil)).Elem()
}

type ClusterRoleInput interface {
	pulumi.Input

	ToClusterRoleOutput() ClusterRoleOutput
	ToClusterRoleOutputWithContext(ctx context.Context) ClusterRoleOutput
}

func (*ClusterRole) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRole)(nil))
}

func (i *ClusterRole) ToClusterRoleOutput() ClusterRoleOutput {
	return i.ToClusterRoleOutputWithContext(context.Background())
}

func (i *ClusterRole) ToClusterRoleOutputWithContext(ctx context.Context) ClusterRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleOutput)
}

func (i *ClusterRole) ToClusterRolePtrOutput() ClusterRolePtrOutput {
	return i.ToClusterRolePtrOutputWithContext(context.Background())
}

func (i *ClusterRole) ToClusterRolePtrOutputWithContext(ctx context.Context) ClusterRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRolePtrOutput)
}

type ClusterRolePtrInput interface {
	pulumi.Input

	ToClusterRolePtrOutput() ClusterRolePtrOutput
	ToClusterRolePtrOutputWithContext(ctx context.Context) ClusterRolePtrOutput
}

type clusterRolePtrType ClusterRoleArgs

func (*clusterRolePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterRole)(nil))
}

func (i *clusterRolePtrType) ToClusterRolePtrOutput() ClusterRolePtrOutput {
	return i.ToClusterRolePtrOutputWithContext(context.Background())
}

func (i *clusterRolePtrType) ToClusterRolePtrOutputWithContext(ctx context.Context) ClusterRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRolePtrOutput)
}

// ClusterRoleArrayInput is an input type that accepts ClusterRoleArray and ClusterRoleArrayOutput values.
// You can construct a concrete instance of `ClusterRoleArrayInput` via:
//
//          ClusterRoleArray{ ClusterRoleArgs{...} }
type ClusterRoleArrayInput interface {
	pulumi.Input

	ToClusterRoleArrayOutput() ClusterRoleArrayOutput
	ToClusterRoleArrayOutputWithContext(context.Context) ClusterRoleArrayOutput
}

type ClusterRoleArray []ClusterRoleInput

func (ClusterRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterRole)(nil)).Elem()
}

func (i ClusterRoleArray) ToClusterRoleArrayOutput() ClusterRoleArrayOutput {
	return i.ToClusterRoleArrayOutputWithContext(context.Background())
}

func (i ClusterRoleArray) ToClusterRoleArrayOutputWithContext(ctx context.Context) ClusterRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleArrayOutput)
}

// ClusterRoleMapInput is an input type that accepts ClusterRoleMap and ClusterRoleMapOutput values.
// You can construct a concrete instance of `ClusterRoleMapInput` via:
//
//          ClusterRoleMap{ "key": ClusterRoleArgs{...} }
type ClusterRoleMapInput interface {
	pulumi.Input

	ToClusterRoleMapOutput() ClusterRoleMapOutput
	ToClusterRoleMapOutputWithContext(context.Context) ClusterRoleMapOutput
}

type ClusterRoleMap map[string]ClusterRoleInput

func (ClusterRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterRole)(nil)).Elem()
}

func (i ClusterRoleMap) ToClusterRoleMapOutput() ClusterRoleMapOutput {
	return i.ToClusterRoleMapOutputWithContext(context.Background())
}

func (i ClusterRoleMap) ToClusterRoleMapOutputWithContext(ctx context.Context) ClusterRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoleMapOutput)
}

type ClusterRoleOutput struct{ *pulumi.OutputState }

func (ClusterRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRole)(nil))
}

func (o ClusterRoleOutput) ToClusterRoleOutput() ClusterRoleOutput {
	return o
}

func (o ClusterRoleOutput) ToClusterRoleOutputWithContext(ctx context.Context) ClusterRoleOutput {
	return o
}

func (o ClusterRoleOutput) ToClusterRolePtrOutput() ClusterRolePtrOutput {
	return o.ToClusterRolePtrOutputWithContext(context.Background())
}

func (o ClusterRoleOutput) ToClusterRolePtrOutputWithContext(ctx context.Context) ClusterRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterRole) *ClusterRole {
		return &v
	}).(ClusterRolePtrOutput)
}

type ClusterRolePtrOutput struct{ *pulumi.OutputState }

func (ClusterRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterRole)(nil))
}

func (o ClusterRolePtrOutput) ToClusterRolePtrOutput() ClusterRolePtrOutput {
	return o
}

func (o ClusterRolePtrOutput) ToClusterRolePtrOutputWithContext(ctx context.Context) ClusterRolePtrOutput {
	return o
}

func (o ClusterRolePtrOutput) Elem() ClusterRoleOutput {
	return o.ApplyT(func(v *ClusterRole) ClusterRole {
		if v != nil {
			return *v
		}
		var ret ClusterRole
		return ret
	}).(ClusterRoleOutput)
}

type ClusterRoleArrayOutput struct{ *pulumi.OutputState }

func (ClusterRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterRole)(nil))
}

func (o ClusterRoleArrayOutput) ToClusterRoleArrayOutput() ClusterRoleArrayOutput {
	return o
}

func (o ClusterRoleArrayOutput) ToClusterRoleArrayOutputWithContext(ctx context.Context) ClusterRoleArrayOutput {
	return o
}

func (o ClusterRoleArrayOutput) Index(i pulumi.IntInput) ClusterRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterRole {
		return vs[0].([]ClusterRole)[vs[1].(int)]
	}).(ClusterRoleOutput)
}

type ClusterRoleMapOutput struct{ *pulumi.OutputState }

func (ClusterRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClusterRole)(nil))
}

func (o ClusterRoleMapOutput) ToClusterRoleMapOutput() ClusterRoleMapOutput {
	return o
}

func (o ClusterRoleMapOutput) ToClusterRoleMapOutputWithContext(ctx context.Context) ClusterRoleMapOutput {
	return o
}

func (o ClusterRoleMapOutput) MapIndex(k pulumi.StringInput) ClusterRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClusterRole {
		return vs[0].(map[string]ClusterRole)[vs[1].(string)]
	}).(ClusterRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoleInput)(nil)).Elem(), &ClusterRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRolePtrInput)(nil)).Elem(), &ClusterRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoleArrayInput)(nil)).Elem(), ClusterRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoleMapInput)(nil)).Elem(), ClusterRoleMap{})
	pulumi.RegisterOutputType(ClusterRoleOutput{})
	pulumi.RegisterOutputType(ClusterRolePtrOutput{})
	pulumi.RegisterOutputType(ClusterRoleArrayOutput{})
	pulumi.RegisterOutputType(ClusterRoleMapOutput{})
}
