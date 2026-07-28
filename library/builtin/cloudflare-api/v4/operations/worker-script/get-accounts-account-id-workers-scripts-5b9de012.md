---
title: List Workers
page_id: operation-get-accounts-account-id-workers-scripts-49296461
path: operations/worker-script
description: Fetch a list of uploaded workers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts
operation_ids:
    - worker-script-list-workers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Workers

`GET /accounts/{account_id}/workers/scripts`

Operation ID: `worker-script-list-workers`

Fetch a list of uploaded workers.

## Definition

```yaml
{"operationId": "worker-script-list-workers", "summary": "List Workers", "description": "Fetch a list of uploaded workers.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "tags", "in": "query", "description": "Filter scripts by tags. Format: comma-separated list of tag:allowed pairs where allowed is 'yes' or 'no'.", "schema": {"type": "string", "example": "production:yes,staging:no"}}], "responses": {"200": {"description": "List Workers response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_script-response-collection"}}}}, "4XX": {"description": "List Workers response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Script"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
