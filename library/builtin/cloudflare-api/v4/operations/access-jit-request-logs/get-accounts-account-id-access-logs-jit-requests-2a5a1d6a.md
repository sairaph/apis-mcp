---
title: List Access JIT request logs
page_id: operation-get-accounts-account-id-access-logs-jit-requests-52c264c7
path: operations/access-jit-request-logs
description: Lists account-wide Access JIT request logs reconstructed from request lifecycle events.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/logs/jit_requests
operation_ids:
    - access-jit-request-logs-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Access JIT request logs

`GET /accounts/{account_id}/access/logs/jit_requests`

Operation ID: `access-jit-request-logs-list`

Lists account-wide Access JIT request logs reconstructed from request lifecycle events.

## Definition

```yaml
{"operationId": "access-jit-request-logs-list", "summary": "List Access JIT request logs", "description": "Lists account-wide Access JIT request logs reconstructed from request lifecycle events.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"$ref": "#/components/parameters/access_page"}, {"name": "per_page", "in": "query", "schema": {"description": "Number of results per page.", "type": "integer", "default": 25, "maximum": 1000}}, {"name": "status", "in": "query", "schema": {"$ref": "#/components/schemas/access_jit_request_status"}}, {"name": "search", "in": "query", "description": "Case-insensitive search over request ID, requester email, application audience, and application hostname.", "schema": {"type": "string"}}, {"name": "since", "in": "query", "description": "The earliest request timestamp to query. Defaults to 366 days before the current time.", "schema": {"type": "string", "format": "date-time"}}, {"name": "until", "in": "query", "description": "The latest request timestamp to query. Defaults to the current time.", "schema": {"type": "string", "format": "date-time"}}], "responses": {"200": {"description": "Access JIT request logs response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-12"}}}}, "4XX": {"description": "Access JIT request logs response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Access JIT request logs"]}
```
