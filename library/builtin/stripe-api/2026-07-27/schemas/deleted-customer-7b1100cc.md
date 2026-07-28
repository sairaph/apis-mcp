---
title: deleted_customer
page_id: schema-deleted-customer-7b1100cc
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# deleted_customer

```yaml
{"title": "DeletedCustomer", "required": ["deleted", "id", "object"], "type": "object", "properties": {"deleted": {"type": "boolean", "description": "Always true for a deleted object", "enum": [true]}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["customer"]}}, "description": "", "x-expandableFields": [], "x-resourceId": "deleted_customer"}
```
