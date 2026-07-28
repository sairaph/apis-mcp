---
title: List Target Environments
page_id: operation-get-accounts-account-id-vuln-scanner-target-environments-3cc44839
path: operations/target-environments
description: Returns all target environments for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/target_environments
operation_ids:
    - list-target-environments
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Target Environments

`GET /accounts/{account_id}/vuln_scanner/target_environments`

Operation ID: `list-target-environments`

Returns all target environments for the account.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}]
```

## Definition

```yaml
{"operationId": "list-target-environments", "summary": "List Target Environments", "description": "Returns all target environments for the account.", "parameters": [{"$ref": "#/components/parameters/vuln_scanner_page"}, {"$ref": "#/components/parameters/vuln_scanner_per_page"}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/vuln_scanner_target-environment"}}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Target Environments"]}
```
