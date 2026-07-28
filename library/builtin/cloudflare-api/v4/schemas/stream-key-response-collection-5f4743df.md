---
title: stream_key_response_collection
page_id: schema-stream-key-response-collection-5f4743df
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_key_response_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/stream_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"properties": {"created": {"$ref": "#/components/schemas/stream_signing_key_created"}, "id": {"$ref": "#/components/schemas/stream_identifier-2"}, "key_id": {"description": "The unique identifier for the signing key.", "type": "string", "example": "e9db990a82666dd571c77f944a5c5c8d"}}, "type": "object"}}}, "type": "object"}]}
```
