---
title: access_create_response
page_id: schema-access-create-response-299ea96e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_create_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"client_id": {"$ref": "#/components/schemas/access_client_id"}, "client_secret": {"$ref": "#/components/schemas/access_client_secret"}, "created_at": {"$ref": "#/components/schemas/access_created_at"}, "duration": {"$ref": "#/components/schemas/access_duration"}, "id": {"description": "The ID of the service token.", "type": "string"}, "name": {"$ref": "#/components/schemas/access_name-2"}, "updated_at": {"$ref": "#/components/schemas/access_updated_at"}}}}}]}
```
