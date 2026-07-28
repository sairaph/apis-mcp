---
title: quotes_resource_transfer_data
page_id: schema-quotes-resource-transfer-data-22a435f3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# quotes_resource_transfer_data

```yaml
{"title": "QuotesResourceTransferData", "required": ["destination"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount in cents (or local equivalent) that will be transferred to the destination account when the invoice is paid. By default, the entire amount is transferred to the destination.", "nullable": true}, "amount_percent": {"type": "number", "description": "A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the destination account. By default, the entire amount will be transferred to the destination.", "nullable": true}, "destination": {"description": "The account where funds from the payment will be transferred to upon payment success.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/account"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/account"}]}}}, "description": "", "x-expandableFields": ["destination"]}
```
