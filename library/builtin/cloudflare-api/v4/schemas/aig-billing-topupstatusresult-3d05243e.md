---
title: aig-billing_TopupStatusResult
page_id: schema-aig-billing-topupstatusresult-3d05243e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aig-billing_TopupStatusResult

```yaml
{"type": "object", "properties": {"payment_intent_id": {"type": "string"}, "status": {"type": "string", "enum": ["completed", "pending"]}}, "required": ["status", "payment_intent_id"]}
```
