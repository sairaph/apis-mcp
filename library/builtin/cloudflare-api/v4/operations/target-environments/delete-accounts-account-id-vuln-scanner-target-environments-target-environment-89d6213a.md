---
title: Delete Target Environment
page_id: operation-delete-accounts-account-id-vuln-scanner-target-environments-target-envir-170cae98
path: operations/target-environments
description: Removes a target environment.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/target_environments/{target_environment_id}
operation_ids:
    - delete-target-environment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Target Environment

`DELETE /accounts/{account_id}/vuln_scanner/target_environments/{target_environment_id}`

Operation ID: `delete-target-environment`

Removes a target environment.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_target_environment_id"}]
```

## Definition

```yaml
{"operationId": "delete-target-environment", "summary": "Delete Target Environment", "description": "Removes a target environment.", "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vuln_scanner_empty-response"}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Target Environments"]}
```
