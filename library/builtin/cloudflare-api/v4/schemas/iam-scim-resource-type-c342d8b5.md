---
title: iam_scim_resource_type
page_id: schema-iam-scim-resource-type-c342d8b5
path: schemas
description: A SCIM ResourceType resource (RFC 7643 Section 6). Describes a category of SCIM resource (e.g. User, Group) and its associated schema.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_resource_type

A SCIM ResourceType resource (RFC 7643 Section 6). Describes a category of SCIM resource (e.g. User, Group) and its associated schema.

```yaml
{"description": "A SCIM ResourceType resource (RFC 7643 Section 6). Describes a category of SCIM resource (e.g. User, Group) and its associated schema.\n", "type": "object", "properties": {"description": {"description": "The resource type's human-readable description.", "type": "string", "example": "User Account"}, "endpoint": {"description": "The resource type's HTTP-addressable endpoint relative to the base URL.", "type": "string", "example": "/Users"}, "id": {"description": "The resource type's server unique id.", "type": "string", "example": "User"}, "meta": {"$ref": "#/components/schemas/iam_scim_resource_type_meta"}, "name": {"description": "The resource type name.", "type": "string", "example": "User"}, "schema": {"description": "The resource type's primary/base schema URI.", "type": "string", "example": "urn:ietf:params:scim:schemas:core:2.0:User"}, "schemaExtensions": {"description": "A list of URIs of the resource type's schema extensions.", "type": "array", "items": {"$ref": "#/components/schemas/iam_scim_schema_extension"}}, "schemas": {"type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:schemas:core:2.0:ResourceType"]}}, "required": ["schemas", "id", "name", "endpoint", "schema"], "title": "SCIM Resource Type"}
```
