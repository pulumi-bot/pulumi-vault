// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Outputs
{

    [OutputType]
    public sealed class SecretBackendConnectionElasticsearch
    {
        /// <summary>
        /// The password to authenticate with.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The URL for Elasticsearch's API. https requires certificate
        /// by trusted CA if used.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// The username to authenticate with.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private SecretBackendConnectionElasticsearch(
            string password,

            string url,

            string username)
        {
            Password = password;
            Url = url;
            Username = username;
        }
    }
}
