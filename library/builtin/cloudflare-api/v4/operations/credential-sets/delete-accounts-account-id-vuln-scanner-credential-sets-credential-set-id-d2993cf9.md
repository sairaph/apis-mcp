---
title: Delete Credential Set
page_id: operation-delete-accounts-account-id-vuln-scanner-credential-sets-credential-set-i-bb8b2ab5
path: operations/credential-sets
description: Deletes a credential set and all of its credentials.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}
operation_ids:
    - delete-credential-set
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Credential Set

`DELETE /accounts/{account_id}/vuln_scanner/credential_sets/{credential_set_id}`

Operation ID: `delete-credential-set`

Deletes a credential set and all of its credentials.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/vuln_scanner_account_id"}, {"$ref": "#/components/parameters/vuln_scanner_credential_set_id"}]
```

## Definition

```yaml
{"operationId": "delete-credential-set", "summary": "Delete Credential Set", "description": "Deletes a credential set and all of its credentials.", "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/vuln_scanner_empty-response"}}}}, "4XX": {"$ref": "#/components/responses/vuln_scanner_4XX"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Credential Sets"]}
```
