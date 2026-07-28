---
title: api-shield_operation_schemas
page_id: schema-api-shield-operation-schemas-eb02bec4
path: schemas
description: OpenAPI JSON schemas for an operation, including both user-uploaded and Cloudflare-learned schemas.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_operation_schemas

OpenAPI JSON schemas for an operation, including both user-uploaded and Cloudflare-learned schemas.

```yaml
{"description": "OpenAPI JSON schemas for an operation, including both user-uploaded and Cloudflare-learned schemas.", "type": "object", "properties": {"learned": {"$ref": "#/components/schemas/api-shield_operation_schema_fragment"}, "uploaded": {"$ref": "#/components/schemas/api-shield_operation_schema_fragment"}}, "readOnly": true}
```
