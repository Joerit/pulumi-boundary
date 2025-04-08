// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Boundary
{
    [BoundaryResourceType("boundary:index/hostCatalogPlugin:HostCatalogPlugin")]
    public partial class HostCatalogPlugin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a
        /// file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
        /// </summary>
        [Output("attributesJson")]
        public Output<string?> AttributesJson { get; private set; } = null!;

        /// <summary>
        /// The host catalog description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Internal only. Used to force update so that we can always check the value of secrets.
        /// </summary>
        [Output("internalForceUpdate")]
        public Output<string> InternalForceUpdate { get; private set; } = null!;

        /// <summary>
        /// Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift
        /// detection.
        /// </summary>
        [Output("internalHmacUsedForSecretsConfigHmac")]
        public Output<string> InternalHmacUsedForSecretsConfigHmac { get; private set; } = null!;

        /// <summary>
        /// Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
        /// </summary>
        [Output("internalSecretsConfigHmac")]
        public Output<string> InternalSecretsConfigHmac { get; private set; } = null!;

        /// <summary>
        /// The host catalog name. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the plugin that should back the resource. This or plugin_name must be defined.
        /// </summary>
        [Output("pluginId")]
        public Output<string> PluginId { get; private set; } = null!;

        /// <summary>
        /// The name of the plugin that should back the resource. This or plugin_id must be defined.
        /// </summary>
        [Output("pluginName")]
        public Output<string> PluginName { get; private set; } = null!;

        /// <summary>
        /// The scope ID in which the resource is created.
        /// </summary>
        [Output("scopeId")]
        public Output<string> ScopeId { get; private set; } = null!;

        /// <summary>
        /// The HMAC'd secrets value returned from the server.
        /// </summary>
        [Output("secretsHmac")]
        public Output<string> SecretsHmac { get; private set; } = null!;

        /// <summary>
        /// The secrets for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a
        /// file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing
        /// this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them
        /// for storage.
        /// </summary>
        [Output("secretsJson")]
        public Output<string?> SecretsJson { get; private set; } = null!;

        /// <summary>
        /// HCP Only. A filter used to control which PKI workers can handle dynamic host catalog requests.
        /// </summary>
        [Output("workerFilter")]
        public Output<string?> WorkerFilter { get; private set; } = null!;


        /// <summary>
        /// Create a HostCatalogPlugin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HostCatalogPlugin(string name, HostCatalogPluginArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/hostCatalogPlugin:HostCatalogPlugin", name, args ?? new HostCatalogPluginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HostCatalogPlugin(string name, Input<string> id, HostCatalogPluginState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/hostCatalogPlugin:HostCatalogPlugin", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "secretsJson",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HostCatalogPlugin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HostCatalogPlugin Get(string name, Input<string> id, HostCatalogPluginState? state = null, CustomResourceOptions? options = null)
        {
            return new HostCatalogPlugin(name, id, state, options);
        }
    }

    public sealed class HostCatalogPluginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a
        /// file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
        /// </summary>
        [Input("attributesJson")]
        public Input<string>? AttributesJson { get; set; }

        /// <summary>
        /// The host catalog description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Internal only. Used to force update so that we can always check the value of secrets.
        /// </summary>
        [Input("internalForceUpdate")]
        public Input<string>? InternalForceUpdate { get; set; }

        /// <summary>
        /// Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift
        /// detection.
        /// </summary>
        [Input("internalHmacUsedForSecretsConfigHmac")]
        public Input<string>? InternalHmacUsedForSecretsConfigHmac { get; set; }

        /// <summary>
        /// Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
        /// </summary>
        [Input("internalSecretsConfigHmac")]
        public Input<string>? InternalSecretsConfigHmac { get; set; }

        /// <summary>
        /// The host catalog name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the plugin that should back the resource. This or plugin_name must be defined.
        /// </summary>
        [Input("pluginId")]
        public Input<string>? PluginId { get; set; }

        /// <summary>
        /// The name of the plugin that should back the resource. This or plugin_id must be defined.
        /// </summary>
        [Input("pluginName")]
        public Input<string>? PluginName { get; set; }

        /// <summary>
        /// The scope ID in which the resource is created.
        /// </summary>
        [Input("scopeId", required: true)]
        public Input<string> ScopeId { get; set; } = null!;

        /// <summary>
        /// The HMAC'd secrets value returned from the server.
        /// </summary>
        [Input("secretsHmac")]
        public Input<string>? SecretsHmac { get; set; }

        [Input("secretsJson")]
        private Input<string>? _secretsJson;

        /// <summary>
        /// The secrets for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a
        /// file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing
        /// this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them
        /// for storage.
        /// </summary>
        public Input<string>? SecretsJson
        {
            get => _secretsJson;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretsJson = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// HCP Only. A filter used to control which PKI workers can handle dynamic host catalog requests.
        /// </summary>
        [Input("workerFilter")]
        public Input<string>? WorkerFilter { get; set; }

        public HostCatalogPluginArgs()
        {
        }
        public static new HostCatalogPluginArgs Empty => new HostCatalogPluginArgs();
    }

    public sealed class HostCatalogPluginState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The attributes for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a
        /// file:// or env:// path. Set to a string "null" or remove the block to clear all attributes in the host catalog.
        /// </summary>
        [Input("attributesJson")]
        public Input<string>? AttributesJson { get; set; }

        /// <summary>
        /// The host catalog description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Internal only. Used to force update so that we can always check the value of secrets.
        /// </summary>
        [Input("internalForceUpdate")]
        public Input<string>? InternalForceUpdate { get; set; }

        /// <summary>
        /// Internal only. The Boundary-provided HMAC used to calculate the current value of the HMAC'd config. Used for drift
        /// detection.
        /// </summary>
        [Input("internalHmacUsedForSecretsConfigHmac")]
        public Input<string>? InternalHmacUsedForSecretsConfigHmac { get; set; }

        /// <summary>
        /// Internal only. HMAC of (serverSecretsHmac + config secrets). Used for proper secrets handling.
        /// </summary>
        [Input("internalSecretsConfigHmac")]
        public Input<string>? InternalSecretsConfigHmac { get; set; }

        /// <summary>
        /// The host catalog name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the plugin that should back the resource. This or plugin_name must be defined.
        /// </summary>
        [Input("pluginId")]
        public Input<string>? PluginId { get; set; }

        /// <summary>
        /// The name of the plugin that should back the resource. This or plugin_id must be defined.
        /// </summary>
        [Input("pluginName")]
        public Input<string>? PluginName { get; set; }

        /// <summary>
        /// The scope ID in which the resource is created.
        /// </summary>
        [Input("scopeId")]
        public Input<string>? ScopeId { get; set; }

        /// <summary>
        /// The HMAC'd secrets value returned from the server.
        /// </summary>
        [Input("secretsHmac")]
        public Input<string>? SecretsHmac { get; set; }

        [Input("secretsJson")]
        private Input<string>? _secretsJson;

        /// <summary>
        /// The secrets for the host catalog. Either values encoded with the "jsonencode" function, pre-escaped JSON string, or a
        /// file:// or env:// path. Set to a string "null" to clear any existing values. NOTE: Unlike "attributes_json", removing
        /// this block will NOT clear secrets from the host catalog; this allows injecting secrets for one call, then removing them
        /// for storage.
        /// </summary>
        public Input<string>? SecretsJson
        {
            get => _secretsJson;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretsJson = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// HCP Only. A filter used to control which PKI workers can handle dynamic host catalog requests.
        /// </summary>
        [Input("workerFilter")]
        public Input<string>? WorkerFilter { get; set; }

        public HostCatalogPluginState()
        {
        }
        public static new HostCatalogPluginState Empty => new HostCatalogPluginState();
    }
}
