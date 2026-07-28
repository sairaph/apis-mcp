---
title: Get Domain History
page_id: operation-get-accounts-account-id-intel-domain-history-42e20de1
path: operations/domain-history
description: Gets historical security threat and content categories currently and previously assigned to a domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/domain-history
operation_ids:
    - domain-history-get-domain-history
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Domain History

`GET /accounts/{account_id}/intel/domain-history`

Operation ID: `domain-history-get-domain-history`

Gets historical security threat and content categories currently and previously assigned to a domain.

## Definition

```yaml
{"operationId": "domain-history-get-domain-history", "summary": "Get Domain History", "description": "Gets historical security threat and content categories currently and previously assigned to a domain.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/intel_identifier"}}, {"name": "domain", "in": "query", "schema": {"type": "string", "example": "example.com"}}], "responses": {"200": {"description": "Get Domain History response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/intel_response"}}}}, "4XX": {"description": "Get Domain History response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/intel_response"}, {"$ref": "#/components/schemas/intel_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Domain History"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
