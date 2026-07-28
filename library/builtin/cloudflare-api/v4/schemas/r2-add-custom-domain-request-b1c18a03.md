---
title: r2_add_custom_domain_request
page_id: schema-r2-add-custom-domain-request-b1c18a03
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_add_custom_domain_request

```yaml
{"type": "object", "properties": {"ciphers": {"description": "An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.", "type": "array", "items": {"type": "string"}, "x-auditable": true}, "domain": {"description": "Name of the custom domain to be added.", "type": "string", "x-auditable": true}, "enabled": {"description": "Whether to enable public bucket access at the custom domain. If undefined, the domain will be enabled.", "type": "boolean", "x-auditable": true}, "minTLS": {"description": "Minimum TLS Version the custom domain will accept for incoming connections. If not set, defaults to 1.0.", "type": "string", "enum": ["1.0", "1.1", "1.2", "1.3"], "x-auditable": true}, "zoneId": {"description": "Zone ID of the custom domain.", "type": "string", "x-auditable": true}}, "example": {"domain": "prefix.example-domain.com", "enabled": true, "zoneId": "36ca64a6d92827b8a6b90be344bb1bfd"}, "required": ["domain", "zoneId", "enabled"]}
```
