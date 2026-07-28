---
title: resource-tagging_tagged_resource_object_account_level_base
page_id: schema-resource-tagging-tagged-resource-object-account-level-base-17d9ba5a
path: schemas
description: Base schema for account-level resources (without type discriminator)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_tagged_resource_object_account_level_base

Base schema for account-level resources (without type discriminator)

```yaml
{"description": "Base schema for account-level resources (without type discriminator)", "type": "object", "properties": {"etag": {"$ref": "#/components/schemas/resource-tagging_etag"}, "id": {"$ref": "#/components/schemas/resource-tagging_resource_id"}, "name": {"$ref": "#/components/schemas/resource-tagging_resource_name"}, "tags": {"$ref": "#/components/schemas/resource-tagging_tags"}}, "required": ["id", "tags", "name", "etag"]}
```
