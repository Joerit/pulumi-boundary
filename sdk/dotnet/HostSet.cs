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
    /// Deprecated: use `boundary.HostSetStatic` instead.
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
    ///     var project = new Boundary.Scope("project", new()
    ///     {
    ///         Name = "project_one",
    ///         Description = "My first scope!",
    ///         ScopeId = org.Id,
    ///         AutoCreateAdminRole = true,
    ///     });
    /// 
    ///     var @static = new Boundary.HostCatalog("static", new()
    ///     {
    ///         Type = "static",
    ///         ScopeId = project.Id,
    ///     });
    /// 
    ///     var first = new Boundary.Host("first", new()
    ///     {
    ///         Type = "static",
    ///         Name = "host_1",
    ///         Description = "My first host!",
    ///         Address = "10.0.0.1",
    ///         HostCatalogId = @static.Id,
    ///     });
    /// 
    ///     var second = new Boundary.Host("second", new()
    ///     {
    ///         Type = "static",
    ///         Name = "host_2",
    ///         Description = "My second host!",
    ///         Address = "10.0.0.2",
    ///         HostCatalogId = @static.Id,
    ///     });
    /// 
    ///     var web = new Boundary.HostSet("web", new()
    ///     {
    ///         HostCatalogId = @static.Id,
    ///         Type = "static",
    ///         HostIds = new[]
    ///         {
    ///             first.Id,
    ///             second.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import boundary:index/hostSet:HostSet foo &lt;my-id&gt;
    /// ```
    /// </summary>
    [BoundaryResourceType("boundary:index/hostSet:HostSet")]
    public partial class HostSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The host set description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The catalog for the host set.
        /// </summary>
        [Output("hostCatalogId")]
        public Output<string> HostCatalogId { get; private set; } = null!;

        /// <summary>
        /// The list of host IDs contained in this set.
        /// </summary>
        [Output("hostIds")]
        public Output<ImmutableArray<string>> HostIds { get; private set; } = null!;

        /// <summary>
        /// The host set name. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of host set
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a HostSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HostSet(string name, HostSetArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/hostSet:HostSet", name, args ?? new HostSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HostSet(string name, Input<string> id, HostSetState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/hostSet:HostSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HostSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HostSet Get(string name, Input<string> id, HostSetState? state = null, CustomResourceOptions? options = null)
        {
            return new HostSet(name, id, state, options);
        }
    }

    public sealed class HostSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The host set description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The catalog for the host set.
        /// </summary>
        [Input("hostCatalogId", required: true)]
        public Input<string> HostCatalogId { get; set; } = null!;

        [Input("hostIds")]
        private InputList<string>? _hostIds;

        /// <summary>
        /// The list of host IDs contained in this set.
        /// </summary>
        public InputList<string> HostIds
        {
            get => _hostIds ?? (_hostIds = new InputList<string>());
            set => _hostIds = value;
        }

        /// <summary>
        /// The host set name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of host set
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public HostSetArgs()
        {
        }
        public static new HostSetArgs Empty => new HostSetArgs();
    }

    public sealed class HostSetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The host set description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The catalog for the host set.
        /// </summary>
        [Input("hostCatalogId")]
        public Input<string>? HostCatalogId { get; set; }

        [Input("hostIds")]
        private InputList<string>? _hostIds;

        /// <summary>
        /// The list of host IDs contained in this set.
        /// </summary>
        public InputList<string> HostIds
        {
            get => _hostIds ?? (_hostIds = new InputList<string>());
            set => _hostIds = value;
        }

        /// <summary>
        /// The host set name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of host set
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public HostSetState()
        {
        }
        public static new HostSetState Empty => new HostSetState();
    }
}
