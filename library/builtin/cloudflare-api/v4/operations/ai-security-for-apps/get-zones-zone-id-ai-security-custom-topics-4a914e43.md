---
title: Get AI Security for Apps Custom Topics
page_id: operation-get-zones-zone-id-ai-security-custom-topics-ebd02940
path: operations/ai-security-for-apps
description: Get the AI Security for Apps custom topic categories for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/ai-security/custom-topics
operation_ids:
    - ai-security-custom-topics-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get AI Security for Apps Custom Topics

`GET /zones/{zone_id}/ai-security/custom-topics`

Operation ID: `ai-security-custom-topics-get`

Get the AI Security for Apps custom topic categories for a zone.

## Definition

```yaml
{"operationId": "ai-security-custom-topics-get", "summary": "Get AI Security for Apps Custom Topics", "description": "Get the AI Security for Apps custom topic categories for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_zone_id"}}], "responses": {"200": {"description": "Get AI Security for Apps custom topics response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-topics"}}}}, "4XX": {"description": "Get AI Security for Apps custom topics failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-topics"}, {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure-3"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["AI Security for Apps"]}
```
