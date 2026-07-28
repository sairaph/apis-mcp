---
title: Start a migration
page_id: operation-patch-accounts-account-id-images-v2-sourcingkit-migrations-migration-id-cf878d6e
path: operations/cloudflare-images-sourcing-kit
description: Start a pending migration. The migration will begin importing objects from the configured source.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/migrations/{migration_id}/lifecycle/start
operation_ids:
    - cloudflare-images-sourcingkit-start-migration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Start a migration

`PATCH /accounts/{account_id}/images/v2/sourcingkit/migrations/{migration_id}/lifecycle/start`

Operation ID: `cloudflare-images-sourcingkit-start-migration`

Start a pending migration. The migration will begin importing objects from the configured source.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-start-migration", "summary": "Start a migration", "description": "Start a pending migration. The migration will begin importing objects from the configured source.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}, {"name": "migration_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_sourcingkit_identifier"}}], "responses": {"200": {"description": "Start migration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_api-response-single"}}}}, "4XX": {"description": "Start migration response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.migrations.lifecycle", "x-fern-sdk-method-name": "start", "x-forge-hidden": true}
```
