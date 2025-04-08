// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package boundary

import (
	"context"
	"reflect"

	"errors"
	"github.com/joerit/pulumi-boundary/sdk/go/boundary/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The role resource allows you to configure a Boundary role.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/joerit/pulumi-boundary/sdk/go/boundary"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			org, err := boundary.NewScope(ctx, "org", &boundary.ScopeArgs{
//				Name:                  pulumi.String("organization_one"),
//				Description:           pulumi.String("My first scope!"),
//				ScopeId:               pulumi.String("global"),
//				AutoCreateAdminRole:   pulumi.Bool(true),
//				AutoCreateDefaultRole: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewRole(ctx, "example", &boundary.RoleArgs{
//				Name:        pulumi.String("My role"),
//				Description: pulumi.String("My first role!"),
//				ScopeId:     org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Usage with a user resource:
//
// ```go
// package main
//
// import (
//
//	"github.com/joerit/pulumi-boundary/sdk/go/boundary"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			org, err := boundary.NewScope(ctx, "org", &boundary.ScopeArgs{
//				Name:                  pulumi.String("organization_one"),
//				Description:           pulumi.String("My first scope!"),
//				ScopeId:               pulumi.String("global"),
//				AutoCreateAdminRole:   pulumi.Bool(true),
//				AutoCreateDefaultRole: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := boundary.NewUser(ctx, "foo", &boundary.UserArgs{
//				Name:    pulumi.String("User 1"),
//				ScopeId: org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			bar, err := boundary.NewUser(ctx, "bar", &boundary.UserArgs{
//				Name:    pulumi.String("User 2"),
//				ScopeId: org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewRole(ctx, "example", &boundary.RoleArgs{
//				Name:        pulumi.String("My role"),
//				Description: pulumi.String("My first role!"),
//				PrincipalIds: pulumi.StringArray{
//					foo.ID(),
//					bar.ID(),
//				},
//				ScopeId: org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Usage with user and grants resource:
//
// ```go
// package main
//
// import (
//
//	"github.com/joerit/pulumi-boundary/sdk/go/boundary"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			org, err := boundary.NewScope(ctx, "org", &boundary.ScopeArgs{
//				Name:                  pulumi.String("organization_one"),
//				Description:           pulumi.String("My first scope!"),
//				ScopeId:               pulumi.String("global"),
//				AutoCreateAdminRole:   pulumi.Bool(true),
//				AutoCreateDefaultRole: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			readonly, err := boundary.NewUser(ctx, "readonly", &boundary.UserArgs{
//				Name:        pulumi.String("readonly"),
//				Description: pulumi.String("A readonly user"),
//				ScopeId:     org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewRole(ctx, "readonly", &boundary.RoleArgs{
//				Name:        pulumi.String("readonly"),
//				Description: pulumi.String("A readonly role"),
//				PrincipalIds: pulumi.StringArray{
//					readonly.ID(),
//				},
//				GrantStrings: pulumi.StringArray{
//					pulumi.String("ids=*;type=*;actions=read"),
//				},
//				ScopeId: org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Usage for a project-specific role:
//
// ```go
// package main
//
// import (
//
//	"github.com/joerit/pulumi-boundary/sdk/go/boundary"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			org, err := boundary.NewScope(ctx, "org", &boundary.ScopeArgs{
//				Name:                  pulumi.String("organization_one"),
//				Description:           pulumi.String("My first scope!"),
//				ScopeId:               pulumi.String("global"),
//				AutoCreateAdminRole:   pulumi.Bool(true),
//				AutoCreateDefaultRole: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			project, err := boundary.NewScope(ctx, "project", &boundary.ScopeArgs{
//				Name:                pulumi.String("project_one"),
//				Description:         pulumi.String("My first scope!"),
//				ScopeId:             org.ID(),
//				AutoCreateAdminRole: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			readonly, err := boundary.NewUser(ctx, "readonly", &boundary.UserArgs{
//				Name:        pulumi.String("readonly"),
//				Description: pulumi.String("A readonly user"),
//				ScopeId:     org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewRole(ctx, "readonly", &boundary.RoleArgs{
//				Name:        pulumi.String("readonly"),
//				Description: pulumi.String("A readonly role"),
//				PrincipalIds: pulumi.StringArray{
//					readonly.ID(),
//				},
//				GrantStrings: pulumi.StringArray{
//					pulumi.String("ids=*;type=*;actions=read"),
//				},
//				ScopeId: project.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
// $ pulumi import boundary:index/role:Role foo <my-id>
// ```
type Role struct {
	pulumi.CustomResourceState

	// The role description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants"
	GrantScopeIds pulumi.StringArrayOutput `pulumi:"grantScopeIds"`
	// A list of stringified grants for the role.
	GrantStrings pulumi.StringArrayOutput `pulumi:"grantStrings"`
	// The role name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of principal (user or group) IDs to add as principals on the role.
	PrincipalIds pulumi.StringArrayOutput `pulumi:"principalIds"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Role
	err := ctx.RegisterResource("boundary:index/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("boundary:index/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// The role description.
	Description *string `pulumi:"description"`
	// A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants"
	GrantScopeIds []string `pulumi:"grantScopeIds"`
	// A list of stringified grants for the role.
	GrantStrings []string `pulumi:"grantStrings"`
	// The role name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// A list of principal (user or group) IDs to add as principals on the role.
	PrincipalIds []string `pulumi:"principalIds"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId *string `pulumi:"scopeId"`
}

type RoleState struct {
	// The role description.
	Description pulumi.StringPtrInput
	// A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants"
	GrantScopeIds pulumi.StringArrayInput
	// A list of stringified grants for the role.
	GrantStrings pulumi.StringArrayInput
	// The role name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// A list of principal (user or group) IDs to add as principals on the role.
	PrincipalIds pulumi.StringArrayInput
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// The role description.
	Description *string `pulumi:"description"`
	// A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants"
	GrantScopeIds []string `pulumi:"grantScopeIds"`
	// A list of stringified grants for the role.
	GrantStrings []string `pulumi:"grantStrings"`
	// The role name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// A list of principal (user or group) IDs to add as principals on the role.
	PrincipalIds []string `pulumi:"principalIds"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId string `pulumi:"scopeId"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// The role description.
	Description pulumi.StringPtrInput
	// A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants"
	GrantScopeIds pulumi.StringArrayInput
	// A list of stringified grants for the role.
	GrantStrings pulumi.StringArrayInput
	// The role name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// A list of principal (user or group) IDs to add as principals on the role.
	PrincipalIds pulumi.StringArrayInput
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

// RoleArrayInput is an input type that accepts RoleArray and RoleArrayOutput values.
// You can construct a concrete instance of `RoleArrayInput` via:
//
//	RoleArray{ RoleArgs{...} }
type RoleArrayInput interface {
	pulumi.Input

	ToRoleArrayOutput() RoleArrayOutput
	ToRoleArrayOutputWithContext(context.Context) RoleArrayOutput
}

type RoleArray []RoleInput

func (RoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (i RoleArray) ToRoleArrayOutput() RoleArrayOutput {
	return i.ToRoleArrayOutputWithContext(context.Background())
}

func (i RoleArray) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleArrayOutput)
}

// RoleMapInput is an input type that accepts RoleMap and RoleMapOutput values.
// You can construct a concrete instance of `RoleMapInput` via:
//
//	RoleMap{ "key": RoleArgs{...} }
type RoleMapInput interface {
	pulumi.Input

	ToRoleMapOutput() RoleMapOutput
	ToRoleMapOutputWithContext(context.Context) RoleMapOutput
}

type RoleMap map[string]RoleInput

func (RoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (i RoleMap) ToRoleMapOutput() RoleMapOutput {
	return i.ToRoleMapOutputWithContext(context.Background())
}

func (i RoleMap) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapOutput)
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

// The role description.
func (o RoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A list of scopes for which the grants in this role should apply, which can include the special values "this", "children", or "descendants"
func (o RoleOutput) GrantScopeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Role) pulumi.StringArrayOutput { return v.GrantScopeIds }).(pulumi.StringArrayOutput)
}

// A list of stringified grants for the role.
func (o RoleOutput) GrantStrings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Role) pulumi.StringArrayOutput { return v.GrantStrings }).(pulumi.StringArrayOutput)
}

// The role name. Defaults to the resource name.
func (o RoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of principal (user or group) IDs to add as principals on the role.
func (o RoleOutput) PrincipalIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Role) pulumi.StringArrayOutput { return v.PrincipalIds }).(pulumi.StringArrayOutput)
}

// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
func (o RoleOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

type RoleArrayOutput struct{ *pulumi.OutputState }

func (RoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (o RoleArrayOutput) ToRoleArrayOutput() RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) Index(i pulumi.IntInput) RoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Role {
		return vs[0].([]*Role)[vs[1].(int)]
	}).(RoleOutput)
}

type RoleMapOutput struct{ *pulumi.OutputState }

func (RoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (o RoleMapOutput) ToRoleMapOutput() RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return o
}

func (o RoleMapOutput) MapIndex(k pulumi.StringInput) RoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Role {
		return vs[0].(map[string]*Role)[vs[1].(string)]
	}).(RoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleInput)(nil)).Elem(), &Role{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleArrayInput)(nil)).Elem(), RoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapInput)(nil)).Elem(), RoleMap{})
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RoleArrayOutput{})
	pulumi.RegisterOutputType(RoleMapOutput{})
}
