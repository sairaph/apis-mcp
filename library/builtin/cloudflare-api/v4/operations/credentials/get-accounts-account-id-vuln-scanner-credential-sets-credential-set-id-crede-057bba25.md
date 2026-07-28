---
title: List Credentials
page_id: operation-get-accounts-account-id-vuln-scanner-credential-sets-credential-set-id-c-46e8e622
path: operations/credentials
description: Returns all credentials within a credential set.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}/credentials
operation_ids:
    - list-credentials
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Credentials

`GET /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}/credentials`

Operation ID: `list-credentials`

Returns all credentials within a credential set.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_credential_set_id"}]
```

## Definition

```yaml
{"operationId": "list-credentials", "summary": "List Credentials", "description": "Returns all credentials within a credential set.", "parameters": [{"$ref": "#/components/parameters/vuln_scanner_page"}, {"$ref": "#/components/parameters/vuln_scanner_per_page"}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/vuln_scanner_credential"}}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Credentials"]}
```
