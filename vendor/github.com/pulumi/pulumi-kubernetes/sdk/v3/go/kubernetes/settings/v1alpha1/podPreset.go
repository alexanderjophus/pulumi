// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// PodPreset is a policy resource that defines additional runtime requirements for a Pod.
type PodPreset struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	Spec     PodPresetSpecPtrOutput     `pulumi:"spec"`
}

// NewPodPreset registers a new resource with the given unique name, arguments, and options.
func NewPodPreset(ctx *pulumi.Context,
	name string, args *PodPresetArgs, opts ...pulumi.ResourceOption) (*PodPreset, error) {
	if args == nil {
		args = &PodPresetArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("settings.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("PodPreset")
	var resource PodPreset
	err := ctx.RegisterResource("kubernetes:settings.k8s.io/v1alpha1:PodPreset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPodPreset gets an existing PodPreset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPodPreset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PodPresetState, opts ...pulumi.ResourceOption) (*PodPreset, error) {
	var resource PodPreset
	err := ctx.ReadResource("kubernetes:settings.k8s.io/v1alpha1:PodPreset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PodPreset resources.
type podPresetState struct {
}

type PodPresetState struct {
}

func (PodPresetState) ElementType() reflect.Type {
	return reflect.TypeOf((*podPresetState)(nil)).Elem()
}

type podPresetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	Spec     *PodPresetSpec     `pulumi:"spec"`
}

// The set of arguments for constructing a PodPreset resource.
type PodPresetArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	Spec     PodPresetSpecPtrInput
}

func (PodPresetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*podPresetArgs)(nil)).Elem()
}

type PodPresetInput interface {
	pulumi.Input

	ToPodPresetOutput() PodPresetOutput
	ToPodPresetOutputWithContext(ctx context.Context) PodPresetOutput
}

func (*PodPreset) ElementType() reflect.Type {
	return reflect.TypeOf((*PodPreset)(nil))
}

func (i *PodPreset) ToPodPresetOutput() PodPresetOutput {
	return i.ToPodPresetOutputWithContext(context.Background())
}

func (i *PodPreset) ToPodPresetOutputWithContext(ctx context.Context) PodPresetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodPresetOutput)
}

func (i *PodPreset) ToPodPresetPtrOutput() PodPresetPtrOutput {
	return i.ToPodPresetPtrOutputWithContext(context.Background())
}

func (i *PodPreset) ToPodPresetPtrOutputWithContext(ctx context.Context) PodPresetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodPresetPtrOutput)
}

type PodPresetPtrInput interface {
	pulumi.Input

	ToPodPresetPtrOutput() PodPresetPtrOutput
	ToPodPresetPtrOutputWithContext(ctx context.Context) PodPresetPtrOutput
}

type podPresetPtrType PodPresetArgs

func (*podPresetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PodPreset)(nil))
}

func (i *podPresetPtrType) ToPodPresetPtrOutput() PodPresetPtrOutput {
	return i.ToPodPresetPtrOutputWithContext(context.Background())
}

func (i *podPresetPtrType) ToPodPresetPtrOutputWithContext(ctx context.Context) PodPresetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodPresetPtrOutput)
}

// PodPresetArrayInput is an input type that accepts PodPresetArray and PodPresetArrayOutput values.
// You can construct a concrete instance of `PodPresetArrayInput` via:
//
//          PodPresetArray{ PodPresetArgs{...} }
type PodPresetArrayInput interface {
	pulumi.Input

	ToPodPresetArrayOutput() PodPresetArrayOutput
	ToPodPresetArrayOutputWithContext(context.Context) PodPresetArrayOutput
}

type PodPresetArray []PodPresetInput

func (PodPresetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PodPreset)(nil)).Elem()
}

func (i PodPresetArray) ToPodPresetArrayOutput() PodPresetArrayOutput {
	return i.ToPodPresetArrayOutputWithContext(context.Background())
}

func (i PodPresetArray) ToPodPresetArrayOutputWithContext(ctx context.Context) PodPresetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodPresetArrayOutput)
}

// PodPresetMapInput is an input type that accepts PodPresetMap and PodPresetMapOutput values.
// You can construct a concrete instance of `PodPresetMapInput` via:
//
//          PodPresetMap{ "key": PodPresetArgs{...} }
type PodPresetMapInput interface {
	pulumi.Input

	ToPodPresetMapOutput() PodPresetMapOutput
	ToPodPresetMapOutputWithContext(context.Context) PodPresetMapOutput
}

type PodPresetMap map[string]PodPresetInput

func (PodPresetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PodPreset)(nil)).Elem()
}

func (i PodPresetMap) ToPodPresetMapOutput() PodPresetMapOutput {
	return i.ToPodPresetMapOutputWithContext(context.Background())
}

func (i PodPresetMap) ToPodPresetMapOutputWithContext(ctx context.Context) PodPresetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PodPresetMapOutput)
}

type PodPresetOutput struct{ *pulumi.OutputState }

func (PodPresetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PodPreset)(nil))
}

func (o PodPresetOutput) ToPodPresetOutput() PodPresetOutput {
	return o
}

func (o PodPresetOutput) ToPodPresetOutputWithContext(ctx context.Context) PodPresetOutput {
	return o
}

func (o PodPresetOutput) ToPodPresetPtrOutput() PodPresetPtrOutput {
	return o.ToPodPresetPtrOutputWithContext(context.Background())
}

func (o PodPresetOutput) ToPodPresetPtrOutputWithContext(ctx context.Context) PodPresetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PodPreset) *PodPreset {
		return &v
	}).(PodPresetPtrOutput)
}

type PodPresetPtrOutput struct{ *pulumi.OutputState }

func (PodPresetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PodPreset)(nil))
}

func (o PodPresetPtrOutput) ToPodPresetPtrOutput() PodPresetPtrOutput {
	return o
}

func (o PodPresetPtrOutput) ToPodPresetPtrOutputWithContext(ctx context.Context) PodPresetPtrOutput {
	return o
}

func (o PodPresetPtrOutput) Elem() PodPresetOutput {
	return o.ApplyT(func(v *PodPreset) PodPreset {
		if v != nil {
			return *v
		}
		var ret PodPreset
		return ret
	}).(PodPresetOutput)
}

type PodPresetArrayOutput struct{ *pulumi.OutputState }

func (PodPresetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PodPreset)(nil))
}

func (o PodPresetArrayOutput) ToPodPresetArrayOutput() PodPresetArrayOutput {
	return o
}

func (o PodPresetArrayOutput) ToPodPresetArrayOutputWithContext(ctx context.Context) PodPresetArrayOutput {
	return o
}

func (o PodPresetArrayOutput) Index(i pulumi.IntInput) PodPresetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PodPreset {
		return vs[0].([]PodPreset)[vs[1].(int)]
	}).(PodPresetOutput)
}

type PodPresetMapOutput struct{ *pulumi.OutputState }

func (PodPresetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PodPreset)(nil))
}

func (o PodPresetMapOutput) ToPodPresetMapOutput() PodPresetMapOutput {
	return o
}

func (o PodPresetMapOutput) ToPodPresetMapOutputWithContext(ctx context.Context) PodPresetMapOutput {
	return o
}

func (o PodPresetMapOutput) MapIndex(k pulumi.StringInput) PodPresetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PodPreset {
		return vs[0].(map[string]PodPreset)[vs[1].(string)]
	}).(PodPresetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PodPresetInput)(nil)).Elem(), &PodPreset{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodPresetPtrInput)(nil)).Elem(), &PodPreset{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodPresetArrayInput)(nil)).Elem(), PodPresetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PodPresetMapInput)(nil)).Elem(), PodPresetMap{})
	pulumi.RegisterOutputType(PodPresetOutput{})
	pulumi.RegisterOutputType(PodPresetPtrOutput{})
	pulumi.RegisterOutputType(PodPresetArrayOutput{})
	pulumi.RegisterOutputType(PodPresetMapOutput{})
}
