---
title: List Credential Sets
page_id: operation-get-accounts-account-id-vuln-scanner-credential-sets-1a1b4f8b
path: operations/credential-sets
description: Returns all credential sets for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/credential_sets
operation_ids:
    - list-credential-sets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Credential Sets

`GET /accounts/{account_id}/vuln_scanner/credential_sets`

Operation ID: `list-credential-sets`

Returns all credential sets for the account.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}]
```

## Definition

```yaml
{"operationId": "list-credential-sets", "summary": "List Credential Sets", "description": "Returns all credential sets for the account.", "parameters": [{"$ref": "#/components/parameters/vuln_scanner_page"}, {"$ref": "#/components/parameters/vuln_scanner_per_page"}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/vuln_scanner_credential-set"}}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Credential Sets"]}
```
