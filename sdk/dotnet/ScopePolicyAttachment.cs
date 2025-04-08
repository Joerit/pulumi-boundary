// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Boundary
{
    [BoundaryResourceType("boundary:index/scopePolicyAttachment:ScopePolicyAttachment")]
    public partial class ScopePolicyAttachment : global::Pulumi.CustomResource
    {
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        [Output("scopeId")]
        public Output<string> ScopeId { get; private set; } = null!;


        /// <summary>
        /// Create a ScopePolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScopePolicyAttachment(string name, ScopePolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/scopePolicyAttachment:ScopePolicyAttachment", name, args ?? new ScopePolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScopePolicyAttachment(string name, Input<string> id, ScopePolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/scopePolicyAttachment:ScopePolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ScopePolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScopePolicyAttachment Get(string name, Input<string> id, ScopePolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ScopePolicyAttachment(name, id, state, options);
        }
    }

    public sealed class ScopePolicyAttachmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        [Input("scopeId", required: true)]
        public Input<string> ScopeId { get; set; } = null!;

        public ScopePolicyAttachmentArgs()
        {
        }
        public static new ScopePolicyAttachmentArgs Empty => new ScopePolicyAttachmentArgs();
    }

    public sealed class ScopePolicyAttachmentState : global::Pulumi.ResourceArgs
    {
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        [Input("scopeId")]
        public Input<string>? ScopeId { get; set; }

        public ScopePolicyAttachmentState()
        {
        }
        public static new ScopePolicyAttachmentState Empty => new ScopePolicyAttachmentState();
    }
}
