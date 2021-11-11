// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a certificate signing request
type CertificateSigningRequest struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// The certificate request itself and any additional information.
	Spec CertificateSigningRequestSpecPtrOutput `pulumi:"spec"`
	// Derived information about the request.
	Status CertificateSigningRequestStatusPtrOutput `pulumi:"status"`
}

// NewCertificateSigningRequest registers a new resource with the given unique name, arguments, and options.
func NewCertificateSigningRequest(ctx *pulumi.Context,
	name string, args *CertificateSigningRequestArgs, opts ...pulumi.ResourceOption) (*CertificateSigningRequest, error) {
	if args == nil {
		args = &CertificateSigningRequestArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("certificates.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("CertificateSigningRequest")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:certificates.k8s.io/v1:CertificateSigningRequest"),
		},
	})
	opts = append(opts, aliases)
	var resource CertificateSigningRequest
	err := ctx.RegisterResource("kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateSigningRequest gets an existing CertificateSigningRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateSigningRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateSigningRequestState, opts ...pulumi.ResourceOption) (*CertificateSigningRequest, error) {
	var resource CertificateSigningRequest
	err := ctx.ReadResource("kubernetes:certificates.k8s.io/v1beta1:CertificateSigningRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateSigningRequest resources.
type certificateSigningRequestState struct {
}

type CertificateSigningRequestState struct {
}

func (CertificateSigningRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestState)(nil)).Elem()
}

type certificateSigningRequestArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// The certificate request itself and any additional information.
	Spec *CertificateSigningRequestSpec `pulumi:"spec"`
}

// The set of arguments for constructing a CertificateSigningRequest resource.
type CertificateSigningRequestArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// The certificate request itself and any additional information.
	Spec CertificateSigningRequestSpecPtrInput
}

func (CertificateSigningRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateSigningRequestArgs)(nil)).Elem()
}

type CertificateSigningRequestInput interface {
	pulumi.Input

	ToCertificateSigningRequestOutput() CertificateSigningRequestOutput
	ToCertificateSigningRequestOutputWithContext(ctx context.Context) CertificateSigningRequestOutput
}

func (*CertificateSigningRequest) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateSigningRequest)(nil))
}

func (i *CertificateSigningRequest) ToCertificateSigningRequestOutput() CertificateSigningRequestOutput {
	return i.ToCertificateSigningRequestOutputWithContext(context.Background())
}

func (i *CertificateSigningRequest) ToCertificateSigningRequestOutputWithContext(ctx context.Context) CertificateSigningRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestOutput)
}

func (i *CertificateSigningRequest) ToCertificateSigningRequestPtrOutput() CertificateSigningRequestPtrOutput {
	return i.ToCertificateSigningRequestPtrOutputWithContext(context.Background())
}

func (i *CertificateSigningRequest) ToCertificateSigningRequestPtrOutputWithContext(ctx context.Context) CertificateSigningRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestPtrOutput)
}

type CertificateSigningRequestPtrInput interface {
	pulumi.Input

	ToCertificateSigningRequestPtrOutput() CertificateSigningRequestPtrOutput
	ToCertificateSigningRequestPtrOutputWithContext(ctx context.Context) CertificateSigningRequestPtrOutput
}

type certificateSigningRequestPtrType CertificateSigningRequestArgs

func (*certificateSigningRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateSigningRequest)(nil))
}

func (i *certificateSigningRequestPtrType) ToCertificateSigningRequestPtrOutput() CertificateSigningRequestPtrOutput {
	return i.ToCertificateSigningRequestPtrOutputWithContext(context.Background())
}

func (i *certificateSigningRequestPtrType) ToCertificateSigningRequestPtrOutputWithContext(ctx context.Context) CertificateSigningRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestPtrOutput)
}

// CertificateSigningRequestArrayInput is an input type that accepts CertificateSigningRequestArray and CertificateSigningRequestArrayOutput values.
// You can construct a concrete instance of `CertificateSigningRequestArrayInput` via:
//
//          CertificateSigningRequestArray{ CertificateSigningRequestArgs{...} }
type CertificateSigningRequestArrayInput interface {
	pulumi.Input

	ToCertificateSigningRequestArrayOutput() CertificateSigningRequestArrayOutput
	ToCertificateSigningRequestArrayOutputWithContext(context.Context) CertificateSigningRequestArrayOutput
}

type CertificateSigningRequestArray []CertificateSigningRequestInput

func (CertificateSigningRequestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateSigningRequest)(nil)).Elem()
}

func (i CertificateSigningRequestArray) ToCertificateSigningRequestArrayOutput() CertificateSigningRequestArrayOutput {
	return i.ToCertificateSigningRequestArrayOutputWithContext(context.Background())
}

func (i CertificateSigningRequestArray) ToCertificateSigningRequestArrayOutputWithContext(ctx context.Context) CertificateSigningRequestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestArrayOutput)
}

// CertificateSigningRequestMapInput is an input type that accepts CertificateSigningRequestMap and CertificateSigningRequestMapOutput values.
// You can construct a concrete instance of `CertificateSigningRequestMapInput` via:
//
//          CertificateSigningRequestMap{ "key": CertificateSigningRequestArgs{...} }
type CertificateSigningRequestMapInput interface {
	pulumi.Input

	ToCertificateSigningRequestMapOutput() CertificateSigningRequestMapOutput
	ToCertificateSigningRequestMapOutputWithContext(context.Context) CertificateSigningRequestMapOutput
}

type CertificateSigningRequestMap map[string]CertificateSigningRequestInput

func (CertificateSigningRequestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateSigningRequest)(nil)).Elem()
}

func (i CertificateSigningRequestMap) ToCertificateSigningRequestMapOutput() CertificateSigningRequestMapOutput {
	return i.ToCertificateSigningRequestMapOutputWithContext(context.Background())
}

func (i CertificateSigningRequestMap) ToCertificateSigningRequestMapOutputWithContext(ctx context.Context) CertificateSigningRequestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateSigningRequestMapOutput)
}

type CertificateSigningRequestOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateSigningRequest)(nil))
}

func (o CertificateSigningRequestOutput) ToCertificateSigningRequestOutput() CertificateSigningRequestOutput {
	return o
}

func (o CertificateSigningRequestOutput) ToCertificateSigningRequestOutputWithContext(ctx context.Context) CertificateSigningRequestOutput {
	return o
}

func (o CertificateSigningRequestOutput) ToCertificateSigningRequestPtrOutput() CertificateSigningRequestPtrOutput {
	return o.ToCertificateSigningRequestPtrOutputWithContext(context.Background())
}

func (o CertificateSigningRequestOutput) ToCertificateSigningRequestPtrOutputWithContext(ctx context.Context) CertificateSigningRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateSigningRequest) *CertificateSigningRequest {
		return &v
	}).(CertificateSigningRequestPtrOutput)
}

type CertificateSigningRequestPtrOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateSigningRequest)(nil))
}

func (o CertificateSigningRequestPtrOutput) ToCertificateSigningRequestPtrOutput() CertificateSigningRequestPtrOutput {
	return o
}

func (o CertificateSigningRequestPtrOutput) ToCertificateSigningRequestPtrOutputWithContext(ctx context.Context) CertificateSigningRequestPtrOutput {
	return o
}

func (o CertificateSigningRequestPtrOutput) Elem() CertificateSigningRequestOutput {
	return o.ApplyT(func(v *CertificateSigningRequest) CertificateSigningRequest {
		if v != nil {
			return *v
		}
		var ret CertificateSigningRequest
		return ret
	}).(CertificateSigningRequestOutput)
}

type CertificateSigningRequestArrayOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateSigningRequest)(nil))
}

func (o CertificateSigningRequestArrayOutput) ToCertificateSigningRequestArrayOutput() CertificateSigningRequestArrayOutput {
	return o
}

func (o CertificateSigningRequestArrayOutput) ToCertificateSigningRequestArrayOutputWithContext(ctx context.Context) CertificateSigningRequestArrayOutput {
	return o
}

func (o CertificateSigningRequestArrayOutput) Index(i pulumi.IntInput) CertificateSigningRequestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateSigningRequest {
		return vs[0].([]CertificateSigningRequest)[vs[1].(int)]
	}).(CertificateSigningRequestOutput)
}

type CertificateSigningRequestMapOutput struct{ *pulumi.OutputState }

func (CertificateSigningRequestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CertificateSigningRequest)(nil))
}

func (o CertificateSigningRequestMapOutput) ToCertificateSigningRequestMapOutput() CertificateSigningRequestMapOutput {
	return o
}

func (o CertificateSigningRequestMapOutput) ToCertificateSigningRequestMapOutputWithContext(ctx context.Context) CertificateSigningRequestMapOutput {
	return o
}

func (o CertificateSigningRequestMapOutput) MapIndex(k pulumi.StringInput) CertificateSigningRequestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CertificateSigningRequest {
		return vs[0].(map[string]CertificateSigningRequest)[vs[1].(string)]
	}).(CertificateSigningRequestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestInput)(nil)).Elem(), &CertificateSigningRequest{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestPtrInput)(nil)).Elem(), &CertificateSigningRequest{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestArrayInput)(nil)).Elem(), CertificateSigningRequestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateSigningRequestMapInput)(nil)).Elem(), CertificateSigningRequestMap{})
	pulumi.RegisterOutputType(CertificateSigningRequestOutput{})
	pulumi.RegisterOutputType(CertificateSigningRequestPtrOutput{})
	pulumi.RegisterOutputType(CertificateSigningRequestArrayOutput{})
	pulumi.RegisterOutputType(CertificateSigningRequestMapOutput{})
}
