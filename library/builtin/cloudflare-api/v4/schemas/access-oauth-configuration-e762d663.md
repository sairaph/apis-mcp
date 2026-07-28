---
title: access_oauth_configuration
page_id: schema-access-oauth-configuration-e762d663
path: schemas
description: '**Beta:** Optional configuration for managing an OAuth authorization flow controlled by Access. When set, Access will act as the OAuth authorization server for this application. Only compatible with OAuth clients that support [RFC 8707](https://datatracker.ietf.org/doc/html/rfc8707) (Resource Indicators for OAuth 2.0). This feature is currently in beta.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_oauth_configuration

**Beta:** Optional configuration for managing an OAuth authorization flow controlled by Access. When set, Access will act as the OAuth authorization server for this application. Only compatible with OAuth clients that support [RFC 8707](https://datatracker.ietf.org/doc/html/rfc8707) (Resource Indicators for OAuth 2.0). This feature is currently in beta.

```yaml
{"description": "**Beta:** Optional configuration for managing an OAuth authorization flow controlled by Access. When set, Access will act as the OAuth authorization server for this application. Only compatible with OAuth clients that support [RFC 8707](https://datatracker.ietf.org/doc/html/rfc8707) (Resource Indicators for OAuth 2.0). This feature is currently in beta.\n", "type": "object", "properties": {"dynamic_client_registration": {"description": "Settings for OAuth dynamic client registration.", "type": "object", "properties": {"allow_any_on_localhost": {"description": "Allows any client with redirect URIs on localhost.", "type": "boolean"}, "allow_any_on_loopback": {"description": "Allows any client with redirect URIs on 127.0.0.1.", "type": "boolean"}, "allowed_uris": {"description": "The URIs that are allowed as redirect URIs for dynamically registered clients. HTTP and HTTPS paths may end in `/*` to match all sub-paths. Custom-scheme URIs must be explicitly configured and match exactly.\n", "type": "array", "items": {"type": "string"}, "example": ["https://example.com/callback", "com.example.app:/oauth/callback"]}, "enabled": {"description": "Whether dynamic client registration is enabled.", "type": "boolean"}}}, "enabled": {"description": "Whether the OAuth configuration is enabled for this application. When set to `false`, Access will not handle OAuth for this application. Defaults to `true` if omitted.\n", "type": "boolean", "default": true}, "grant": {"description": "Settings for OAuth grant behavior.", "type": "object", "properties": {"access_token_lifetime": {"description": "The lifetime of the access token. Must be in the format `300ms` or `2h45m`. Valid time units are ns, us (or µs), ms, s, m, h.", "type": "string", "example": "5m"}, "session_duration": {"description": "The duration of the OAuth session. Must be in the format `300ms` or `2h45m`. Valid time units are ns, us (or µs), ms, s, m, h.", "type": "string", "example": "24h"}}}}}
```
