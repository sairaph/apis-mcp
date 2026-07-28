---
title: Update Target Environment
page_id: operation-put-accounts-account-id-vuln-scanner-target-environments-target-environm-c3762b1d
path: operations/target-environments
description: Replaces a target environment. All fields must be provided.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/target_environments/{target_environment_id}
operation_ids:
    - update-target-environment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Target Environment

`PUT /accounts/{account_id}/vuln_scanner/target_environments/{target_environment_id}`

Operation ID: `update-target-environment`

Replaces a target environment. All fields must be provided.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_target_environment_id"}]
```

## Definition

```yaml
{"operationId": "update-target-environment", "summary": "Update Target Environment", "description": "Replaces a target environment. All fields must be provided.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vuln_scanner_update-target-environment-request"}}}}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/vuln_scanner_target-environment"}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Target Environments"]}
```
