---
title: access_custom-claims-support
page_id: schema-access-custom-claims-support-914d64fd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_custom-claims-support

```yaml
{"type": "object", "properties": {"claims": {"description": "Custom claims", "type": "array", "items": {"type": "string"}, "example": ["email_verified", "preferred_username", "custom_claim_name"], "x-auditable": true}, "email_claim_name": {"description": "The claim name for email in the id_token response.", "type": "string", "example": "custom_claim_name", "x-auditable": true}}}
```
