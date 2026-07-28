---
title: iam_request_create_resource_group
page_id: schema-iam-request-create-resource-group-77778887
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_request_create_resource_group

```yaml
{"type": "object", "properties": {"name": {"description": "Name of the resource group", "type": "string", "example": "NewResourceGroup"}, "scope": {"$ref": "#/components/schemas/iam_create-scope"}}, "required": ["name", "scope"], "title": "Create resource group"}
```
