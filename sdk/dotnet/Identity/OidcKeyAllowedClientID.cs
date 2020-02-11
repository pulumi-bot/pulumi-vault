// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_oidc_key_allowed_client_id.html.markdown.
    /// </summary>
    public partial class OidcKeyAllowedClientID : Pulumi.CustomResource
    {
        /// <summary>
        /// Client ID to allow usage with the OIDC named key
        /// </summary>
        [Output("allowedClientId")]
        public Output<string> AllowedClientId { get; private set; } = null!;

        /// <summary>
        /// Name of the OIDC Key allow the Client ID.
        /// </summary>
        [Output("keyName")]
        public Output<string> KeyName { get; private set; } = null!;


        /// <summary>
        /// Create a OidcKeyAllowedClientID resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OidcKeyAllowedClientID(string name, OidcKeyAllowedClientIDArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private OidcKeyAllowedClientID(string name, Input<string> id, OidcKeyAllowedClientIDState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OidcKeyAllowedClientID resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OidcKeyAllowedClientID Get(string name, Input<string> id, OidcKeyAllowedClientIDState? state = null, CustomResourceOptions? options = null)
        {
            return new OidcKeyAllowedClientID(name, id, state, options);
        }
    }

    public sealed class OidcKeyAllowedClientIDArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client ID to allow usage with the OIDC named key
        /// </summary>
        [Input("allowedClientId", required: true)]
        public Input<string> AllowedClientId { get; set; } = null!;

        /// <summary>
        /// Name of the OIDC Key allow the Client ID.
        /// </summary>
        [Input("keyName", required: true)]
        public Input<string> KeyName { get; set; } = null!;

        public OidcKeyAllowedClientIDArgs()
        {
        }
    }

    public sealed class OidcKeyAllowedClientIDState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client ID to allow usage with the OIDC named key
        /// </summary>
        [Input("allowedClientId")]
        public Input<string>? AllowedClientId { get; set; }

        /// <summary>
        /// Name of the OIDC Key allow the Client ID.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        public OidcKeyAllowedClientIDState()
        {
        }
    }
}
