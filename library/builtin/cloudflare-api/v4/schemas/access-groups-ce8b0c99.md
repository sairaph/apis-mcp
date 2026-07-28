---
title: access_groups
page_id: schema-access-groups-ce8b0c99
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_groups

```yaml
{"type": "object", "properties": {"displayName": {"description": "The display name of the SCIM Group resource.", "type": "string", "example": "ALL EMPLOYEES"}, "externalId": {"$ref": "#/components/schemas/access_externalId"}, "id": {"$ref": "#/components/schemas/access_id"}, "meta": {"$ref": "#/components/schemas/access_meta"}, "schemas": {"description": "The list of URIs which indicate the attributes contained within a SCIM resource.", "type": "array", "items": {"type": "string"}, "example": ["urn:ietf:params:scim:schemas:core:2.0:Group"]}}}
```
