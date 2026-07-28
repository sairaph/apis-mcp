---
title: Attach Domain
page_id: operation-put-accounts-account-id-workers-domains-96a40e9b
path: operations/domains
description: Attaches a domain that routes traffic to a Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/workers/domains
operation_ids:
    - workers.domains.update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Attach Domain

`PUT /accounts/{account_id}/workers/domains`

Operation ID: `workers.domains.update`

Attaches a domain that routes traffic to a Worker.

## Definition

```yaml
{"operationId": "workers.domains.update", "summary": "Attach Domain", "description": "Attaches a domain that routes traffic to a Worker.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_Domain"}, {"required": ["hostname", "service"], "type": "object"}]}}}}, "responses": {"200": {"description": "Attach domain response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_Domain"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Attach domain failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Domains"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.domains", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
