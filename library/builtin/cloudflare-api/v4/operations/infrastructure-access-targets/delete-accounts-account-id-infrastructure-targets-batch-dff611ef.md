---
title: Delete targets (Deprecated)
page_id: operation-delete-accounts-account-id-infrastructure-targets-batch-0cadab8b
path: operations/infrastructure-access-targets
description: Removes one or more targets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/infrastructure/targets/batch
operation_ids:
    - infra-targets-delete-batch
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete targets (Deprecated)

`DELETE /accounts/{account_id}/infrastructure/targets/batch`

Operation ID: `infra-targets-delete-batch`

Removes one or more targets.

## Definition

```yaml
{"operationId": "infra-targets-delete-batch", "summary": "Delete targets (Deprecated)", "description": "Removes one or more targets.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/infra_AccountTag"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"target_ids": {"description": "List of target IDs to bulk delete", "type": "array", "items": {"$ref": "#/components/schemas/infra_TargetId"}}}, "required": ["target_ids"]}}}}, "responses": {"200": {"description": "Successfully deleted the targets"}, "4XX": {"description": "Failed to delete the targets", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/infra_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Infrastructure Access Targets"]}
```
