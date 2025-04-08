// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Boundary
{
    [BoundaryResourceType("boundary:index/accountPassword:AccountPassword")]
    public partial class AccountPassword : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource ID for the auth method.
        /// </summary>
        [Output("authMethodId")]
        public Output<string> AuthMethodId { get; private set; } = null!;

        /// <summary>
        /// The account description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The login name for this account.
        /// </summary>
        [Output("loginName")]
        public Output<string?> LoginName { get; private set; } = null!;

        /// <summary>
        /// The account name. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The account password. Only set on create, changes will not be reflected when updating account.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AccountPassword resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountPassword(string name, AccountPasswordArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/accountPassword:AccountPassword", name, args ?? new AccountPasswordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountPassword(string name, Input<string> id, AccountPasswordState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/accountPassword:AccountPassword", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AccountPassword resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountPassword Get(string name, Input<string> id, AccountPasswordState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountPassword(name, id, state, options);
        }
    }

    public sealed class AccountPasswordArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID for the auth method.
        /// </summary>
        [Input("authMethodId", required: true)]
        public Input<string> AuthMethodId { get; set; } = null!;

        /// <summary>
        /// The account description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The login name for this account.
        /// </summary>
        [Input("loginName")]
        public Input<string>? LoginName { get; set; }

        /// <summary>
        /// The account name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The account password. Only set on create, changes will not be reflected when updating account.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The resource type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AccountPasswordArgs()
        {
        }
        public static new AccountPasswordArgs Empty => new AccountPasswordArgs();
    }

    public sealed class AccountPasswordState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID for the auth method.
        /// </summary>
        [Input("authMethodId")]
        public Input<string>? AuthMethodId { get; set; }

        /// <summary>
        /// The account description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The login name for this account.
        /// </summary>
        [Input("loginName")]
        public Input<string>? LoginName { get; set; }

        /// <summary>
        /// The account name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The account password. Only set on create, changes will not be reflected when updating account.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The resource type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AccountPasswordState()
        {
        }
        public static new AccountPasswordState Empty => new AccountPasswordState();
    }
}
