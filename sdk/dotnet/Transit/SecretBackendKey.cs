// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Transit
{
    /// <summary>
    /// Creates an Encryption Keyring on a Transit Secret Backend for Vault.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/transit_secret_backend_key.html.md.
    /// </summary>
    public partial class SecretBackendKey : Pulumi.CustomResource
    {
        /// <summary>
        /// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
        /// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api/secret/transit/index.html#backup-key)
        /// </summary>
        [Output("allowPlaintextBackup")]
        public Output<bool?> AllowPlaintextBackup { get; private set; } = null!;

        /// <summary>
        /// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
        /// </summary>
        [Output("convergentEncryption")]
        public Output<bool?> ConvergentEncryption { get; private set; } = null!;

        /// <summary>
        /// Specifies if the key is allowed to be deleted.
        /// </summary>
        [Output("deletionAllowed")]
        public Output<bool?> DeletionAllowed { get; private set; } = null!;

        /// <summary>
        /// Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
        /// </summary>
        [Output("derived")]
        public Output<bool?> Derived { get; private set; } = null!;

        /// <summary>
        /// Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
        /// </summary>
        [Output("exportable")]
        public Output<bool?> Exportable { get; private set; } = null!;

        /// <summary>
        /// List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
        /// * for key types `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
        /// * for key types `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`, each key version will be a map of the following:
        /// </summary>
        [Output("keys")]
        public Output<ImmutableArray<ImmutableDictionary<string, object>>> Keys { get; private set; } = null!;

        /// <summary>
        /// Latest key version available. This value is 1-indexed, so if `latest_version` is `1`, then the key's information can be referenced from `keys` by selecting element `0`
        /// </summary>
        [Output("latestVersion")]
        public Output<int> LatestVersion { get; private set; } = null!;

        /// <summary>
        /// Minimum key version available for use. If keys have been archived by increasing `min_decryption_version`, this attribute will reflect that change.
        /// </summary>
        [Output("minAvailableVersion")]
        public Output<int> MinAvailableVersion { get; private set; } = null!;

        /// <summary>
        /// Minimum key version to use for decryption.
        /// </summary>
        [Output("minDecryptionVersion")]
        public Output<int?> MinDecryptionVersion { get; private set; } = null!;

        /// <summary>
        /// Minimum key version to use for encryption
        /// </summary>
        [Output("minEncryptionVersion")]
        public Output<int?> MinEncryptionVersion { get; private set; } = null!;

        /// <summary>
        /// The name to identify this key within the backend. Must be unique within the backend.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether or not the key supports decryption, based on key type.
        /// </summary>
        [Output("supportsDecryption")]
        public Output<bool> SupportsDecryption { get; private set; } = null!;

        /// <summary>
        /// Whether or not the key supports derivation, based on key type.
        /// </summary>
        [Output("supportsDerivation")]
        public Output<bool> SupportsDerivation { get; private set; } = null!;

        /// <summary>
        /// Whether or not the key supports encryption, based on key type.
        /// </summary>
        [Output("supportsEncryption")]
        public Output<bool> SupportsEncryption { get; private set; } = null!;

        /// <summary>
        /// Whether or not the key supports signing, based on key type.
        /// </summary>
        [Output("supportsSigning")]
        public Output<bool> SupportsSigning { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of key to create. The currently-supported types are: `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`. 
        /// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit/index.html#key-types)
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendKey(string name, SecretBackendKeyArgs args, CustomResourceOptions? options = null)
            : base("vault:transit/secretBackendKey:SecretBackendKey", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendKey(string name, Input<string> id, SecretBackendKeyState? state = null, CustomResourceOptions? options = null)
            : base("vault:transit/secretBackendKey:SecretBackendKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendKey Get(string name, Input<string> id, SecretBackendKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendKey(name, id, state, options);
        }
    }

    public sealed class SecretBackendKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
        /// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api/secret/transit/index.html#backup-key)
        /// </summary>
        [Input("allowPlaintextBackup")]
        public Input<bool>? AllowPlaintextBackup { get; set; }

        /// <summary>
        /// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
        /// </summary>
        [Input("convergentEncryption")]
        public Input<bool>? ConvergentEncryption { get; set; }

        /// <summary>
        /// Specifies if the key is allowed to be deleted.
        /// </summary>
        [Input("deletionAllowed")]
        public Input<bool>? DeletionAllowed { get; set; }

        /// <summary>
        /// Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
        /// </summary>
        [Input("derived")]
        public Input<bool>? Derived { get; set; }

        /// <summary>
        /// Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
        /// </summary>
        [Input("exportable")]
        public Input<bool>? Exportable { get; set; }

        /// <summary>
        /// Minimum key version to use for decryption.
        /// </summary>
        [Input("minDecryptionVersion")]
        public Input<int>? MinDecryptionVersion { get; set; }

        /// <summary>
        /// Minimum key version to use for encryption
        /// </summary>
        [Input("minEncryptionVersion")]
        public Input<int>? MinEncryptionVersion { get; set; }

        /// <summary>
        /// The name to identify this key within the backend. Must be unique within the backend.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the type of key to create. The currently-supported types are: `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`. 
        /// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit/index.html#key-types)
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SecretBackendKeyArgs()
        {
        }
    }

    public sealed class SecretBackendKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
        /// * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api/secret/transit/index.html#backup-key)
        /// </summary>
        [Input("allowPlaintextBackup")]
        public Input<bool>? AllowPlaintextBackup { get; set; }

        /// <summary>
        /// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
        /// </summary>
        [Input("convergentEncryption")]
        public Input<bool>? ConvergentEncryption { get; set; }

        /// <summary>
        /// Specifies if the key is allowed to be deleted.
        /// </summary>
        [Input("deletionAllowed")]
        public Input<bool>? DeletionAllowed { get; set; }

        /// <summary>
        /// Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
        /// </summary>
        [Input("derived")]
        public Input<bool>? Derived { get; set; }

        /// <summary>
        /// Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
        /// </summary>
        [Input("exportable")]
        public Input<bool>? Exportable { get; set; }

        [Input("keys")]
        private InputList<ImmutableDictionary<string, object>>? _keys;

        /// <summary>
        /// List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
        /// * for key types `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
        /// * for key types `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`, each key version will be a map of the following:
        /// </summary>
        public InputList<ImmutableDictionary<string, object>> Keys
        {
            get => _keys ?? (_keys = new InputList<ImmutableDictionary<string, object>>());
            set => _keys = value;
        }

        /// <summary>
        /// Latest key version available. This value is 1-indexed, so if `latest_version` is `1`, then the key's information can be referenced from `keys` by selecting element `0`
        /// </summary>
        [Input("latestVersion")]
        public Input<int>? LatestVersion { get; set; }

        /// <summary>
        /// Minimum key version available for use. If keys have been archived by increasing `min_decryption_version`, this attribute will reflect that change.
        /// </summary>
        [Input("minAvailableVersion")]
        public Input<int>? MinAvailableVersion { get; set; }

        /// <summary>
        /// Minimum key version to use for decryption.
        /// </summary>
        [Input("minDecryptionVersion")]
        public Input<int>? MinDecryptionVersion { get; set; }

        /// <summary>
        /// Minimum key version to use for encryption
        /// </summary>
        [Input("minEncryptionVersion")]
        public Input<int>? MinEncryptionVersion { get; set; }

        /// <summary>
        /// The name to identify this key within the backend. Must be unique within the backend.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether or not the key supports decryption, based on key type.
        /// </summary>
        [Input("supportsDecryption")]
        public Input<bool>? SupportsDecryption { get; set; }

        /// <summary>
        /// Whether or not the key supports derivation, based on key type.
        /// </summary>
        [Input("supportsDerivation")]
        public Input<bool>? SupportsDerivation { get; set; }

        /// <summary>
        /// Whether or not the key supports encryption, based on key type.
        /// </summary>
        [Input("supportsEncryption")]
        public Input<bool>? SupportsEncryption { get; set; }

        /// <summary>
        /// Whether or not the key supports signing, based on key type.
        /// </summary>
        [Input("supportsSigning")]
        public Input<bool>? SupportsSigning { get; set; }

        /// <summary>
        /// Specifies the type of key to create. The currently-supported types are: `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`. 
        /// * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit/index.html#key-types)
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SecretBackendKeyState()
        {
        }
    }
}
