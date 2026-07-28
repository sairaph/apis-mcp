---
title: iam_scim_schema_list_response
page_id: schema-iam-scim-schema-list-response-660d15d4
path: schemas
description: SCIM ListResponse envelope for Schema resources (RFC 7644 Section 4).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_schema_list_response

SCIM ListResponse envelope for Schema resources (RFC 7644 Section 4).

```yaml
{"description": "SCIM ListResponse envelope for Schema resources (RFC 7644 Section 4).\n", "type": "object", "properties": {"Resources": {"type": "array", "items": {"$ref": "#/components/schemas/iam_scim_schema"}}, "itemsPerPage": {"type": "integer", "example": 2}, "schemas": {"type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:api:messages:2.0:ListResponse"]}, "startIndex": {"type": "integer", "example": 1}, "totalResults": {"type": "integer", "example": 2}}, "required": ["schemas", "totalResults", "Resources"], "title": "SCIM Schema List Response"}
```
