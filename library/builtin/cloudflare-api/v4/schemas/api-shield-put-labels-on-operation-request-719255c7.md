---
title: api-shield_put_labels_on_operation_request
page_id: schema-api-shield-put-labels-on-operation-request-719255c7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_put_labels_on_operation_request

```yaml
{"type": "object", "properties": {"managed": {"description": "List of managed label names. Omitting this property or passing an empty array will result in all managed labels being removed from the operation", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_label_name"}}, "user": {"description": "List of user label names. Omitting this property or passing an empty array will result in all user labels being removed from the operation", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_label_name"}}}}
```
