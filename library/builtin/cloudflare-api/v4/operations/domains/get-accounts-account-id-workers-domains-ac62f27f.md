---
title: List Domains
page_id: operation-get-accounts-account-id-workers-domains-e6e29e64
path: operations/domains
description: Lists all domains for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/domains
operation_ids:
    - workers.domains.list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Domains

`GET /accounts/{account_id}/workers/domains`

Operation ID: `workers.domains.list`

Lists all domains for an account.

## Definition

```yaml
{"operationId": "workers.domains.list", "summary": "List Domains", "description": "Lists all domains for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "zone_id", "in": "query", "schema": {"description": "ID of the zone containing the domain hostname.", "type": "string", "example": "593c9c94de529bbbfaac7c53ced0447d"}}, {"name": "zone_name", "in": "query", "schema": {"description": "Name of the zone containing the domain hostname.", "type": "string", "example": "example.com"}}, {"name": "service", "in": "query", "schema": {"description": "Name of the Worker associated with the domain.", "type": "string", "example": "my-worker"}}, {"name": "hostname", "in": "query", "schema": {"description": "Hostname of the domain.", "type": "string", "example": "app.example.com"}}, {"name": "environment", "in": "query", "schema": {"description": "Worker environment associated with the domain.", "type": "string", "example": "production"}, "deprecated": true}], "responses": {"200": {"description": "List domains response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/workers_Domain"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "List domains failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Domains"], "x-api-token-group": ["Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.domains", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
