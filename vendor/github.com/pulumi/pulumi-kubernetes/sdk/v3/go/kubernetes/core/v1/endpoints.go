// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Endpoints is a collection of endpoints that implement the actual service. Example:
//   Name: "mysvc",
//   Subsets: [
//     {
//       Addresses: [{"ip": "10.10.1.1"}, {"ip": "10.10.2.2"}],
//       Ports: [{"name": "a", "port": 8675}, {"name": "b", "port": 309}]
//     },
//     {
//       Addresses: [{"ip": "10.10.3.3"}],
//       Ports: [{"name": "a", "port": 93}, {"name": "b", "port": 76}]
//     },
//  ]
type Endpoints struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service.
	Subsets EndpointSubsetArrayOutput `pulumi:"subsets"`
}

// NewEndpoints registers a new resource with the given unique name, arguments, and options.
func NewEndpoints(ctx *pulumi.Context,
	name string, args *EndpointsArgs, opts ...pulumi.ResourceOption) (*Endpoints, error) {
	if args == nil {
		args = &EndpointsArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("Endpoints")
	var resource Endpoints
	err := ctx.RegisterResource("kubernetes:core/v1:Endpoints", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoints gets an existing Endpoints resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoints(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointsState, opts ...pulumi.ResourceOption) (*Endpoints, error) {
	var resource Endpoints
	err := ctx.ReadResource("kubernetes:core/v1:Endpoints", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoints resources.
type endpointsState struct {
}

type EndpointsState struct {
}

func (EndpointsState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointsState)(nil)).Elem()
}

type endpointsArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service.
	Subsets []EndpointSubset `pulumi:"subsets"`
}

// The set of arguments for constructing a Endpoints resource.
type EndpointsArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service.
	Subsets EndpointSubsetArrayInput
}

func (EndpointsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointsArgs)(nil)).Elem()
}

type EndpointsInput interface {
	pulumi.Input

	ToEndpointsOutput() EndpointsOutput
	ToEndpointsOutputWithContext(ctx context.Context) EndpointsOutput
}

func (*Endpoints) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoints)(nil))
}

func (i *Endpoints) ToEndpointsOutput() EndpointsOutput {
	return i.ToEndpointsOutputWithContext(context.Background())
}

func (i *Endpoints) ToEndpointsOutputWithContext(ctx context.Context) EndpointsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsOutput)
}

func (i *Endpoints) ToEndpointsPtrOutput() EndpointsPtrOutput {
	return i.ToEndpointsPtrOutputWithContext(context.Background())
}

func (i *Endpoints) ToEndpointsPtrOutputWithContext(ctx context.Context) EndpointsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsPtrOutput)
}

type EndpointsPtrInput interface {
	pulumi.Input

	ToEndpointsPtrOutput() EndpointsPtrOutput
	ToEndpointsPtrOutputWithContext(ctx context.Context) EndpointsPtrOutput
}

type endpointsPtrType EndpointsArgs

func (*endpointsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoints)(nil))
}

func (i *endpointsPtrType) ToEndpointsPtrOutput() EndpointsPtrOutput {
	return i.ToEndpointsPtrOutputWithContext(context.Background())
}

func (i *endpointsPtrType) ToEndpointsPtrOutputWithContext(ctx context.Context) EndpointsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsPtrOutput)
}

// EndpointsArrayInput is an input type that accepts EndpointsArray and EndpointsArrayOutput values.
// You can construct a concrete instance of `EndpointsArrayInput` via:
//
//          EndpointsArray{ EndpointsArgs{...} }
type EndpointsArrayInput interface {
	pulumi.Input

	ToEndpointsArrayOutput() EndpointsArrayOutput
	ToEndpointsArrayOutputWithContext(context.Context) EndpointsArrayOutput
}

type EndpointsArray []EndpointsInput

func (EndpointsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoints)(nil)).Elem()
}

func (i EndpointsArray) ToEndpointsArrayOutput() EndpointsArrayOutput {
	return i.ToEndpointsArrayOutputWithContext(context.Background())
}

func (i EndpointsArray) ToEndpointsArrayOutputWithContext(ctx context.Context) EndpointsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsArrayOutput)
}

// EndpointsMapInput is an input type that accepts EndpointsMap and EndpointsMapOutput values.
// You can construct a concrete instance of `EndpointsMapInput` via:
//
//          EndpointsMap{ "key": EndpointsArgs{...} }
type EndpointsMapInput interface {
	pulumi.Input

	ToEndpointsMapOutput() EndpointsMapOutput
	ToEndpointsMapOutputWithContext(context.Context) EndpointsMapOutput
}

type EndpointsMap map[string]EndpointsInput

func (EndpointsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoints)(nil)).Elem()
}

func (i EndpointsMap) ToEndpointsMapOutput() EndpointsMapOutput {
	return i.ToEndpointsMapOutputWithContext(context.Background())
}

func (i EndpointsMap) ToEndpointsMapOutputWithContext(ctx context.Context) EndpointsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointsMapOutput)
}

type EndpointsOutput struct{ *pulumi.OutputState }

func (EndpointsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoints)(nil))
}

func (o EndpointsOutput) ToEndpointsOutput() EndpointsOutput {
	return o
}

func (o EndpointsOutput) ToEndpointsOutputWithContext(ctx context.Context) EndpointsOutput {
	return o
}

func (o EndpointsOutput) ToEndpointsPtrOutput() EndpointsPtrOutput {
	return o.ToEndpointsPtrOutputWithContext(context.Background())
}

func (o EndpointsOutput) ToEndpointsPtrOutputWithContext(ctx context.Context) EndpointsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Endpoints) *Endpoints {
		return &v
	}).(EndpointsPtrOutput)
}

type EndpointsPtrOutput struct{ *pulumi.OutputState }

func (EndpointsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoints)(nil))
}

func (o EndpointsPtrOutput) ToEndpointsPtrOutput() EndpointsPtrOutput {
	return o
}

func (o EndpointsPtrOutput) ToEndpointsPtrOutputWithContext(ctx context.Context) EndpointsPtrOutput {
	return o
}

func (o EndpointsPtrOutput) Elem() EndpointsOutput {
	return o.ApplyT(func(v *Endpoints) Endpoints {
		if v != nil {
			return *v
		}
		var ret Endpoints
		return ret
	}).(EndpointsOutput)
}

type EndpointsArrayOutput struct{ *pulumi.OutputState }

func (EndpointsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Endpoints)(nil))
}

func (o EndpointsArrayOutput) ToEndpointsArrayOutput() EndpointsArrayOutput {
	return o
}

func (o EndpointsArrayOutput) ToEndpointsArrayOutputWithContext(ctx context.Context) EndpointsArrayOutput {
	return o
}

func (o EndpointsArrayOutput) Index(i pulumi.IntInput) EndpointsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Endpoints {
		return vs[0].([]Endpoints)[vs[1].(int)]
	}).(EndpointsOutput)
}

type EndpointsMapOutput struct{ *pulumi.OutputState }

func (EndpointsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Endpoints)(nil))
}

func (o EndpointsMapOutput) ToEndpointsMapOutput() EndpointsMapOutput {
	return o
}

func (o EndpointsMapOutput) ToEndpointsMapOutputWithContext(ctx context.Context) EndpointsMapOutput {
	return o
}

func (o EndpointsMapOutput) MapIndex(k pulumi.StringInput) EndpointsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Endpoints {
		return vs[0].(map[string]Endpoints)[vs[1].(string)]
	}).(EndpointsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointsInput)(nil)).Elem(), &Endpoints{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointsPtrInput)(nil)).Elem(), &Endpoints{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointsArrayInput)(nil)).Elem(), EndpointsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointsMapInput)(nil)).Elem(), EndpointsMap{})
	pulumi.RegisterOutputType(EndpointsOutput{})
	pulumi.RegisterOutputType(EndpointsPtrOutput{})
	pulumi.RegisterOutputType(EndpointsArrayOutput{})
	pulumi.RegisterOutputType(EndpointsMapOutput{})
}