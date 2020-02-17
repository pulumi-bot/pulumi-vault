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
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_oidc_key.html.md.
    /// </summary>
    public partial class OidcKey : Pulumi.CustomResource
    {
        /// <summary>
        /// Signing algorithm to use. Signing algorithm to use.
        /// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
        /// </summary>
        [Output("algorithm")]
        public Output<string?> Algorithm { get; private set; } = null!;

        /// <summary>
        /// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all
        /// roles are allowed.
        /// </summary>
        [Output("allowedClientIds")]
        public Output<ImmutableArray<string>> AllowedClientIds { get; private set; } = null!;

        /// <summary>
        /// Name of the OIDC Key to create.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// How often to generate a new signing key in number of seconds
        /// </summary>
        [Output("rotationPeriod")]
        public Output<int?> RotationPeriod { get; private set; } = null!;

        /// <summary>
        /// "Controls how long the public portion of a signing key will be
        /// available for verification after being rotated in seconds.
        /// </summary>
        [Output("verificationTtl")]
        public Output<int?> VerificationTtl { get; private set; } = null!;


        /// <summary>
        /// Create a OidcKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OidcKey(string name, OidcKeyArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:identity/oidcKey:OidcKey", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private OidcKey(string name, Input<string> id, OidcKeyState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/oidcKey:OidcKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OidcKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OidcKey Get(string name, Input<string> id, OidcKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new OidcKey(name, id, state, options);
        }
    }

    public sealed class OidcKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Signing algorithm to use. Signing algorithm to use.
        /// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("allowedClientIds")]
        private InputList<string>? _allowedClientIds;

        /// <summary>
        /// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all
        /// roles are allowed.
        /// </summary>
        public InputList<string> AllowedClientIds
        {
            get => _allowedClientIds ?? (_allowedClientIds = new InputList<string>());
            set => _allowedClientIds = value;
        }

        /// <summary>
        /// Name of the OIDC Key to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// How often to generate a new signing key in number of seconds
        /// </summary>
        [Input("rotationPeriod")]
        public Input<int>? RotationPeriod { get; set; }

        /// <summary>
        /// "Controls how long the public portion of a signing key will be
        /// available for verification after being rotated in seconds.
        /// </summary>
        [Input("verificationTtl")]
        public Input<int>? VerificationTtl { get; set; }

        public OidcKeyArgs()
        {
        }
    }

    public sealed class OidcKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Signing algorithm to use. Signing algorithm to use.
        /// Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("allowedClientIds")]
        private InputList<string>? _allowedClientIds;

        /// <summary>
        /// Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all
        /// roles are allowed.
        /// </summary>
        public InputList<string> AllowedClientIds
        {
            get => _allowedClientIds ?? (_allowedClientIds = new InputList<string>());
            set => _allowedClientIds = value;
        }

        /// <summary>
        /// Name of the OIDC Key to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// How often to generate a new signing key in number of seconds
        /// </summary>
        [Input("rotationPeriod")]
        public Input<int>? RotationPeriod { get; set; }

        /// <summary>
        /// "Controls how long the public portion of a signing key will be
        /// available for verification after being rotated in seconds.
        /// </summary>
        [Input("verificationTtl")]
        public Input<int>? VerificationTtl { get; set; }

        public OidcKeyState()
        {
        }
    }
}
