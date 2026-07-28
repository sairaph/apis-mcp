---
title: resource-sharing_share_recipient_object
page_id: schema-resource-sharing-share-recipient-object-87ae900a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-sharing_share_recipient_object

```yaml
{"properties": {"account_id": {"$ref": "#/components/schemas/resource-sharing_account_id"}, "association_status": {"$ref": "#/components/schemas/resource-sharing_recipient_association_status"}, "created": {"$ref": "#/components/schemas/resource-sharing_created"}, "id": {"$ref": "#/components/schemas/resource-sharing_recipient_id"}, "modified": {"$ref": "#/components/schemas/resource-sharing_modified"}, "resources": {"type": "array", "items": {"$ref": "#/components/schemas/resource-sharing_share_recipient_resource_object"}}}, "required": ["id", "account_id", "association_status", "created", "modified"]}
```
