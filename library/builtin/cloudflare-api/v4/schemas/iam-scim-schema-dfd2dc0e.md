---
title: iam_scim_schema
page_id: schema-iam-scim-schema-dfd2dc0e
path: schemas
description: A SCIM Schema resource (RFC 7643 Section 7). Defines the attributes of a SCIM resource type (e.g. User or Group).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_schema

A SCIM Schema resource (RFC 7643 Section 7). Defines the attributes of a SCIM resource type (e.g. User or Group).

```yaml
{"description": "A SCIM Schema resource (RFC 7643 Section 7). Defines the attributes of a SCIM resource type (e.g. User or Group).\n", "type": "object", "properties": {"attributes": {"description": "A complex attribute that includes the attributes of a schema.", "type": "array", "items": {"$ref": "#/components/schemas/iam_scim_schema_attr"}}, "description": {"description": "The schema's human-readable description.", "type": "string", "example": "User Account"}, "id": {"description": "The unique URI of the schema.", "type": "string", "example": "urn:ietf:params:scim:schemas:core:2.0:User"}, "meta": {"$ref": "#/components/schemas/iam_scim_schema_meta"}, "name": {"description": "The schema's human-readable name.", "type": "string", "example": "User"}, "schemas": {"type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:schemas:core:2.0:Schema"]}}, "required": ["schemas", "id", "name", "attributes"], "title": "SCIM Schema"}
```
