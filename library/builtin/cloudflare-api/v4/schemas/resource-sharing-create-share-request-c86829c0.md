---
title: resource-sharing_create_share_request
page_id: schema-resource-sharing-create-share-request-c86829c0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-sharing_create_share_request

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/resource-sharing_share_name"}, "recipients": {"type": "array", "items": {"$ref": "#/components/schemas/resource-sharing_create_share_recipient_request"}}, "resources": {"type": "array", "items": {"$ref": "#/components/schemas/resource-sharing_create_share_resource_request"}}}, "required": ["name", "resources", "recipients"]}
```
