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

// The group resource allows you to configure a Boundary group.
//
// ## Example Usage
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
//				Description: pulumi.String("foo user"),
//				ScopeId:     org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewGroup(ctx, "example", &boundary.GroupArgs{
//				Name:        pulumi.String("My group"),
//				Description: pulumi.String("My first group!"),
//				MemberIds: pulumi.StringArray{
//					foo.ID(),
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
// Usage for project-specific group:
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
//			foo, err := boundary.NewUser(ctx, "foo", &boundary.UserArgs{
//				Description: pulumi.String("foo user"),
//				ScopeId:     org.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewGroup(ctx, "example", &boundary.GroupArgs{
//				Name:        pulumi.String("My group"),
//				Description: pulumi.String("My first group!"),
//				MemberIds: pulumi.StringArray{
//					foo.ID(),
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
// $ pulumi import boundary:index/group:Group foo <my-id>
// ```
type Group struct {
	pulumi.CustomResourceState

	// The group description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource IDs for group members, these are most likely boundary users.
	MemberIds pulumi.StringArrayOutput `pulumi:"memberIds"`
	// The group name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Group
	err := ctx.RegisterResource("boundary:index/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("boundary:index/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The group description.
	Description *string `pulumi:"description"`
	// Resource IDs for group members, these are most likely boundary users.
	MemberIds []string `pulumi:"memberIds"`
	// The group name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId *string `pulumi:"scopeId"`
}

type GroupState struct {
	// The group description.
	Description pulumi.StringPtrInput
	// Resource IDs for group members, these are most likely boundary users.
	MemberIds pulumi.StringArrayInput
	// The group name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The group description.
	Description *string `pulumi:"description"`
	// Resource IDs for group members, these are most likely boundary users.
	MemberIds []string `pulumi:"memberIds"`
	// The group name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId string `pulumi:"scopeId"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The group description.
	Description pulumi.StringPtrInput
	// Resource IDs for group members, these are most likely boundary users.
	MemberIds pulumi.StringArrayInput
	// The group name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
	ScopeId pulumi.StringInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//	GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//	GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

// The group description.
func (o GroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource IDs for group members, these are most likely boundary users.
func (o GroupOutput) MemberIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Group) pulumi.StringArrayOutput { return v.MemberIds }).(pulumi.StringArrayOutput)
}

// The group name. Defaults to the resource name.
func (o GroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The scope ID in which the resource is created. Defaults to the provider's `defaultScope` if unset.
func (o GroupOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Group)(nil)).Elem()
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Group {
		return vs[0].([]*Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Group)(nil)).Elem()
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Group {
		return vs[0].(map[string]*Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupArrayInput)(nil)).Elem(), GroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupMapInput)(nil)).Elem(), GroupMap{})
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
