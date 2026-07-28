---
title: bill-subs-api_billing-history
page_id: schema-bill-subs-api-billing-history-d7186055
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# bill-subs-api_billing-history

```yaml
{"type": "object", "properties": {"action": {"$ref": "#/components/schemas/bill-subs-api_action"}, "amount": {"$ref": "#/components/schemas/bill-subs-api_amount"}, "currency": {"$ref": "#/components/schemas/bill-subs-api_currency"}, "description": {"$ref": "#/components/schemas/bill-subs-api_description"}, "id": {"$ref": "#/components/schemas/bill-subs-api_components-schemas-identifier"}, "occurred_at": {"$ref": "#/components/schemas/bill-subs-api_occurred_at"}, "type": {"$ref": "#/components/schemas/bill-subs-api_type"}, "zone": {"$ref": "#/components/schemas/bill-subs-api_schemas-zone"}}, "required": ["id", "type", "action", "description", "occurred_at", "amount", "currency", "zone"]}
```
