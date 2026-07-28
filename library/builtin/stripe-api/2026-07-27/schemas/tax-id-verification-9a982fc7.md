---
title: tax_id_verification
page_id: schema-tax-id-verification-9a982fc7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_id_verification

```yaml
{"title": "tax_id_verification", "required": ["status"], "type": "object", "properties": {"status": {"type": "string", "description": "Verification status, one of `pending`, `verified`, `unverified`, or `unavailable`.", "enum": ["pending", "unavailable", "unverified", "verified"]}, "verified_address": {"maxLength": 5000, "type": "string", "description": "Verified address.", "nullable": true}, "verified_name": {"maxLength": 5000, "type": "string", "description": "Verified name.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
