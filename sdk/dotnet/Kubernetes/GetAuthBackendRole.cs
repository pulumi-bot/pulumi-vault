// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Kubernetes
{
    public static partial class Invokes
    {
        /// <summary>
        /// Reads the Role of an Kubernetes from a Vault server. See the [Vault
        /// documentation](https://www.vaultproject.io/api/auth/kubernetes/index.html#read-role) for more
        /// information.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/kubernetes_auth_backend_role.html.markdown.
        /// </summary>
        public static Task<GetAuthBackendRoleResult> GetAuthBackendRole(GetAuthBackendRoleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthBackendRoleResult>("vault:kubernetes/getAuthBackendRole:getAuthBackendRole", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetAuthBackendRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name for the Kubernetes backend the role to
        /// retrieve Role attributes for resides in. Defaults to "kubernetes".
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("boundCidrs")]
        private InputList<string>? _boundCidrs;
        public InputList<string> BoundCidrs
        {
            get => _boundCidrs ?? (_boundCidrs = new InputList<string>());
            set => _boundCidrs = value;
        }

        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        [Input("numUses")]
        public Input<int>? NumUses { get; set; }

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The name of the role to retrieve the Role attributes for.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        [Input("tokenBoundCidrs")]
        private InputList<string>? _tokenBoundCidrs;
        public InputList<string> TokenBoundCidrs
        {
            get => _tokenBoundCidrs ?? (_tokenBoundCidrs = new InputList<string>());
            set => _tokenBoundCidrs = value;
        }

        [Input("tokenExplicitMaxTtl")]
        public Input<int>? TokenExplicitMaxTtl { get; set; }

        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        [Input("tokenNoDefaultPolicy")]
        public Input<bool>? TokenNoDefaultPolicy { get; set; }

        [Input("tokenNumUses")]
        public Input<int>? TokenNumUses { get; set; }

        [Input("tokenPeriod")]
        public Input<int>? TokenPeriod { get; set; }

        [Input("tokenPolicies")]
        private InputList<string>? _tokenPolicies;
        public InputList<string> TokenPolicies
        {
            get => _tokenPolicies ?? (_tokenPolicies = new InputList<string>());
            set => _tokenPolicies = value;
        }

        [Input("tokenTtl")]
        public Input<int>? TokenTtl { get; set; }

        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public GetAuthBackendRoleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAuthBackendRoleResult
    {
        public readonly string? Backend;
        public readonly ImmutableArray<string> BoundCidrs;
        /// <summary>
        /// List of service account names able to access this role. If set to "*" all names are allowed, both this and bound_service_account_namespaces can not be "*".
        /// </summary>
        public readonly ImmutableArray<string> BoundServiceAccountNames;
        /// <summary>
        /// List of namespaces allowed to access this role. If set to "*" all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
        /// </summary>
        public readonly ImmutableArray<string> BoundServiceAccountNamespaces;
        public readonly int? MaxTtl;
        public readonly int? NumUses;
        public readonly int? Period;
        public readonly ImmutableArray<string> Policies;
        public readonly string RoleName;
        /// <summary>
        /// List of CIDR blocks; if set, specifies blocks of IP
        /// addresses which can authenticate successfully, and ties the resulting token to these blocks
        /// as well.
        /// </summary>
        public readonly ImmutableArray<string> TokenBoundCidrs;
        /// <summary>
        /// If set, will encode an
        /// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
        /// onto the token in number of seconds. This is a hard cap even if `token_ttl` and
        /// `token_max_ttl` would otherwise allow a renewal.
        /// </summary>
        public readonly int? TokenExplicitMaxTtl;
        /// <summary>
        /// The maximum lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        public readonly int? TokenMaxTtl;
        /// <summary>
        /// If set, the default policy will not be set on
        /// generated tokens; otherwise it will be added to the policies set in token_policies.
        /// </summary>
        public readonly bool? TokenNoDefaultPolicy;
        /// <summary>
        /// The
        /// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
        /// if any, in number of seconds to set on the token.
        /// </summary>
        public readonly int? TokenNumUses;
        public readonly int? TokenPeriod;
        /// <summary>
        /// List of policies to encode onto generated tokens. Depending
        /// on the auth method, this list may be supplemented by user/group/other values.
        /// </summary>
        public readonly ImmutableArray<string> TokenPolicies;
        /// <summary>
        /// The incremental lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        public readonly int? TokenTtl;
        /// <summary>
        /// The type of token that should be generated. Can be `service`,
        /// `batch`, or `default` to use the mount's tuned default (which unless changed will be
        /// `service` tokens). For token store roles, there are two additional possibilities:
        /// `default-service` and `default-batch` which specify the type to return unless the client
        /// requests a different type at generation time.
        /// </summary>
        public readonly string? TokenType;
        public readonly int? Ttl;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAuthBackendRoleResult(
            string? backend,
            ImmutableArray<string> boundCidrs,
            ImmutableArray<string> boundServiceAccountNames,
            ImmutableArray<string> boundServiceAccountNamespaces,
            int? maxTtl,
            int? numUses,
            int? period,
            ImmutableArray<string> policies,
            string roleName,
            ImmutableArray<string> tokenBoundCidrs,
            int? tokenExplicitMaxTtl,
            int? tokenMaxTtl,
            bool? tokenNoDefaultPolicy,
            int? tokenNumUses,
            int? tokenPeriod,
            ImmutableArray<string> tokenPolicies,
            int? tokenTtl,
            string? tokenType,
            int? ttl,
            string id)
        {
            Backend = backend;
            BoundCidrs = boundCidrs;
            BoundServiceAccountNames = boundServiceAccountNames;
            BoundServiceAccountNamespaces = boundServiceAccountNamespaces;
            MaxTtl = maxTtl;
            NumUses = numUses;
            Period = period;
            Policies = policies;
            RoleName = roleName;
            TokenBoundCidrs = tokenBoundCidrs;
            TokenExplicitMaxTtl = tokenExplicitMaxTtl;
            TokenMaxTtl = tokenMaxTtl;
            TokenNoDefaultPolicy = tokenNoDefaultPolicy;
            TokenNumUses = tokenNumUses;
            TokenPeriod = tokenPeriod;
            TokenPolicies = tokenPolicies;
            TokenTtl = tokenTtl;
            TokenType = tokenType;
            Ttl = ttl;
            Id = id;
        }
    }
}
