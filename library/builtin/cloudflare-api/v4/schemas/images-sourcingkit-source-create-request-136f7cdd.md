---
title: images_sourcingkit_source_create_request
page_id: schema-images-sourcingkit-source-create-request-136f7cdd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_sourcingkit_source_create_request

```yaml
{"type": "object", "properties": {"account": {"description": "Account identifier for the bucket (required for R2 vendor).", "type": "string"}, "bucket": {"description": "The name of the storage bucket.", "type": "string", "example": "my-images-bucket", "maxLength": 128, "minLength": 1}, "name": {"description": "A human-readable name for the source.", "type": "string", "example": "my-s3-bucket", "maxLength": 128, "minLength": 1}, "secret": {"description": "Storage credentials for accessing the bucket. Shape depends on vendor.", "type": "object", "writeOnly": true}, "vendor": {"$ref": "#/components/schemas/images_sourcingkit_vendor"}}, "required": ["name", "bucket", "vendor", "secret"]}
```
