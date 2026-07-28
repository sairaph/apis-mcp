---
title: Get URL Intelligence
page_id: operation-get-accounts-account-id-intel-url-a8314460
path: operations/url-intelligence
description: Gets security information about a URL, including content categories and risk types. The URL must be provided as a query parameter.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/url
operation_ids:
    - url-intelligence-get-url-intelligence
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get URL Intelligence

`GET /accounts/{account_id}/intel/url`

Operation ID: `url-intelligence-get-url-intelligence`

Gets security information about a URL, including content categories and risk types. The URL must be provided as a query parameter.

## Definition

```yaml
{"operationId": "url-intelligence-get-url-intelligence", "summary": "Get URL Intelligence", "description": "Gets security information about a URL, including content categories and risk types. The URL must be provided as a query parameter.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/intel_identifier"}}, {"name": "url", "in": "query", "description": "The URL to look up.", "required": true, "schema": {"type": "string"}, "example": "https://example.com/path"}], "responses": {"200": {"description": "Get URL Intelligence response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel_url_intelligence_single_response"}}}}, "4XX": {"description": "Get URL Intelligence response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/intel_url_intelligence_single_response"}, {"$ref": "#/components/schemas/intel_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["URL Intelligence"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
