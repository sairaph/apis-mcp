---
title: iam_scim_user_list_response
page_id: schema-iam-scim-user-list-response-97820aad
path: schemas
description: SCIM ListResponse envelope for User resources (RFC 7644 Section 3.4.2).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_user_list_response

SCIM ListResponse envelope for User resources (RFC 7644 Section 3.4.2).

```yaml
{"description": "SCIM ListResponse envelope for User resources (RFC 7644 Section 3.4.2).\n", "type": "object", "properties": {"Resources": {"type": "array", "items": {"$ref": "#/components/schemas/iam_scim_user"}}, "itemsPerPage": {"description": "The number of resources returned in this page.", "type": "integer", "example": 3}, "schemas": {"type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:api:messages:2.0:ListResponse"]}, "startIndex": {"description": "The 1-based index of the first result in this set.", "type": "integer", "example": 1}, "totalResults": {"description": "The total number of results matching the query.", "type": "integer", "example": 3}}, "required": ["schemas", "totalResults", "Resources"], "title": "SCIM User List Response"}
```
