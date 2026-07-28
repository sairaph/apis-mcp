---
title: magic_create_redundancy_group_request
page_id: schema-magic-create-redundancy-group-request-60462b48
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_create_redundancy_group_request

```yaml
{"type": "object", "properties": {"description": {"description": "Optional description", "type": "string", "default": "", "maxLength": 1024}, "members": {"description": "Tunnels to add to the group", "type": "array", "items": {"$ref": "#/components/schemas/magic_redundancy_member_request"}, "default": []}, "name": {"description": "Human-readable name for the redundancy group", "type": "string", "example": "primary-group", "maxLength": 255, "minLength": 1}}, "required": ["name"]}
```
