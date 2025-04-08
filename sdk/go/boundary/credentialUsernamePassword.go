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

type CredentialUsernamePassword struct {
	pulumi.CustomResourceState

	// The credential store in which to save this username/password credential.
	CredentialStoreId pulumi.StringOutput `pulumi:"credentialStoreId"`
	// The description of this username/password credential.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of this username/password credential. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password of this username/password credential.
	Password pulumi.StringOutput `pulumi:"password"`
	// The password hmac.
	PasswordHmac pulumi.StringOutput `pulumi:"passwordHmac"`
	// The username of this username/password credential.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewCredentialUsernamePassword registers a new resource with the given unique name, arguments, and options.
func NewCredentialUsernamePassword(ctx *pulumi.Context,
	name string, args *CredentialUsernamePasswordArgs, opts ...pulumi.ResourceOption) (*CredentialUsernamePassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CredentialStoreId == nil {
		return nil, errors.New("invalid value for required argument 'CredentialStoreId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CredentialUsernamePassword
	err := ctx.RegisterResource("boundary:index/credentialUsernamePassword:CredentialUsernamePassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCredentialUsernamePassword gets an existing CredentialUsernamePassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCredentialUsernamePassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CredentialUsernamePasswordState, opts ...pulumi.ResourceOption) (*CredentialUsernamePassword, error) {
	var resource CredentialUsernamePassword
	err := ctx.ReadResource("boundary:index/credentialUsernamePassword:CredentialUsernamePassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CredentialUsernamePassword resources.
type credentialUsernamePasswordState struct {
	// The credential store in which to save this username/password credential.
	CredentialStoreId *string `pulumi:"credentialStoreId"`
	// The description of this username/password credential.
	Description *string `pulumi:"description"`
	// The name of this username/password credential. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The password of this username/password credential.
	Password *string `pulumi:"password"`
	// The password hmac.
	PasswordHmac *string `pulumi:"passwordHmac"`
	// The username of this username/password credential.
	Username *string `pulumi:"username"`
}

type CredentialUsernamePasswordState struct {
	// The credential store in which to save this username/password credential.
	CredentialStoreId pulumi.StringPtrInput
	// The description of this username/password credential.
	Description pulumi.StringPtrInput
	// The name of this username/password credential. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The password of this username/password credential.
	Password pulumi.StringPtrInput
	// The password hmac.
	PasswordHmac pulumi.StringPtrInput
	// The username of this username/password credential.
	Username pulumi.StringPtrInput
}

func (CredentialUsernamePasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialUsernamePasswordState)(nil)).Elem()
}

type credentialUsernamePasswordArgs struct {
	// The credential store in which to save this username/password credential.
	CredentialStoreId string `pulumi:"credentialStoreId"`
	// The description of this username/password credential.
	Description *string `pulumi:"description"`
	// The name of this username/password credential. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The password of this username/password credential.
	Password string `pulumi:"password"`
	// The username of this username/password credential.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a CredentialUsernamePassword resource.
type CredentialUsernamePasswordArgs struct {
	// The credential store in which to save this username/password credential.
	CredentialStoreId pulumi.StringInput
	// The description of this username/password credential.
	Description pulumi.StringPtrInput
	// The name of this username/password credential. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The password of this username/password credential.
	Password pulumi.StringInput
	// The username of this username/password credential.
	Username pulumi.StringInput
}

func (CredentialUsernamePasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialUsernamePasswordArgs)(nil)).Elem()
}

type CredentialUsernamePasswordInput interface {
	pulumi.Input

	ToCredentialUsernamePasswordOutput() CredentialUsernamePasswordOutput
	ToCredentialUsernamePasswordOutputWithContext(ctx context.Context) CredentialUsernamePasswordOutput
}

func (*CredentialUsernamePassword) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialUsernamePassword)(nil)).Elem()
}

func (i *CredentialUsernamePassword) ToCredentialUsernamePasswordOutput() CredentialUsernamePasswordOutput {
	return i.ToCredentialUsernamePasswordOutputWithContext(context.Background())
}

func (i *CredentialUsernamePassword) ToCredentialUsernamePasswordOutputWithContext(ctx context.Context) CredentialUsernamePasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialUsernamePasswordOutput)
}

// CredentialUsernamePasswordArrayInput is an input type that accepts CredentialUsernamePasswordArray and CredentialUsernamePasswordArrayOutput values.
// You can construct a concrete instance of `CredentialUsernamePasswordArrayInput` via:
//
//	CredentialUsernamePasswordArray{ CredentialUsernamePasswordArgs{...} }
type CredentialUsernamePasswordArrayInput interface {
	pulumi.Input

	ToCredentialUsernamePasswordArrayOutput() CredentialUsernamePasswordArrayOutput
	ToCredentialUsernamePasswordArrayOutputWithContext(context.Context) CredentialUsernamePasswordArrayOutput
}

type CredentialUsernamePasswordArray []CredentialUsernamePasswordInput

func (CredentialUsernamePasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CredentialUsernamePassword)(nil)).Elem()
}

func (i CredentialUsernamePasswordArray) ToCredentialUsernamePasswordArrayOutput() CredentialUsernamePasswordArrayOutput {
	return i.ToCredentialUsernamePasswordArrayOutputWithContext(context.Background())
}

func (i CredentialUsernamePasswordArray) ToCredentialUsernamePasswordArrayOutputWithContext(ctx context.Context) CredentialUsernamePasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialUsernamePasswordArrayOutput)
}

// CredentialUsernamePasswordMapInput is an input type that accepts CredentialUsernamePasswordMap and CredentialUsernamePasswordMapOutput values.
// You can construct a concrete instance of `CredentialUsernamePasswordMapInput` via:
//
//	CredentialUsernamePasswordMap{ "key": CredentialUsernamePasswordArgs{...} }
type CredentialUsernamePasswordMapInput interface {
	pulumi.Input

	ToCredentialUsernamePasswordMapOutput() CredentialUsernamePasswordMapOutput
	ToCredentialUsernamePasswordMapOutputWithContext(context.Context) CredentialUsernamePasswordMapOutput
}

type CredentialUsernamePasswordMap map[string]CredentialUsernamePasswordInput

func (CredentialUsernamePasswordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CredentialUsernamePassword)(nil)).Elem()
}

func (i CredentialUsernamePasswordMap) ToCredentialUsernamePasswordMapOutput() CredentialUsernamePasswordMapOutput {
	return i.ToCredentialUsernamePasswordMapOutputWithContext(context.Background())
}

func (i CredentialUsernamePasswordMap) ToCredentialUsernamePasswordMapOutputWithContext(ctx context.Context) CredentialUsernamePasswordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialUsernamePasswordMapOutput)
}

type CredentialUsernamePasswordOutput struct{ *pulumi.OutputState }

func (CredentialUsernamePasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialUsernamePassword)(nil)).Elem()
}

func (o CredentialUsernamePasswordOutput) ToCredentialUsernamePasswordOutput() CredentialUsernamePasswordOutput {
	return o
}

func (o CredentialUsernamePasswordOutput) ToCredentialUsernamePasswordOutputWithContext(ctx context.Context) CredentialUsernamePasswordOutput {
	return o
}

// The credential store in which to save this username/password credential.
func (o CredentialUsernamePasswordOutput) CredentialStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialUsernamePassword) pulumi.StringOutput { return v.CredentialStoreId }).(pulumi.StringOutput)
}

// The description of this username/password credential.
func (o CredentialUsernamePasswordOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CredentialUsernamePassword) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of this username/password credential. Defaults to the resource name.
func (o CredentialUsernamePasswordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialUsernamePassword) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The password of this username/password credential.
func (o CredentialUsernamePasswordOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialUsernamePassword) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The password hmac.
func (o CredentialUsernamePasswordOutput) PasswordHmac() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialUsernamePassword) pulumi.StringOutput { return v.PasswordHmac }).(pulumi.StringOutput)
}

// The username of this username/password credential.
func (o CredentialUsernamePasswordOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialUsernamePassword) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type CredentialUsernamePasswordArrayOutput struct{ *pulumi.OutputState }

func (CredentialUsernamePasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CredentialUsernamePassword)(nil)).Elem()
}

func (o CredentialUsernamePasswordArrayOutput) ToCredentialUsernamePasswordArrayOutput() CredentialUsernamePasswordArrayOutput {
	return o
}

func (o CredentialUsernamePasswordArrayOutput) ToCredentialUsernamePasswordArrayOutputWithContext(ctx context.Context) CredentialUsernamePasswordArrayOutput {
	return o
}

func (o CredentialUsernamePasswordArrayOutput) Index(i pulumi.IntInput) CredentialUsernamePasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CredentialUsernamePassword {
		return vs[0].([]*CredentialUsernamePassword)[vs[1].(int)]
	}).(CredentialUsernamePasswordOutput)
}

type CredentialUsernamePasswordMapOutput struct{ *pulumi.OutputState }

func (CredentialUsernamePasswordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CredentialUsernamePassword)(nil)).Elem()
}

func (o CredentialUsernamePasswordMapOutput) ToCredentialUsernamePasswordMapOutput() CredentialUsernamePasswordMapOutput {
	return o
}

func (o CredentialUsernamePasswordMapOutput) ToCredentialUsernamePasswordMapOutputWithContext(ctx context.Context) CredentialUsernamePasswordMapOutput {
	return o
}

func (o CredentialUsernamePasswordMapOutput) MapIndex(k pulumi.StringInput) CredentialUsernamePasswordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CredentialUsernamePassword {
		return vs[0].(map[string]*CredentialUsernamePassword)[vs[1].(string)]
	}).(CredentialUsernamePasswordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialUsernamePasswordInput)(nil)).Elem(), &CredentialUsernamePassword{})
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialUsernamePasswordArrayInput)(nil)).Elem(), CredentialUsernamePasswordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialUsernamePasswordMapInput)(nil)).Elem(), CredentialUsernamePasswordMap{})
	pulumi.RegisterOutputType(CredentialUsernamePasswordOutput{})
	pulumi.RegisterOutputType(CredentialUsernamePasswordArrayOutput{})
	pulumi.RegisterOutputType(CredentialUsernamePasswordMapOutput{})
}
