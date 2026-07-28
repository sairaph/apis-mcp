---
title: issuing.physical_bundle
page_id: schema-issuing-physical-bundle-1c64d505
path: schemas
description: A Physical Bundle represents the bundle of physical items - card stock, carrier letter, and envelope - that is shipped to a cardholder when you create a physical card.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing.physical_bundle

A Physical Bundle represents the bundle of physical items - card stock, carrier letter, and envelope - that is shipped to a cardholder when you create a physical card.

```yaml
{"title": "IssuingPhysicalBundle", "required": ["features", "id", "livemode", "name", "object", "status", "type"], "type": "object", "properties": {"features": {"$ref": "#/components/schemas/issuing_physical_bundle_features"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "name": {"maxLength": 5000, "type": "string", "description": "Friendly display name."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["issuing.physical_bundle"]}, "status": {"type": "string", "description": "Whether this physical bundle can be used to create cards.", "enum": ["active", "inactive", "review"]}, "type": {"type": "string", "description": "Whether this physical bundle is a standard Stripe offering or custom-made for you.", "enum": ["custom", "standard"]}}, "description": "A Physical Bundle represents the bundle of physical items - card stock, carrier letter, and envelope - that is shipped to a cardholder when you create a physical card.", "x-expandableFields": ["features"], "x-resourceId": "issuing.physical_bundle"}
```
