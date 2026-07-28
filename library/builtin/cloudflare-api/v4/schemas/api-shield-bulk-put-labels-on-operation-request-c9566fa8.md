---
title: api-shield_bulk_put_labels_on_operation_request
page_id: schema-api-shield-bulk-put-labels-on-operation-request-c9566fa8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_bulk_put_labels_on_operation_request

```yaml
{"type": "object", "properties": {"managed": {"description": "Managed labels to replace for all affected operations", "type": "object", "properties": {"labels": {"description": "List of managed label names. Providing an empty array will result in all managed labels being removed from all affected operations", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_label_name"}}}, "required": ["labels"]}, "selector": {"$ref": "#/components/schemas/api-shield_operation_id_selector"}, "user": {"description": "User labels to replace for all affected operations", "type": "object", "properties": {"labels": {"description": "List of user label names. Providing an empty array will result in all user labels being removed from all affected operations", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_label_name"}}}, "required": ["labels"]}}, "required": ["selector", "user", "managed"]}
```
