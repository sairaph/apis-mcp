---
title: images_sourcingkit_migration_progress_response
page_id: schema-images-sourcingkit-migration-progress-response-506df436
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_sourcingkit_migration_progress_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/images_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"endedAt": {"type": "string", "format": "date-time", "nullable": true}, "importErrorsCount": {"description": "Objects that failed to import.", "type": "integer"}, "importSuccessCount": {"description": "Objects successfully imported.", "type": "integer"}, "isAccurate": {"description": "Whether the counts are accurate or estimated.", "type": "boolean"}, "scannedCount": {"description": "Total number of objects scanned in the source.", "type": "integer"}, "skippedConflictCount": {"description": "Objects skipped due to existing images at destination.", "type": "integer"}, "skippedExcludedContentTypeCount": {"description": "Objects skipped because their content type was excluded.", "type": "integer"}, "skippedInvalidMediaCount": {"description": "Objects skipped because the media could not be decoded.", "type": "integer"}, "skippedInvalidNameCount": {"description": "Objects skipped because their name is not valid.", "type": "integer"}, "skippedOversizeCount": {"description": "Objects skipped because they exceed the size limit.", "type": "integer"}, "skippedStorageClassCount": {"description": "Objects skipped due to unsupported storage class.", "type": "integer"}, "skippedUnsupportedContentTypeCount": {"description": "Objects skipped because their content type is not supported.", "type": "integer"}, "startedAt": {"type": "string", "format": "date-time", "nullable": true}, "status": {"$ref": "#/components/schemas/images_sourcingkit_migration_status"}}}}}]}
```
