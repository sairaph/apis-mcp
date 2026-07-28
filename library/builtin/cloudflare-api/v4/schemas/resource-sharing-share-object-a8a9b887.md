---
title: resource-sharing_share_object
page_id: schema-resource-sharing-share-object-a8a9b887
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-sharing_share_object

```yaml
{"properties": {"account_id": {"$ref": "#/components/schemas/resource-sharing_account_id"}, "account_name": {"$ref": "#/components/schemas/resource-sharing_account_name"}, "associated_recipient_count": {"description": "The number of recipients in the 'associated' state. This field is only included when requested via the 'include_recipient_counts' parameter.", "type": "integer", "example": 10}, "associating_recipient_count": {"description": "The number of recipients in the 'associating' state. This field is only included when requested via the 'include_recipient_counts' parameter.", "type": "integer", "example": 1}, "created": {"$ref": "#/components/schemas/resource-sharing_created"}, "disassociated_recipient_count": {"description": "The number of recipients in the 'disassociated' state. This field is only included when requested via the 'include_recipient_counts' parameter.", "type": "integer", "example": 0}, "disassociating_recipient_count": {"description": "The number of recipients in the 'disassociating' state. This field is only included when requested via the 'include_recipient_counts' parameter.", "type": "integer", "example": 0}, "id": {"$ref": "#/components/schemas/resource-sharing_share_id"}, "kind": {"$ref": "#/components/schemas/resource-sharing_share_kind"}, "modified": {"$ref": "#/components/schemas/resource-sharing_modified"}, "name": {"$ref": "#/components/schemas/resource-sharing_share_name"}, "organization_id": {"$ref": "#/components/schemas/resource-sharing_organization_id"}, "resources": {"description": "A list of resources that are part of the share. This field is only included when requested via the 'include_resources' parameter.", "type": "array", "items": {"$ref": "#/components/schemas/resource-sharing_share_resource_object"}}, "status": {"$ref": "#/components/schemas/resource-sharing_share_status"}, "target_type": {"$ref": "#/components/schemas/resource-sharing_share_target_type"}}, "required": ["id", "name", "account_id", "account_name", "organization_id", "created", "modified", "status", "target_type"]}
```
