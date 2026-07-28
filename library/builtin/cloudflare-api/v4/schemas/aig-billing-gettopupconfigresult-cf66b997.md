---
title: aig-billing_GetTopupConfigResult
page_id: schema-aig-billing-gettopupconfigresult-cf66b997
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aig-billing_GetTopupConfigResult

```yaml
{"type": "object", "properties": {"amount": {"type": "number", "nullable": true}, "disabledReason": {"type": "string", "nullable": true}, "error": {"type": "string", "nullable": true}, "lastFailedAt": {"type": "number", "nullable": true}, "threshold": {"type": "number", "nullable": true}}, "required": ["threshold", "amount", "error", "disabledReason", "lastFailedAt"]}
```
