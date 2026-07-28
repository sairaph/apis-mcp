---
title: mandate
page_id: schema-mandate-d959b638
path: schemas
description: A Mandate is a record of the permission that your customer gives you to debit their payment method.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# mandate

A Mandate is a record of the permission that your customer gives you to debit their payment method.

```yaml
{"title": "Mandate", "required": ["customer_acceptance", "id", "livemode", "object", "payment_method", "payment_method_details", "status", "type"], "type": "object", "properties": {"customer_acceptance": {"$ref": "#/components/schemas/customer_acceptance"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "multi_use": {"$ref": "#/components/schemas/mandate_multi_use"}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["mandate"]}, "on_behalf_of": {"maxLength": 5000, "type": "string", "description": "The account (if any) that the mandate is intended for."}, "payment_method": {"description": "ID of the payment method associated with this mandate.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/payment_method"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/payment_method"}]}}, "payment_method_details": {"$ref": "#/components/schemas/mandate_payment_method_details"}, "single_use": {"$ref": "#/components/schemas/mandate_single_use"}, "status": {"type": "string", "description": "The mandate status indicates whether or not you can use it to initiate a payment.", "enum": ["active", "inactive", "pending"]}, "type": {"type": "string", "description": "The type of the mandate.", "enum": ["multi_use", "single_use"]}}, "description": "A Mandate is a record of the permission that your customer gives you to debit their payment method.", "x-expandableFields": ["customer_acceptance", "multi_use", "payment_method", "payment_method_details", "single_use"], "x-resourceId": "mandate"}
```
