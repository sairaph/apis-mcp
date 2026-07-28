---
title: r2_edit_custom_domain_response
page_id: schema-r2-edit-custom-domain-response-10572753
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_edit_custom_domain_response

```yaml
{"type": "object", "properties": {"ciphers": {"description": "An allowlist of ciphers for TLS termination. These ciphers must be in the BoringSSL format.", "type": "array", "items": {"type": "string"}, "x-auditable": true}, "domain": {"description": "Domain name of the affected custom domain.", "type": "string", "x-auditable": true}, "enabled": {"description": "Whether this bucket is publicly accessible at the specified custom domain.", "type": "boolean", "x-auditable": true}, "minTLS": {"description": "Minimum TLS Version the custom domain will accept for incoming connections. If not set, defaults to 1.0.", "type": "string", "enum": ["1.0", "1.1", "1.2", "1.3"], "x-auditable": true}}, "example": {"domain": "example-domain.com", "enabled": true}, "required": ["domain"]}
```
