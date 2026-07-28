---
title: Retrieve data about a specific document fingerprint.
page_id: operation-get-accounts-account-id-dlp-document-fingerprints-document-fingerprint-i-c3c63bbe
path: operations/dlp-document-fingerprints
description: Gets a document fingerprint and its latest upload status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/document_fingerprints/{document_fingerprint_id}
operation_ids:
    - dlp-document-fingerprints-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve data about a specific document fingerprint.

`GET /accounts/{account_id}/dlp/document_fingerprints/{document_fingerprint_id}`

Operation ID: `dlp-document-fingerprints-read`

Gets a document fingerprint and its latest upload status.

## Definition

```yaml
{"operationId": "dlp-document-fingerprints-read", "summary": "Retrieve data about a specific document fingerprint.", "description": "Gets a document fingerprint and its latest upload status.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "document_fingerprint_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Document fingerprint read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DocumentFingerprint"}}, "type": "object"}]}}}}, "4XX": {"description": "Document fingerprint read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Document Fingerprints"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
