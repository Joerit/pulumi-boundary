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

// The static host catalog resource allows you to configure a Boundary static-type host catalog. Host catalogs are always part of a project, so a project resource should be used inline or you should have the project ID in hand to successfully configure a host catalog.
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
//				ScopeId:               pulumi.Any(global.Id),
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
//			_, err = boundary.NewHostCatalogStatic(ctx, "example", &boundary.HostCatalogStaticArgs{
//				Name:        pulumi.String("My catalog"),
//				Description: pulumi.String("My first host catalog!"),
//				ScopeId:     project.ID(),
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
// $ pulumi import boundary:index/hostCatalogStatic:HostCatalogStatic foo <my-id>
// ```
type HostCatalogStatic struct {
	pulumi.CustomResourceState

	// The host catalog description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The host catalog name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The scope ID in which the resource is created.
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
}

// NewHostCatalogStatic registers a new resource with the given unique name, arguments, and options.
func NewHostCatalogStatic(ctx *pulumi.Context,
	name string, args *HostCatalogStaticArgs, opts ...pulumi.ResourceOption) (*HostCatalogStatic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HostCatalogStatic
	err := ctx.RegisterResource("boundary:index/hostCatalogStatic:HostCatalogStatic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostCatalogStatic gets an existing HostCatalogStatic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostCatalogStatic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostCatalogStaticState, opts ...pulumi.ResourceOption) (*HostCatalogStatic, error) {
	var resource HostCatalogStatic
	err := ctx.ReadResource("boundary:index/hostCatalogStatic:HostCatalogStatic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostCatalogStatic resources.
type hostCatalogStaticState struct {
	// The host catalog description.
	Description *string `pulumi:"description"`
	// The host catalog name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID in which the resource is created.
	ScopeId *string `pulumi:"scopeId"`
}

type HostCatalogStaticState struct {
	// The host catalog description.
	Description pulumi.StringPtrInput
	// The host catalog name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID in which the resource is created.
	ScopeId pulumi.StringPtrInput
}

func (HostCatalogStaticState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostCatalogStaticState)(nil)).Elem()
}

type hostCatalogStaticArgs struct {
	// The host catalog description.
	Description *string `pulumi:"description"`
	// The host catalog name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The scope ID in which the resource is created.
	ScopeId string `pulumi:"scopeId"`
}

// The set of arguments for constructing a HostCatalogStatic resource.
type HostCatalogStaticArgs struct {
	// The host catalog description.
	Description pulumi.StringPtrInput
	// The host catalog name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The scope ID in which the resource is created.
	ScopeId pulumi.StringInput
}

func (HostCatalogStaticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostCatalogStaticArgs)(nil)).Elem()
}

type HostCatalogStaticInput interface {
	pulumi.Input

	ToHostCatalogStaticOutput() HostCatalogStaticOutput
	ToHostCatalogStaticOutputWithContext(ctx context.Context) HostCatalogStaticOutput
}

func (*HostCatalogStatic) ElementType() reflect.Type {
	return reflect.TypeOf((**HostCatalogStatic)(nil)).Elem()
}

func (i *HostCatalogStatic) ToHostCatalogStaticOutput() HostCatalogStaticOutput {
	return i.ToHostCatalogStaticOutputWithContext(context.Background())
}

func (i *HostCatalogStatic) ToHostCatalogStaticOutputWithContext(ctx context.Context) HostCatalogStaticOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostCatalogStaticOutput)
}

// HostCatalogStaticArrayInput is an input type that accepts HostCatalogStaticArray and HostCatalogStaticArrayOutput values.
// You can construct a concrete instance of `HostCatalogStaticArrayInput` via:
//
//	HostCatalogStaticArray{ HostCatalogStaticArgs{...} }
type HostCatalogStaticArrayInput interface {
	pulumi.Input

	ToHostCatalogStaticArrayOutput() HostCatalogStaticArrayOutput
	ToHostCatalogStaticArrayOutputWithContext(context.Context) HostCatalogStaticArrayOutput
}

type HostCatalogStaticArray []HostCatalogStaticInput

func (HostCatalogStaticArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostCatalogStatic)(nil)).Elem()
}

func (i HostCatalogStaticArray) ToHostCatalogStaticArrayOutput() HostCatalogStaticArrayOutput {
	return i.ToHostCatalogStaticArrayOutputWithContext(context.Background())
}

func (i HostCatalogStaticArray) ToHostCatalogStaticArrayOutputWithContext(ctx context.Context) HostCatalogStaticArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostCatalogStaticArrayOutput)
}

// HostCatalogStaticMapInput is an input type that accepts HostCatalogStaticMap and HostCatalogStaticMapOutput values.
// You can construct a concrete instance of `HostCatalogStaticMapInput` via:
//
//	HostCatalogStaticMap{ "key": HostCatalogStaticArgs{...} }
type HostCatalogStaticMapInput interface {
	pulumi.Input

	ToHostCatalogStaticMapOutput() HostCatalogStaticMapOutput
	ToHostCatalogStaticMapOutputWithContext(context.Context) HostCatalogStaticMapOutput
}

type HostCatalogStaticMap map[string]HostCatalogStaticInput

func (HostCatalogStaticMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostCatalogStatic)(nil)).Elem()
}

func (i HostCatalogStaticMap) ToHostCatalogStaticMapOutput() HostCatalogStaticMapOutput {
	return i.ToHostCatalogStaticMapOutputWithContext(context.Background())
}

func (i HostCatalogStaticMap) ToHostCatalogStaticMapOutputWithContext(ctx context.Context) HostCatalogStaticMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostCatalogStaticMapOutput)
}

type HostCatalogStaticOutput struct{ *pulumi.OutputState }

func (HostCatalogStaticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostCatalogStatic)(nil)).Elem()
}

func (o HostCatalogStaticOutput) ToHostCatalogStaticOutput() HostCatalogStaticOutput {
	return o
}

func (o HostCatalogStaticOutput) ToHostCatalogStaticOutputWithContext(ctx context.Context) HostCatalogStaticOutput {
	return o
}

// The host catalog description.
func (o HostCatalogStaticOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostCatalogStatic) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The host catalog name. Defaults to the resource name.
func (o HostCatalogStaticOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HostCatalogStatic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The scope ID in which the resource is created.
func (o HostCatalogStaticOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostCatalogStatic) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

type HostCatalogStaticArrayOutput struct{ *pulumi.OutputState }

func (HostCatalogStaticArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostCatalogStatic)(nil)).Elem()
}

func (o HostCatalogStaticArrayOutput) ToHostCatalogStaticArrayOutput() HostCatalogStaticArrayOutput {
	return o
}

func (o HostCatalogStaticArrayOutput) ToHostCatalogStaticArrayOutputWithContext(ctx context.Context) HostCatalogStaticArrayOutput {
	return o
}

func (o HostCatalogStaticArrayOutput) Index(i pulumi.IntInput) HostCatalogStaticOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HostCatalogStatic {
		return vs[0].([]*HostCatalogStatic)[vs[1].(int)]
	}).(HostCatalogStaticOutput)
}

type HostCatalogStaticMapOutput struct{ *pulumi.OutputState }

func (HostCatalogStaticMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostCatalogStatic)(nil)).Elem()
}

func (o HostCatalogStaticMapOutput) ToHostCatalogStaticMapOutput() HostCatalogStaticMapOutput {
	return o
}

func (o HostCatalogStaticMapOutput) ToHostCatalogStaticMapOutputWithContext(ctx context.Context) HostCatalogStaticMapOutput {
	return o
}

func (o HostCatalogStaticMapOutput) MapIndex(k pulumi.StringInput) HostCatalogStaticOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HostCatalogStatic {
		return vs[0].(map[string]*HostCatalogStatic)[vs[1].(string)]
	}).(HostCatalogStaticOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostCatalogStaticInput)(nil)).Elem(), &HostCatalogStatic{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostCatalogStaticArrayInput)(nil)).Elem(), HostCatalogStaticArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostCatalogStaticMapInput)(nil)).Elem(), HostCatalogStaticMap{})
	pulumi.RegisterOutputType(HostCatalogStaticOutput{})
	pulumi.RegisterOutputType(HostCatalogStaticArrayOutput{})
	pulumi.RegisterOutputType(HostCatalogStaticMapOutput{})
}
