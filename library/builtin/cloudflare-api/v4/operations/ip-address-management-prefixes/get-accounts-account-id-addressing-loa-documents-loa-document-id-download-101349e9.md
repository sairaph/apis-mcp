---
title: Download LOA Document
page_id: operation-get-accounts-account-id-addressing-loa-documents-loa-document-id-downloa-745b38aa
path: operations/ip-address-management-prefixes
description: Download specified LOA document under the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/addressing/loa_documents/{loa_document_id}/download
operation_ids:
    - ip-address-management-prefixes-download-loa-document
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Download LOA Document

`GET /accounts/{account_id}/addressing/loa_documents/{loa_document_id}/download`

Operation ID: `ip-address-management-prefixes-download-loa-document`

Download specified LOA document under the account.

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-download-loa-document", "summary": "Download LOA Document", "description": "Download specified LOA document under the account.", "parameters": [{"name": "loa_document_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_loa_document_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "responses": {"200": {"description": "Download LOA Document response", "content": {"application/pdf": {"schema": {"type": "string", "format": "binary"}}}}, "4XX": {"description": "Download LOA Document response failure", "content": {"application/json": {"schema": {"allOf": [{}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Prefixes"], "x-api-token-group": ["Magic Transit Read", "Magic Transit Write"]}
```
