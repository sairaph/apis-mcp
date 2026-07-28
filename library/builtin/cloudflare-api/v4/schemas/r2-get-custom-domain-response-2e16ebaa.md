---
title: r2_get_custom_domain_response
page_id: schema-r2-get-custom-domain-response-2e16ebaa
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_get_custom_domain_response

```yaml
{"type": "object", "properties": {"ciphers": {"description": "An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.", "type": "array", "items": {"type": "string"}, "x-auditable": true}, "domain": {"description": "Domain name of the custom domain to be added.", "type": "string", "x-auditable": true}, "enabled": {"description": "Whether this bucket is publicly accessible at the specified custom domain.", "type": "boolean", "x-auditable": true}, "minTLS": {"description": "Minimum TLS Version the custom domain will accept for incoming connections. If not set, defaults to 1.0.", "type": "string", "enum": ["1.0", "1.1", "1.2", "1.3"], "x-auditable": true}, "status": {"type": "object", "properties": {"ownership": {"description": "Ownership status of the domain.", "type": "string", "enum": ["pending", "active", "deactivated", "blocked", "error", "unknown"], "x-auditable": true}, "ssl": {"description": "SSL certificate status.", "type": "string", "enum": ["initializing", "pending", "active", "deactivated", "error", "unknown"], "x-auditable": true}}, "required": ["ssl", "ownership"]}, "zoneId": {"description": "Zone ID of the custom domain resides in.", "type": "string", "x-auditable": true}, "zoneName": {"description": "Zone that the custom domain resides in.", "type": "string", "x-auditable": true}}, "example": {"domain": "prefix.example-domain.one.com", "enabled": false, "status": {"ownership": "deactivated", "ssl": "pending"}, "zoneId": "36ca64a6d92827b8a6b90be344bb1bfd", "zoneName": "example-domain.one.com"}, "required": ["domain", "status", "enabled"]}
```
