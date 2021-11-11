// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it.
type ValidatingWebhookConfiguration struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks ValidatingWebhookArrayOutput `pulumi:"webhooks"`
}

// NewValidatingWebhookConfiguration registers a new resource with the given unique name, arguments, and options.
func NewValidatingWebhookConfiguration(ctx *pulumi.Context,
	name string, args *ValidatingWebhookConfigurationArgs, opts ...pulumi.ResourceOption) (*ValidatingWebhookConfiguration, error) {
	if args == nil {
		args = &ValidatingWebhookConfigurationArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("admissionregistration.k8s.io/v1")
	args.Kind = pulumi.StringPtr("ValidatingWebhookConfiguration")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingWebhookConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource ValidatingWebhookConfiguration
	err := ctx.RegisterResource("kubernetes:admissionregistration.k8s.io/v1:ValidatingWebhookConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetValidatingWebhookConfiguration gets an existing ValidatingWebhookConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetValidatingWebhookConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ValidatingWebhookConfigurationState, opts ...pulumi.ResourceOption) (*ValidatingWebhookConfiguration, error) {
	var resource ValidatingWebhookConfiguration
	err := ctx.ReadResource("kubernetes:admissionregistration.k8s.io/v1:ValidatingWebhookConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ValidatingWebhookConfiguration resources.
type validatingWebhookConfigurationState struct {
}

type ValidatingWebhookConfigurationState struct {
}

func (ValidatingWebhookConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*validatingWebhookConfigurationState)(nil)).Elem()
}

type validatingWebhookConfigurationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks []ValidatingWebhook `pulumi:"webhooks"`
}

// The set of arguments for constructing a ValidatingWebhookConfiguration resource.
type ValidatingWebhookConfigurationArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata metav1.ObjectMetaPtrInput
	// Webhooks is a list of webhooks and the affected resources and operations.
	Webhooks ValidatingWebhookArrayInput
}

func (ValidatingWebhookConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*validatingWebhookConfigurationArgs)(nil)).Elem()
}

type ValidatingWebhookConfigurationInput interface {
	pulumi.Input

	ToValidatingWebhookConfigurationOutput() ValidatingWebhookConfigurationOutput
	ToValidatingWebhookConfigurationOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationOutput
}

func (*ValidatingWebhookConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*ValidatingWebhookConfiguration)(nil))
}

func (i *ValidatingWebhookConfiguration) ToValidatingWebhookConfigurationOutput() ValidatingWebhookConfigurationOutput {
	return i.ToValidatingWebhookConfigurationOutputWithContext(context.Background())
}

func (i *ValidatingWebhookConfiguration) ToValidatingWebhookConfigurationOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingWebhookConfigurationOutput)
}

func (i *ValidatingWebhookConfiguration) ToValidatingWebhookConfigurationPtrOutput() ValidatingWebhookConfigurationPtrOutput {
	return i.ToValidatingWebhookConfigurationPtrOutputWithContext(context.Background())
}

func (i *ValidatingWebhookConfiguration) ToValidatingWebhookConfigurationPtrOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingWebhookConfigurationPtrOutput)
}

type ValidatingWebhookConfigurationPtrInput interface {
	pulumi.Input

	ToValidatingWebhookConfigurationPtrOutput() ValidatingWebhookConfigurationPtrOutput
	ToValidatingWebhookConfigurationPtrOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationPtrOutput
}

type validatingWebhookConfigurationPtrType ValidatingWebhookConfigurationArgs

func (*validatingWebhookConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ValidatingWebhookConfiguration)(nil))
}

func (i *validatingWebhookConfigurationPtrType) ToValidatingWebhookConfigurationPtrOutput() ValidatingWebhookConfigurationPtrOutput {
	return i.ToValidatingWebhookConfigurationPtrOutputWithContext(context.Background())
}

func (i *validatingWebhookConfigurationPtrType) ToValidatingWebhookConfigurationPtrOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingWebhookConfigurationPtrOutput)
}

// ValidatingWebhookConfigurationArrayInput is an input type that accepts ValidatingWebhookConfigurationArray and ValidatingWebhookConfigurationArrayOutput values.
// You can construct a concrete instance of `ValidatingWebhookConfigurationArrayInput` via:
//
//          ValidatingWebhookConfigurationArray{ ValidatingWebhookConfigurationArgs{...} }
type ValidatingWebhookConfigurationArrayInput interface {
	pulumi.Input

	ToValidatingWebhookConfigurationArrayOutput() ValidatingWebhookConfigurationArrayOutput
	ToValidatingWebhookConfigurationArrayOutputWithContext(context.Context) ValidatingWebhookConfigurationArrayOutput
}

type ValidatingWebhookConfigurationArray []ValidatingWebhookConfigurationInput

func (ValidatingWebhookConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ValidatingWebhookConfiguration)(nil)).Elem()
}

func (i ValidatingWebhookConfigurationArray) ToValidatingWebhookConfigurationArrayOutput() ValidatingWebhookConfigurationArrayOutput {
	return i.ToValidatingWebhookConfigurationArrayOutputWithContext(context.Background())
}

func (i ValidatingWebhookConfigurationArray) ToValidatingWebhookConfigurationArrayOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingWebhookConfigurationArrayOutput)
}

// ValidatingWebhookConfigurationMapInput is an input type that accepts ValidatingWebhookConfigurationMap and ValidatingWebhookConfigurationMapOutput values.
// You can construct a concrete instance of `ValidatingWebhookConfigurationMapInput` via:
//
//          ValidatingWebhookConfigurationMap{ "key": ValidatingWebhookConfigurationArgs{...} }
type ValidatingWebhookConfigurationMapInput interface {
	pulumi.Input

	ToValidatingWebhookConfigurationMapOutput() ValidatingWebhookConfigurationMapOutput
	ToValidatingWebhookConfigurationMapOutputWithContext(context.Context) ValidatingWebhookConfigurationMapOutput
}

type ValidatingWebhookConfigurationMap map[string]ValidatingWebhookConfigurationInput

func (ValidatingWebhookConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ValidatingWebhookConfiguration)(nil)).Elem()
}

func (i ValidatingWebhookConfigurationMap) ToValidatingWebhookConfigurationMapOutput() ValidatingWebhookConfigurationMapOutput {
	return i.ToValidatingWebhookConfigurationMapOutputWithContext(context.Background())
}

func (i ValidatingWebhookConfigurationMap) ToValidatingWebhookConfigurationMapOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ValidatingWebhookConfigurationMapOutput)
}

type ValidatingWebhookConfigurationOutput struct{ *pulumi.OutputState }

func (ValidatingWebhookConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ValidatingWebhookConfiguration)(nil))
}

func (o ValidatingWebhookConfigurationOutput) ToValidatingWebhookConfigurationOutput() ValidatingWebhookConfigurationOutput {
	return o
}

func (o ValidatingWebhookConfigurationOutput) ToValidatingWebhookConfigurationOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationOutput {
	return o
}

func (o ValidatingWebhookConfigurationOutput) ToValidatingWebhookConfigurationPtrOutput() ValidatingWebhookConfigurationPtrOutput {
	return o.ToValidatingWebhookConfigurationPtrOutputWithContext(context.Background())
}

func (o ValidatingWebhookConfigurationOutput) ToValidatingWebhookConfigurationPtrOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ValidatingWebhookConfiguration) *ValidatingWebhookConfiguration {
		return &v
	}).(ValidatingWebhookConfigurationPtrOutput)
}

type ValidatingWebhookConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ValidatingWebhookConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ValidatingWebhookConfiguration)(nil))
}

func (o ValidatingWebhookConfigurationPtrOutput) ToValidatingWebhookConfigurationPtrOutput() ValidatingWebhookConfigurationPtrOutput {
	return o
}

func (o ValidatingWebhookConfigurationPtrOutput) ToValidatingWebhookConfigurationPtrOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationPtrOutput {
	return o
}

func (o ValidatingWebhookConfigurationPtrOutput) Elem() ValidatingWebhookConfigurationOutput {
	return o.ApplyT(func(v *ValidatingWebhookConfiguration) ValidatingWebhookConfiguration {
		if v != nil {
			return *v
		}
		var ret ValidatingWebhookConfiguration
		return ret
	}).(ValidatingWebhookConfigurationOutput)
}

type ValidatingWebhookConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ValidatingWebhookConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ValidatingWebhookConfiguration)(nil))
}

func (o ValidatingWebhookConfigurationArrayOutput) ToValidatingWebhookConfigurationArrayOutput() ValidatingWebhookConfigurationArrayOutput {
	return o
}

func (o ValidatingWebhookConfigurationArrayOutput) ToValidatingWebhookConfigurationArrayOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationArrayOutput {
	return o
}

func (o ValidatingWebhookConfigurationArrayOutput) Index(i pulumi.IntInput) ValidatingWebhookConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ValidatingWebhookConfiguration {
		return vs[0].([]ValidatingWebhookConfiguration)[vs[1].(int)]
	}).(ValidatingWebhookConfigurationOutput)
}

type ValidatingWebhookConfigurationMapOutput struct{ *pulumi.OutputState }

func (ValidatingWebhookConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ValidatingWebhookConfiguration)(nil))
}

func (o ValidatingWebhookConfigurationMapOutput) ToValidatingWebhookConfigurationMapOutput() ValidatingWebhookConfigurationMapOutput {
	return o
}

func (o ValidatingWebhookConfigurationMapOutput) ToValidatingWebhookConfigurationMapOutputWithContext(ctx context.Context) ValidatingWebhookConfigurationMapOutput {
	return o
}

func (o ValidatingWebhookConfigurationMapOutput) MapIndex(k pulumi.StringInput) ValidatingWebhookConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ValidatingWebhookConfiguration {
		return vs[0].(map[string]ValidatingWebhookConfiguration)[vs[1].(string)]
	}).(ValidatingWebhookConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ValidatingWebhookConfigurationInput)(nil)).Elem(), &ValidatingWebhookConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*ValidatingWebhookConfigurationPtrInput)(nil)).Elem(), &ValidatingWebhookConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*ValidatingWebhookConfigurationArrayInput)(nil)).Elem(), ValidatingWebhookConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ValidatingWebhookConfigurationMapInput)(nil)).Elem(), ValidatingWebhookConfigurationMap{})
	pulumi.RegisterOutputType(ValidatingWebhookConfigurationOutput{})
	pulumi.RegisterOutputType(ValidatingWebhookConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ValidatingWebhookConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ValidatingWebhookConfigurationMapOutput{})
}