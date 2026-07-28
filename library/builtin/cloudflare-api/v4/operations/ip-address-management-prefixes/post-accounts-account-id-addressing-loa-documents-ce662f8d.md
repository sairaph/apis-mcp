---
title: Upload LOA Document
page_id: operation-post-accounts-account-id-addressing-loa-documents-acc7f259
path: operations/ip-address-management-prefixes
description: Submit LOA document (pdf format) under the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/addressing/loa_documents
operation_ids:
    - ip-address-management-prefixes-upload-loa-document
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload LOA Document

`POST /accounts/{account_id}/addressing/loa_documents`

Operation ID: `ip-address-management-prefixes-upload-loa-document`

Submit LOA document (pdf format) under the account.

## Definition

```yaml
{"operationId": "ip-address-management-prefixes-upload-loa-document", "summary": "Upload LOA Document", "description": "Submit LOA document (pdf format) under the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/addressing_account_identifier"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"schema": {"type": "object", "properties": {"loa_document": {"description": "LOA document to upload.", "type": "string", "example": "@document.pdf"}}, "required": ["loa_document"]}}}}, "responses": {"201": {"description": "Upload LOA Document response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/addressing_loa_upload_response"}}}}, "4XX": {"description": "Upload LOA Document response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/addressing_loa_upload_response"}, {"$ref": "#/components/schemas/addressing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["IP Address Management Prefixes"], "x-api-token-group": ["Magic Transit Write"]}
```
