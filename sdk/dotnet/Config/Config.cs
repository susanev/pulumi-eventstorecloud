// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.EventStoreCloud
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly Pulumi.Config __config = new Pulumi.Config("eventstorecloud");

        private static readonly __Value<string?> _clientId = new __Value<string?>(() => __config.Get("clientId"));
        public static string? ClientId
        {
            get => _clientId.Get();
            set => _clientId.Set(value);
        }

        private static readonly __Value<string?> _identityProviderUrl = new __Value<string?>(() => __config.Get("identityProviderUrl"));
        public static string? IdentityProviderUrl
        {
            get => _identityProviderUrl.Get();
            set => _identityProviderUrl.Set(value);
        }

        private static readonly __Value<string?> _organizationId = new __Value<string?>(() => __config.Get("organizationId"));
        public static string? OrganizationId
        {
            get => _organizationId.Get();
            set => _organizationId.Set(value);
        }

        private static readonly __Value<string?> _token = new __Value<string?>(() => __config.Get("token"));
        public static string? Token
        {
            get => _token.Get();
            set => _token.Set(value);
        }

        private static readonly __Value<string?> _tokenStore = new __Value<string?>(() => __config.Get("tokenStore"));
        public static string? TokenStore
        {
            get => _tokenStore.Get();
            set => _tokenStore.Set(value);
        }

        private static readonly __Value<string?> _url = new __Value<string?>(() => __config.Get("url"));
        public static string? Url
        {
            get => _url.Get();
            set => _url.Set(value);
        }

    }
}
