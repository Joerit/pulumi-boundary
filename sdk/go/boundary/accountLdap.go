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

type AccountLdap struct {
	pulumi.CustomResourceState

	// The resource ID for the auth method.
	AuthMethodId pulumi.StringOutput `pulumi:"authMethodId"`
	// The account description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The login name for this account.
	LoginName pulumi.StringPtrOutput `pulumi:"loginName"`
	// The account name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource type.
	//
	// Deprecated: The value for this field will be infered since 'ldap' is the only possible value.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewAccountLdap registers a new resource with the given unique name, arguments, and options.
func NewAccountLdap(ctx *pulumi.Context,
	name string, args *AccountLdapArgs, opts ...pulumi.ResourceOption) (*AccountLdap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthMethodId == nil {
		return nil, errors.New("invalid value for required argument 'AuthMethodId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccountLdap
	err := ctx.RegisterResource("boundary:index/accountLdap:AccountLdap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountLdap gets an existing AccountLdap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountLdap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountLdapState, opts ...pulumi.ResourceOption) (*AccountLdap, error) {
	var resource AccountLdap
	err := ctx.ReadResource("boundary:index/accountLdap:AccountLdap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountLdap resources.
type accountLdapState struct {
	// The resource ID for the auth method.
	AuthMethodId *string `pulumi:"authMethodId"`
	// The account description.
	Description *string `pulumi:"description"`
	// The login name for this account.
	LoginName *string `pulumi:"loginName"`
	// The account name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The resource type.
	//
	// Deprecated: The value for this field will be infered since 'ldap' is the only possible value.
	Type *string `pulumi:"type"`
}

type AccountLdapState struct {
	// The resource ID for the auth method.
	AuthMethodId pulumi.StringPtrInput
	// The account description.
	Description pulumi.StringPtrInput
	// The login name for this account.
	LoginName pulumi.StringPtrInput
	// The account name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The resource type.
	//
	// Deprecated: The value for this field will be infered since 'ldap' is the only possible value.
	Type pulumi.StringPtrInput
}

func (AccountLdapState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountLdapState)(nil)).Elem()
}

type accountLdapArgs struct {
	// The resource ID for the auth method.
	AuthMethodId string `pulumi:"authMethodId"`
	// The account description.
	Description *string `pulumi:"description"`
	// The login name for this account.
	LoginName *string `pulumi:"loginName"`
	// The account name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The resource type.
	//
	// Deprecated: The value for this field will be infered since 'ldap' is the only possible value.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a AccountLdap resource.
type AccountLdapArgs struct {
	// The resource ID for the auth method.
	AuthMethodId pulumi.StringInput
	// The account description.
	Description pulumi.StringPtrInput
	// The login name for this account.
	LoginName pulumi.StringPtrInput
	// The account name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The resource type.
	//
	// Deprecated: The value for this field will be infered since 'ldap' is the only possible value.
	Type pulumi.StringPtrInput
}

func (AccountLdapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountLdapArgs)(nil)).Elem()
}

type AccountLdapInput interface {
	pulumi.Input

	ToAccountLdapOutput() AccountLdapOutput
	ToAccountLdapOutputWithContext(ctx context.Context) AccountLdapOutput
}

func (*AccountLdap) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountLdap)(nil)).Elem()
}

func (i *AccountLdap) ToAccountLdapOutput() AccountLdapOutput {
	return i.ToAccountLdapOutputWithContext(context.Background())
}

func (i *AccountLdap) ToAccountLdapOutputWithContext(ctx context.Context) AccountLdapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountLdapOutput)
}

// AccountLdapArrayInput is an input type that accepts AccountLdapArray and AccountLdapArrayOutput values.
// You can construct a concrete instance of `AccountLdapArrayInput` via:
//
//	AccountLdapArray{ AccountLdapArgs{...} }
type AccountLdapArrayInput interface {
	pulumi.Input

	ToAccountLdapArrayOutput() AccountLdapArrayOutput
	ToAccountLdapArrayOutputWithContext(context.Context) AccountLdapArrayOutput
}

type AccountLdapArray []AccountLdapInput

func (AccountLdapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountLdap)(nil)).Elem()
}

func (i AccountLdapArray) ToAccountLdapArrayOutput() AccountLdapArrayOutput {
	return i.ToAccountLdapArrayOutputWithContext(context.Background())
}

func (i AccountLdapArray) ToAccountLdapArrayOutputWithContext(ctx context.Context) AccountLdapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountLdapArrayOutput)
}

// AccountLdapMapInput is an input type that accepts AccountLdapMap and AccountLdapMapOutput values.
// You can construct a concrete instance of `AccountLdapMapInput` via:
//
//	AccountLdapMap{ "key": AccountLdapArgs{...} }
type AccountLdapMapInput interface {
	pulumi.Input

	ToAccountLdapMapOutput() AccountLdapMapOutput
	ToAccountLdapMapOutputWithContext(context.Context) AccountLdapMapOutput
}

type AccountLdapMap map[string]AccountLdapInput

func (AccountLdapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountLdap)(nil)).Elem()
}

func (i AccountLdapMap) ToAccountLdapMapOutput() AccountLdapMapOutput {
	return i.ToAccountLdapMapOutputWithContext(context.Background())
}

func (i AccountLdapMap) ToAccountLdapMapOutputWithContext(ctx context.Context) AccountLdapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountLdapMapOutput)
}

type AccountLdapOutput struct{ *pulumi.OutputState }

func (AccountLdapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountLdap)(nil)).Elem()
}

func (o AccountLdapOutput) ToAccountLdapOutput() AccountLdapOutput {
	return o
}

func (o AccountLdapOutput) ToAccountLdapOutputWithContext(ctx context.Context) AccountLdapOutput {
	return o
}

// The resource ID for the auth method.
func (o AccountLdapOutput) AuthMethodId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountLdap) pulumi.StringOutput { return v.AuthMethodId }).(pulumi.StringOutput)
}

// The account description.
func (o AccountLdapOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountLdap) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The login name for this account.
func (o AccountLdapOutput) LoginName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountLdap) pulumi.StringPtrOutput { return v.LoginName }).(pulumi.StringPtrOutput)
}

// The account name. Defaults to the resource name.
func (o AccountLdapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountLdap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource type.
//
// Deprecated: The value for this field will be infered since 'ldap' is the only possible value.
func (o AccountLdapOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccountLdap) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type AccountLdapArrayOutput struct{ *pulumi.OutputState }

func (AccountLdapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountLdap)(nil)).Elem()
}

func (o AccountLdapArrayOutput) ToAccountLdapArrayOutput() AccountLdapArrayOutput {
	return o
}

func (o AccountLdapArrayOutput) ToAccountLdapArrayOutputWithContext(ctx context.Context) AccountLdapArrayOutput {
	return o
}

func (o AccountLdapArrayOutput) Index(i pulumi.IntInput) AccountLdapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountLdap {
		return vs[0].([]*AccountLdap)[vs[1].(int)]
	}).(AccountLdapOutput)
}

type AccountLdapMapOutput struct{ *pulumi.OutputState }

func (AccountLdapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountLdap)(nil)).Elem()
}

func (o AccountLdapMapOutput) ToAccountLdapMapOutput() AccountLdapMapOutput {
	return o
}

func (o AccountLdapMapOutput) ToAccountLdapMapOutputWithContext(ctx context.Context) AccountLdapMapOutput {
	return o
}

func (o AccountLdapMapOutput) MapIndex(k pulumi.StringInput) AccountLdapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountLdap {
		return vs[0].(map[string]*AccountLdap)[vs[1].(string)]
	}).(AccountLdapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountLdapInput)(nil)).Elem(), &AccountLdap{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountLdapArrayInput)(nil)).Elem(), AccountLdapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountLdapMapInput)(nil)).Elem(), AccountLdapMap{})
	pulumi.RegisterOutputType(AccountLdapOutput{})
	pulumi.RegisterOutputType(AccountLdapArrayOutput{})
	pulumi.RegisterOutputType(AccountLdapMapOutput{})
}
