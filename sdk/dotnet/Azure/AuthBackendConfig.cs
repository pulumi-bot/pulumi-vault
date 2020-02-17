// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Azure
{
    public partial class AuthBackendConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// Unique name of the auth backend to configure.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources
        /// are required.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The client secret for credentials to query the Azure APIs
        /// </summary>
        [Output("clientSecret")]
        public Output<string?> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
        /// AzureGermanCloud.
        /// </summary>
        [Output("environment")]
        public Output<string?> Environment { get; private set; } = null!;

        /// <summary>
        /// The configured URL for the application registered in Azure Active Directory.
        /// </summary>
        [Output("resource")]
        public Output<string> Resource { get; private set; } = null!;

        /// <summary>
        /// The tenant id for the Azure Active Directory organization.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendConfig(string name, AuthBackendConfigArgs args, CustomResourceOptions? options = null)
            : base("vault:azure/authBackendConfig:AuthBackendConfig", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendConfig(string name, Input<string> id, AuthBackendConfigState? state = null, CustomResourceOptions? options = null)
            : base("vault:azure/authBackendConfig:AuthBackendConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthBackendConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendConfig Get(string name, Input<string> id, AuthBackendConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendConfig(name, id, state, options);
        }
    }

    public sealed class AuthBackendConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique name of the auth backend to configure.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources
        /// are required.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client secret for credentials to query the Azure APIs
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
        /// AzureGermanCloud.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The configured URL for the application registered in Azure Active Directory.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// The tenant id for the Azure Active Directory organization.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public AuthBackendConfigArgs()
        {
        }
    }

    public sealed class AuthBackendConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique name of the auth backend to configure.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources
        /// are required.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client secret for credentials to query the Azure APIs
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
        /// AzureGermanCloud.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The configured URL for the application registered in Azure Active Directory.
        /// </summary>
        [Input("resource")]
        public Input<string>? Resource { get; set; }

        /// <summary>
        /// The tenant id for the Azure Active Directory organization.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public AuthBackendConfigState()
        {
        }
    }
}
