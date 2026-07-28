---
title: Fetch all datasets
page_id: operation-get-accounts-account-id-dlp-datasets-b9975af6
path: operations/dlp-datasets
description: Lists all DLP datasets configured for the account, including custom word lists and EDM datasets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/datasets
operation_ids:
    - dlp-datasets-read-all
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch all datasets

`GET /accounts/{account_id}/dlp/datasets`

Operation ID: `dlp-datasets-read-all`

Lists all DLP datasets configured for the account, including custom word lists and EDM datasets.

## Definition

```yaml
{"operationId": "dlp-datasets-read-all", "summary": "Fetch all datasets", "description": "Lists all DLP datasets configured for the account, including custom word lists and EDM datasets.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Datasets read successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DatasetArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Datasets read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Datasets"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
