---
title: Update Credential Set
page_id: operation-put-accounts-account-id-vuln-scanner-credential-sets-credential-set-id-180cd1b0
path: operations/credential-sets
description: Replaces a credential set. All fields must be provided.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}
operation_ids:
    - update-credential-set
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Credential Set

`PUT /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}`

Operation ID: `update-credential-set`

Replaces a credential set. All fields must be provided.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_credential_set_id"}]
```

## Definition

```yaml
{"operationId": "update-credential-set", "summary": "Update Credential Set", "description": "Replaces a credential set. All fields must be provided.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vuln_scanner_update-credential-set-request"}}}}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/vuln_scanner_credential-set"}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Credential Sets"]}
```
