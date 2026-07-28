---
title: Edit Credential
page_id: operation-patch-accounts-account-id-vuln-scanner-credential-sets-credential-set-id-0245be29
path: operations/credentials
description: Updates a credential with only the provided fields; omitted fields remain unchanged.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}/credentials/{credential_id}
operation_ids:
    - edit-credential
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit Credential

`PATCH /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}/credentials/{credential_id}`

Operation ID: `edit-credential`

Updates a credential with only the provided fields; omitted fields remain unchanged.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_credential_set_id"}, {"$ref": "#/components/parameters/vuln_scanner_credential_id"}]
```

## Definition

```yaml
{"operationId": "edit-credential", "summary": "Edit Credential", "description": "Updates a credential with only the provided fields; omitted fields remain unchanged.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vuln_scanner_patch-credential-request"}}}}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/vuln_scanner_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/vuln_scanner_credential"}, "result_info": {"type": "object", "nullable": true}}, "type": "object"}]}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Credentials"]}
```
