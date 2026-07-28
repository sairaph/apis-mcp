---
title: Create Subdomain
page_id: operation-put-accounts-account-id-workers-subdomain-2ba8678d
path: operations/worker-subdomain
description: Creates a Workers subdomain for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workers/subdomain
operation_ids:
    - worker-subdomain-create-subdomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Subdomain

`PUT /accounts/{account_id}/workers/subdomain`

Operation ID: `worker-subdomain-create-subdomain`

Creates a Workers subdomain for an account.

## Definition

```yaml
{"operationId": "worker-subdomain-create-subdomain", "summary": "Create Subdomain", "description": "Creates a Workers subdomain for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_subdomain-2"}}}}, "responses": {"200": {"description": "Create Subdomain response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_subdomain-2"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Create Subdomain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Subdomain"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.subdomains", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
