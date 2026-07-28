---
title: Update Credential
page_id: operation-put-accounts-account-id-vuln-scanner-credential-sets-credential-set-id-c-de75d8e8
path: operations/credentials
description: Replaces a credential. All fields must be provided.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}/credentials/{credential_id}
operation_ids:
    - update-credential
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Credential

`PUT /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}/credentials/{credential_id}`

Operation ID: `update-credential`

Replaces a credential. All fields must be provided.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_credential_set_id"}, {"$ref": "#/components/parameters/vuln_scanner_credential_id"}]
```

## Definition

```yaml
{"operationId": "update-credential", "summary": "Update Credential", "description": "Replaces a credential. All fields must be provided.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vuln_scanner_update-credential-request"}}}}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/vuln_scanner_credential"}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Credentials"]}
```
