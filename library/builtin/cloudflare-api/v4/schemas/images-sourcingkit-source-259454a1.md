---
title: images_sourcingkit_source
page_id: schema-images-sourcingkit-source-259454a1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_sourcingkit_source

```yaml
{"type": "object", "properties": {"bucket": {"description": "The name of the storage bucket.", "type": "string", "example": "my-images-bucket"}, "createdAt": {"description": "When the source was created.", "type": "string", "format": "date-time", "readOnly": true}, "id": {"description": "Source unique identifier.", "type": "string", "format": "uuid", "example": "5ca167e6-4e40-4ced-a209-ca1b9dd6a265", "readOnly": true}, "name": {"description": "A human-readable name for the source.", "type": "string", "example": "my-s3-bucket"}, "region": {"description": "The region of the storage bucket (S3 sources only).", "type": "string", "example": "us-east-1", "nullable": true}, "updatedAt": {"description": "When the source was last updated.", "type": "string", "format": "date-time", "readOnly": true}, "vendor": {"$ref": "#/components/schemas/images_sourcingkit_vendor"}}}
```
