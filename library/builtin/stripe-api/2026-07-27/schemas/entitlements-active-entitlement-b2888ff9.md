---
title: entitlements.active_entitlement
page_id: schema-entitlements-active-entitlement-b2888ff9
path: schemas
description: An active entitlement describes access to a feature for a customer.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# entitlements.active_entitlement

An active entitlement describes access to a feature for a customer.

```yaml
{"title": "ActiveEntitlement", "required": ["feature", "id", "livemode", "lookup_key", "object"], "type": "object", "properties": {"feature": {"description": "The [Feature](https://docs.stripe.com/api/entitlements/feature) that the customer is entitled to.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/entitlements.feature"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/entitlements.feature"}]}}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "lookup_key": {"maxLength": 5000, "type": "string", "description": "A unique key you provide as your own system identifier. This may be up to 80 characters."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["entitlements.active_entitlement"]}}, "description": "An active entitlement describes access to a feature for a customer.", "x-expandableFields": ["feature"], "x-resourceId": "entitlements.active_entitlement"}
```
