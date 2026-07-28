---
title: Get migration progress
page_id: operation-get-accounts-account-id-images-v2-sourcingkit-migrations-migration-id-li-5c34dbdf
path: operations/cloudflare-images-sourcing-kit
description: |-
    Get the current progress of a migration including counts of scanned, imported,
    skipped, and errored objects.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/migrations/{migration_id}/lifecycle
operation_ids:
    - cloudflare-images-sourcingkit-get-migration-progress
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get migration progress

`GET /accounts/{account_id}/images/v2/sourcingkit/migrations/{migration_id}/lifecycle`

Operation ID: `cloudflare-images-sourcingkit-get-migration-progress`

Get the current progress of a migration including counts of scanned, imported,
skipped, and errored objects.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-get-migration-progress", "summary": "Get migration progress", "description": "Get the current progress of a migration including counts of scanned, imported,\nskipped, and errored objects.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}, {"name": "migration_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_sourcingkit_identifier"}}], "responses": {"200": {"description": "Migration progress response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_migration_progress_response"}}}}, "4XX": {"description": "Migration progress response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_sourcingkit_migration_progress_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.migrations.lifecycle", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
