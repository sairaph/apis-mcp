---
title: Check source connectivity
page_id: operation-put-accounts-account-id-slurper-source-connectivity-precheck-31e33a48
path: operations/r2-super-slurper
description: Check whether tokens are valid against the source bucket
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/slurper/source/connectivity-precheck
operation_ids:
    - slurper-check-source-connectivity
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Check source connectivity

`PUT /accounts/{account_id}/slurper/source/connectivity-precheck`

Operation ID: `slurper-check-source-connectivity`

Check whether tokens are valid against the source bucket

## Definition

```yaml
{"operationId": "slurper-check-source-connectivity", "summary": "Check source connectivity", "description": "Check whether tokens are valid against the source bucket", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_SourceJobSchema"}}}}, "responses": {"200": {"description": "Source connectivity checked", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/r2-slurper_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-slurper_ConnectivityResponse"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["R2 Super Slurper"]}
```
