---
title: Create Credential Set
page_id: operation-post-accounts-account-id-vuln-scanner-credential-sets-a6eda6e2
path: operations/credential-sets
description: Creates a new credential set.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/credential_sets
operation_ids:
    - create-credential-set
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Credential Set

`POST /accounts/{account_id}/vuln_scanner/credential_sets`

Operation ID: `create-credential-set`

Creates a new credential set.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}]
```

## Definition

```yaml
{"operationId": "create-credential-set", "summary": "Create Credential Set", "description": "Creates a new credential set.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vuln_scanner_create-credential-set-request"}}}}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/vuln_scanner_credential-set"}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Credential Sets"]}
```
