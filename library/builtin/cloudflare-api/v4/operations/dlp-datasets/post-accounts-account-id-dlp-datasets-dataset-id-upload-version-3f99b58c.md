---
title: Upload a new version of a dataset
page_id: operation-post-accounts-account-id-dlp-datasets-dataset-id-upload-version-0bd9434a
path: operations/dlp-datasets
description: |-
    This is used for single-column EDMv1 and Custom Word Lists. The EDM format
    can only be created in the Cloudflare dashboard. For other clients, this
    operation can only be used for non-secret Custom Word Lists. The body must
    be a UTF-8 encoded, newline (NL or CRNL) separated list of words to be matched.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/datasets/{dataset_id}/upload/{version}
operation_ids:
    - dlp-datasets-upload-version
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload a new version of a dataset

`POST /accounts/{account_id}/dlp/datasets/{dataset_id}/upload/{version}`

Operation ID: `dlp-datasets-upload-version`

This is used for single-column EDMv1 and Custom Word Lists. The EDM format
can only be created in the Cloudflare dashboard. For other clients, this
operation can only be used for non-secret Custom Word Lists. The body must
be a UTF-8 encoded, newline (NL or CRNL) separated list of words to be matched.

## Definition

```yaml
{"operationId": "dlp-datasets-upload-version", "summary": "Upload a new version of a dataset", "description": "This is used for single-column EDMv1 and Custom Word Lists. The EDM format\ncan only be created in the Cloudflare dashboard. For other clients, this\noperation can only be used for non-secret Custom Word Lists. The body must\nbe a UTF-8 encoded, newline (NL or CRNL) separated list of words to be matched.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "dataset_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "version", "in": "path", "required": true, "schema": {"type": "integer", "format": "int64"}}], "requestBody": {"description": "Dataset. For custom wordlists this contains UTF-8 patterns separated by newline characters.", "required": true, "content": {"application/octet-stream": {"schema": {"type": "string", "format": "binary"}}}}, "responses": {"200": {"description": "Dataset version uploaded successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Dataset"}}, "type": "object"}]}}}}, "4XX": {"description": "Dataset version upload failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Datasets"], "x-api-token-group": ["Zero Trust Write"]}
```
