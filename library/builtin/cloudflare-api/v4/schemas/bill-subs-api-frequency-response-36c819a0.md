---
title: bill-subs-api_frequency_response
page_id: schema-bill-subs-api-frequency-response-36c819a0
path: schemas
description: How often the subscription is renewed automatically (computed field).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# bill-subs-api_frequency_response

How often the subscription is renewed automatically (computed field).

```yaml
{"description": "How often the subscription is renewed automatically (computed field).", "allOf": [{"$ref": "#/components/schemas/bill-subs-api_frequency"}, {"enum": ["weekly", "monthly", "quarterly", "yearly", "not-applicable"], "type": "string"}], "readOnly": true, "x-stainless-terraform-configurability": "computed"}
```
