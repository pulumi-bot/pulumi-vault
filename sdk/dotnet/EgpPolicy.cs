// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    /// <summary>
    /// Provides a resource to manage Endpoint Governing Policy (EGP) via [Sentinel](https://www.vaultproject.io/docs/enterprise/sentinel/index.html).
    /// 
    /// **Note** this feature is available only with Vault Enterprise.
    /// 
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/egp_policy.html.md.
    /// </summary>
    public partial class EgpPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        /// </summary>
        [Output("enforcementLevel")]
        public Output<string> EnforcementLevel { get; private set; } = null!;

        /// <summary>
        /// The name of the policy
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of paths to which the policy will be applied to
        /// </summary>
        [Output("paths")]
        public Output<ImmutableArray<string>> Paths { get; private set; } = null!;

        /// <summary>
        /// String containing a Sentinel policy
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a EgpPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EgpPolicy(string name, EgpPolicyArgs args, CustomResourceOptions? options = null)
            : base("vault:index/egpPolicy:EgpPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private EgpPolicy(string name, Input<string> id, EgpPolicyState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/egpPolicy:EgpPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EgpPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EgpPolicy Get(string name, Input<string> id, EgpPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new EgpPolicy(name, id, state, options);
        }
    }

    public sealed class EgpPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        /// </summary>
        [Input("enforcementLevel", required: true)]
        public Input<string> EnforcementLevel { get; set; } = null!;

        /// <summary>
        /// The name of the policy
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("paths", required: true)]
        private InputList<string>? _paths;

        /// <summary>
        /// List of paths to which the policy will be applied to
        /// </summary>
        public InputList<string> Paths
        {
            get => _paths ?? (_paths = new InputList<string>());
            set => _paths = value;
        }

        /// <summary>
        /// String containing a Sentinel policy
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public EgpPolicyArgs()
        {
        }
    }

    public sealed class EgpPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        /// </summary>
        [Input("enforcementLevel")]
        public Input<string>? EnforcementLevel { get; set; }

        /// <summary>
        /// The name of the policy
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("paths")]
        private InputList<string>? _paths;

        /// <summary>
        /// List of paths to which the policy will be applied to
        /// </summary>
        public InputList<string> Paths
        {
            get => _paths ?? (_paths = new InputList<string>());
            set => _paths = value;
        }

        /// <summary>
        /// String containing a Sentinel policy
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public EgpPolicyState()
        {
        }
    }
}
