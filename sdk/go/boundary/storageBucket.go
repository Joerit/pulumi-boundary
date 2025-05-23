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

// The storage bucket resource allows you to configure a Boundary storage bucket. A storage bucket can only belong to the Global scope or an Org scope. At this time, the only supported storage for storage buckets is AWS S3. This feature requires Boundary Enterprise or Boundary HCP.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
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
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"region": "us-east-1",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"access_key_id":     "aws_access_key_id_value",
//				"secret_access_key": "aws_secret_access_key_value",
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			_, err = boundary.NewStorageBucket(ctx, "aws_static_credentials_example", &boundary.StorageBucketArgs{
//				Name:           pulumi.String("My aws storage bucket with static credentials"),
//				Description:    pulumi.String("My first storage bucket!"),
//				ScopeId:        org.ID(),
//				PluginName:     pulumi.String("aws"),
//				BucketName:     pulumi.String("mybucket"),
//				AttributesJson: pulumi.String(json0),
//				SecretsJson:    pulumi.String(json1),
//				WorkerFilter:   pulumi.String("\"pki\" in \"/tags/type\""),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON2, err := json.Marshal(map[string]interface{}{
//				"region":                      "us-east-1",
//				"role_arn":                    "arn:aws:iam::123456789012:role/S3Access",
//				"disable_credential_rotation": "true",
//			})
//			if err != nil {
//				return err
//			}
//			json2 := string(tmpJSON2)
//			_, err = boundary.NewStorageBucket(ctx, "aws_dynamic_credentials_example", &boundary.StorageBucketArgs{
//				Name:           pulumi.String("My aws storage bucket with dynamic credentials"),
//				Description:    pulumi.String("My first storage bucket!"),
//				ScopeId:        org.ID(),
//				PluginName:     pulumi.String("aws"),
//				BucketName:     pulumi.String("mybucket"),
//				AttributesJson: pulumi.String(json2),
//				WorkerFilter:   pulumi.String("\"pki\" in \"/tags/type\""),
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
// $ pulumi import boundary:index/storageBucket:StorageBucket foo <my-id>
// ```
type StorageBucket struct {
	pulumi.CustomResourceState

	// The attributes for the storage bucket. The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
	AttributesJson pulumi.StringPtrOutput `pulumi:"attributesJson"`
	// The name of the bucket within the external object store service.
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// The prefix used to organize the data held within the external object store.
	BucketPrefix pulumi.StringPtrOutput `pulumi:"bucketPrefix"`
	// The storage bucket description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Internal only. Used to force update so that we can always check the value of secrets.
	InternalForceUpdate pulumi.StringOutput `pulumi:"internalForceUpdate"`
	// Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
	InternalHmacUsedForSecretsConfigHmac pulumi.StringOutput `pulumi:"internalHmacUsedForSecretsConfigHmac"`
	// Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
	InternalSecretsConfigHmac pulumi.StringOutput `pulumi:"internalSecretsConfigHmac"`
	// The storage bucket name. Defaults to the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the plugin that should back the resource. This or pluginName must be defined.
	PluginId pulumi.StringOutput `pulumi:"pluginId"`
	// The name of the plugin that should back the resource. This or pluginId must be defined.
	PluginName pulumi.StringPtrOutput `pulumi:"pluginName"`
	// The scope for this storage bucket.
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
	// The HMAC'd secrets value returned from the server.
	SecretsHmac pulumi.StringOutput `pulumi:"secretsHmac"`
	// The secrets for the storage bucket. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributesJson", removing this block will NOT clear secrets from the storage bucket; this allows injecting secrets for one call, then removing them for storage.
	SecretsJson pulumi.StringPtrOutput `pulumi:"secretsJson"`
	// Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
	WorkerFilter pulumi.StringOutput `pulumi:"workerFilter"`
}

// NewStorageBucket registers a new resource with the given unique name, arguments, and options.
func NewStorageBucket(ctx *pulumi.Context,
	name string, args *StorageBucketArgs, opts ...pulumi.ResourceOption) (*StorageBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	if args.WorkerFilter == nil {
		return nil, errors.New("invalid value for required argument 'WorkerFilter'")
	}
	if args.SecretsJson != nil {
		args.SecretsJson = pulumi.ToSecret(args.SecretsJson).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretsJson",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource StorageBucket
	err := ctx.RegisterResource("boundary:index/storageBucket:StorageBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageBucket gets an existing StorageBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageBucketState, opts ...pulumi.ResourceOption) (*StorageBucket, error) {
	var resource StorageBucket
	err := ctx.ReadResource("boundary:index/storageBucket:StorageBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageBucket resources.
type storageBucketState struct {
	// The attributes for the storage bucket. The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
	AttributesJson *string `pulumi:"attributesJson"`
	// The name of the bucket within the external object store service.
	BucketName *string `pulumi:"bucketName"`
	// The prefix used to organize the data held within the external object store.
	BucketPrefix *string `pulumi:"bucketPrefix"`
	// The storage bucket description.
	Description *string `pulumi:"description"`
	// Internal only. Used to force update so that we can always check the value of secrets.
	InternalForceUpdate *string `pulumi:"internalForceUpdate"`
	// Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
	InternalHmacUsedForSecretsConfigHmac *string `pulumi:"internalHmacUsedForSecretsConfigHmac"`
	// Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
	InternalSecretsConfigHmac *string `pulumi:"internalSecretsConfigHmac"`
	// The storage bucket name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The ID of the plugin that should back the resource. This or pluginName must be defined.
	PluginId *string `pulumi:"pluginId"`
	// The name of the plugin that should back the resource. This or pluginId must be defined.
	PluginName *string `pulumi:"pluginName"`
	// The scope for this storage bucket.
	ScopeId *string `pulumi:"scopeId"`
	// The HMAC'd secrets value returned from the server.
	SecretsHmac *string `pulumi:"secretsHmac"`
	// The secrets for the storage bucket. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributesJson", removing this block will NOT clear secrets from the storage bucket; this allows injecting secrets for one call, then removing them for storage.
	SecretsJson *string `pulumi:"secretsJson"`
	// Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
	WorkerFilter *string `pulumi:"workerFilter"`
}

type StorageBucketState struct {
	// The attributes for the storage bucket. The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
	AttributesJson pulumi.StringPtrInput
	// The name of the bucket within the external object store service.
	BucketName pulumi.StringPtrInput
	// The prefix used to organize the data held within the external object store.
	BucketPrefix pulumi.StringPtrInput
	// The storage bucket description.
	Description pulumi.StringPtrInput
	// Internal only. Used to force update so that we can always check the value of secrets.
	InternalForceUpdate pulumi.StringPtrInput
	// Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
	InternalHmacUsedForSecretsConfigHmac pulumi.StringPtrInput
	// Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
	InternalSecretsConfigHmac pulumi.StringPtrInput
	// The storage bucket name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The ID of the plugin that should back the resource. This or pluginName must be defined.
	PluginId pulumi.StringPtrInput
	// The name of the plugin that should back the resource. This or pluginId must be defined.
	PluginName pulumi.StringPtrInput
	// The scope for this storage bucket.
	ScopeId pulumi.StringPtrInput
	// The HMAC'd secrets value returned from the server.
	SecretsHmac pulumi.StringPtrInput
	// The secrets for the storage bucket. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributesJson", removing this block will NOT clear secrets from the storage bucket; this allows injecting secrets for one call, then removing them for storage.
	SecretsJson pulumi.StringPtrInput
	// Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
	WorkerFilter pulumi.StringPtrInput
}

func (StorageBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketState)(nil)).Elem()
}

type storageBucketArgs struct {
	// The attributes for the storage bucket. The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
	AttributesJson *string `pulumi:"attributesJson"`
	// The name of the bucket within the external object store service.
	BucketName string `pulumi:"bucketName"`
	// The prefix used to organize the data held within the external object store.
	BucketPrefix *string `pulumi:"bucketPrefix"`
	// The storage bucket description.
	Description *string `pulumi:"description"`
	// The storage bucket name. Defaults to the resource name.
	Name *string `pulumi:"name"`
	// The ID of the plugin that should back the resource. This or pluginName must be defined.
	PluginId *string `pulumi:"pluginId"`
	// The name of the plugin that should back the resource. This or pluginId must be defined.
	PluginName *string `pulumi:"pluginName"`
	// The scope for this storage bucket.
	ScopeId string `pulumi:"scopeId"`
	// The secrets for the storage bucket. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributesJson", removing this block will NOT clear secrets from the storage bucket; this allows injecting secrets for one call, then removing them for storage.
	SecretsJson *string `pulumi:"secretsJson"`
	// Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
	WorkerFilter string `pulumi:"workerFilter"`
}

// The set of arguments for constructing a StorageBucket resource.
type StorageBucketArgs struct {
	// The attributes for the storage bucket. The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
	AttributesJson pulumi.StringPtrInput
	// The name of the bucket within the external object store service.
	BucketName pulumi.StringInput
	// The prefix used to organize the data held within the external object store.
	BucketPrefix pulumi.StringPtrInput
	// The storage bucket description.
	Description pulumi.StringPtrInput
	// The storage bucket name. Defaults to the resource name.
	Name pulumi.StringPtrInput
	// The ID of the plugin that should back the resource. This or pluginName must be defined.
	PluginId pulumi.StringPtrInput
	// The name of the plugin that should back the resource. This or pluginId must be defined.
	PluginName pulumi.StringPtrInput
	// The scope for this storage bucket.
	ScopeId pulumi.StringInput
	// The secrets for the storage bucket. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributesJson", removing this block will NOT clear secrets from the storage bucket; this allows injecting secrets for one call, then removing them for storage.
	SecretsJson pulumi.StringPtrInput
	// Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
	WorkerFilter pulumi.StringInput
}

func (StorageBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketArgs)(nil)).Elem()
}

type StorageBucketInput interface {
	pulumi.Input

	ToStorageBucketOutput() StorageBucketOutput
	ToStorageBucketOutputWithContext(ctx context.Context) StorageBucketOutput
}

func (*StorageBucket) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucket)(nil)).Elem()
}

func (i *StorageBucket) ToStorageBucketOutput() StorageBucketOutput {
	return i.ToStorageBucketOutputWithContext(context.Background())
}

func (i *StorageBucket) ToStorageBucketOutputWithContext(ctx context.Context) StorageBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketOutput)
}

// StorageBucketArrayInput is an input type that accepts StorageBucketArray and StorageBucketArrayOutput values.
// You can construct a concrete instance of `StorageBucketArrayInput` via:
//
//	StorageBucketArray{ StorageBucketArgs{...} }
type StorageBucketArrayInput interface {
	pulumi.Input

	ToStorageBucketArrayOutput() StorageBucketArrayOutput
	ToStorageBucketArrayOutputWithContext(context.Context) StorageBucketArrayOutput
}

type StorageBucketArray []StorageBucketInput

func (StorageBucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageBucket)(nil)).Elem()
}

func (i StorageBucketArray) ToStorageBucketArrayOutput() StorageBucketArrayOutput {
	return i.ToStorageBucketArrayOutputWithContext(context.Background())
}

func (i StorageBucketArray) ToStorageBucketArrayOutputWithContext(ctx context.Context) StorageBucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketArrayOutput)
}

// StorageBucketMapInput is an input type that accepts StorageBucketMap and StorageBucketMapOutput values.
// You can construct a concrete instance of `StorageBucketMapInput` via:
//
//	StorageBucketMap{ "key": StorageBucketArgs{...} }
type StorageBucketMapInput interface {
	pulumi.Input

	ToStorageBucketMapOutput() StorageBucketMapOutput
	ToStorageBucketMapOutputWithContext(context.Context) StorageBucketMapOutput
}

type StorageBucketMap map[string]StorageBucketInput

func (StorageBucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageBucket)(nil)).Elem()
}

func (i StorageBucketMap) ToStorageBucketMapOutput() StorageBucketMapOutput {
	return i.ToStorageBucketMapOutputWithContext(context.Background())
}

func (i StorageBucketMap) ToStorageBucketMapOutputWithContext(ctx context.Context) StorageBucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketMapOutput)
}

type StorageBucketOutput struct{ *pulumi.OutputState }

func (StorageBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucket)(nil)).Elem()
}

func (o StorageBucketOutput) ToStorageBucketOutput() StorageBucketOutput {
	return o
}

func (o StorageBucketOutput) ToStorageBucketOutputWithContext(ctx context.Context) StorageBucketOutput {
	return o
}

// The attributes for the storage bucket. The "region" attribute field is required when creating an AWS storage bucket. Values are either encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the storage bucket.
func (o StorageBucketOutput) AttributesJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringPtrOutput { return v.AttributesJson }).(pulumi.StringPtrOutput)
}

// The name of the bucket within the external object store service.
func (o StorageBucketOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.BucketName }).(pulumi.StringOutput)
}

// The prefix used to organize the data held within the external object store.
func (o StorageBucketOutput) BucketPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringPtrOutput { return v.BucketPrefix }).(pulumi.StringPtrOutput)
}

// The storage bucket description.
func (o StorageBucketOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Internal only. Used to force update so that we can always check the value of secrets.
func (o StorageBucketOutput) InternalForceUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.InternalForceUpdate }).(pulumi.StringOutput)
}

// Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift detection.
func (o StorageBucketOutput) InternalHmacUsedForSecretsConfigHmac() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.InternalHmacUsedForSecretsConfigHmac }).(pulumi.StringOutput)
}

// Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
func (o StorageBucketOutput) InternalSecretsConfigHmac() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.InternalSecretsConfigHmac }).(pulumi.StringOutput)
}

// The storage bucket name. Defaults to the resource name.
func (o StorageBucketOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the plugin that should back the resource. This or pluginName must be defined.
func (o StorageBucketOutput) PluginId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.PluginId }).(pulumi.StringOutput)
}

// The name of the plugin that should back the resource. This or pluginId must be defined.
func (o StorageBucketOutput) PluginName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringPtrOutput { return v.PluginName }).(pulumi.StringPtrOutput)
}

// The scope for this storage bucket.
func (o StorageBucketOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

// The HMAC'd secrets value returned from the server.
func (o StorageBucketOutput) SecretsHmac() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.SecretsHmac }).(pulumi.StringOutput)
}

// The secrets for the storage bucket. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributesJson", removing this block will NOT clear secrets from the storage bucket; this allows injecting secrets for one call, then removing them for storage.
func (o StorageBucketOutput) SecretsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringPtrOutput { return v.SecretsJson }).(pulumi.StringPtrOutput)
}

// Filters to the worker(s) that can handle requests for this storage bucket. The filter must match an existing worker in order to create a storage bucket.
func (o StorageBucketOutput) WorkerFilter() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucket) pulumi.StringOutput { return v.WorkerFilter }).(pulumi.StringOutput)
}

type StorageBucketArrayOutput struct{ *pulumi.OutputState }

func (StorageBucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageBucket)(nil)).Elem()
}

func (o StorageBucketArrayOutput) ToStorageBucketArrayOutput() StorageBucketArrayOutput {
	return o
}

func (o StorageBucketArrayOutput) ToStorageBucketArrayOutputWithContext(ctx context.Context) StorageBucketArrayOutput {
	return o
}

func (o StorageBucketArrayOutput) Index(i pulumi.IntInput) StorageBucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StorageBucket {
		return vs[0].([]*StorageBucket)[vs[1].(int)]
	}).(StorageBucketOutput)
}

type StorageBucketMapOutput struct{ *pulumi.OutputState }

func (StorageBucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageBucket)(nil)).Elem()
}

func (o StorageBucketMapOutput) ToStorageBucketMapOutput() StorageBucketMapOutput {
	return o
}

func (o StorageBucketMapOutput) ToStorageBucketMapOutputWithContext(ctx context.Context) StorageBucketMapOutput {
	return o
}

func (o StorageBucketMapOutput) MapIndex(k pulumi.StringInput) StorageBucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StorageBucket {
		return vs[0].(map[string]*StorageBucket)[vs[1].(string)]
	}).(StorageBucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketInput)(nil)).Elem(), &StorageBucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketArrayInput)(nil)).Elem(), StorageBucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketMapInput)(nil)).Elem(), StorageBucketMap{})
	pulumi.RegisterOutputType(StorageBucketOutput{})
	pulumi.RegisterOutputType(StorageBucketArrayOutput{})
	pulumi.RegisterOutputType(StorageBucketMapOutput{})
}
