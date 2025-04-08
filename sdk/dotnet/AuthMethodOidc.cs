// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Boundary
{
    [BoundaryResourceType("boundary:index/authMethodOidc:AuthMethodOidc")]
    public partial class AuthMethodOidc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Account claim maps for the to_claim of sub.
        /// </summary>
        [Output("accountClaimMaps")]
        public Output<ImmutableArray<string>> AccountClaimMaps { get; private set; } = null!;

        /// <summary>
        /// Audiences for which the provider responses will be allowed
        /// </summary>
        [Output("allowedAudiences")]
        public Output<ImmutableArray<string>> AllowedAudiences { get; private set; } = null!;

        /// <summary>
        /// The API prefix to use when generating callback URLs for the provider. Should be set to an address at which the provider
        /// can reach back to the controller.
        /// </summary>
        [Output("apiUrlPrefix")]
        public Output<string?> ApiUrlPrefix { get; private set; } = null!;

        /// <summary>
        /// The URL that should be provided to the IdP for callbacks.
        /// </summary>
        [Output("callbackUrl")]
        public Output<string> CallbackUrl { get; private set; } = null!;

        /// <summary>
        /// Claims scopes.
        /// </summary>
        [Output("claimsScopes")]
        public Output<ImmutableArray<string>> ClaimsScopes { get; private set; } = null!;

        /// <summary>
        /// The client ID assigned to this auth method from the provider.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The secret key assigned to this auth method from the provider. Once set, only the hash will be kept and the original
        /// value can be removed from configuration.
        /// </summary>
        [Output("clientSecret")]
        public Output<string?> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// The HMAC of the client secret returned by the Boundary controller, which is used for comparison after initial setting of
        /// the value.
        /// </summary>
        [Output("clientSecretHmac")]
        public Output<string> ClientSecretHmac { get; private set; } = null!;

        /// <summary>
        /// The auth method description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Disables validation logic ensuring that the OIDC provider's information from its discovery endpoint matches the
        /// information here. The validation is only performed at create or update time.
        /// </summary>
        [Output("disableDiscoveredConfigValidation")]
        public Output<bool?> DisableDiscoveredConfigValidation { get; private set; } = null!;

        /// <summary>
        /// A list of CA certificates to trust when validating the IdP's token signatures.
        /// </summary>
        [Output("idpCaCerts")]
        public Output<ImmutableArray<string>> IdpCaCerts { get; private set; } = null!;

        /// <summary>
        /// When true, makes this auth method the primary auth method for the scope in which it resides. The primary auth method for
        /// a scope means the user will be automatically created when they login using an OIDC account.
        /// </summary>
        [Output("isPrimaryForScope")]
        public Output<bool?> IsPrimaryForScope { get; private set; } = null!;

        /// <summary>
        /// The issuer corresponding to the provider, which must match the issuer field in generated tokens.
        /// </summary>
        [Output("issuer")]
        public Output<string?> Issuer { get; private set; } = null!;

        [Output("maxAge")]
        public Output<int?> MaxAge { get; private set; } = null!;

        /// <summary>
        /// The auth method name. Defaults to the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The prompts passed to the identity provider to determine whether to prompt the end-user for reauthentication, account
        /// selection or consent. Please note the values passed are case-sensitive. The valid values are: `none`, `login`, `consent`
        /// and `select_account`.
        /// </summary>
        [Output("prompts")]
        public Output<ImmutableArray<string>> Prompts { get; private set; } = null!;

        /// <summary>
        /// The scope ID.
        /// </summary>
        [Output("scopeId")]
        public Output<string> ScopeId { get; private set; } = null!;

        /// <summary>
        /// Allowed signing algorithms for the provider's issued tokens.
        /// </summary>
        [Output("signingAlgorithms")]
        public Output<ImmutableArray<string>> SigningAlgorithms { get; private set; } = null!;

        /// <summary>
        /// Can be one of 'inactive', 'active-private', or 'active-public'. Currently automatically set to active-public.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The type of auth method; hardcoded.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AuthMethodOidc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthMethodOidc(string name, AuthMethodOidcArgs args, CustomResourceOptions? options = null)
            : base("boundary:index/authMethodOidc:AuthMethodOidc", name, args ?? new AuthMethodOidcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthMethodOidc(string name, Input<string> id, AuthMethodOidcState? state = null, CustomResourceOptions? options = null)
            : base("boundary:index/authMethodOidc:AuthMethodOidc", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "clientSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthMethodOidc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthMethodOidc Get(string name, Input<string> id, AuthMethodOidcState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthMethodOidc(name, id, state, options);
        }
    }

    public sealed class AuthMethodOidcArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountClaimMaps")]
        private InputList<string>? _accountClaimMaps;

        /// <summary>
        /// Account claim maps for the to_claim of sub.
        /// </summary>
        public InputList<string> AccountClaimMaps
        {
            get => _accountClaimMaps ?? (_accountClaimMaps = new InputList<string>());
            set => _accountClaimMaps = value;
        }

        [Input("allowedAudiences")]
        private InputList<string>? _allowedAudiences;

        /// <summary>
        /// Audiences for which the provider responses will be allowed
        /// </summary>
        public InputList<string> AllowedAudiences
        {
            get => _allowedAudiences ?? (_allowedAudiences = new InputList<string>());
            set => _allowedAudiences = value;
        }

        /// <summary>
        /// The API prefix to use when generating callback URLs for the provider. Should be set to an address at which the provider
        /// can reach back to the controller.
        /// </summary>
        [Input("apiUrlPrefix")]
        public Input<string>? ApiUrlPrefix { get; set; }

        /// <summary>
        /// The URL that should be provided to the IdP for callbacks.
        /// </summary>
        [Input("callbackUrl")]
        public Input<string>? CallbackUrl { get; set; }

        [Input("claimsScopes")]
        private InputList<string>? _claimsScopes;

        /// <summary>
        /// Claims scopes.
        /// </summary>
        public InputList<string> ClaimsScopes
        {
            get => _claimsScopes ?? (_claimsScopes = new InputList<string>());
            set => _claimsScopes = value;
        }

        /// <summary>
        /// The client ID assigned to this auth method from the provider.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// The secret key assigned to this auth method from the provider. Once set, only the hash will be kept and the original
        /// value can be removed from configuration.
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The HMAC of the client secret returned by the Boundary controller, which is used for comparison after initial setting of
        /// the value.
        /// </summary>
        [Input("clientSecretHmac")]
        public Input<string>? ClientSecretHmac { get; set; }

        /// <summary>
        /// The auth method description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Disables validation logic ensuring that the OIDC provider's information from its discovery endpoint matches the
        /// information here. The validation is only performed at create or update time.
        /// </summary>
        [Input("disableDiscoveredConfigValidation")]
        public Input<bool>? DisableDiscoveredConfigValidation { get; set; }

        [Input("idpCaCerts")]
        private InputList<string>? _idpCaCerts;

        /// <summary>
        /// A list of CA certificates to trust when validating the IdP's token signatures.
        /// </summary>
        public InputList<string> IdpCaCerts
        {
            get => _idpCaCerts ?? (_idpCaCerts = new InputList<string>());
            set => _idpCaCerts = value;
        }

        /// <summary>
        /// When true, makes this auth method the primary auth method for the scope in which it resides. The primary auth method for
        /// a scope means the user will be automatically created when they login using an OIDC account.
        /// </summary>
        [Input("isPrimaryForScope")]
        public Input<bool>? IsPrimaryForScope { get; set; }

        /// <summary>
        /// The issuer corresponding to the provider, which must match the issuer field in generated tokens.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        [Input("maxAge")]
        public Input<int>? MaxAge { get; set; }

        /// <summary>
        /// The auth method name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("prompts")]
        private InputList<string>? _prompts;

        /// <summary>
        /// The prompts passed to the identity provider to determine whether to prompt the end-user for reauthentication, account
        /// selection or consent. Please note the values passed are case-sensitive. The valid values are: `none`, `login`, `consent`
        /// and `select_account`.
        /// </summary>
        public InputList<string> Prompts
        {
            get => _prompts ?? (_prompts = new InputList<string>());
            set => _prompts = value;
        }

        /// <summary>
        /// The scope ID.
        /// </summary>
        [Input("scopeId", required: true)]
        public Input<string> ScopeId { get; set; } = null!;

        [Input("signingAlgorithms")]
        private InputList<string>? _signingAlgorithms;

        /// <summary>
        /// Allowed signing algorithms for the provider's issued tokens.
        /// </summary>
        public InputList<string> SigningAlgorithms
        {
            get => _signingAlgorithms ?? (_signingAlgorithms = new InputList<string>());
            set => _signingAlgorithms = value;
        }

        /// <summary>
        /// Can be one of 'inactive', 'active-private', or 'active-public'. Currently automatically set to active-public.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The type of auth method; hardcoded.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AuthMethodOidcArgs()
        {
        }
        public static new AuthMethodOidcArgs Empty => new AuthMethodOidcArgs();
    }

    public sealed class AuthMethodOidcState : global::Pulumi.ResourceArgs
    {
        [Input("accountClaimMaps")]
        private InputList<string>? _accountClaimMaps;

        /// <summary>
        /// Account claim maps for the to_claim of sub.
        /// </summary>
        public InputList<string> AccountClaimMaps
        {
            get => _accountClaimMaps ?? (_accountClaimMaps = new InputList<string>());
            set => _accountClaimMaps = value;
        }

        [Input("allowedAudiences")]
        private InputList<string>? _allowedAudiences;

        /// <summary>
        /// Audiences for which the provider responses will be allowed
        /// </summary>
        public InputList<string> AllowedAudiences
        {
            get => _allowedAudiences ?? (_allowedAudiences = new InputList<string>());
            set => _allowedAudiences = value;
        }

        /// <summary>
        /// The API prefix to use when generating callback URLs for the provider. Should be set to an address at which the provider
        /// can reach back to the controller.
        /// </summary>
        [Input("apiUrlPrefix")]
        public Input<string>? ApiUrlPrefix { get; set; }

        /// <summary>
        /// The URL that should be provided to the IdP for callbacks.
        /// </summary>
        [Input("callbackUrl")]
        public Input<string>? CallbackUrl { get; set; }

        [Input("claimsScopes")]
        private InputList<string>? _claimsScopes;

        /// <summary>
        /// Claims scopes.
        /// </summary>
        public InputList<string> ClaimsScopes
        {
            get => _claimsScopes ?? (_claimsScopes = new InputList<string>());
            set => _claimsScopes = value;
        }

        /// <summary>
        /// The client ID assigned to this auth method from the provider.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// The secret key assigned to this auth method from the provider. Once set, only the hash will be kept and the original
        /// value can be removed from configuration.
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The HMAC of the client secret returned by the Boundary controller, which is used for comparison after initial setting of
        /// the value.
        /// </summary>
        [Input("clientSecretHmac")]
        public Input<string>? ClientSecretHmac { get; set; }

        /// <summary>
        /// The auth method description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Disables validation logic ensuring that the OIDC provider's information from its discovery endpoint matches the
        /// information here. The validation is only performed at create or update time.
        /// </summary>
        [Input("disableDiscoveredConfigValidation")]
        public Input<bool>? DisableDiscoveredConfigValidation { get; set; }

        [Input("idpCaCerts")]
        private InputList<string>? _idpCaCerts;

        /// <summary>
        /// A list of CA certificates to trust when validating the IdP's token signatures.
        /// </summary>
        public InputList<string> IdpCaCerts
        {
            get => _idpCaCerts ?? (_idpCaCerts = new InputList<string>());
            set => _idpCaCerts = value;
        }

        /// <summary>
        /// When true, makes this auth method the primary auth method for the scope in which it resides. The primary auth method for
        /// a scope means the user will be automatically created when they login using an OIDC account.
        /// </summary>
        [Input("isPrimaryForScope")]
        public Input<bool>? IsPrimaryForScope { get; set; }

        /// <summary>
        /// The issuer corresponding to the provider, which must match the issuer field in generated tokens.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        [Input("maxAge")]
        public Input<int>? MaxAge { get; set; }

        /// <summary>
        /// The auth method name. Defaults to the resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("prompts")]
        private InputList<string>? _prompts;

        /// <summary>
        /// The prompts passed to the identity provider to determine whether to prompt the end-user for reauthentication, account
        /// selection or consent. Please note the values passed are case-sensitive. The valid values are: `none`, `login`, `consent`
        /// and `select_account`.
        /// </summary>
        public InputList<string> Prompts
        {
            get => _prompts ?? (_prompts = new InputList<string>());
            set => _prompts = value;
        }

        /// <summary>
        /// The scope ID.
        /// </summary>
        [Input("scopeId")]
        public Input<string>? ScopeId { get; set; }

        [Input("signingAlgorithms")]
        private InputList<string>? _signingAlgorithms;

        /// <summary>
        /// Allowed signing algorithms for the provider's issued tokens.
        /// </summary>
        public InputList<string> SigningAlgorithms
        {
            get => _signingAlgorithms ?? (_signingAlgorithms = new InputList<string>());
            set => _signingAlgorithms = value;
        }

        /// <summary>
        /// Can be one of 'inactive', 'active-private', or 'active-public'. Currently automatically set to active-public.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The type of auth method; hardcoded.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AuthMethodOidcState()
        {
        }
        public static new AuthMethodOidcState Empty => new AuthMethodOidcState();
    }
}
