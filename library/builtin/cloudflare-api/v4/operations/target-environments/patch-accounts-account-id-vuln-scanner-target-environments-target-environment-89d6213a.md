---
title: Edit Target Environment
page_id: operation-patch-accounts-account-id-vuln-scanner-target-environments-target-enviro-2e00b385
path: operations/target-environments
description: Updates a target environment with only the provided fields; omitted fields remain unchanged.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/target_environments/{target_environment_id}
operation_ids:
    - edit-target-environment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit Target Environment

`PATCH /accounts/{account_id}/vuln_scanner/target_environments/{target_environment_id}`

Operation ID: `edit-target-environment`

Updates a target environment with only the provided fields; omitted fields remain unchanged.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_target_environment_id"}]
```

## Definition

```yaml
{"operationId": "edit-target-environment", "summary": "Edit Target Environment", "description": "Updates a target environment with only the provided fields; omitted fields remain unchanged.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vuln_scanner_patch-target-environment-request"}}}}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/vuln_scanner_target-environment"}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Target Environments"]}
```
