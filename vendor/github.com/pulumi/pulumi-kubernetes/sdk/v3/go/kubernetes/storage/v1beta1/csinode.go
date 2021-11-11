// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CSINode holds information about all CSI drivers installed on a node. CSI drivers do not need to create the CSINode object directly. As long as they use the node-driver-registrar sidecar container, the kubelet will automatically populate the CSINode object for the CSI driver as part of kubelet plugin registration. CSINode has the same name as a node. If the object is missing, it means either there are no CSI Drivers available on the node, or the Kubelet version is low enough that it doesn't create this object. CSINode has an OwnerReference that points to the corresponding node object.
//
// Deprecated: storage/v1beta1/CSINode is deprecated by storage.k8s.io/v1/CSINode.
type CSINode struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// metadata.name must be the Kubernetes node name.
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// spec is the specification of CSINode
	Spec CSINodeSpecOutput `pulumi:"spec"`
}

// NewCSINode registers a new resource with the given unique name, arguments, and options.
func NewCSINode(ctx *pulumi.Context,
	name string, args *CSINodeArgs, opts ...pulumi.ResourceOption) (*CSINode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("storage.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("CSINode")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:storage.k8s.io/v1:CSINode"),
		},
	})
	opts = append(opts, aliases)
	var resource CSINode
	err := ctx.RegisterResource("kubernetes:storage.k8s.io/v1beta1:CSINode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCSINode gets an existing CSINode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCSINode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CSINodeState, opts ...pulumi.ResourceOption) (*CSINode, error) {
	var resource CSINode
	err := ctx.ReadResource("kubernetes:storage.k8s.io/v1beta1:CSINode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CSINode resources.
type csinodeState struct {
}

type CSINodeState struct {
}

func (CSINodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*csinodeState)(nil)).Elem()
}

type csinodeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// metadata.name must be the Kubernetes node name.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// spec is the specification of CSINode
	Spec CSINodeSpec `pulumi:"spec"`
}

// The set of arguments for constructing a CSINode resource.
type CSINodeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// metadata.name must be the Kubernetes node name.
	Metadata metav1.ObjectMetaPtrInput
	// spec is the specification of CSINode
	Spec CSINodeSpecInput
}

func (CSINodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*csinodeArgs)(nil)).Elem()
}

type CSINodeInput interface {
	pulumi.Input

	ToCSINodeOutput() CSINodeOutput
	ToCSINodeOutputWithContext(ctx context.Context) CSINodeOutput
}

func (*CSINode) ElementType() reflect.Type {
	return reflect.TypeOf((*CSINode)(nil))
}

func (i *CSINode) ToCSINodeOutput() CSINodeOutput {
	return i.ToCSINodeOutputWithContext(context.Background())
}

func (i *CSINode) ToCSINodeOutputWithContext(ctx context.Context) CSINodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSINodeOutput)
}

func (i *CSINode) ToCSINodePtrOutput() CSINodePtrOutput {
	return i.ToCSINodePtrOutputWithContext(context.Background())
}

func (i *CSINode) ToCSINodePtrOutputWithContext(ctx context.Context) CSINodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSINodePtrOutput)
}

type CSINodePtrInput interface {
	pulumi.Input

	ToCSINodePtrOutput() CSINodePtrOutput
	ToCSINodePtrOutputWithContext(ctx context.Context) CSINodePtrOutput
}

type csinodePtrType CSINodeArgs

func (*csinodePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CSINode)(nil))
}

func (i *csinodePtrType) ToCSINodePtrOutput() CSINodePtrOutput {
	return i.ToCSINodePtrOutputWithContext(context.Background())
}

func (i *csinodePtrType) ToCSINodePtrOutputWithContext(ctx context.Context) CSINodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSINodePtrOutput)
}

// CSINodeArrayInput is an input type that accepts CSINodeArray and CSINodeArrayOutput values.
// You can construct a concrete instance of `CSINodeArrayInput` via:
//
//          CSINodeArray{ CSINodeArgs{...} }
type CSINodeArrayInput interface {
	pulumi.Input

	ToCSINodeArrayOutput() CSINodeArrayOutput
	ToCSINodeArrayOutputWithContext(context.Context) CSINodeArrayOutput
}

type CSINodeArray []CSINodeInput

func (CSINodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CSINode)(nil)).Elem()
}

func (i CSINodeArray) ToCSINodeArrayOutput() CSINodeArrayOutput {
	return i.ToCSINodeArrayOutputWithContext(context.Background())
}

func (i CSINodeArray) ToCSINodeArrayOutputWithContext(ctx context.Context) CSINodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSINodeArrayOutput)
}

// CSINodeMapInput is an input type that accepts CSINodeMap and CSINodeMapOutput values.
// You can construct a concrete instance of `CSINodeMapInput` via:
//
//          CSINodeMap{ "key": CSINodeArgs{...} }
type CSINodeMapInput interface {
	pulumi.Input

	ToCSINodeMapOutput() CSINodeMapOutput
	ToCSINodeMapOutputWithContext(context.Context) CSINodeMapOutput
}

type CSINodeMap map[string]CSINodeInput

func (CSINodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CSINode)(nil)).Elem()
}

func (i CSINodeMap) ToCSINodeMapOutput() CSINodeMapOutput {
	return i.ToCSINodeMapOutputWithContext(context.Background())
}

func (i CSINodeMap) ToCSINodeMapOutputWithContext(ctx context.Context) CSINodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSINodeMapOutput)
}

type CSINodeOutput struct{ *pulumi.OutputState }

func (CSINodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CSINode)(nil))
}

func (o CSINodeOutput) ToCSINodeOutput() CSINodeOutput {
	return o
}

func (o CSINodeOutput) ToCSINodeOutputWithContext(ctx context.Context) CSINodeOutput {
	return o
}

func (o CSINodeOutput) ToCSINodePtrOutput() CSINodePtrOutput {
	return o.ToCSINodePtrOutputWithContext(context.Background())
}

func (o CSINodeOutput) ToCSINodePtrOutputWithContext(ctx context.Context) CSINodePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CSINode) *CSINode {
		return &v
	}).(CSINodePtrOutput)
}

type CSINodePtrOutput struct{ *pulumi.OutputState }

func (CSINodePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CSINode)(nil))
}

func (o CSINodePtrOutput) ToCSINodePtrOutput() CSINodePtrOutput {
	return o
}

func (o CSINodePtrOutput) ToCSINodePtrOutputWithContext(ctx context.Context) CSINodePtrOutput {
	return o
}

func (o CSINodePtrOutput) Elem() CSINodeOutput {
	return o.ApplyT(func(v *CSINode) CSINode {
		if v != nil {
			return *v
		}
		var ret CSINode
		return ret
	}).(CSINodeOutput)
}

type CSINodeArrayOutput struct{ *pulumi.OutputState }

func (CSINodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CSINode)(nil))
}

func (o CSINodeArrayOutput) ToCSINodeArrayOutput() CSINodeArrayOutput {
	return o
}

func (o CSINodeArrayOutput) ToCSINodeArrayOutputWithContext(ctx context.Context) CSINodeArrayOutput {
	return o
}

func (o CSINodeArrayOutput) Index(i pulumi.IntInput) CSINodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CSINode {
		return vs[0].([]CSINode)[vs[1].(int)]
	}).(CSINodeOutput)
}

type CSINodeMapOutput struct{ *pulumi.OutputState }

func (CSINodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CSINode)(nil))
}

func (o CSINodeMapOutput) ToCSINodeMapOutput() CSINodeMapOutput {
	return o
}

func (o CSINodeMapOutput) ToCSINodeMapOutputWithContext(ctx context.Context) CSINodeMapOutput {
	return o
}

func (o CSINodeMapOutput) MapIndex(k pulumi.StringInput) CSINodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CSINode {
		return vs[0].(map[string]CSINode)[vs[1].(string)]
	}).(CSINodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CSINodeInput)(nil)).Elem(), &CSINode{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSINodePtrInput)(nil)).Elem(), &CSINode{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSINodeArrayInput)(nil)).Elem(), CSINodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CSINodeMapInput)(nil)).Elem(), CSINodeMap{})
	pulumi.RegisterOutputType(CSINodeOutput{})
	pulumi.RegisterOutputType(CSINodePtrOutput{})
	pulumi.RegisterOutputType(CSINodeArrayOutput{})
	pulumi.RegisterOutputType(CSINodeMapOutput{})
}
