---
title: Get Credential Set
page_id: operation-get-accounts-account-id-vuln-scanner-credential-sets-credential-set-id-54915db1
path: operations/credential-sets
description: Returns a single credential set by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}
operation_ids:
    - get-credential-set
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Credential Set

`GET /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}`

Operation ID: `get-credential-set`

Returns a single credential set by ID.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_credential_set_id"}]
```

## Definition

```yaml
{"operationId": "get-credential-set", "summary": "Get Credential Set", "description": "Returns a single credential set by ID.", "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/vuln_scanner_credential-set"}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Credential Sets"]}
```
