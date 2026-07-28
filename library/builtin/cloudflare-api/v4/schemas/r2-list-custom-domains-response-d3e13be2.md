---
title: r2_list_custom_domains_response
page_id: schema-r2-list-custom-domains-response-d3e13be2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_list_custom_domains_response

```yaml
{"type": "object", "properties": {"domains": {"type": "array", "items": {"properties": {"ciphers": {"description": "An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.", "type": "array", "items": {"type": "string"}, "x-auditable": true}, "domain": {"description": "Domain name of the custom domain to be added.", "type": "string", "x-auditable": true}, "enabled": {"description": "Whether this bucket is publicly accessible at the specified custom domain.", "type": "boolean", "x-auditable": true}, "minTLS": {"description": "Minimum TLS Version the custom domain will accept for incoming connections. If not set, defaults to 1.0.", "type": "string", "enum": ["1.0", "1.1", "1.2", "1.3"], "x-auditable": true}, "status": {"type": "object", "properties": {"ownership": {"description": "Ownership status of the domain.", "type": "string", "enum": ["pending", "active", "deactivated", "blocked", "error", "unknown"], "x-auditable": true}, "ssl": {"description": "SSL certificate status.", "type": "string", "enum": ["initializing", "pending", "active", "deactivated", "error", "unknown"], "x-auditable": true}}, "required": ["ssl", "ownership"]}, "zoneId": {"description": "Zone ID of the custom domain resides in.", "type": "string", "x-auditable": true}, "zoneName": {"description": "Zone that the custom domain resides in.", "type": "string", "x-auditable": true}}, "required": ["domain", "status", "enabled"], "type": "object"}}}, "example": {"domains": [{"domain": "prefix.example-domain.one.com", "enabled": false, "status": {"ownership": "deactivated", "ssl": "pending"}, "zoneId": "36ca64a6d92827b8a6b90be344bb1bfd", "zoneName": "example-domain.one.com"}, {"domain": "prefix.example-domain.two.com", "enabled": true, "status": {"ownership": "active", "ssl": "active"}, "zoneId": "d9d28585d5f8f5b0f857b055bf574f19"}]}, "required": ["domains"]}
```
