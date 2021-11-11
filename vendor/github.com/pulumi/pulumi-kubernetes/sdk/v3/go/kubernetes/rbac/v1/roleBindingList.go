// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// RoleBindingList is a collection of RoleBindings
type RoleBindingList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is a list of RoleBindings
	Items RoleBindingTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewRoleBindingList registers a new resource with the given unique name, arguments, and options.
func NewRoleBindingList(ctx *pulumi.Context,
	name string, args *RoleBindingListArgs, opts ...pulumi.ResourceOption) (*RoleBindingList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1")
	args.Kind = pulumi.StringPtr("RoleBindingList")
	var resource RoleBindingList
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1:RoleBindingList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleBindingList gets an existing RoleBindingList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleBindingList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleBindingListState, opts ...pulumi.ResourceOption) (*RoleBindingList, error) {
	var resource RoleBindingList
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1:RoleBindingList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleBindingList resources.
type roleBindingListState struct {
}

type RoleBindingListState struct {
}

func (RoleBindingListState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleBindingListState)(nil)).Elem()
}

type roleBindingListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of RoleBindings
	Items []RoleBindingType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a RoleBindingList resource.
type RoleBindingListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of RoleBindings
	Items RoleBindingTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ListMetaPtrInput
}

func (RoleBindingListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleBindingListArgs)(nil)).Elem()
}

type RoleBindingListInput interface {
	pulumi.Input

	ToRoleBindingListOutput() RoleBindingListOutput
	ToRoleBindingListOutputWithContext(ctx context.Context) RoleBindingListOutput
}

func (*RoleBindingList) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleBindingList)(nil))
}

func (i *RoleBindingList) ToRoleBindingListOutput() RoleBindingListOutput {
	return i.ToRoleBindingListOutputWithContext(context.Background())
}

func (i *RoleBindingList) ToRoleBindingListOutputWithContext(ctx context.Context) RoleBindingListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleBindingListOutput)
}

func (i *RoleBindingList) ToRoleBindingListPtrOutput() RoleBindingListPtrOutput {
	return i.ToRoleBindingListPtrOutputWithContext(context.Background())
}

func (i *RoleBindingList) ToRoleBindingListPtrOutputWithContext(ctx context.Context) RoleBindingListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleBindingListPtrOutput)
}

type RoleBindingListPtrInput interface {
	pulumi.Input

	ToRoleBindingListPtrOutput() RoleBindingListPtrOutput
	ToRoleBindingListPtrOutputWithContext(ctx context.Context) RoleBindingListPtrOutput
}

type roleBindingListPtrType RoleBindingListArgs

func (*roleBindingListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleBindingList)(nil))
}

func (i *roleBindingListPtrType) ToRoleBindingListPtrOutput() RoleBindingListPtrOutput {
	return i.ToRoleBindingListPtrOutputWithContext(context.Background())
}

func (i *roleBindingListPtrType) ToRoleBindingListPtrOutputWithContext(ctx context.Context) RoleBindingListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleBindingListPtrOutput)
}

// RoleBindingListArrayInput is an input type that accepts RoleBindingListArray and RoleBindingListArrayOutput values.
// You can construct a concrete instance of `RoleBindingListArrayInput` via:
//
//          RoleBindingListArray{ RoleBindingListArgs{...} }
type RoleBindingListArrayInput interface {
	pulumi.Input

	ToRoleBindingListArrayOutput() RoleBindingListArrayOutput
	ToRoleBindingListArrayOutputWithContext(context.Context) RoleBindingListArrayOutput
}

type RoleBindingListArray []RoleBindingListInput

func (RoleBindingListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleBindingList)(nil)).Elem()
}

func (i RoleBindingListArray) ToRoleBindingListArrayOutput() RoleBindingListArrayOutput {
	return i.ToRoleBindingListArrayOutputWithContext(context.Background())
}

func (i RoleBindingListArray) ToRoleBindingListArrayOutputWithContext(ctx context.Context) RoleBindingListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleBindingListArrayOutput)
}

// RoleBindingListMapInput is an input type that accepts RoleBindingListMap and RoleBindingListMapOutput values.
// You can construct a concrete instance of `RoleBindingListMapInput` via:
//
//          RoleBindingListMap{ "key": RoleBindingListArgs{...} }
type RoleBindingListMapInput interface {
	pulumi.Input

	ToRoleBindingListMapOutput() RoleBindingListMapOutput
	ToRoleBindingListMapOutputWithContext(context.Context) RoleBindingListMapOutput
}

type RoleBindingListMap map[string]RoleBindingListInput

func (RoleBindingListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleBindingList)(nil)).Elem()
}

func (i RoleBindingListMap) ToRoleBindingListMapOutput() RoleBindingListMapOutput {
	return i.ToRoleBindingListMapOutputWithContext(context.Background())
}

func (i RoleBindingListMap) ToRoleBindingListMapOutputWithContext(ctx context.Context) RoleBindingListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleBindingListMapOutput)
}

type RoleBindingListOutput struct{ *pulumi.OutputState }

func (RoleBindingListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleBindingList)(nil))
}

func (o RoleBindingListOutput) ToRoleBindingListOutput() RoleBindingListOutput {
	return o
}

func (o RoleBindingListOutput) ToRoleBindingListOutputWithContext(ctx context.Context) RoleBindingListOutput {
	return o
}

func (o RoleBindingListOutput) ToRoleBindingListPtrOutput() RoleBindingListPtrOutput {
	return o.ToRoleBindingListPtrOutputWithContext(context.Background())
}

func (o RoleBindingListOutput) ToRoleBindingListPtrOutputWithContext(ctx context.Context) RoleBindingListPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoleBindingList) *RoleBindingList {
		return &v
	}).(RoleBindingListPtrOutput)
}

type RoleBindingListPtrOutput struct{ *pulumi.OutputState }

func (RoleBindingListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleBindingList)(nil))
}

func (o RoleBindingListPtrOutput) ToRoleBindingListPtrOutput() RoleBindingListPtrOutput {
	return o
}

func (o RoleBindingListPtrOutput) ToRoleBindingListPtrOutputWithContext(ctx context.Context) RoleBindingListPtrOutput {
	return o
}

func (o RoleBindingListPtrOutput) Elem() RoleBindingListOutput {
	return o.ApplyT(func(v *RoleBindingList) RoleBindingList {
		if v != nil {
			return *v
		}
		var ret RoleBindingList
		return ret
	}).(RoleBindingListOutput)
}

type RoleBindingListArrayOutput struct{ *pulumi.OutputState }

func (RoleBindingListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoleBindingList)(nil))
}

func (o RoleBindingListArrayOutput) ToRoleBindingListArrayOutput() RoleBindingListArrayOutput {
	return o
}

func (o RoleBindingListArrayOutput) ToRoleBindingListArrayOutputWithContext(ctx context.Context) RoleBindingListArrayOutput {
	return o
}

func (o RoleBindingListArrayOutput) Index(i pulumi.IntInput) RoleBindingListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoleBindingList {
		return vs[0].([]RoleBindingList)[vs[1].(int)]
	}).(RoleBindingListOutput)
}

type RoleBindingListMapOutput struct{ *pulumi.OutputState }

func (RoleBindingListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RoleBindingList)(nil))
}

func (o RoleBindingListMapOutput) ToRoleBindingListMapOutput() RoleBindingListMapOutput {
	return o
}

func (o RoleBindingListMapOutput) ToRoleBindingListMapOutputWithContext(ctx context.Context) RoleBindingListMapOutput {
	return o
}

func (o RoleBindingListMapOutput) MapIndex(k pulumi.StringInput) RoleBindingListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RoleBindingList {
		return vs[0].(map[string]RoleBindingList)[vs[1].(string)]
	}).(RoleBindingListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleBindingListInput)(nil)).Elem(), &RoleBindingList{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleBindingListPtrInput)(nil)).Elem(), &RoleBindingList{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleBindingListArrayInput)(nil)).Elem(), RoleBindingListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleBindingListMapInput)(nil)).Elem(), RoleBindingListMap{})
	pulumi.RegisterOutputType(RoleBindingListOutput{})
	pulumi.RegisterOutputType(RoleBindingListPtrOutput{})
	pulumi.RegisterOutputType(RoleBindingListArrayOutput{})
	pulumi.RegisterOutputType(RoleBindingListMapOutput{})
}
