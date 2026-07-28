---
title: bill-subs-api_state
page_id: schema-bill-subs-api-state-efdeffc0
path: schemas
description: The state that the subscription is in.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# bill-subs-api_state

The state that the subscription is in.

```yaml
{"description": "The state that the subscription is in.", "type": "string", "example": "Paid", "enum": ["Trial", "Provisioned", "Paid", "AwaitingPayment", "Cancelled", "Failed", "Expired"], "readOnly": true, "x-auditable": true}
```
