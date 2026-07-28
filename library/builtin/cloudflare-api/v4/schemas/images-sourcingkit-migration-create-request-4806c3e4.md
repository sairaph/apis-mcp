---
title: images_sourcingkit_migration_create_request
page_id: schema-images-sourcingkit-migration-create-request-4806c3e4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_sourcingkit_migration_create_request

```yaml
{"type": "object", "properties": {"conflictBehaviour": {"$ref": "#/components/schemas/images_sourcingkit_conflict_behaviour"}, "excludedContentTypes": {"description": "Content types to skip during migration.", "type": "array", "items": {"$ref": "#/components/schemas/images_sourcingkit_acceptable_content_type"}}, "pathPrefix": {"description": "Prefix to prepend to image custom IDs.", "type": "string", "maxLength": 128, "minLength": 1}, "rootDirectory": {"description": "Only import objects under this prefix in the source bucket.", "type": "string", "maxLength": 128, "minLength": 1}, "sourceId": {"description": "The identifier of the source to migrate from.", "type": "string", "format": "uuid"}}, "required": ["sourceId"]}
```
