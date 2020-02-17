// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Aws
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_auth_backend_sts_role.html.md.
    /// </summary>
    public partial class AuthBackendStsRole : Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS account ID to configure the STS role for.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The path the AWS auth backend being configured was
        /// mounted at.  Defaults to `aws`.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// The STS role to assume when verifying requests made
        /// by EC2 instances in the account specified by `account_id`.
        /// </summary>
        [Output("stsRole")]
        public Output<string> StsRole { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendStsRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendStsRole(string name, AuthBackendStsRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendStsRole:AuthBackendStsRole", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendStsRole(string name, Input<string> id, AuthBackendStsRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendStsRole:AuthBackendStsRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthBackendStsRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendStsRole Get(string name, Input<string> id, AuthBackendStsRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendStsRole(name, id, state, options);
        }
    }

    public sealed class AuthBackendStsRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS account ID to configure the STS role for.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// The path the AWS auth backend being configured was
        /// mounted at.  Defaults to `aws`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The STS role to assume when verifying requests made
        /// by EC2 instances in the account specified by `account_id`.
        /// </summary>
        [Input("stsRole", required: true)]
        public Input<string> StsRole { get; set; } = null!;

        public AuthBackendStsRoleArgs()
        {
        }
    }

    public sealed class AuthBackendStsRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS account ID to configure the STS role for.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The path the AWS auth backend being configured was
        /// mounted at.  Defaults to `aws`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The STS role to assume when verifying requests made
        /// by EC2 instances in the account specified by `account_id`.
        /// </summary>
        [Input("stsRole")]
        public Input<string>? StsRole { get; set; }

        public AuthBackendStsRoleState()
        {
        }
    }
}
