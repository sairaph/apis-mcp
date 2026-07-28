---
title: Delete Credential
page_id: operation-delete-accounts-account-id-vuln-scanner-credential-sets-credential-set-i-0c7d50e7
path: operations/credentials
description: Deletes a credential.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}/credentials/{credential_id}
operation_ids:
    - delete-credential
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Credential

`DELETE /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}/credentials/{credential_id}`

Operation ID: `delete-credential`

Deletes a credential.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_credential_set_id"}, {"$ref": "#/components/parameters/vuln_scanner_credential_id"}]
```

## Definition

```yaml
{"operationId": "delete-credential", "summary": "Delete Credential", "description": "Deletes a credential.", "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vuln_scanner_empty-response"}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Credentials"]}
```
