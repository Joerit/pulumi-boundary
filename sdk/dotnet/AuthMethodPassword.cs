// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Boundary
{
    /// <summary>
    /// The auth method resource allows you to configure a Boundary auth_method_password.
    /// </summary>
    [BoundaryResourceType("boundary:index/authMethodPassword:AuthMethodPassword")]
    public partial class AuthMethodPassword : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The auth method description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The minimum login name length.
        /// </summary>
        [Output("minLoginNameLength")]
        public Output<int> MinLoginNameLength { get; private set; } = null!;

        /// <summary>
        /// The minimum password length.
        /// </summary>
        [Output("minPasswordLength")]
        public Output<int> MinPasswordLength { get; private set; } = null!;

        /// <summary>
        /// The auth method name. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The scope ID.
        /// </summary>
        [Output("scopeId")]
        public Output<string> ScopeId { get; private set; } = null!;

        /// <summary>
        /// The resource type, hardcoded per resource
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AuthMethodPassword resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthMethodPassword(string name, AuthMethodPasswordArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/authMethodPassword:AuthMethodPassword", name, args ?? new AuthMethodPasswordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthMethodPassword(string name, Input<string> id, AuthMethodPasswordState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/authMethodPassword:AuthMethodPassword", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthMethodPassword resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthMethodPassword Get(string name, Input<string> id, AuthMethodPasswordState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthMethodPassword(name, id, state, options);
        }
    }

    public sealed class AuthMethodPasswordArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The auth method description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The minimum login name length.
        /// </summary>
        [Input("minLoginNameLength")]
        public Input<int>? MinLoginNameLength { get; set; }

        /// <summary>
        /// The minimum password length.
        /// </summary>
        [Input("minPasswordLength")]
        public Input<int>? MinPasswordLength { get; set; }

        /// <summary>
        /// The auth method name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The scope ID.
        /// </summary>
        [Input("scopeId", required: true)]
        public Input<string> ScopeId { get; set; } = null!;

        /// <summary>
        /// The resource type, hardcoded per resource
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AuthMethodPasswordArgs()
        {
        }
        public static new AuthMethodPasswordArgs Empty => new AuthMethodPasswordArgs();
    }

    public sealed class AuthMethodPasswordState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The auth method description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The minimum login name length.
        /// </summary>
        [Input("minLoginNameLength")]
        public Input<int>? MinLoginNameLength { get; set; }

        /// <summary>
        /// The minimum password length.
        /// </summary>
        [Input("minPasswordLength")]
        public Input<int>? MinPasswordLength { get; set; }

        /// <summary>
        /// The auth method name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The scope ID.
        /// </summary>
        [Input("scopeId")]
        public Input<string>? ScopeId { get; set; }

        /// <summary>
        /// The resource type, hardcoded per resource
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AuthMethodPasswordState()
        {
        }
        public static new AuthMethodPasswordState Empty => new AuthMethodPasswordState();
    }
}
