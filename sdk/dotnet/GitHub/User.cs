// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.GitHub
{
    /// <summary>
    /// Manages policy mappings for Github Users authenticated via Github. See the [Vault
    /// documentation](https://www.vaultproject.io/docs/auth/github.html) for more
    /// information.
    /// </summary>
    public partial class User : Pulumi.CustomResource
    {
        /// <summary>
        /// Path where the github auth backend is mounted. Defaults to `github`
        /// if not specified.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// An array of strings specifying the policies to be set on tokens issued
        /// using this role.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// Specifies the blocks of IP addresses which are allowed to use the generated token
        /// </summary>
        [Output("tokenBoundCidrs")]
        public Output<ImmutableArray<string>> TokenBoundCidrs { get; private set; } = null!;

        /// <summary>
        /// Generated Token's Explicit Maximum TTL in seconds
        /// </summary>
        [Output("tokenExplicitMaxTtl")]
        public Output<int?> TokenExplicitMaxTtl { get; private set; } = null!;

        /// <summary>
        /// The maximum lifetime of the generated token
        /// </summary>
        [Output("tokenMaxTtl")]
        public Output<int?> TokenMaxTtl { get; private set; } = null!;

        /// <summary>
        /// If true, the 'default' policy will not automatically be added to generated tokens
        /// </summary>
        [Output("tokenNoDefaultPolicy")]
        public Output<bool?> TokenNoDefaultPolicy { get; private set; } = null!;

        /// <summary>
        /// The maximum number of times a token may be used, a value of zero means unlimited
        /// </summary>
        [Output("tokenNumUses")]
        public Output<int?> TokenNumUses { get; private set; } = null!;

        /// <summary>
        /// Generated Token's Period
        /// </summary>
        [Output("tokenPeriod")]
        public Output<int?> TokenPeriod { get; private set; } = null!;

        /// <summary>
        /// Generated Token's Policies
        /// </summary>
        [Output("tokenPolicies")]
        public Output<ImmutableArray<string>> TokenPolicies { get; private set; } = null!;

        /// <summary>
        /// The initial ttl of the token to generate in seconds
        /// </summary>
        [Output("tokenTtl")]
        public Output<int?> TokenTtl { get; private set; } = null!;

        /// <summary>
        /// The type of token to generate, service or batch
        /// </summary>
        [Output("tokenType")]
        public Output<string?> TokenType { get; private set; } = null!;

        /// <summary>
        /// GitHub user name.
        /// </summary>
        [Output("user")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("vault:github/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("vault:github/user:User", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path where the github auth backend is mounted. Defaults to `github`
        /// if not specified.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// An array of strings specifying the policies to be set on tokens issued
        /// using this role.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        [Input("tokenBoundCidrs")]
        private InputList<string>? _tokenBoundCidrs;

        /// <summary>
        /// Specifies the blocks of IP addresses which are allowed to use the generated token
        /// </summary>
        [Obsolete(@"This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")]
        public InputList<string> TokenBoundCidrs
        {
            get => _tokenBoundCidrs ?? (_tokenBoundCidrs = new InputList<string>());
            set => _tokenBoundCidrs = value;
        }

        /// <summary>
        /// Generated Token's Explicit Maximum TTL in seconds
        /// </summary>
        [Input("tokenExplicitMaxTtl")]
        public Input<int>? TokenExplicitMaxTtl { get; set; }

        /// <summary>
        /// The maximum lifetime of the generated token
        /// </summary>
        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        /// <summary>
        /// If true, the 'default' policy will not automatically be added to generated tokens
        /// </summary>
        [Input("tokenNoDefaultPolicy")]
        public Input<bool>? TokenNoDefaultPolicy { get; set; }

        /// <summary>
        /// The maximum number of times a token may be used, a value of zero means unlimited
        /// </summary>
        [Input("tokenNumUses")]
        public Input<int>? TokenNumUses { get; set; }

        /// <summary>
        /// Generated Token's Period
        /// </summary>
        [Input("tokenPeriod")]
        public Input<int>? TokenPeriod { get; set; }

        [Input("tokenPolicies")]
        private InputList<string>? _tokenPolicies;

        /// <summary>
        /// Generated Token's Policies
        /// </summary>
        [Obsolete(@"This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")]
        public InputList<string> TokenPolicies
        {
            get => _tokenPolicies ?? (_tokenPolicies = new InputList<string>());
            set => _tokenPolicies = value;
        }

        /// <summary>
        /// The initial ttl of the token to generate in seconds
        /// </summary>
        [Input("tokenTtl")]
        public Input<int>? TokenTtl { get; set; }

        /// <summary>
        /// The type of token to generate, service or batch
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// GitHub user name.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public UserArgs()
        {
        }
    }

    public sealed class UserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path where the github auth backend is mounted. Defaults to `github`
        /// if not specified.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// An array of strings specifying the policies to be set on tokens issued
        /// using this role.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        [Input("tokenBoundCidrs")]
        private InputList<string>? _tokenBoundCidrs;

        /// <summary>
        /// Specifies the blocks of IP addresses which are allowed to use the generated token
        /// </summary>
        [Obsolete(@"This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")]
        public InputList<string> TokenBoundCidrs
        {
            get => _tokenBoundCidrs ?? (_tokenBoundCidrs = new InputList<string>());
            set => _tokenBoundCidrs = value;
        }

        /// <summary>
        /// Generated Token's Explicit Maximum TTL in seconds
        /// </summary>
        [Input("tokenExplicitMaxTtl")]
        public Input<int>? TokenExplicitMaxTtl { get; set; }

        /// <summary>
        /// The maximum lifetime of the generated token
        /// </summary>
        [Input("tokenMaxTtl")]
        public Input<int>? TokenMaxTtl { get; set; }

        /// <summary>
        /// If true, the 'default' policy will not automatically be added to generated tokens
        /// </summary>
        [Input("tokenNoDefaultPolicy")]
        public Input<bool>? TokenNoDefaultPolicy { get; set; }

        /// <summary>
        /// The maximum number of times a token may be used, a value of zero means unlimited
        /// </summary>
        [Input("tokenNumUses")]
        public Input<int>? TokenNumUses { get; set; }

        /// <summary>
        /// Generated Token's Period
        /// </summary>
        [Input("tokenPeriod")]
        public Input<int>? TokenPeriod { get; set; }

        [Input("tokenPolicies")]
        private InputList<string>? _tokenPolicies;

        /// <summary>
        /// Generated Token's Policies
        /// </summary>
        [Obsolete(@"This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")]
        public InputList<string> TokenPolicies
        {
            get => _tokenPolicies ?? (_tokenPolicies = new InputList<string>());
            set => _tokenPolicies = value;
        }

        /// <summary>
        /// The initial ttl of the token to generate in seconds
        /// </summary>
        [Input("tokenTtl")]
        public Input<int>? TokenTtl { get; set; }

        /// <summary>
        /// The type of token to generate, service or batch
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        /// <summary>
        /// GitHub user name.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public UserState()
        {
        }
    }
}
