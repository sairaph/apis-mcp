---
title: Creates a new document fingerprint.
page_id: operation-post-accounts-account-id-dlp-document-fingerprints-e5b95240
path: operations/dlp-document-fingerprints
description: Creates a new document fingerprint for DLP scanning. Document fingerprints detect documents that are structurally similar to the uploaded sample.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/document_fingerprints
operation_ids:
    - dlp-document-fingerprints-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates a new document fingerprint.

`POST /accounts/{account_id}/dlp/document_fingerprints`

Operation ID: `dlp-document-fingerprints-create`

Creates a new document fingerprint for DLP scanning. Document fingerprints detect documents that are structurally similar to the uploaded sample.

## Definition

```yaml
{"operationId": "dlp-document-fingerprints-create", "summary": "Creates a new document fingerprint.", "description": "Creates a new document fingerprint for DLP scanning. Document fingerprints detect documents that are structurally similar to the uploaded sample.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "Attributes of the new document fingerprint.", "required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"type": "string", "default": ""}, "match_percent": {"type": "integer", "format": "int32"}, "name": {"type": "string"}}, "required": ["name", "match_percent"]}}}}, "responses": {"200": {"description": "Document fingerprint created successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DocumentFingerprint"}}, "type": "object"}]}}}}, "4XX": {"description": "Document fingerprint creation failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Document Fingerprints"], "x-api-token-group": ["Zero Trust Write"]}
```
