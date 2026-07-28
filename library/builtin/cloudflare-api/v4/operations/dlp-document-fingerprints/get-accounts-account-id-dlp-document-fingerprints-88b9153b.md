---
title: Retrieve data about all document fingerprints.
page_id: operation-get-accounts-account-id-dlp-document-fingerprints-0ff1550c
path: operations/dlp-document-fingerprints
description: Lists all document fingerprints configured for DLP scanning in the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/document_fingerprints
operation_ids:
    - dlp-document-fingerprints-read-all
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve data about all document fingerprints.

`GET /accounts/{account_id}/dlp/document_fingerprints`

Operation ID: `dlp-document-fingerprints-read-all`

Lists all document fingerprints configured for DLP scanning in the account.

## Definition

```yaml
{"operationId": "dlp-document-fingerprints-read-all", "summary": "Retrieve data about all document fingerprints.", "description": "Lists all document fingerprints configured for DLP scanning in the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Document fingerprint read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DocumentFingerprintArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Document fingerprint read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Document Fingerprints"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
