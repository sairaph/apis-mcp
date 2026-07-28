---
title: access_service-tokens
page_id: schema-access-service-tokens-07096569
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_service-tokens

```yaml
{"type": "object", "properties": {"client_id": {"$ref": "#/components/schemas/access_client_id"}, "created_at": {"$ref": "#/components/schemas/access_created_at"}, "duration": {"$ref": "#/components/schemas/access_duration"}, "expires_at": {"$ref": "#/components/schemas/access_timestamp"}, "id": {"allOf": [{"description": "The ID of the service token."}, {"$ref": "#/components/schemas/access_uuid"}]}, "last_seen_at": {"allOf": [{"x-stainless-skip": true}, {"$ref": "#/components/schemas/access_timestamp"}]}, "name": {"$ref": "#/components/schemas/access_name-2"}, "updated_at": {"$ref": "#/components/schemas/access_updated_at"}}}
```
