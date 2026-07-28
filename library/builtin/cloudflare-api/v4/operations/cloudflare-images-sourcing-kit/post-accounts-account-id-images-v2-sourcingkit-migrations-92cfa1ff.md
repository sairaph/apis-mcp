---
title: Create a sourcing kit migration
page_id: operation-post-accounts-account-id-images-v2-sourcingkit-migrations-d7c5209b
path: operations/cloudflare-images-sourcing-kit
description: |-
    Create a new migration from an existing source. The migration will import
    objects from the source bucket into Cloudflare Images.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/migrations
operation_ids:
    - cloudflare-images-sourcingkit-create-migration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a sourcing kit migration

`POST /accounts/{account_id}/images/v2/sourcingkit/migrations`

Operation ID: `cloudflare-images-sourcingkit-create-migration`

Create a new migration from an existing source. The migration will import
objects from the source bucket into Cloudflare Images.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-create-migration", "summary": "Create a sourcing kit migration", "description": "Create a new migration from an existing source. The migration will import\nobjects from the source bucket into Cloudflare Images.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_migration_create_request"}}}}, "responses": {"200": {"description": "Create sourcing kit migration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_migration_create_response"}}}}, "4XX": {"description": "Create sourcing kit migration response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_sourcingkit_migration_create_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.migrations", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
