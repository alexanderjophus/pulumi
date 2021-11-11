// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CSIStorageCapacity stores the result of one CSI GetCapacity call. For a given StorageClass, this describes the available capacity in a particular topology segment.  This can be used when considering where to instantiate new PersistentVolumes.
//
// For example this can express things like: - StorageClass "standard" has "1234 GiB" available in "topology.kubernetes.io/zone=us-east1" - StorageClass "localssd" has "10 GiB" available in "kubernetes.io/hostname=knode-abc123"
//
// The following three cases all imply that no capacity is available for a certain combination: - no object exists with suitable topology and storage class name - such an object exists, but the capacity is unset - such an object exists, but the capacity is zero
//
// The producer of these objects can decide which approach is more suitable.
//
// They are consumed by the kube-scheduler if the CSIStorageCapacity beta feature gate is enabled there and a CSI driver opts into capacity-aware scheduling with CSIDriver.StorageCapacity.
type CSIStorageCapacity struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Capacity is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// The semantic is currently (CSI spec 1.2) defined as: The available capacity, in bytes, of the storage that can be used to provision volumes. If not set, that information is currently unavailable and treated like zero capacity.
	Capacity pulumi.StringPtrOutput `pulumi:"capacity"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// MaximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// This is defined since CSI spec 1.4.0 as the largest size that may be used in a CreateVolumeRequest.capacity_range.required_bytes field to create a volume with the same parameters as those in GetCapacityRequest. The corresponding value in the Kubernetes API is ResourceRequirements.Requests in a volume claim.
	MaximumVolumeSize pulumi.StringPtrOutput `pulumi:"maximumVolumeSize"`
	// Standard object's metadata. The name has no particular meaning. It must be be a DNS subdomain (dots allowed, 253 characters). To ensure that there are no conflicts with other CSI drivers on the cluster, the recommendation is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends with the unique CSI driver name.
	//
	// Objects are namespaced.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// NodeTopology defines which nodes have access to the storage for which capacity was reported. If not set, the storage is not accessible from any node in the cluster. If empty, the storage is accessible from all nodes. This field is immutable.
	NodeTopology metav1.LabelSelectorPtrOutput `pulumi:"nodeTopology"`
	// The name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.
	StorageClassName pulumi.StringOutput `pulumi:"storageClassName"`
}

// NewCSIStorageCapacity registers a new resource with the given unique name, arguments, and options.
func NewCSIStorageCapacity(ctx *pulumi.Context,
	name string, args *CSIStorageCapacityArgs, opts ...pulumi.ResourceOption) (*CSIStorageCapacity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StorageClassName == nil {
		return nil, errors.New("invalid value for required argument 'StorageClassName'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("CSIStorageCapacity")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1beta1:CSIStorageCapacity"),
		},
	})
	opts = append(opts, aliases)
	var resource CSIStorageCapacity
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1alpha1:CSIStorageCapacity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSIStorageCapacity gets an existing CSIStorageCapacity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSIStorageCapacity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSIStorageCapacityState, opts ...pulumi.ResourceOption) (*CSIStorageCapacity, error) {
	var resource CSIStorageCapacity
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1alpha1:CSIStorageCapacity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSIStorageCapacity resources.
type csistorageCapacityState struct {
}

type CSIStorageCapacityState struct {
}

func (CSIStorageCapacityState) ElementType() reflect.Type {
	return reflect.TypeOf((*csistorageCapacityState)(nil)).Elem()
}

type csistorageCapacityArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Capacity is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// The semantic is currently (CSI spec 1.2) defined as: The available capacity, in bytes, of the storage that can be used to provision volumes. If not set, that information is currently unavailable and treated like zero capacity.
	Capacity *string `pulumi:"capacity"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// MaximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// This is defined since CSI spec 1.4.0 as the largest size that may be used in a CreateVolumeRequest.capacity_range.required_bytes field to create a volume with the same parameters as those in GetCapacityRequest. The corresponding value in the Kubernetes API is ResourceRequirements.Requests in a volume claim.
	MaximumVolumeSize *string `pulumi:"maximumVolumeSize"`
	// Standard object's metadata. The name has no particular meaning. It must be be a DNS subdomain (dots allowed, 253 characters). To ensure that there are no conflicts with other CSI drivers on the cluster, the recommendation is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends with the unique CSI driver name.
	//
	// Objects are namespaced.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// NodeTopology defines which nodes have access to the storage for which capacity was reported. If not set, the storage is not accessible from any node in the cluster. If empty, the storage is accessible from all nodes. This field is immutable.
	NodeTopology *metav1.LabelSelector `pulumi:"nodeTopology"`
	// The name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.
	StorageClassName string `pulumi:"storageClassName"`
}

// The set of arguments for constructing a CSIStorageCapacity resource.
type CSIStorageCapacityArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Capacity is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// The semantic is currently (CSI spec 1.2) defined as: The available capacity, in bytes, of the storage that can be used to provision volumes. If not set, that information is currently unavailable and treated like zero capacity.
	Capacity pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// MaximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.
	//
	// This is defined since CSI spec 1.4.0 as the largest size that may be used in a CreateVolumeRequest.capacity_range.required_bytes field to create a volume with the same parameters as those in GetCapacityRequest. The corresponding value in the Kubernetes API is ResourceRequirements.Requests in a volume claim.
	MaximumVolumeSize pulumi.StringPtrInput
	// Standard object's metadata. The name has no particular meaning. It must be be a DNS subdomain (dots allowed, 253 characters). To ensure that there are no conflicts with other CSI drivers on the cluster, the recommendation is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends with the unique CSI driver name.
	//
	// Objects are namespaced.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// NodeTopology defines which nodes have access to the storage for which capacity was reported. If not set, the storage is not accessible from any node in the cluster. If empty, the storage is accessible from all nodes. This field is immutable.
	NodeTopology metav1.LabelSelectorPtrInput
	// The name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.
	StorageClassName pulumi.StringInput
}

func (CSIStorageCapacityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csistorageCapacityArgs)(nil)).Elem()
}

type CSIStorageCapacityInput interface {
	pulumi.Input

	ToCSIStorageCapacityOutput() CSIStorageCapacityOutput
	ToCSIStorageCapacityOutputWithContext(ctx context.Context) CSIStorageCapacityOutput
}

func (*CSIStorageCapacity) ElementType() reflect.Type {
	return reflect.TypeOf((*CSIStorageCapacity)(nil))
}

func (i *CSIStorageCapacity) ToCSIStorageCapacityOutput() CSIStorageCapacityOutput {
	return i.ToCSIStorageCapacityOutputWithContext(context.Background())
}

func (i *CSIStorageCapacity) ToCSIStorageCapacityOutputWithContext(ctx context.Context) CSIStorageCapacityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityOutput)
}

func (i *CSIStorageCapacity) ToCSIStorageCapacityPtrOutput() CSIStorageCapacityPtrOutput {
	return i.ToCSIStorageCapacityPtrOutputWithContext(context.Background())
}

func (i *CSIStorageCapacity) ToCSIStorageCapacityPtrOutputWithContext(ctx context.Context) CSIStorageCapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityPtrOutput)
}

type CSIStorageCapacityPtrInput interface {
	pulumi.Input

	ToCSIStorageCapacityPtrOutput() CSIStorageCapacityPtrOutput
	ToCSIStorageCapacityPtrOutputWithContext(ctx context.Context) CSIStorageCapacityPtrOutput
}

type csistorageCapacityPtrType CSIStorageCapacityArgs

func (*csistorageCapacityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIStorageCapacity)(nil))
}

func (i *csistorageCapacityPtrType) ToCSIStorageCapacityPtrOutput() CSIStorageCapacityPtrOutput {
	return i.ToCSIStorageCapacityPtrOutputWithContext(context.Background())
}

func (i *csistorageCapacityPtrType) ToCSIStorageCapacityPtrOutputWithContext(ctx context.Context) CSIStorageCapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityPtrOutput)
}

// CSIStorageCapacityArrayInput is an input type that accepts CSIStorageCapacityArray and CSIStorageCapacityArrayOutput values.
// You can construct a concrete instance of `CSIStorageCapacityArrayInput` via:
//
//          CSIStorageCapacityArray{ CSIStorageCapacityArgs{...} }
type CSIStorageCapacityArrayInput interface {
	pulumi.Input

	ToCSIStorageCapacityArrayOutput() CSIStorageCapacityArrayOutput
	ToCSIStorageCapacityArrayOutputWithContext(context.Context) CSIStorageCapacityArrayOutput
}

type CSIStorageCapacityArray []CSIStorageCapacityInput

func (CSIStorageCapacityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSIStorageCapacity)(nil)).Elem()
}

func (i CSIStorageCapacityArray) ToCSIStorageCapacityArrayOutput() CSIStorageCapacityArrayOutput {
	return i.ToCSIStorageCapacityArrayOutputWithContext(context.Background())
}

func (i CSIStorageCapacityArray) ToCSIStorageCapacityArrayOutputWithContext(ctx context.Context) CSIStorageCapacityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityArrayOutput)
}

// CSIStorageCapacityMapInput is an input type that accepts CSIStorageCapacityMap and CSIStorageCapacityMapOutput values.
// You can construct a concrete instance of `CSIStorageCapacityMapInput` via:
//
//          CSIStorageCapacityMap{ "key": CSIStorageCapacityArgs{...} }
type CSIStorageCapacityMapInput interface {
	pulumi.Input

	ToCSIStorageCapacityMapOutput() CSIStorageCapacityMapOutput
	ToCSIStorageCapacityMapOutputWithContext(context.Context) CSIStorageCapacityMapOutput
}

type CSIStorageCapacityMap map[string]CSIStorageCapacityInput

func (CSIStorageCapacityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSIStorageCapacity)(nil)).Elem()
}

func (i CSIStorageCapacityMap) ToCSIStorageCapacityMapOutput() CSIStorageCapacityMapOutput {
	return i.ToCSIStorageCapacityMapOutputWithContext(context.Background())
}

func (i CSIStorageCapacityMap) ToCSIStorageCapacityMapOutputWithContext(ctx context.Context) CSIStorageCapacityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSIStorageCapacityMapOutput)
}

type CSIStorageCapacityOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CSIStorageCapacity)(nil))
}

func (o CSIStorageCapacityOutput) ToCSIStorageCapacityOutput() CSIStorageCapacityOutput {
	return o
}

func (o CSIStorageCapacityOutput) ToCSIStorageCapacityOutputWithContext(ctx context.Context) CSIStorageCapacityOutput {
	return o
}

func (o CSIStorageCapacityOutput) ToCSIStorageCapacityPtrOutput() CSIStorageCapacityPtrOutput {
	return o.ToCSIStorageCapacityPtrOutputWithContext(context.Background())
}

func (o CSIStorageCapacityOutput) ToCSIStorageCapacityPtrOutputWithContext(ctx context.Context) CSIStorageCapacityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CSIStorageCapacity) *CSIStorageCapacity {
		return &v
	}).(CSIStorageCapacityPtrOutput)
}

type CSIStorageCapacityPtrOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CSIStorageCapacity)(nil))
}

func (o CSIStorageCapacityPtrOutput) ToCSIStorageCapacityPtrOutput() CSIStorageCapacityPtrOutput {
	return o
}

func (o CSIStorageCapacityPtrOutput) ToCSIStorageCapacityPtrOutputWithContext(ctx context.Context) CSIStorageCapacityPtrOutput {
	return o
}

func (o CSIStorageCapacityPtrOutput) Elem() CSIStorageCapacityOutput {
	return o.ApplyT(func(v *CSIStorageCapacity) CSIStorageCapacity {
		if v != nil {
			return *v
		}
		var ret CSIStorageCapacity
		return ret
	}).(CSIStorageCapacityOutput)
}

type CSIStorageCapacityArrayOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CSIStorageCapacity)(nil))
}

func (o CSIStorageCapacityArrayOutput) ToCSIStorageCapacityArrayOutput() CSIStorageCapacityArrayOutput {
	return o
}

func (o CSIStorageCapacityArrayOutput) ToCSIStorageCapacityArrayOutputWithContext(ctx context.Context) CSIStorageCapacityArrayOutput {
	return o
}

func (o CSIStorageCapacityArrayOutput) Index(i pulumi.IntInput) CSIStorageCapacityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CSIStorageCapacity {
		return vs[0].([]CSIStorageCapacity)[vs[1].(int)]
	}).(CSIStorageCapacityOutput)
}

type CSIStorageCapacityMapOutput struct{ *pulumi.OutputState }

func (CSIStorageCapacityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CSIStorageCapacity)(nil))
}

func (o CSIStorageCapacityMapOutput) ToCSIStorageCapacityMapOutput() CSIStorageCapacityMapOutput {
	return o
}

func (o CSIStorageCapacityMapOutput) ToCSIStorageCapacityMapOutputWithContext(ctx context.Context) CSIStorageCapacityMapOutput {
	return o
}

func (o CSIStorageCapacityMapOutput) MapIndex(k pulumi.StringInput) CSIStorageCapacityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CSIStorageCapacity {
		return vs[0].(map[string]CSIStorageCapacity)[vs[1].(string)]
	}).(CSIStorageCapacityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityInput)(nil)).Elem(), &CSIStorageCapacity{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityPtrInput)(nil)).Elem(), &CSIStorageCapacity{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityArrayInput)(nil)).Elem(), CSIStorageCapacityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSIStorageCapacityMapInput)(nil)).Elem(), CSIStorageCapacityMap{})
	pulumi.RegisterOutputType(CSIStorageCapacityOutput{})
	pulumi.RegisterOutputType(CSIStorageCapacityPtrOutput{})
	pulumi.RegisterOutputType(CSIStorageCapacityArrayOutput{})
	pulumi.RegisterOutputType(CSIStorageCapacityMapOutput{})
}
