// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Boundary
{
    [BoundaryResourceType("boundary:index/credentialJson:CredentialJson")]
    public partial class CredentialJson : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The credential store in which to save this json credential.
        /// </summary>
        [Output("credentialStoreId")]
        public Output<string> CredentialStoreId { get; private set; } = null!;

        /// <summary>
        /// The description of this json credential.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of this json credential. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The object for the this json credential. Either values encoded with the "jsonencode" function, pre-escaped JSON string,
        /// or a file
        /// </summary>
        [Output("object")]
        public Output<string> Object { get; private set; } = null!;

        /// <summary>
        /// The object hmac.
        /// </summary>
        [Output("objectHmac")]
        public Output<string> ObjectHmac { get; private set; } = null!;


        /// <summary>
        /// Create a CredentialJson resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CredentialJson(string name, CredentialJsonArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/credentialJson:CredentialJson", name, args ?? new CredentialJsonArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CredentialJson(string name, Input<string> id, CredentialJsonState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/credentialJson:CredentialJson", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "object",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CredentialJson resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CredentialJson Get(string name, Input<string> id, CredentialJsonState? state = null, CustomResourceOptions? options = null)
        {
            return new CredentialJson(name, id, state, options);
        }
    }

    public sealed class CredentialJsonArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The credential store in which to save this json credential.
        /// </summary>
        [Input("credentialStoreId", required: true)]
        public Input<string> CredentialStoreId { get; set; } = null!;

        /// <summary>
        /// The description of this json credential.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of this json credential. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("object", required: true)]
        private Input<string>? _object;

        /// <summary>
        /// The object for the this json credential. Either values encoded with the "jsonencode" function, pre-escaped JSON string,
        /// or a file
        /// </summary>
        public Input<string>? Object
        {
            get => _object;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _object = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public CredentialJsonArgs()
        {
        }
        public static new CredentialJsonArgs Empty => new CredentialJsonArgs();
    }

    public sealed class CredentialJsonState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The credential store in which to save this json credential.
        /// </summary>
        [Input("credentialStoreId")]
        public Input<string>? CredentialStoreId { get; set; }

        /// <summary>
        /// The description of this json credential.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of this json credential. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("object")]
        private Input<string>? _object;

        /// <summary>
        /// The object for the this json credential. Either values encoded with the "jsonencode" function, pre-escaped JSON string,
        /// or a file
        /// </summary>
        public Input<string>? Object
        {
            get => _object;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _object = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The object hmac.
        /// </summary>
        [Input("objectHmac")]
        public Input<string>? ObjectHmac { get; set; }

        public CredentialJsonState()
        {
        }
        public static new CredentialJsonState Empty => new CredentialJsonState();
    }
}
