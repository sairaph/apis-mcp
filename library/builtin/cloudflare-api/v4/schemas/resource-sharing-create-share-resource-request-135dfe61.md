---
title: resource-sharing_create_share_resource_request
page_id: schema-resource-sharing-create-share-resource-request-135dfe61
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-sharing_create_share_resource_request

```yaml
{"type": "object", "properties": {"meta": {"$ref": "#/components/schemas/resource-sharing_resource_meta"}, "resource_account_id": {"$ref": "#/components/schemas/resource-sharing_account_id"}, "resource_id": {"$ref": "#/components/schemas/resource-sharing_resource_resource_id"}, "resource_type": {"$ref": "#/components/schemas/resource-sharing_resource_type"}}, "required": ["resource_id", "resource_type", "resource_account_id", "meta"]}
```
