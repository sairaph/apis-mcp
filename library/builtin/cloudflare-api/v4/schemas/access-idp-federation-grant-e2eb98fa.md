---
title: access_idp_federation_grant
page_id: schema-access-idp-federation-grant-e2eb98fa
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_idp_federation_grant

```yaml
{"type": "object", "properties": {"created_at": {"$ref": "#/components/schemas/access_created_at"}, "id": {"description": "UID of the IdP federation grant.", "allOf": [{"$ref": "#/components/schemas/access_identifier"}], "readOnly": true, "x-auditable": true}, "idp_id": {"description": "UID of the identity provider being federated.", "type": "string", "format": "uuid", "example": "a79de439-0e7f-4ebb-8a02-222222222222", "x-auditable": true}}, "required": ["id", "idp_id", "created_at"]}
```
