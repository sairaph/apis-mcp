---
title: Get Target Environment
page_id: operation-get-accounts-account-id-vuln-scanner-target-environments-target-environm-2261b785
path: operations/target-environments
description: Returns a single target environment by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/target_environments/{target_environment_id}
operation_ids:
    - get-target-environment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Target Environment

`GET /accounts/{account_id}/vuln_scanner/target_environments/{target_environment_id}`

Operation ID: `get-target-environment`

Returns a single target environment by ID.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_target_environment_id"}]
```

## Definition

```yaml
{"operationId": "get-target-environment", "summary": "Get Target Environment", "description": "Returns a single target environment by ID.", "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/vuln_scanner_target-environment"}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Target Environments"]}
```
