---
title: addressing_loa_upload_response
page_id: schema-addressing-loa-upload-response-ea1273a8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# addressing_loa_upload_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/addressing_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"account_id": {"$ref": "#/components/schemas/addressing_account_identifier"}, "auto_generated": {"$ref": "#/components/schemas/addressing_auto_generated"}, "created": {"$ref": "#/components/schemas/addressing_timestamp"}, "filename": {"$ref": "#/components/schemas/addressing_filename"}, "id": {"$ref": "#/components/schemas/addressing_loa_document_identifier"}, "size_bytes": {"$ref": "#/components/schemas/addressing_size_bytes"}, "verified": {"$ref": "#/components/schemas/addressing_verified"}, "verified_at": {"$ref": "#/components/schemas/addressing_verified_at"}}}}}]}
```
