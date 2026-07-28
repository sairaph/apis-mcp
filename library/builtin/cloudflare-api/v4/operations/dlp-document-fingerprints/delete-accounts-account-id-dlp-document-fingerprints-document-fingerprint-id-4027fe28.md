---
title: Delete a single document fingerprint.
page_id: operation-delete-accounts-account-id-dlp-document-fingerprints-document-fingerprin-14be97ac
path: operations/dlp-document-fingerprints
description: Removes a document fingerprint from DLP configuration. Documents matching this fingerprint will no longer be detected.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dlp/document_fingerprints/{document_fingerprint_id}
operation_ids:
    - dlp-document-fingerprints-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a single document fingerprint.

`DELETE /accounts/{account_id}/dlp/document_fingerprints/{document_fingerprint_id}`

Operation ID: `dlp-document-fingerprints-delete`

Removes a document fingerprint from DLP configuration. Documents matching this fingerprint will no longer be detected.

## Definition

```yaml
{"operationId": "dlp-document-fingerprints-delete", "summary": "Delete a single document fingerprint.", "description": "Removes a document fingerprint from DLP configuration. Documents matching this fingerprint will no longer be detected.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "document_fingerprint_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Document fingerprint delete was successful."}, "4XX": {"description": "Document fingerprint delete failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Document Fingerprints"], "x-api-token-group": ["Zero Trust Write"]}
```
