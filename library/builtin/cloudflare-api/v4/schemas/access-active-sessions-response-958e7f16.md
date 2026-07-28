---
title: access_active_sessions_response
page_id: schema-access-active-sessions-response-958e7f16
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_active_sessions_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"properties": {"expiration": {"type": "integer", "example": 1694813506}, "metadata": {"type": "object", "properties": {"apps": {"type": "object", "additionalProperties": {"properties": {"hostname": {"type": "string", "example": "test.example.com"}, "name": {"type": "string", "example": "app name"}, "type": {"type": "string", "example": "self_hosted"}, "uid": {"type": "string", "example": "cc2a8145-0128-4429-87f3-872c4d380c4e"}}, "type": "object"}}, "expires": {"type": "integer", "example": 1694813506}, "iat": {"type": "integer", "example": 1694791905}, "nonce": {"type": "string", "example": "X1aXj1lFVcqqyoXF"}, "ttl": {"type": "integer", "example": 21600}}}, "name": {"type": "string"}}, "type": "object"}}}}]}
```
