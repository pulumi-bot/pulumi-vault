// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.AppRole
{
    /// <summary>
    /// Logs into Vault using the AppRole auth backend. See the [Vault
    /// documentation](https://www.vaultproject.io/docs/auth/approle.html) for more
    /// information.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/approle_auth_backend_login.html.markdown.
    /// </summary>
    public partial class AuthBackendLogin : Pulumi.CustomResource
    {
        /// <summary>
        /// The accessor for the token.
        /// </summary>
        [Output("accessor")]
        public Output<string> Accessor { get; private set; } = null!;

        /// <summary>
        /// The unique path of the Vault backend to log in with.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// The Vault token created.
        /// </summary>
        [Output("clientToken")]
        public Output<string> ClientToken { get; private set; } = null!;

        /// <summary>
        /// How long the token is valid for, in seconds.
        /// </summary>
        [Output("leaseDuration")]
        public Output<int> LeaseDuration { get; private set; } = null!;

        /// <summary>
        /// The date and time the lease started, in RFC 3339 format.
        /// </summary>
        [Output("leaseStarted")]
        public Output<string> LeaseStarted { get; private set; } = null!;

        /// <summary>
        /// The metadata associated with the token.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        /// <summary>
        /// A list of policies applied to the token.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// Whether the token is renewable or not.
        /// </summary>
        [Output("renewable")]
        public Output<bool> Renewable { get; private set; } = null!;

        /// <summary>
        /// The ID of the role to log in with.
        /// </summary>
        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;

        /// <summary>
        /// The secret ID of the role to log in with. Required
        /// unless `bind_secret_id` is set to false on the role.
        /// </summary>
        [Output("secretId")]
        public Output<string?> SecretId { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendLogin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendLogin(string name, AuthBackendLoginArgs args, CustomResourceOptions? options = null)
            : base("vault:appRole/authBackendLogin:AuthBackendLogin", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendLogin(string name, Input<string> id, AuthBackendLoginState? state = null, CustomResourceOptions? options = null)
            : base("vault:appRole/authBackendLogin:AuthBackendLogin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthBackendLogin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendLogin Get(string name, Input<string> id, AuthBackendLoginState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendLogin(name, id, state, options);
        }
    }

    public sealed class AuthBackendLoginArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique path of the Vault backend to log in with.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The ID of the role to log in with.
        /// </summary>
        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        /// <summary>
        /// The secret ID of the role to log in with. Required
        /// unless `bind_secret_id` is set to false on the role.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        public AuthBackendLoginArgs()
        {
        }
    }

    public sealed class AuthBackendLoginState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The accessor for the token.
        /// </summary>
        [Input("accessor")]
        public Input<string>? Accessor { get; set; }

        /// <summary>
        /// The unique path of the Vault backend to log in with.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The Vault token created.
        /// </summary>
        [Input("clientToken")]
        public Input<string>? ClientToken { get; set; }

        /// <summary>
        /// How long the token is valid for, in seconds.
        /// </summary>
        [Input("leaseDuration")]
        public Input<int>? LeaseDuration { get; set; }

        /// <summary>
        /// The date and time the lease started, in RFC 3339 format.
        /// </summary>
        [Input("leaseStarted")]
        public Input<string>? LeaseStarted { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// The metadata associated with the token.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// A list of policies applied to the token.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// Whether the token is renewable or not.
        /// </summary>
        [Input("renewable")]
        public Input<bool>? Renewable { get; set; }

        /// <summary>
        /// The ID of the role to log in with.
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        /// <summary>
        /// The secret ID of the role to log in with. Required
        /// unless `bind_secret_id` is set to false on the role.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        public AuthBackendLoginState()
        {
        }
    }
}
