---
title: Delete a sourcing kit migration
page_id: operation-delete-accounts-account-id-images-v2-sourcingkit-migrations-migration-id-3c59d750
path: operations/cloudflare-images-sourcing-kit
description: Delete an existing migration. Only completed, errored, or aborted migrations can be deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/migrations/{migration_id}
operation_ids:
    - cloudflare-images-sourcingkit-delete-migration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a sourcing kit migration

`DELETE /accounts/{account_id}/images/v2/sourcingkit/migrations/{migration_id}`

Operation ID: `cloudflare-images-sourcingkit-delete-migration`

Delete an existing migration. Only completed, errored, or aborted migrations can be deleted.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-delete-migration", "summary": "Delete a sourcing kit migration", "description": "Delete an existing migration. Only completed, errored, or aborted migrations can be deleted.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}, {"name": "migration_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_sourcingkit_identifier"}}], "responses": {"200": {"description": "Delete sourcing kit migration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_deleted_response"}}}}, "4XX": {"description": "Delete sourcing kit migration response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_deleted_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.migrations", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
