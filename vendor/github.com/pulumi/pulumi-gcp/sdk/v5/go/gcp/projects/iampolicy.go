// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Four different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
//
// * `projects.IAMPolicy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
// * `projects.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
// * `projects.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
// * `projects.IAMAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
//
// > **Note:** `projects.IAMPolicy` **cannot** be used in conjunction with `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig` or they will fight over what your policy should be.
//
// > **Note:** `projects.IAMBinding` resources **can be** used in conjunction with `projects.IAMMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:** The underlying API method `projects.setIamPolicy` has a lot of constraints which are documented [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy). In addition to these constraints,
//    IAM Conditions cannot be used with Basic Roles such as Owner. Violating these constraints will result in the API returning 400 error code so please review these if you encounter errors with this resource.
//
// ## google\_project\_iam\_policy
//
// > **Be careful!** You can accidentally lock yourself out of your project
//    using this resource. Deleting a `projects.IAMPolicy` removes access
//    from anyone without organization-level access to the project. Proceed with caution.
//    It's not recommended to use `projects.IAMPolicy` with your provider project
//    to avoid locking yourself out, and it should generally only be used with projects
//    fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
//    applying the change.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
// 			Project:    pulumi.String("your-project-id"),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 						Title:       "expires_after_2019_12_31",
// 					},
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Role: "roles/compute.admin",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
// 			PolicyData: pulumi.String(admin.PolicyData),
// 			Project:    pulumi.String("your-project-id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_project\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewIAMBinding(ctx, "project", &projects.IAMBindingArgs{
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewIAMBinding(ctx, "project", &projects.IAMBindingArgs{
// 			Condition: &projects.IAMBindingConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/container.admin"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_project\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewIAMMember(ctx, "project", &projects.IAMMemberArgs{
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewIAMMember(ctx, "project", &projects.IAMMemberArgs{
// 			Condition: &projects.IAMMemberConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/firebase.admin"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_project\_iam\_audit\_config
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewIAMAuditConfig(ctx, "project", &projects.IAMAuditConfigArgs{
// 			AuditLogConfigs: projects.IAMAuditConfigAuditLogConfigArray{
// 				&projects.IAMAuditConfigAuditLogConfigArgs{
// 					LogType: pulumi.String("ADMIN_READ"),
// 				},
// 				&projects.IAMAuditConfigAuditLogConfigArgs{
// 					ExemptedMembers: pulumi.StringArray{
// 						pulumi.String("user:joebloggs@hashicorp.com"),
// 					},
// 					LogType: pulumi.String("DATA_READ"),
// 				},
// 			},
// 			Project: pulumi.String("your-project-id"),
// 			Service: pulumi.String("allServices"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `project_id`, role, and member e.g.
//
// ```sh
//  $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project "your-project-id roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `project_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project "your-project-id roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `project_id`.
//
// ```sh
//  $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project your-project-id
// ```
//
//  IAM audit config imports use the identifier of the resource in question and the service, e.g.
//
// ```sh
//  $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project "your-project-id foo.googleapis.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type IAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the project's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewIAMPolicy(ctx *pulumi.Context,
	name string, args *IAMPolicyArgs, opts ...pulumi.ResourceOption) (*IAMPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource IAMPolicy
	err := ctx.RegisterResource("gcp:projects/iAMPolicy:IAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMPolicy gets an existing IAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMPolicyState, opts ...pulumi.ResourceOption) (*IAMPolicy, error) {
	var resource IAMPolicy
	err := ctx.ReadResource("gcp:projects/iAMPolicy:IAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMPolicy resources.
type iampolicyState struct {
	// (Computed) The etag of the project's IAM policy.
	Etag *string `pulumi:"etag"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	PolicyData *string `pulumi:"policyData"`
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project *string `pulumi:"project"`
}

type IAMPolicyState struct {
	// (Computed) The etag of the project's IAM policy.
	Etag pulumi.StringPtrInput
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	PolicyData pulumi.StringPtrInput
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project pulumi.StringPtrInput
}

func (IAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iampolicyState)(nil)).Elem()
}

type iampolicyArgs struct {
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	PolicyData string `pulumi:"policyData"`
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a IAMPolicy resource.
type IAMPolicyArgs struct {
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the project. The policy will be
	// merged with any existing policy applied to the project.
	PolicyData pulumi.StringInput
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project pulumi.StringInput
}

func (IAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iampolicyArgs)(nil)).Elem()
}

type IAMPolicyInput interface {
	pulumi.Input

	ToIAMPolicyOutput() IAMPolicyOutput
	ToIAMPolicyOutputWithContext(ctx context.Context) IAMPolicyOutput
}

func (*IAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*IAMPolicy)(nil))
}

func (i *IAMPolicy) ToIAMPolicyOutput() IAMPolicyOutput {
	return i.ToIAMPolicyOutputWithContext(context.Background())
}

func (i *IAMPolicy) ToIAMPolicyOutputWithContext(ctx context.Context) IAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyOutput)
}

func (i *IAMPolicy) ToIAMPolicyPtrOutput() IAMPolicyPtrOutput {
	return i.ToIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *IAMPolicy) ToIAMPolicyPtrOutputWithContext(ctx context.Context) IAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyPtrOutput)
}

type IAMPolicyPtrInput interface {
	pulumi.Input

	ToIAMPolicyPtrOutput() IAMPolicyPtrOutput
	ToIAMPolicyPtrOutputWithContext(ctx context.Context) IAMPolicyPtrOutput
}

type iampolicyPtrType IAMPolicyArgs

func (*iampolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMPolicy)(nil))
}

func (i *iampolicyPtrType) ToIAMPolicyPtrOutput() IAMPolicyPtrOutput {
	return i.ToIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *iampolicyPtrType) ToIAMPolicyPtrOutputWithContext(ctx context.Context) IAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyPtrOutput)
}

// IAMPolicyArrayInput is an input type that accepts IAMPolicyArray and IAMPolicyArrayOutput values.
// You can construct a concrete instance of `IAMPolicyArrayInput` via:
//
//          IAMPolicyArray{ IAMPolicyArgs{...} }
type IAMPolicyArrayInput interface {
	pulumi.Input

	ToIAMPolicyArrayOutput() IAMPolicyArrayOutput
	ToIAMPolicyArrayOutputWithContext(context.Context) IAMPolicyArrayOutput
}

type IAMPolicyArray []IAMPolicyInput

func (IAMPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMPolicy)(nil)).Elem()
}

func (i IAMPolicyArray) ToIAMPolicyArrayOutput() IAMPolicyArrayOutput {
	return i.ToIAMPolicyArrayOutputWithContext(context.Background())
}

func (i IAMPolicyArray) ToIAMPolicyArrayOutputWithContext(ctx context.Context) IAMPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyArrayOutput)
}

// IAMPolicyMapInput is an input type that accepts IAMPolicyMap and IAMPolicyMapOutput values.
// You can construct a concrete instance of `IAMPolicyMapInput` via:
//
//          IAMPolicyMap{ "key": IAMPolicyArgs{...} }
type IAMPolicyMapInput interface {
	pulumi.Input

	ToIAMPolicyMapOutput() IAMPolicyMapOutput
	ToIAMPolicyMapOutputWithContext(context.Context) IAMPolicyMapOutput
}

type IAMPolicyMap map[string]IAMPolicyInput

func (IAMPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMPolicy)(nil)).Elem()
}

func (i IAMPolicyMap) ToIAMPolicyMapOutput() IAMPolicyMapOutput {
	return i.ToIAMPolicyMapOutputWithContext(context.Background())
}

func (i IAMPolicyMap) ToIAMPolicyMapOutputWithContext(ctx context.Context) IAMPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyMapOutput)
}

type IAMPolicyOutput struct{ *pulumi.OutputState }

func (IAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IAMPolicy)(nil))
}

func (o IAMPolicyOutput) ToIAMPolicyOutput() IAMPolicyOutput {
	return o
}

func (o IAMPolicyOutput) ToIAMPolicyOutputWithContext(ctx context.Context) IAMPolicyOutput {
	return o
}

func (o IAMPolicyOutput) ToIAMPolicyPtrOutput() IAMPolicyPtrOutput {
	return o.ToIAMPolicyPtrOutputWithContext(context.Background())
}

func (o IAMPolicyOutput) ToIAMPolicyPtrOutputWithContext(ctx context.Context) IAMPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IAMPolicy) *IAMPolicy {
		return &v
	}).(IAMPolicyPtrOutput)
}

type IAMPolicyPtrOutput struct{ *pulumi.OutputState }

func (IAMPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMPolicy)(nil))
}

func (o IAMPolicyPtrOutput) ToIAMPolicyPtrOutput() IAMPolicyPtrOutput {
	return o
}

func (o IAMPolicyPtrOutput) ToIAMPolicyPtrOutputWithContext(ctx context.Context) IAMPolicyPtrOutput {
	return o
}

func (o IAMPolicyPtrOutput) Elem() IAMPolicyOutput {
	return o.ApplyT(func(v *IAMPolicy) IAMPolicy {
		if v != nil {
			return *v
		}
		var ret IAMPolicy
		return ret
	}).(IAMPolicyOutput)
}

type IAMPolicyArrayOutput struct{ *pulumi.OutputState }

func (IAMPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IAMPolicy)(nil))
}

func (o IAMPolicyArrayOutput) ToIAMPolicyArrayOutput() IAMPolicyArrayOutput {
	return o
}

func (o IAMPolicyArrayOutput) ToIAMPolicyArrayOutputWithContext(ctx context.Context) IAMPolicyArrayOutput {
	return o
}

func (o IAMPolicyArrayOutput) Index(i pulumi.IntInput) IAMPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IAMPolicy {
		return vs[0].([]IAMPolicy)[vs[1].(int)]
	}).(IAMPolicyOutput)
}

type IAMPolicyMapOutput struct{ *pulumi.OutputState }

func (IAMPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IAMPolicy)(nil))
}

func (o IAMPolicyMapOutput) ToIAMPolicyMapOutput() IAMPolicyMapOutput {
	return o
}

func (o IAMPolicyMapOutput) ToIAMPolicyMapOutputWithContext(ctx context.Context) IAMPolicyMapOutput {
	return o
}

func (o IAMPolicyMapOutput) MapIndex(k pulumi.StringInput) IAMPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IAMPolicy {
		return vs[0].(map[string]IAMPolicy)[vs[1].(string)]
	}).(IAMPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(IAMPolicyOutput{})
	pulumi.RegisterOutputType(IAMPolicyPtrOutput{})
	pulumi.RegisterOutputType(IAMPolicyArrayOutput{})
	pulumi.RegisterOutputType(IAMPolicyMapOutput{})
}
