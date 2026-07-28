---
title: organizations-api_Organization
page_id: schema-organizations-api-organization-74350a22
path: schemas
description: References an Organization in the Cloudflare data model.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# organizations-api_Organization

References an Organization in the Cloudflare data model.

```yaml
{"description": "References an Organization in the Cloudflare data model.", "type": "object", "properties": {"create_time": {"type": "string", "format": "date-time", "readOnly": true}, "id": {"allOf": [{"$ref": "#/components/schemas/organizations-api_OrganizationID"}], "readOnly": true}, "meta": {"type": "object", "additionalProperties": {"type": "object"}, "properties": {"flags": {"allOf": [{"$ref": "#/components/schemas/organizations-api_OrganizationFlags"}]}, "hierarchy_tags": {"description": "Ordered chain of organization tags from the root organization down to\n(and including) this organization itself. Root organizations return a\nsingle-element array containing their own tag; sub-organizations return\n`[rootTag, ...intermediateTags, parentTag, selfTag]`. Useful for\nconstructing authorization scopes that need to cover every ancestor\nin the hierarchy.", "type": "array", "items": {"type": "string"}}, "managed_by": {"type": "string"}}, "readOnly": true}, "name": {"type": "string"}, "parent": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/organizations-api_OrganizationID"}, "name": {"type": "string", "readOnly": true}}, "required": ["id", "name"]}, "profile": {"$ref": "#/components/schemas/organizations-api_Profile"}}, "required": ["id", "name", "create_time", "meta"]}
```
