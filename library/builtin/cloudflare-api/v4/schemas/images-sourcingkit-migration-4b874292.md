---
title: images_sourcingkit_migration
page_id: schema-images-sourcingkit-migration-4b874292
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_sourcingkit_migration

```yaml
{"type": "object", "properties": {"conflictBehaviour": {"$ref": "#/components/schemas/images_sourcingkit_conflict_behaviour"}, "createdAt": {"type": "string", "format": "date-time", "readOnly": true}, "endedAt": {"type": "string", "format": "date-time", "nullable": true, "readOnly": true}, "excludedContentTypes": {"description": "Content types to skip during migration.", "type": "array", "items": {"$ref": "#/components/schemas/images_sourcingkit_acceptable_content_type"}, "nullable": true}, "id": {"description": "Migration unique identifier.", "type": "string", "format": "uuid", "readOnly": true}, "imagesPathPrefix": {"description": "Alias for pathPrefix (deprecated).", "type": "string", "nullable": true}, "logStopped": {"description": "Whether logging was stopped due to excessive log volume.", "type": "boolean", "readOnly": true}, "pathPrefix": {"description": "Prefix to prepend to image custom IDs.", "type": "string", "nullable": true}, "rootDirectory": {"description": "Only import objects under this prefix in the source bucket.", "type": "string", "nullable": true}, "sourceId": {"description": "The source this migration imports from.", "type": "string", "format": "uuid", "nullable": true}, "startedAt": {"type": "string", "format": "date-time", "nullable": true, "readOnly": true}, "status": {"$ref": "#/components/schemas/images_sourcingkit_migration_status"}}}
```
