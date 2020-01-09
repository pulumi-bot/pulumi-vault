// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Gcp
{
    /// <summary>
    /// Provides a resource to create a role in an [GCP auth backend within Vault](https://www.vaultproject.io/docs/auth/gcp.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/gcp_auth_backend_role.html.markdown.
    /// </summary>
    public partial class AuthBackendRole : Pulumi.CustomResource
    {
        [Output("addGroupAliases")]
        public Output<bool> AddGroupAliases { get; private set; } = null!;

        /// <summary>
        /// A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.
        /// </summary>
        [Output("allowGceInference")]
        public Output<bool> AllowGceInference { get; private set; } = null!;

        /// <summary>
        /// Path to the mounted GCP auth backend
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either `bound_zones` or `bound_regions` must be set too.
        /// </summary>
        [Output("boundInstanceGroups")]
        public Output<ImmutableArray<string>> BoundInstanceGroups { get; private set; } = null!;

        /// <summary>
        /// A comma-separated list of GCP labels formatted as `"key:value"` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.
        /// </summary>
        [Output("boundLabels")]
        public Output<ImmutableArray<string>> BoundLabels { get; private set; } = null!;

        /// <summary>
        /// GCP Projects that the role exists within
        /// </summary>
        [Output("boundProjects")]
        public Output<ImmutableArray<string>> BoundProjects { get; private set; } = null!;

        /// <summary>
        /// The list of regions that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored.
        /// </summary>
        [Output("boundRegions")]
        public Output<ImmutableArray<string>> BoundRegions { get; private set; } = null!;

        /// <summary>
        /// GCP Service Accounts allowed to issue tokens under this role. (Note: **Required** if role is `iam`)
        /// </summary>
        [Output("boundServiceAccounts")]
        public Output<ImmutableArray<string>> BoundServiceAccounts { get; private set; } = null!;

        /// <summary>
        /// The list of zones that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone.
        /// </summary>
        [Output("boundZones")]
        public Output<ImmutableArray<string>> BoundZones { get; private set; } = null!;

        /// <summary>
        /// The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter `exp`. The GCE metadata tokens currently do not allow the `exp` claim to be customized.
        /// </summary>
        [Output("maxJwtExp")]
        public Output<string> MaxJwtExp { get; private set; } = null!;

        /// <summary>
        /// The maximum allowed lifetime of tokens
        /// issued using this role, provided as a number of seconds.
        /// </summary>
        [Output("maxTtl")]
        public Output<string> MaxTtl { get; private set; } = null!;

        /// <summary>
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Output("period")]
        public Output<string> Period { get; private set; } = null!;

        /// <summary>
        /// An array of strings
        /// specifying the policies to be set on tokens issued using this role.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// Name of the GCP role
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// List of CIDR blocks; if set, specifies blocks of IP
        /// addresses which can authenticate successfully, and ties the resulting token to these blocks
        /// as well.
        /// </summary>
        [Output("tokenBoundCidrs")]
        public Output<ImmutableArray<string>> TokenBoundCidrs { get; private set; } = null!;

        /// <summary>
        /// If set, will encode an
        /// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
        /// onto the token in number of seconds. This is a hard cap even if `token_ttl` and
        /// `token_max_ttl` would otherwise allow a renewal.
        /// </summary>
        [Output("tokenExplicitMaxTtl")]
        public Output<int?> TokenExplicitMaxTtl { get; private set; } = null!;

        /// <summary>
        /// The maximum lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Output("tokenMaxTtl")]
        public Output<int?> TokenMaxTtl { get; private set; } = null!;

        /// <summary>
        /// If set, the default policy will not be set on
        /// generated tokens; otherwise it will be added to the policies set in token_policies.
        /// </summary>
        [Output("tokenNoDefaultPolicy")]
        public Output<bool?> TokenNoDefaultPolicy { get; private set; } = null!;

        /// <summary>
        /// The
        /// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
        /// if any, in number of seconds to set on the token.
        /// </summary>
        [Output("tokenNumUses")]
        public Output<int?> TokenNumUses { get; private set; } = null!;

        /// <summary>
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Output("tokenPeriod")]
        public Output<int?> TokenPeriod { get; private set; } = null!;

        /// <summary>
        /// List of policies to encode onto generated tokens. Depending
        /// on the auth method, this list may be supplemented by user/group/other values.
        /// </summary>
        [Output("tokenPolicies")]
        public Output<ImmutableArray<string>> TokenPolicies { get; private set; } = null!;

        /// <summary>
        /// The incremental lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Output("tokenTtl")]
        public Output<int?> TokenTtl { get; private set; } = null!;

        /// <summary>
        /// The type of token that should be generated. Can be `service`,
        /// `batch`, or `default` to use the mount's tuned default (which unless changed will be
        /// `service` tokens). For token store roles, there are two additional possibilities:
        /// `default-service` and `default-batch` which specify the type to return unless the client
        /// requests a different type at generation time.
        /// </summary>
        [Output("tokenType")]
        public Output<string?> TokenType { get; private set; } = null!;

        /// <summary>
        /// The TTL period of tokens issued
        /// using this role, provided as a number of seconds.
        /// </summary>
        [Output("ttl")]
        public Output<string> Ttl { get; private set; } = null!;

        /// <summary>
        /// Type of GCP authentication role (either `gce` or `iam`)
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendRole(string name, AuthBackendRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:gcp/authBackendRole:AuthBackendRole", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendRole(string name, Input<string> id, AuthBackendRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:gcp/authBackendRole:AuthBackendRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthBackendRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendRole Get(string name, Input<string> id, AuthBackendRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendRole(name, id, state, options);
        }
    }

    public sealed class AuthBackendRoleArgs : Pulumi.ResourceArgs
    {
        [Input("addGroupAliases")]
        public Input<bool>? AddGroupAliases { get; set; }

        /// <summary>
        /// A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.
        /// </summary>
        [Input("allowGceInference")]
        public Input<bool>? AllowGceInference { get; set; }

        /// <summary>
        /// Path to the mounted GCP auth backend
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("boundInstanceGroups")]
        private InputList<string>? _boundInstanceGroups;

        /// <summary>
        /// The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either `bound_zones` or `bound_regions` must be set too.
        /// </summary>
        public InputList<string> BoundInstanceGroups
        {
            get => _boundInstanceGroups ?? (_boundInstanceGroups = new InputList<string>());
            set => _boundInstanceGroups = value;
        }

        [Input("boundLabels")]
        private InputList<string>? _boundLabels;

        /// <summary>
        /// A comma-separated list of GCP labels formatted as `"key:value"` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.
        /// </summary>
        public InputList<string> BoundLabels
        {
            get => _boundLabels ?? (_boundLabels = new InputList<string>());
            set => _boundLabels = value;
        }

        [Input("boundProjects")]
        private InputList<string>? _boundProjects;

        /// <summary>
        /// GCP Projects that the role exists within
        /// </summary>
        public InputList<string> BoundProjects
        {
            get => _boundProjects ?? (_boundProjects = new InputList<string>());
            set => _boundProjects = value;
        }

        [Input("boundRegions")]
        private InputList<string>? _boundRegions;

        /// <summary>
        /// The list of regions that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored.
        /// </summary>
        public InputList<string> BoundRegions
        {
            get => _boundRegions ?? (_boundRegions = new InputList<string>());
            set => _boundRegions = value;
        }

        [Input("boundServiceAccounts")]
        private InputList<string>? _boundServiceAccounts;

        /// <summary>
        /// GCP Service Accounts allowed to issue tokens under this role. (Note: **Required** if role is `iam`)
        /// </summary>
        public InputList<string> BoundServiceAccounts
        {
            get => _boundServiceAccounts ?? (_boundServiceAccounts = new InputList<string>());
            set => _boundServiceAccounts = value;
        }

        [Input("boundZones")]
        private InputList<string>? _boundZones;

        /// <summary>
        /// The list of zones that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone.
        /// </summary>
        public InputList<string> BoundZones
        {
            get => _boundZones ?? (_boundZones = new InputList<string>());
            set => _boundZones = value;
        }

        /// <summary>
        /// The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter `exp`. The GCE metadata tokens currently do not allow the `exp` claim to be customized.
        /// </summary>
        [Input("maxJwtExp")]
        public Input<string>? MaxJwtExp { get; set; }

        /// <summary>
        /// The maximum allowed lifetime of tokens
        /// issued using this role, provided as a number of seconds.
        /// </summary>
        [Input("maxTtl")]
        public Input<string>? MaxTtl { get; set; }

        /// <summary>
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// An array of strings
        /// specifying the policies to be set on tokens issued using this role.
        /// </summary>
        [Obsolete("use `token_policies` instead if you are running Vault >= 1.2")]
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// Name of the GCP role
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        [Input("tokenBoundCidrs")]
        private InputList<string>? _tokenBoundCidrs;

        /// <summary>
        /// List of CIDR blocks; if set, specifies blocks of IP
        /// addresses which can authenticate successfully, and ties the resulting token to these blocks
        /// as well.
        /// </summary>
        public InputList<string> TokenBoundCidrs
        {
            get => _tokenBoundCidrs ?? (_tokenBoundCidrs = new InputList<string>());
            set => _tokenBoundCidrs = value;
        }

        /// <summary>
        /// If set, will encode an
        /// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
        /// onto the token in number of seconds. This is a hard cap even if `token_ttl` and
        /// `token_max_ttl` would otherwise allow a renewal.
        /// </summary>
        [Input("tokenExplicitMaxTtl")]
        public Input<int>? TokenExplicitMaxTtl { get; set; }

        /// <summary>
        /// The maximum lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        /// <summary>
        /// If set, the default policy will not be set on
        /// generated tokens; otherwise it will be added to the policies set in token_policies.
        /// </summary>
        [Input("tokenNoDefaultPolicy")]
        public Input<bool>? TokenNoDefaultPolicy { get; set; }

        /// <summary>
        /// The
        /// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
        /// if any, in number of seconds to set on the token.
        /// </summary>
        [Input("tokenNumUses")]
        public Input<int>? TokenNumUses { get; set; }

        /// <summary>
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Input("tokenPeriod")]
        public Input<int>? TokenPeriod { get; set; }

        [Input("tokenPolicies")]
        private InputList<string>? _tokenPolicies;

        /// <summary>
        /// List of policies to encode onto generated tokens. Depending
        /// on the auth method, this list may be supplemented by user/group/other values.
        /// </summary>
        public InputList<string> TokenPolicies
        {
            get => _tokenPolicies ?? (_tokenPolicies = new InputList<string>());
            set => _tokenPolicies = value;
        }

        /// <summary>
        /// The incremental lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Input("tokenTtl")]
        public Input<int>? TokenTtl { get; set; }

        /// <summary>
        /// The type of token that should be generated. Can be `service`,
        /// `batch`, or `default` to use the mount's tuned default (which unless changed will be
        /// `service` tokens). For token store roles, there are two additional possibilities:
        /// `default-service` and `default-batch` which specify the type to return unless the client
        /// requests a different type at generation time.
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// The TTL period of tokens issued
        /// using this role, provided as a number of seconds.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// Type of GCP authentication role (either `gce` or `iam`)
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AuthBackendRoleArgs()
        {
        }
    }

    public sealed class AuthBackendRoleState : Pulumi.ResourceArgs
    {
        [Input("addGroupAliases")]
        public Input<bool>? AddGroupAliases { get; set; }

        /// <summary>
        /// A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.
        /// </summary>
        [Input("allowGceInference")]
        public Input<bool>? AllowGceInference { get; set; }

        /// <summary>
        /// Path to the mounted GCP auth backend
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("boundInstanceGroups")]
        private InputList<string>? _boundInstanceGroups;

        /// <summary>
        /// The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either `bound_zones` or `bound_regions` must be set too.
        /// </summary>
        public InputList<string> BoundInstanceGroups
        {
            get => _boundInstanceGroups ?? (_boundInstanceGroups = new InputList<string>());
            set => _boundInstanceGroups = value;
        }

        [Input("boundLabels")]
        private InputList<string>? _boundLabels;

        /// <summary>
        /// A comma-separated list of GCP labels formatted as `"key:value"` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.
        /// </summary>
        public InputList<string> BoundLabels
        {
            get => _boundLabels ?? (_boundLabels = new InputList<string>());
            set => _boundLabels = value;
        }

        [Input("boundProjects")]
        private InputList<string>? _boundProjects;

        /// <summary>
        /// GCP Projects that the role exists within
        /// </summary>
        public InputList<string> BoundProjects
        {
            get => _boundProjects ?? (_boundProjects = new InputList<string>());
            set => _boundProjects = value;
        }

        [Input("boundRegions")]
        private InputList<string>? _boundRegions;

        /// <summary>
        /// The list of regions that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored.
        /// </summary>
        public InputList<string> BoundRegions
        {
            get => _boundRegions ?? (_boundRegions = new InputList<string>());
            set => _boundRegions = value;
        }

        [Input("boundServiceAccounts")]
        private InputList<string>? _boundServiceAccounts;

        /// <summary>
        /// GCP Service Accounts allowed to issue tokens under this role. (Note: **Required** if role is `iam`)
        /// </summary>
        public InputList<string> BoundServiceAccounts
        {
            get => _boundServiceAccounts ?? (_boundServiceAccounts = new InputList<string>());
            set => _boundServiceAccounts = value;
        }

        [Input("boundZones")]
        private InputList<string>? _boundZones;

        /// <summary>
        /// The list of zones that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone.
        /// </summary>
        public InputList<string> BoundZones
        {
            get => _boundZones ?? (_boundZones = new InputList<string>());
            set => _boundZones = value;
        }

        /// <summary>
        /// The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter `exp`. The GCE metadata tokens currently do not allow the `exp` claim to be customized.
        /// </summary>
        [Input("maxJwtExp")]
        public Input<string>? MaxJwtExp { get; set; }

        /// <summary>
        /// The maximum allowed lifetime of tokens
        /// issued using this role, provided as a number of seconds.
        /// </summary>
        [Input("maxTtl")]
        public Input<string>? MaxTtl { get; set; }

        /// <summary>
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// An array of strings
        /// specifying the policies to be set on tokens issued using this role.
        /// </summary>
        [Obsolete("use `token_policies` instead if you are running Vault >= 1.2")]
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// Name of the GCP role
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("tokenBoundCidrs")]
        private InputList<string>? _tokenBoundCidrs;

        /// <summary>
        /// List of CIDR blocks; if set, specifies blocks of IP
        /// addresses which can authenticate successfully, and ties the resulting token to these blocks
        /// as well.
        /// </summary>
        public InputList<string> TokenBoundCidrs
        {
            get => _tokenBoundCidrs ?? (_tokenBoundCidrs = new InputList<string>());
            set => _tokenBoundCidrs = value;
        }

        /// <summary>
        /// If set, will encode an
        /// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
        /// onto the token in number of seconds. This is a hard cap even if `token_ttl` and
        /// `token_max_ttl` would otherwise allow a renewal.
        /// </summary>
        [Input("tokenExplicitMaxTtl")]
        public Input<int>? TokenExplicitMaxTtl { get; set; }

        /// <summary>
        /// The maximum lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        /// <summary>
        /// If set, the default policy will not be set on
        /// generated tokens; otherwise it will be added to the policies set in token_policies.
        /// </summary>
        [Input("tokenNoDefaultPolicy")]
        public Input<bool>? TokenNoDefaultPolicy { get; set; }

        /// <summary>
        /// The
        /// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
        /// if any, in number of seconds to set on the token.
        /// </summary>
        [Input("tokenNumUses")]
        public Input<int>? TokenNumUses { get; set; }

        /// <summary>
        /// If set, indicates that the
        /// token generated using this role should never expire. The token should be renewed within the
        /// duration specified by this value. At each renewal, the token's TTL will be set to the
        /// value of this field. Specified in seconds.
        /// </summary>
        [Input("tokenPeriod")]
        public Input<int>? TokenPeriod { get; set; }

        [Input("tokenPolicies")]
        private InputList<string>? _tokenPolicies;

        /// <summary>
        /// List of policies to encode onto generated tokens. Depending
        /// on the auth method, this list may be supplemented by user/group/other values.
        /// </summary>
        public InputList<string> TokenPolicies
        {
            get => _tokenPolicies ?? (_tokenPolicies = new InputList<string>());
            set => _tokenPolicies = value;
        }

        /// <summary>
        /// The incremental lifetime for generated tokens in number of seconds.
        /// Its current value will be referenced at renewal time.
        /// </summary>
        [Input("tokenTtl")]
        public Input<int>? TokenTtl { get; set; }

        /// <summary>
        /// The type of token that should be generated. Can be `service`,
        /// `batch`, or `default` to use the mount's tuned default (which unless changed will be
        /// `service` tokens). For token store roles, there are two additional possibilities:
        /// `default-service` and `default-batch` which specify the type to return unless the client
        /// requests a different type at generation time.
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// The TTL period of tokens issued
        /// using this role, provided as a number of seconds.
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// Type of GCP authentication role (either `gce` or `iam`)
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AuthBackendRoleState()
        {
        }
    }
}
