---
title: Get Domain
page_id: operation-get-accounts-account-id-workers-domains-domain-id-96f22ab7
path: operations/domains
description: Gets information about a domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/domains/{domain_id}
operation_ids:
    - workers.domains.get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Domain

`GET /accounts/{account_id}/workers/domains/{domain_id}`

Operation ID: `workers.domains.get`

Gets information about a domain.

## Definition

```yaml
{"operationId": "workers.domains.get", "summary": "Get Domain", "description": "Gets information about a domain.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "domain_id", "in": "path", "required": true, "schema": {"description": "ID of the domain.", "type": "string", "example": "dbe10b4bc17c295377eabd600e1787fd"}}], "responses": {"200": {"description": "Get domain response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_Domain"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get domain failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Domains"], "x-api-token-group": ["Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.domains", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
