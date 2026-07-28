---
title: api-shield_post_labels_on_operation_request
page_id: schema-api-shield-post-labels-on-operation-request-59e7fdb8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_post_labels_on_operation_request

```yaml
{"type": "object", "properties": {"managed": {"description": "List of managed label names.", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_label_name"}, "minItems": 1}, "user": {"description": "List of user label names.", "type": "array", "items": {"$ref": "#/components/schemas/api-shield_label_name"}, "minItems": 1}}}
```
