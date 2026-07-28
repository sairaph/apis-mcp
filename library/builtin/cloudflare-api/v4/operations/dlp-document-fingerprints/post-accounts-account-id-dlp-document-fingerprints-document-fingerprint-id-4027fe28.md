---
title: Update the attributes of a single document fingerprint.
page_id: operation-post-accounts-account-id-dlp-document-fingerprints-document-fingerprint-baa41b01
path: operations/dlp-document-fingerprints
description: Updates metadata for an existing document fingerprint, such as its name or description.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/document_fingerprints/{document_fingerprint_id}
operation_ids:
    - dlp-document-fingerprints-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update the attributes of a single document fingerprint.

`POST /accounts/{account_id}/dlp/document_fingerprints/{document_fingerprint_id}`

Operation ID: `dlp-document-fingerprints-update`

Updates metadata for an existing document fingerprint, such as its name or description.

## Definition

```yaml
{"operationId": "dlp-document-fingerprints-update", "summary": "Update the attributes of a single document fingerprint.", "description": "Updates metadata for an existing document fingerprint, such as its name or description.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "document_fingerprint_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Attributes of the document fingerprint to update.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_UpdateDocumentFingerprint"}}}}, "responses": {"200": {"description": "Document fingerprint read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DocumentFingerprint"}}, "type": "object"}]}}}}, "4XX": {"description": "Document fingerprint read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Document Fingerprints"], "x-api-token-group": ["Zero Trust Write"]}
```
