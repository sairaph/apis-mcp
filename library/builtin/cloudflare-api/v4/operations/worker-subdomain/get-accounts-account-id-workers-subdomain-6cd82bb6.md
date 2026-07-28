---
title: Get Subdomain
page_id: operation-get-accounts-account-id-workers-subdomain-b63bc18d
path: operations/worker-subdomain
description: Returns a Workers subdomain for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/subdomain
operation_ids:
    - worker-subdomain-get-subdomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Subdomain

`GET /accounts/{account_id}/workers/subdomain`

Operation ID: `worker-subdomain-get-subdomain`

Returns a Workers subdomain for an account.

## Definition

```yaml
{"operationId": "worker-subdomain-get-subdomain", "summary": "Get Subdomain", "description": "Returns a Workers subdomain for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "responses": {"200": {"description": "Get Subdomain response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_subdomain-2"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get Subdomain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Subdomain"], "x-api-token-group": ["Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.subdomains", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
