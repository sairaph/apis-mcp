---
title: Create Credential
page_id: operation-post-accounts-account-id-vuln-scanner-credential-sets-credential-set-id-c5e59f2d
path: operations/credentials
description: Creates a new credential within a credential set.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}/credentials
operation_ids:
    - create-credential
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Credential

`POST /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}/credentials`

Operation ID: `create-credential`

Creates a new credential within a credential set.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_credential_set_id"}]
```

## Definition

```yaml
{"operationId": "create-credential", "summary": "Create Credential", "description": "Creates a new credential within a credential set.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vuln_scanner_create-credential-request"}}}}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/vuln_scanner_credential"}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Credentials"]}
```
