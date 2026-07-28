---
title: email-security_ReleaseResponse
page_id: schema-email-security-releaseresponse-7eed8db0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_ReleaseResponse

```yaml
{"type": "object", "properties": {"delivered": {"type": "array", "items": {"type": "string"}, "nullable": true}, "failed": {"type": "array", "items": {"type": "string"}, "nullable": true}, "id": {"allOf": [{"$ref": "#/components/schemas/email-security_InvestigateId"}], "x-auditable": true}, "postfix_id": {"description": "Deprecated, use `id` instead. End of life: November 1, 2026.", "allOf": [{"$ref": "#/components/schemas/email-security_PostfixId"}], "deprecated": true, "x-stainless-deprecation-message": "Use `id` instead."}, "undelivered": {"type": "array", "items": {"type": "string"}, "nullable": true}}, "required": ["id"]}
```
