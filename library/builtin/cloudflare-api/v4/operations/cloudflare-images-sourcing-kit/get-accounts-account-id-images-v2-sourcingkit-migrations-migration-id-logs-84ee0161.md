---
title: List migration logs
page_id: operation-get-accounts-account-id-images-v2-sourcingkit-migrations-migration-id-lo-fe22d1bb
path: operations/cloudflare-images-sourcing-kit
description: List log entries for a specific migration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/migrations/{migration_id}/logs
operation_ids:
    - cloudflare-images-sourcingkit-list-migration-logs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List migration logs

`GET /accounts/{account_id}/images/v2/sourcingkit/migrations/{migration_id}/logs`

Operation ID: `cloudflare-images-sourcingkit-list-migration-logs`

List log entries for a specific migration.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-list-migration-logs", "summary": "List migration logs", "description": "List log entries for a specific migration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}, {"name": "migration_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_sourcingkit_identifier"}}, {"name": "offset", "in": "query", "schema": {"description": "Number of items to skip before returning results.", "type": "integer", "default": 0, "minimum": 0}}, {"name": "limit", "in": "query", "schema": {"description": "Maximum number of items to return.", "type": "integer", "default": 25, "maximum": 100, "minimum": 1}}], "responses": {"200": {"description": "List migration logs response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_migration_log_list_response"}}}}, "4XX": {"description": "List migration logs response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_sourcingkit_migration_log_list_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.migrations.lifecycle", "x-fern-sdk-method-name": "logs", "x-forge-hidden": true}
```
