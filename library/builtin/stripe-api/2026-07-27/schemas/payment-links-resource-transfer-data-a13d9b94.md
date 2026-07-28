---
title: payment_links_resource_transfer_data
page_id: schema-payment-links-resource-transfer-data-a13d9b94
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_transfer_data

```yaml
{"title": "PaymentLinksResourceTransferData", "required": ["destination"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount in cents (or local equivalent) that will be transferred to the destination account. By default, the entire amount is transferred to the destination.", "nullable": true}, "destination": {"description": "The connected account receiving the transfer.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/account"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/account"}]}}}, "description": "", "x-expandableFields": ["destination"]}
```
