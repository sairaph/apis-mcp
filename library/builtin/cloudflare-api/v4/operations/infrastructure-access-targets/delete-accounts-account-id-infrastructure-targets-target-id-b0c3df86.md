---
title: Delete target
page_id: operation-delete-accounts-account-id-infrastructure-targets-target-id-b5a954fd
path: operations/infrastructure-access-targets
description: Delete target
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/infrastructure/targets/{target_id}
operation_ids:
    - infra-targets-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete target

`DELETE /accounts/{account_id}/infrastructure/targets/{target_id}`

Operation ID: `infra-targets-delete`

## Definition

```yaml
{"operationId": "infra-targets-delete", "summary": "Delete target", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/infra_AccountTag"}}, {"name": "target_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/infra_TargetId"}}], "responses": {"200": {"description": "Successfully deleted the target"}, "4XX": {"description": "Failed to delete the target", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/infra_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Infrastructure Access Targets"]}
```
