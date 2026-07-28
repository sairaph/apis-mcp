---
title: Abort a migration
page_id: operation-patch-accounts-account-id-images-v2-sourcingkit-migrations-migration-id-5ca9b75a
path: operations/cloudflare-images-sourcing-kit
description: Abort a running migration. Objects already imported will not be removed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/migrations/{migration_id}/lifecycle/abort
operation_ids:
    - cloudflare-images-sourcingkit-abort-migration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Abort a migration

`PATCH /accounts/{account_id}/images/v2/sourcingkit/migrations/{migration_id}/lifecycle/abort`

Operation ID: `cloudflare-images-sourcingkit-abort-migration`

Abort a running migration. Objects already imported will not be removed.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-abort-migration", "summary": "Abort a migration", "description": "Abort a running migration. Objects already imported will not be removed.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}, {"name": "migration_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_sourcingkit_identifier"}}], "responses": {"200": {"description": "Abort migration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_api-response-single"}}}}, "4XX": {"description": "Abort migration response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.migrations.lifecycle", "x-fern-sdk-method-name": "abort", "x-forge-hidden": true}
```
