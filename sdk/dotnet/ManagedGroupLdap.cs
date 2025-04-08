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
    /// The managed group resource allows you to configure a Boundary group.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Boundary = Pulumi.Boundary;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var org = new Boundary.Scope("org", new()
    ///     {
    ///         Name = "organization_one",
    ///         Description = "My first scope!",
    ///         ScopeId = "global",
    ///         AutoCreateAdminRole = true,
    ///         AutoCreateDefaultRole = true,
    ///     });
    /// 
    ///     var forumsysLdap = new Boundary.AuthMethodLdap("forumsys_ldap", new()
    ///     {
    ///         Name = "forumsys public LDAP",
    ///         ScopeId = "global",
    ///         Urls = new[]
    ///         {
    ///             "ldap://ldap.forumsys.com",
    ///         },
    ///         UserDn = "dc=example,dc=com",
    ///         UserAttr = "uid",
    ///         GroupDn = "dc=example,dc=com",
    ///         BindDn = "cn=read-only-admin,dc=example,dc=com",
    ///         BindPassword = "password",
    ///         State = "active-public",
    ///         EnableGroups = true,
    ///         DiscoverDn = true,
    ///     });
    /// 
    ///     var forumsysScientists = new Boundary.ManagedGroupLdap("forumsys_scientists", new()
    ///     {
    ///         Name = "scientists",
    ///         Description = "forumsys scientists managed group",
    ///         AuthMethodId = forumsysLdap.Id,
    ///         GroupNames = new[]
    ///         {
    ///             "Scientists",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import boundary:index/managedGroupLdap:ManagedGroupLdap foo &lt;my-id&gt;
    /// ```
    /// </summary>
    [BoundaryResourceType("boundary:index/managedGroupLdap:ManagedGroupLdap")]
    public partial class ManagedGroupLdap : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource ID for the auth method.
        /// </summary>
        [Output("authMethodId")]
        public Output<string> AuthMethodId { get; private set; } = null!;

        /// <summary>
        /// The managed group description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The list of groups that make up the managed group.
        /// </summary>
        [Output("groupNames")]
        public Output<ImmutableArray<string>> GroupNames { get; private set; } = null!;

        /// <summary>
        /// The managed group name. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedGroupLdap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedGroupLdap(string name, ManagedGroupLdapArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/managedGroupLdap:ManagedGroupLdap", name, args ?? new ManagedGroupLdapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedGroupLdap(string name, Input<string> id, ManagedGroupLdapState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/managedGroupLdap:ManagedGroupLdap", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ManagedGroupLdap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedGroupLdap Get(string name, Input<string> id, ManagedGroupLdapState? state = null, CustomResourceOptions? options = null)
        {
            return new ManagedGroupLdap(name, id, state, options);
        }
    }

    public sealed class ManagedGroupLdapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID for the auth method.
        /// </summary>
        [Input("authMethodId", required: true)]
        public Input<string> AuthMethodId { get; set; } = null!;

        /// <summary>
        /// The managed group description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("groupNames", required: true)]
        private InputList<string>? _groupNames;

        /// <summary>
        /// The list of groups that make up the managed group.
        /// </summary>
        public InputList<string> GroupNames
        {
            get => _groupNames ?? (_groupNames = new InputList<string>());
            set => _groupNames = value;
        }

        /// <summary>
        /// The managed group name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ManagedGroupLdapArgs()
        {
        }
        public static new ManagedGroupLdapArgs Empty => new ManagedGroupLdapArgs();
    }

    public sealed class ManagedGroupLdapState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID for the auth method.
        /// </summary>
        [Input("authMethodId")]
        public Input<string>? AuthMethodId { get; set; }

        /// <summary>
        /// The managed group description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("groupNames")]
        private InputList<string>? _groupNames;

        /// <summary>
        /// The list of groups that make up the managed group.
        /// </summary>
        public InputList<string> GroupNames
        {
            get => _groupNames ?? (_groupNames = new InputList<string>());
            set => _groupNames = value;
        }

        /// <summary>
        /// The managed group name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ManagedGroupLdapState()
        {
        }
        public static new ManagedGroupLdapState Empty => new ManagedGroupLdapState();
    }
}
