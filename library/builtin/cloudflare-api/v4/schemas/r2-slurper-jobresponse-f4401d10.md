---
title: r2-slurper_JobResponse
page_id: schema-r2-slurper-jobresponse-f4401d10
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-slurper_JobResponse

```yaml
{"type": "object", "properties": {"createdAt": {"type": "string"}, "finishedAt": {"type": "string", "nullable": true}, "id": {"type": "string"}, "overwrite": {"type": "boolean"}, "source": {"oneOf": [{"properties": {"bucket": {"type": "string"}, "endpoint": {"type": "string", "format": "uri", "nullable": true}, "keys": {"type": "array", "items": {"type": "string"}, "nullable": true}, "pathPrefix": {"type": "string", "nullable": true}, "vendor": {"type": "string", "enum": ["s3"]}}, "title": "S3SourceResponseSchema", "type": "object"}, {"properties": {"bucket": {"type": "string"}, "keys": {"type": "array", "items": {"type": "string"}, "nullable": true}, "pathPrefix": {"type": "string", "nullable": true}, "vendor": {"type": "string", "enum": ["gcs"]}}, "title": "GCSSourceResponseSchema", "type": "object"}, {"properties": {"bucket": {"type": "string"}, "jurisdiction": {"$ref": "#/components/schemas/r2-slurper_Jurisdiction"}, "keys": {"type": "array", "items": {"type": "string"}, "nullable": true}, "pathPrefix": {"type": "string", "nullable": true}, "vendor": {"type": "string", "enum": ["r2"]}}, "title": "R2SourceResponseSchema", "type": "object"}]}, "status": {"$ref": "#/components/schemas/r2-slurper_JobStatus"}, "target": {"type": "object", "properties": {"bucket": {"type": "string"}, "jurisdiction": {"$ref": "#/components/schemas/r2-slurper_Jurisdiction"}, "vendor": {"type": "string", "enum": ["r2"]}}, "title": "R2TargetResponseSchema"}}}
```
