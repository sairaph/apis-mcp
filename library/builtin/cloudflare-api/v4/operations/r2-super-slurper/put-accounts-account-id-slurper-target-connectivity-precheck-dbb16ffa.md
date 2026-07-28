---
title: Check target connectivity
page_id: operation-put-accounts-account-id-slurper-target-connectivity-precheck-6d245e24
path: operations/r2-super-slurper
description: Check whether tokens are valid against the target bucket
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/slurper/target/connectivity-precheck
operation_ids:
    - slurper-check-target-connectivity
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Check target connectivity

`PUT /accounts/{account_id}/slurper/target/connectivity-precheck`

Operation ID: `slurper-check-target-connectivity`

Check whether tokens are valid against the target bucket

## Definition

```yaml
{"operationId": "slurper-check-target-connectivity", "summary": "Check target connectivity", "description": "Check whether tokens are valid against the target bucket", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_R2TargetSchema"}}}}, "responses": {"200": {"description": "Target connectivity checked", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/r2-slurper_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/r2-slurper_ConnectivityResponse"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2-slurper_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["R2 Super Slurper"]}
```
