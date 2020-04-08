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
    /// Provides a resource to manage [Duo MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-duo.html).
    /// 
    /// **Note** this feature is available only with Vault Enterprise.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/mfa_duo.html.md.
    /// </summary>
    public partial class MfaDuo : Pulumi.CustomResource
    {
        /// <summary>
        /// API hostname for Duo.
        /// </summary>
        [Output("apiHostname")]
        public Output<string> ApiHostname { get; private set; } = null!;

        /// <summary>
        /// Integration key for Duo.
        /// </summary>
        [Output("integrationKey")]
        public Output<string> IntegrationKey { get; private set; } = null!;

        /// <summary>
        /// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of
        /// Aliases associated with this mount as the username in the mapping.
        /// </summary>
        [Output("mountAccessor")]
        public Output<string> MountAccessor { get; private set; } = null!;

        /// <summary>
        /// Name of the MFA method.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Push information for Duo.
        /// </summary>
        [Output("pushInfo")]
        public Output<string?> PushInfo { get; private set; } = null!;

        /// <summary>
        /// Secret key for Duo.
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;

        /// <summary>
        /// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in
        /// `{{}}`.
        /// </summary>
        [Output("usernameFormat")]
        public Output<string?> UsernameFormat { get; private set; } = null!;


        /// <summary>
        /// Create a MfaDuo resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MfaDuo(string name, MfaDuoArgs args, CustomResourceOptions? options = null)
            : base("vault:index/mfaDuo:MfaDuo", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private MfaDuo(string name, Input<string> id, MfaDuoState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/mfaDuo:MfaDuo", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MfaDuo resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MfaDuo Get(string name, Input<string> id, MfaDuoState? state = null, CustomResourceOptions? options = null)
        {
            return new MfaDuo(name, id, state, options);
        }
    }

    public sealed class MfaDuoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API hostname for Duo.
        /// </summary>
        [Input("apiHostname", required: true)]
        public Input<string> ApiHostname { get; set; } = null!;

        /// <summary>
        /// Integration key for Duo.
        /// </summary>
        [Input("integrationKey", required: true)]
        public Input<string> IntegrationKey { get; set; } = null!;

        /// <summary>
        /// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of
        /// Aliases associated with this mount as the username in the mapping.
        /// </summary>
        [Input("mountAccessor", required: true)]
        public Input<string> MountAccessor { get; set; } = null!;

        /// <summary>
        /// Name of the MFA method.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Push information for Duo.
        /// </summary>
        [Input("pushInfo")]
        public Input<string>? PushInfo { get; set; }

        /// <summary>
        /// Secret key for Duo.
        /// </summary>
        [Input("secretKey", required: true)]
        public Input<string> SecretKey { get; set; } = null!;

        /// <summary>
        /// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in
        /// `{{}}`.
        /// </summary>
        [Input("usernameFormat")]
        public Input<string>? UsernameFormat { get; set; }

        public MfaDuoArgs()
        {
        }
    }

    public sealed class MfaDuoState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API hostname for Duo.
        /// </summary>
        [Input("apiHostname")]
        public Input<string>? ApiHostname { get; set; }

        /// <summary>
        /// Integration key for Duo.
        /// </summary>
        [Input("integrationKey")]
        public Input<string>? IntegrationKey { get; set; }

        /// <summary>
        /// The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of
        /// Aliases associated with this mount as the username in the mapping.
        /// </summary>
        [Input("mountAccessor")]
        public Input<string>? MountAccessor { get; set; }

        /// <summary>
        /// Name of the MFA method.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Push information for Duo.
        /// </summary>
        [Input("pushInfo")]
        public Input<string>? PushInfo { get; set; }

        /// <summary>
        /// Secret key for Duo.
        /// </summary>
        [Input("secretKey")]
        public Input<string>? SecretKey { get; set; }

        /// <summary>
        /// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in
        /// `{{}}`.
        /// </summary>
        [Input("usernameFormat")]
        public Input<string>? UsernameFormat { get; set; }

        public MfaDuoState()
        {
        }
    }
}
