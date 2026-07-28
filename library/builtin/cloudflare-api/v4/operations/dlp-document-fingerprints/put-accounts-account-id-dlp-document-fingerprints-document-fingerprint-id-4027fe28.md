---
title: Uploads a new version for a document fingerprint.
page_id: operation-put-accounts-account-id-dlp-document-fingerprints-document-fingerprint-i-719f7deb
path: operations/dlp-document-fingerprints
description: Uploads a new document to create or update a fingerprint. The document structure is analyzed to enable detection of similar documents.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/document_fingerprints/{document_fingerprint_id}
operation_ids:
    - dlp-document-fingerprints-upload
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Uploads a new version for a document fingerprint.

`PUT /accounts/{account_id}/dlp/document_fingerprints/{document_fingerprint_id}`

Operation ID: `dlp-document-fingerprints-upload`

Uploads a new document to create or update a fingerprint. The document structure is analyzed to enable detection of similar documents.

## Definition

```yaml
{"operationId": "dlp-document-fingerprints-upload", "summary": "Uploads a new version for a document fingerprint.", "description": "Uploads a new document to create or update a fingerprint. The document structure is analyzed to enable detection of similar documents.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "document_fingerprint_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "File used for document fingerprinting.", "required": true, "content": {"multipart/form-data": {"schema": {"type": "object", "properties": {"file": {"type": "string", "format": "binary"}}, "required": ["file"]}}}}, "responses": {"200": {"description": "File uploaded successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DocumentFingerprintUpload"}}, "type": "object"}]}}}}, "4XX": {"description": "Failed to upload file.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Document Fingerprints"], "x-api-token-group": ["Zero Trust Write"]}
```
