---
title: resource-tagging_delete_tags_request_account_level_base
page_id: schema-resource-tagging-delete-tags-request-account-level-base-82b62f7d
path: schemas
description: Request body schema for deleting tags from account-level resources.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_delete_tags_request_account_level_base

Request body schema for deleting tags from account-level resources.

```yaml
{"description": "Request body schema for deleting tags from account-level resources.", "type": "object", "properties": {"resource_id": {"$ref": "#/components/schemas/resource-tagging_resource_id"}, "resource_type": {"$ref": "#/components/schemas/resource-tagging_account_resource_type_base_enum"}}, "required": ["resource_type", "resource_id"]}
```
