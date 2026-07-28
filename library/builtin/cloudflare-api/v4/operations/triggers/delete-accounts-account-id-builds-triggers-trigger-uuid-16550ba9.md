---
title: Delete trigger
page_id: operation-delete-accounts-account-id-builds-triggers-trigger-uuid-8b6dad62
path: operations/triggers
description: Remove a CI/CD trigger
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/builds/triggers/{trigger_uuid}
operation_ids:
    - deleteTrigger
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete trigger

`DELETE /accounts/{account_id}/builds/triggers/{trigger_uuid}`

Operation ID: `deleteTrigger`

Remove a CI/CD trigger

## Definition

```yaml
{"operationId": "deleteTrigger", "summary": "Delete trigger", "description": "Remove a CI/CD trigger", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_TriggerUuid"}], "responses": {"200": {"$ref": "#/components/responses/builds_SuccessEmpty"}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Triggers"], "x-api-token-group": ["Workers CI Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.triggers", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
