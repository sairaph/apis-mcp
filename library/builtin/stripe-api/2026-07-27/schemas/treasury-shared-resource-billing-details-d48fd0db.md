---
title: treasury_shared_resource_billing_details
page_id: schema-treasury-shared-resource-billing-details-d48fd0db
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_shared_resource_billing_details

```yaml
{"title": "TreasurySharedResourceBillingDetails", "required": ["address"], "type": "object", "properties": {"address": {"$ref": "#/components/schemas/address"}, "email": {"maxLength": 5000, "type": "string", "description": "Email address.", "nullable": true}, "name": {"maxLength": 5000, "type": "string", "description": "Full name.", "nullable": true}}, "description": "", "x-expandableFields": ["address"]}
```
