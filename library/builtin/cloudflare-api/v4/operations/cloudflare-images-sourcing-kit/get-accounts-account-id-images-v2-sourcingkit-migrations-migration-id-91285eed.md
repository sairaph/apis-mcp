---
title: Get sourcing kit migration
page_id: operation-get-accounts-account-id-images-v2-sourcingkit-migrations-migration-id-5d64c56f
path: operations/cloudflare-images-sourcing-kit
description: Fetch details for a single migration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/migrations/{migration_id}
operation_ids:
    - cloudflare-images-sourcingkit-get-migration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get sourcing kit migration

`GET /accounts/{account_id}/images/v2/sourcingkit/migrations/{migration_id}`

Operation ID: `cloudflare-images-sourcingkit-get-migration`

Fetch details for a single migration.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-get-migration", "summary": "Get sourcing kit migration", "description": "Fetch details for a single migration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}, {"name": "migration_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_sourcingkit_identifier"}}], "responses": {"200": {"description": "Get sourcing kit migration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_migration_single_response"}}}}, "4XX": {"description": "Get sourcing kit migration response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_sourcingkit_migration_single_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.migrations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
