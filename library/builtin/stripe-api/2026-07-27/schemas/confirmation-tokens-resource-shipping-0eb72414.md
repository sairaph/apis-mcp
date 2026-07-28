---
title: confirmation_tokens_resource_shipping
page_id: schema-confirmation-tokens-resource-shipping-0eb72414
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# confirmation_tokens_resource_shipping

```yaml
{"title": "ConfirmationTokensResourceShipping", "required": ["address", "name"], "type": "object", "properties": {"address": {"$ref": "#/components/schemas/address"}, "name": {"maxLength": 5000, "type": "string", "description": "Recipient name."}, "phone": {"maxLength": 5000, "type": "string", "description": "Recipient phone (including extension).", "nullable": true}}, "description": "", "x-expandableFields": ["address"]}
```
