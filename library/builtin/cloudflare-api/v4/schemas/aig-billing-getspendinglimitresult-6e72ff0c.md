---
title: aig-billing_GetSpendingLimitResult
page_id: schema-aig-billing-getspendinglimitresult-6e72ff0c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aig-billing_GetSpendingLimitResult

```yaml
{"type": "object", "properties": {"config": {"type": "object", "properties": {"amount": {"type": "number", "nullable": true}, "duration": {"type": "string", "nullable": true}, "strategy": {"type": "string", "nullable": true}}, "required": ["amount", "duration", "strategy"]}, "enabled": {"type": "boolean"}}, "required": ["enabled", "config"]}
```
