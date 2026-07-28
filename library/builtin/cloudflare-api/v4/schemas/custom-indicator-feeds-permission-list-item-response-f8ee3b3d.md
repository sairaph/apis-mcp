---
title: custom-indicator-feeds_permission_list_item_response
page_id: schema-custom-indicator-feeds-permission-list-item-response-f8ee3b3d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# custom-indicator-feeds_permission_list_item_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/custom-indicator-feeds_permission_list_item"}, "example": [{"description": "An important indicator list", "id": 1, "is_attributable": false, "is_downloadable": false, "is_public": false, "name": "indicator_list_1"}, {"description": "An even more important indicator list", "id": 2, "is_attributable": true, "is_downloadable": false, "is_public": true, "name": "indicator_list_2"}]}}}]}
```
