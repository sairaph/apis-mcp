---
title: Purge build cache
page_id: operation-post-accounts-account-id-builds-triggers-trigger-uuid-purge-build-cache-dcdb1836
path: operations/triggers
description: Clear the build cache for a specific trigger
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/builds/triggers/{trigger_uuid}/purge_build_cache
operation_ids:
    - purgeBuildCache
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Purge build cache

`POST /accounts/{account_id}/builds/triggers/{trigger_uuid}/purge_build_cache`

Operation ID: `purgeBuildCache`

Clear the build cache for a specific trigger

## Definition

```yaml
{"operationId": "purgeBuildCache", "summary": "Purge build cache", "description": "Clear the build cache for a specific trigger", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_TriggerUuid"}], "responses": {"200": {"$ref": "#/components/responses/builds_SuccessEmpty"}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Triggers"], "x-api-token-group": ["Workers CI Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.triggers", "x-fern-sdk-method-name": "purge-cache", "x-forge-hidden": true, "x-forge-require-confirmation": "This operation wipes the workers-build cache."}
```
