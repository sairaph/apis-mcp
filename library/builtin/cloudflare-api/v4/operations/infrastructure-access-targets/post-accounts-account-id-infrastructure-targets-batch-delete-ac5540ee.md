---
title: Delete targets
page_id: operation-post-accounts-account-id-infrastructure-targets-batch-delete-db1686cc
path: operations/infrastructure-access-targets
description: Removes one or more targets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/infrastructure/targets/batch_delete
operation_ids:
    - infra-targets-delete-batch-post
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete targets

`POST /accounts/{account_id}/infrastructure/targets/batch_delete`

Operation ID: `infra-targets-delete-batch-post`

Removes one or more targets.

## Definition

```yaml
{"operationId": "infra-targets-delete-batch-post", "summary": "Delete targets", "description": "Removes one or more targets.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/infra_AccountTag"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"target_ids": {"description": "List of target IDs to bulk delete", "type": "array", "items": {"$ref": "#/components/schemas/infra_TargetId"}}}, "required": ["target_ids"]}}}}, "responses": {"200": {"description": "Successfully deleted the targets"}, "4XX": {"description": "Failed to delete the targets", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/infra_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Infrastructure Access Targets"]}
```
