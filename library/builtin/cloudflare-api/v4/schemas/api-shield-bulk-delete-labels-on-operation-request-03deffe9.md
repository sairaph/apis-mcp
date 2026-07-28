---
title: api-shield_bulk_delete_labels_on_operation_request
page_id: schema-api-shield-bulk-delete-labels-on-operation-request-03deffe9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_bulk_delete_labels_on_operation_request

```yaml
{"type": "object", "properties": {"managed": {"type": "object", "properties": {"labels": {"description": "List of managed label names.", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_label_name"}}}}, "selector": {"$ref": "#/components/schemas/api-shield_operation_id_selector"}, "user": {"type": "object", "properties": {"labels": {"description": "List of user label names.", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_label_name"}}}}}, "required": ["selector"]}
```
