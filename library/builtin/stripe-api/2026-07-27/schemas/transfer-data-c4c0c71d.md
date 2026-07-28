---
title: transfer_data
page_id: schema-transfer-data-c4c0c71d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# transfer_data

```yaml
{"title": "transfer_data", "required": ["destination"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount transferred to the destination account. This transfer will occur automatically after the payment succeeds. If no amount is specified, by default the entire payment amount is transferred to the destination account.\n The amount must be less than or equal to the [amount](https://docs.stripe.com/api/payment_intents/object#payment_intent_object-amount), and must be a positive integer\n representing how much to transfer in the smallest currency unit (e.g., 100 cents to charge $1.00)."}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the transfer. Often useful for displaying to users."}, "destination": {"description": "The account (if any) that the payment is attributed to for tax reporting, and where funds from the payment are transferred to after payment success.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/account"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/account"}]}}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "payment_data": {"$ref": "#/components/schemas/payment_data"}}, "description": "", "x-expandableFields": ["destination", "payment_data"]}
```
