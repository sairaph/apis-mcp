---
title: api-shield_delete_labels_on_operation_request
page_id: schema-api-shield-delete-labels-on-operation-request-f1f29971
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_delete_labels_on_operation_request

```yaml
{"type": "object", "properties": {"managed": {"description": "List of managed label names.", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_label_name"}}, "user": {"description": "List of user label names.", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_label_name"}}}}
```
