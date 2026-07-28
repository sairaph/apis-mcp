---
title: r2_edit_custom_domain_request
page_id: schema-r2-edit-custom-domain-request-98a1644d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_edit_custom_domain_request

```yaml
{"type": "object", "properties": {"ciphers": {"description": "An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.", "type": "array", "items": {"type": "string"}, "x-auditable": true}, "enabled": {"description": "Whether to enable public bucket access at the specified custom domain.", "type": "boolean", "x-auditable": true}, "minTLS": {"description": "Minimum TLS Version the custom domain will accept for incoming connections. If not set, defaults to previous value.", "type": "string", "enum": ["1.0", "1.1", "1.2", "1.3"], "x-auditable": true}}, "example": {"enabled": true, "minTLS": "1.2"}}
```
