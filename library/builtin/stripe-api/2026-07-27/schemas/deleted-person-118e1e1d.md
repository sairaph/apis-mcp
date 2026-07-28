---
title: deleted_person
page_id: schema-deleted-person-118e1e1d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# deleted_person

```yaml
{"title": "DeletedPerson", "required": ["deleted", "id", "object"], "type": "object", "properties": {"deleted": {"type": "boolean", "description": "Always true for a deleted object", "enum": [true]}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["person"]}}, "description": "", "x-expandableFields": [], "x-resourceId": "deleted_person"}
```
