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

// The credential library for Vault resource allows you to configure a Boundary credential library for Vault.
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
//			project, err := boundary.NewScope(ctx, "project", &boundary.ScopeArgs{
//				Name:                pulumi.String("project_one"),
//				Description:         pulumi.String("My first scope!"),
//				ScopeId:             org.ID(),
//				AutoCreateAdminRole: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := boundary.NewCredentialStoreVault(ctx, "foo", &boundary.CredentialStoreVaultArgs{
//				Name:        pulumi.String("foo"),
//				Description: pulumi.String("My first Vault credential store!"),
//				Address:     pulumi.String("http://127.0.0.1:8200"),
//				Token:       pulumi.String("s.0ufRo6XEGU2jOqnIr7OlFYP5"),
//				ScopeId:     project.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewCredentialLibraryVaultSshCertificate(ctx, "foo", &boundary.CredentialLibraryVaultSshCertificateArgs{
//				Name:              pulumi.String("foo"),
//				Description:       pulumi.String("My first Vault SSH certificate credential library!"),
//				CredentialStoreId: foo.ID(),
//				Path:              pulumi.String("ssh/sign/foo"),
//				Username:          pulumi.String("foo"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewCredentialLibraryVaultSshCertificate(ctx, "bar", &boundary.CredentialLibraryVaultSshCertificateArgs{
//				Name:              pulumi.String("bar"),
//				Description:       pulumi.String("My second Vault SSH certificate credential library!"),
//				CredentialStoreId: foo.ID(),
//				Path:              pulumi.String("ssh/sign/foo"),
//				Username:          pulumi.String("foo"),
//				KeyType:           pulumi.String("ecdsa"),
//				KeyBits:           pulumi.Int(384),
//				Extensions: pulumi.StringMap{
//					"permit-pty": pulumi.String(""),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = boundary.NewCredentialLibraryVaultSshCertificate(ctx, "baz", &boundary.CredentialLibraryVaultSshCertificateArgs{
//				Name:              pulumi.String("baz"),
//				Description:       pulumi.String("vault "),
//				CredentialStoreId: foo.ID(),
//				Path:              pulumi.String("ssh/issue/foo"),
//				Username:          pulumi.String("foo"),
//				KeyType:           pulumi.String("rsa"),
//				KeyBits:           pulumi.Int(4096),
//				Extensions: pulumi.StringMap{
//					"permit-pty":            pulumi.String(""),
//					"permit-X11-forwarding": pulumi.String(""),
//				},
//				CriticalOptions: pulumi.StringMap{
//					"force-command": pulumi.String("/bin/some_script"),
//				},
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
// $ pulumi import boundary:index/credentialLibraryVaultSshCertificate:CredentialLibraryVaultSshCertificate foo <my-id>
// ```
type CredentialLibraryVaultSshCertificate struct {
	pulumi.CustomResourceState

	// Principals to be signed as "validPrinciples" in addition to username.
	AdditionalValidPrincipals pulumi.StringArrayOutput `pulumi:"additionalValidPrincipals"`
	// The ID of the credential store that this library belongs to.
	CredentialStoreId pulumi.StringOutput `pulumi:"credentialStoreId"`
	// Specifies a map of the critical options that the certificate should be signed for.
	CriticalOptions pulumi.StringMapOutput `pulumi:"criticalOptions"`
	// The Vault credential library description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies a map of the extensions that the certificate should be signed for.
	Extensions pulumi.StringMapOutput `pulumi:"extensions"`
	// Specifies the number of bits to use for the generated keys.
	KeyBits pulumi.IntPtrOutput `pulumi:"keyBits"`
	// Specifies the key id a certificate should have.
	KeyId pulumi.StringPtrOutput `pulumi:"keyId"`
	// Specifies the desired key type; must be ed25519, ecdsa, or rsa.
	KeyType pulumi.StringPtrOutput `pulumi:"keyType"`
	// The Vault credential library name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The path in Vault to request credentials from.
	Path pulumi.StringOutput `pulumi:"path"`
	// Specifies the requested time to live for a certificate returned from the library.
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// The username to use with the certificate returned by the library.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewCredentialLibraryVaultSshCertificate registers a new resource with the given unique name, arguments, and options.
func NewCredentialLibraryVaultSshCertificate(ctx *pulumi.Context,
	name string, args *CredentialLibraryVaultSshCertificateArgs, opts ...pulumi.ResourceOption) (*CredentialLibraryVaultSshCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CredentialStoreId == nil {
		return nil, errors.New("invalid value for required argument 'CredentialStoreId'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CredentialLibraryVaultSshCertificate
	err := ctx.RegisterResource("boundary:index/credentialLibraryVaultSshCertificate:CredentialLibraryVaultSshCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCredentialLibraryVaultSshCertificate gets an existing CredentialLibraryVaultSshCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCredentialLibraryVaultSshCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CredentialLibraryVaultSshCertificateState, opts ...pulumi.ResourceOption) (*CredentialLibraryVaultSshCertificate, error) {
	var resource CredentialLibraryVaultSshCertificate
	err := ctx.ReadResource("boundary:index/credentialLibraryVaultSshCertificate:CredentialLibraryVaultSshCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CredentialLibraryVaultSshCertificate resources.
type credentialLibraryVaultSshCertificateState struct {
	// Principals to be signed as "validPrinciples" in addition to username.
	AdditionalValidPrincipals []string `pulumi:"additionalValidPrincipals"`
	// The ID of the credential store that this library belongs to.
	CredentialStoreId *string `pulumi:"credentialStoreId"`
	// Specifies a map of the critical options that the certificate should be signed for.
	CriticalOptions map[string]string `pulumi:"criticalOptions"`
	// The Vault credential library description.
	Description *string `pulumi:"description"`
	// Specifies a map of the extensions that the certificate should be signed for.
	Extensions map[string]string `pulumi:"extensions"`
	// Specifies the number of bits to use for the generated keys.
	KeyBits *int `pulumi:"keyBits"`
	// Specifies the key id a certificate should have.
	KeyId *string `pulumi:"keyId"`
	// Specifies the desired key type; must be ed25519, ecdsa, or rsa.
	KeyType *string `pulumi:"keyType"`
	// The Vault credential library name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The path in Vault to request credentials from.
	Path *string `pulumi:"path"`
	// Specifies the requested time to live for a certificate returned from the library.
	Ttl *string `pulumi:"ttl"`
	// The username to use with the certificate returned by the library.
	Username *string `pulumi:"username"`
}

type CredentialLibraryVaultSshCertificateState struct {
	// Principals to be signed as "validPrinciples" in addition to username.
	AdditionalValidPrincipals pulumi.StringArrayInput
	// The ID of the credential store that this library belongs to.
	CredentialStoreId pulumi.StringPtrInput
	// Specifies a map of the critical options that the certificate should be signed for.
	CriticalOptions pulumi.StringMapInput
	// The Vault credential library description.
	Description pulumi.StringPtrInput
	// Specifies a map of the extensions that the certificate should be signed for.
	Extensions pulumi.StringMapInput
	// Specifies the number of bits to use for the generated keys.
	KeyBits pulumi.IntPtrInput
	// Specifies the key id a certificate should have.
	KeyId pulumi.StringPtrInput
	// Specifies the desired key type; must be ed25519, ecdsa, or rsa.
	KeyType pulumi.StringPtrInput
	// The Vault credential library name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The path in Vault to request credentials from.
	Path pulumi.StringPtrInput
	// Specifies the requested time to live for a certificate returned from the library.
	Ttl pulumi.StringPtrInput
	// The username to use with the certificate returned by the library.
	Username pulumi.StringPtrInput
}

func (CredentialLibraryVaultSshCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialLibraryVaultSshCertificateState)(nil)).Elem()
}

type credentialLibraryVaultSshCertificateArgs struct {
	// Principals to be signed as "validPrinciples" in addition to username.
	AdditionalValidPrincipals []string `pulumi:"additionalValidPrincipals"`
	// The ID of the credential store that this library belongs to.
	CredentialStoreId string `pulumi:"credentialStoreId"`
	// Specifies a map of the critical options that the certificate should be signed for.
	CriticalOptions map[string]string `pulumi:"criticalOptions"`
	// The Vault credential library description.
	Description *string `pulumi:"description"`
	// Specifies a map of the extensions that the certificate should be signed for.
	Extensions map[string]string `pulumi:"extensions"`
	// Specifies the number of bits to use for the generated keys.
	KeyBits *int `pulumi:"keyBits"`
	// Specifies the key id a certificate should have.
	KeyId *string `pulumi:"keyId"`
	// Specifies the desired key type; must be ed25519, ecdsa, or rsa.
	KeyType *string `pulumi:"keyType"`
	// The Vault credential library name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The path in Vault to request credentials from.
	Path string `pulumi:"path"`
	// Specifies the requested time to live for a certificate returned from the library.
	Ttl *string `pulumi:"ttl"`
	// The username to use with the certificate returned by the library.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a CredentialLibraryVaultSshCertificate resource.
type CredentialLibraryVaultSshCertificateArgs struct {
	// Principals to be signed as "validPrinciples" in addition to username.
	AdditionalValidPrincipals pulumi.StringArrayInput
	// The ID of the credential store that this library belongs to.
	CredentialStoreId pulumi.StringInput
	// Specifies a map of the critical options that the certificate should be signed for.
	CriticalOptions pulumi.StringMapInput
	// The Vault credential library description.
	Description pulumi.StringPtrInput
	// Specifies a map of the extensions that the certificate should be signed for.
	Extensions pulumi.StringMapInput
	// Specifies the number of bits to use for the generated keys.
	KeyBits pulumi.IntPtrInput
	// Specifies the key id a certificate should have.
	KeyId pulumi.StringPtrInput
	// Specifies the desired key type; must be ed25519, ecdsa, or rsa.
	KeyType pulumi.StringPtrInput
	// The Vault credential library name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The path in Vault to request credentials from.
	Path pulumi.StringInput
	// Specifies the requested time to live for a certificate returned from the library.
	Ttl pulumi.StringPtrInput
	// The username to use with the certificate returned by the library.
	Username pulumi.StringInput
}

func (CredentialLibraryVaultSshCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialLibraryVaultSshCertificateArgs)(nil)).Elem()
}

type CredentialLibraryVaultSshCertificateInput interface {
	pulumi.Input

	ToCredentialLibraryVaultSshCertificateOutput() CredentialLibraryVaultSshCertificateOutput
	ToCredentialLibraryVaultSshCertificateOutputWithContext(ctx context.Context) CredentialLibraryVaultSshCertificateOutput
}

func (*CredentialLibraryVaultSshCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialLibraryVaultSshCertificate)(nil)).Elem()
}

func (i *CredentialLibraryVaultSshCertificate) ToCredentialLibraryVaultSshCertificateOutput() CredentialLibraryVaultSshCertificateOutput {
	return i.ToCredentialLibraryVaultSshCertificateOutputWithContext(context.Background())
}

func (i *CredentialLibraryVaultSshCertificate) ToCredentialLibraryVaultSshCertificateOutputWithContext(ctx context.Context) CredentialLibraryVaultSshCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialLibraryVaultSshCertificateOutput)
}

// CredentialLibraryVaultSshCertificateArrayInput is an input type that accepts CredentialLibraryVaultSshCertificateArray and CredentialLibraryVaultSshCertificateArrayOutput values.
// You can construct a concrete instance of `CredentialLibraryVaultSshCertificateArrayInput` via:
//
//	CredentialLibraryVaultSshCertificateArray{ CredentialLibraryVaultSshCertificateArgs{...} }
type CredentialLibraryVaultSshCertificateArrayInput interface {
	pulumi.Input

	ToCredentialLibraryVaultSshCertificateArrayOutput() CredentialLibraryVaultSshCertificateArrayOutput
	ToCredentialLibraryVaultSshCertificateArrayOutputWithContext(context.Context) CredentialLibraryVaultSshCertificateArrayOutput
}

type CredentialLibraryVaultSshCertificateArray []CredentialLibraryVaultSshCertificateInput

func (CredentialLibraryVaultSshCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CredentialLibraryVaultSshCertificate)(nil)).Elem()
}

func (i CredentialLibraryVaultSshCertificateArray) ToCredentialLibraryVaultSshCertificateArrayOutput() CredentialLibraryVaultSshCertificateArrayOutput {
	return i.ToCredentialLibraryVaultSshCertificateArrayOutputWithContext(context.Background())
}

func (i CredentialLibraryVaultSshCertificateArray) ToCredentialLibraryVaultSshCertificateArrayOutputWithContext(ctx context.Context) CredentialLibraryVaultSshCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialLibraryVaultSshCertificateArrayOutput)
}

// CredentialLibraryVaultSshCertificateMapInput is an input type that accepts CredentialLibraryVaultSshCertificateMap and CredentialLibraryVaultSshCertificateMapOutput values.
// You can construct a concrete instance of `CredentialLibraryVaultSshCertificateMapInput` via:
//
//	CredentialLibraryVaultSshCertificateMap{ "key": CredentialLibraryVaultSshCertificateArgs{...} }
type CredentialLibraryVaultSshCertificateMapInput interface {
	pulumi.Input

	ToCredentialLibraryVaultSshCertificateMapOutput() CredentialLibraryVaultSshCertificateMapOutput
	ToCredentialLibraryVaultSshCertificateMapOutputWithContext(context.Context) CredentialLibraryVaultSshCertificateMapOutput
}

type CredentialLibraryVaultSshCertificateMap map[string]CredentialLibraryVaultSshCertificateInput

func (CredentialLibraryVaultSshCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CredentialLibraryVaultSshCertificate)(nil)).Elem()
}

func (i CredentialLibraryVaultSshCertificateMap) ToCredentialLibraryVaultSshCertificateMapOutput() CredentialLibraryVaultSshCertificateMapOutput {
	return i.ToCredentialLibraryVaultSshCertificateMapOutputWithContext(context.Background())
}

func (i CredentialLibraryVaultSshCertificateMap) ToCredentialLibraryVaultSshCertificateMapOutputWithContext(ctx context.Context) CredentialLibraryVaultSshCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialLibraryVaultSshCertificateMapOutput)
}

type CredentialLibraryVaultSshCertificateOutput struct{ *pulumi.OutputState }

func (CredentialLibraryVaultSshCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CredentialLibraryVaultSshCertificate)(nil)).Elem()
}

func (o CredentialLibraryVaultSshCertificateOutput) ToCredentialLibraryVaultSshCertificateOutput() CredentialLibraryVaultSshCertificateOutput {
	return o
}

func (o CredentialLibraryVaultSshCertificateOutput) ToCredentialLibraryVaultSshCertificateOutputWithContext(ctx context.Context) CredentialLibraryVaultSshCertificateOutput {
	return o
}

// Principals to be signed as "validPrinciples" in addition to username.
func (o CredentialLibraryVaultSshCertificateOutput) AdditionalValidPrincipals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CredentialLibraryVaultSshCertificate) pulumi.StringArrayOutput {
		return v.AdditionalValidPrincipals
	}).(pulumi.StringArrayOutput)
}

// The ID of the credential store that this library belongs to.
func (o CredentialLibraryVaultSshCertificateOutput) CredentialStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialLibraryVaultSshCertificate) pulumi.StringOutput { return v.CredentialStoreId }).(pulumi.StringOutput)
}

// Specifies a map of the critical options that the certificate should be signed for.
func (o CredentialLibraryVaultSshCertificateOutput) CriticalOptions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CredentialLibraryVaultSshCertificate) pulumi.StringMapOutput { return v.CriticalOptions }).(pulumi.StringMapOutput)
}

// The Vault credential library description.
func (o CredentialLibraryVaultSshCertificateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CredentialLibraryVaultSshCertificate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies a map of the extensions that the certificate should be signed for.
func (o CredentialLibraryVaultSshCertificateOutput) Extensions() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CredentialLibraryVaultSshCertificate) pulumi.StringMapOutput { return v.Extensions }).(pulumi.StringMapOutput)
}

// Specifies the number of bits to use for the generated keys.
func (o CredentialLibraryVaultSshCertificateOutput) KeyBits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CredentialLibraryVaultSshCertificate) pulumi.IntPtrOutput { return v.KeyBits }).(pulumi.IntPtrOutput)
}

// Specifies the key id a certificate should have.
func (o CredentialLibraryVaultSshCertificateOutput) KeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CredentialLibraryVaultSshCertificate) pulumi.StringPtrOutput { return v.KeyId }).(pulumi.StringPtrOutput)
}

// Specifies the desired key type; must be ed25519, ecdsa, or rsa.
func (o CredentialLibraryVaultSshCertificateOutput) KeyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CredentialLibraryVaultSshCertificate) pulumi.StringPtrOutput { return v.KeyType }).(pulumi.StringPtrOutput)
}

// The Vault credential library name. Defaults to the resource name.
func (o CredentialLibraryVaultSshCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialLibraryVaultSshCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The path in Vault to request credentials from.
func (o CredentialLibraryVaultSshCertificateOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialLibraryVaultSshCertificate) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Specifies the requested time to live for a certificate returned from the library.
func (o CredentialLibraryVaultSshCertificateOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CredentialLibraryVaultSshCertificate) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

// The username to use with the certificate returned by the library.
func (o CredentialLibraryVaultSshCertificateOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *CredentialLibraryVaultSshCertificate) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type CredentialLibraryVaultSshCertificateArrayOutput struct{ *pulumi.OutputState }

func (CredentialLibraryVaultSshCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CredentialLibraryVaultSshCertificate)(nil)).Elem()
}

func (o CredentialLibraryVaultSshCertificateArrayOutput) ToCredentialLibraryVaultSshCertificateArrayOutput() CredentialLibraryVaultSshCertificateArrayOutput {
	return o
}

func (o CredentialLibraryVaultSshCertificateArrayOutput) ToCredentialLibraryVaultSshCertificateArrayOutputWithContext(ctx context.Context) CredentialLibraryVaultSshCertificateArrayOutput {
	return o
}

func (o CredentialLibraryVaultSshCertificateArrayOutput) Index(i pulumi.IntInput) CredentialLibraryVaultSshCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CredentialLibraryVaultSshCertificate {
		return vs[0].([]*CredentialLibraryVaultSshCertificate)[vs[1].(int)]
	}).(CredentialLibraryVaultSshCertificateOutput)
}

type CredentialLibraryVaultSshCertificateMapOutput struct{ *pulumi.OutputState }

func (CredentialLibraryVaultSshCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CredentialLibraryVaultSshCertificate)(nil)).Elem()
}

func (o CredentialLibraryVaultSshCertificateMapOutput) ToCredentialLibraryVaultSshCertificateMapOutput() CredentialLibraryVaultSshCertificateMapOutput {
	return o
}

func (o CredentialLibraryVaultSshCertificateMapOutput) ToCredentialLibraryVaultSshCertificateMapOutputWithContext(ctx context.Context) CredentialLibraryVaultSshCertificateMapOutput {
	return o
}

func (o CredentialLibraryVaultSshCertificateMapOutput) MapIndex(k pulumi.StringInput) CredentialLibraryVaultSshCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CredentialLibraryVaultSshCertificate {
		return vs[0].(map[string]*CredentialLibraryVaultSshCertificate)[vs[1].(string)]
	}).(CredentialLibraryVaultSshCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialLibraryVaultSshCertificateInput)(nil)).Elem(), &CredentialLibraryVaultSshCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialLibraryVaultSshCertificateArrayInput)(nil)).Elem(), CredentialLibraryVaultSshCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialLibraryVaultSshCertificateMapInput)(nil)).Elem(), CredentialLibraryVaultSshCertificateMap{})
	pulumi.RegisterOutputType(CredentialLibraryVaultSshCertificateOutput{})
	pulumi.RegisterOutputType(CredentialLibraryVaultSshCertificateArrayOutput{})
	pulumi.RegisterOutputType(CredentialLibraryVaultSshCertificateMapOutput{})
}
