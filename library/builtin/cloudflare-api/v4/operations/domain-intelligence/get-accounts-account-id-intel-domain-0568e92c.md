---
title: Get Domain Details
page_id: operation-get-accounts-account-id-intel-domain-91305411
path: operations/domain-intelligence
description: Gets security details and statistics about a domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/domain
operation_ids:
    - domain-intelligence-get-domain-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Domain Details

`GET /accounts/{account_id}/intel/domain`

Operation ID: `domain-intelligence-get-domain-details`

Gets security details and statistics about a domain.

## Definition

```yaml
{"operationId": "domain-intelligence-get-domain-details", "summary": "Get Domain Details", "description": "Gets security details and statistics about a domain.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/intel_identifier"}}, {"name": "domain", "in": "query", "schema": {"type": "string"}}, {"name": "skip_dns", "in": "query", "description": "Skip DNS resolution lookups for faster response.", "schema": {"type": "boolean"}}, {"name": "skip_ranking", "in": "query", "description": "Skip the domain ranking lookup for faster responses. Defaults to\n`false` (ranking is included). Set to `true` to opt out — primarily\nused by callers like Cloudflare Radar that need to avoid a\ncircular dependency when building the domain details page.\nNote: the bulk endpoint (`/intel/domain/bulk`) uses opposite\ndefaults — see `include_ranking` there.\n", "schema": {"type": "boolean", "default": false}}], "responses": {"200": {"description": "Get Domain Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel_single_response"}}}}, "4XX": {"description": "Get Domain Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/intel_single_response"}, {"$ref": "#/components/schemas/intel_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Domain Intelligence"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
