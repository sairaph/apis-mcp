---
title: entitlements.feature
page_id: schema-entitlements-feature-6fd1576e
path: schemas
description: |-
    A feature represents a monetizable ability or functionality in your system.
    Features can be assigned to products, and when those products are purchased, Stripe will create an entitlement to the feature for the purchasing customer.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# entitlements.feature

A feature represents a monetizable ability or functionality in your system.
Features can be assigned to products, and when those products are purchased, Stripe will create an entitlement to the feature for the purchasing customer.

```yaml
{"title": "Feature", "required": ["active", "id", "livemode", "lookup_key", "metadata", "name", "object"], "type": "object", "properties": {"active": {"type": "boolean", "description": "Inactive features cannot be attached to new products and will not be returned from the features list endpoint."}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "lookup_key": {"maxLength": 5000, "type": "string", "description": "A unique key you provide as your own system identifier. This may be up to 80 characters."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "name": {"maxLength": 80, "type": "string", "description": "The feature's name, for your own purpose, not meant to be displayable to the customer."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["entitlements.feature"]}}, "description": "A feature represents a monetizable ability or functionality in your system.\nFeatures can be assigned to products, and when those products are purchased, Stripe will create an entitlement to the feature for the purchasing customer.", "x-expandableFields": [], "x-resourceId": "entitlements.feature"}
```
