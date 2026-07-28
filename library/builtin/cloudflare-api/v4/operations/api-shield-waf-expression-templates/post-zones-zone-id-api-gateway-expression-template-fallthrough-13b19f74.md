---
title: Generate fallthrough WAF expression template from a set of API hosts
page_id: operation-post-zones-zone-id-api-gateway-expression-template-fallthrough-fa0f517c
path: operations/api-shield-waf-expression-templates
description: Creates an expression template fallthrough rule for API Shield. Used for configuring default behavior when no other expression templates match.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/api_gateway/expression-template/fallthrough
operation_ids:
    - api-shield-expression-templates-fallthrough
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Generate fallthrough WAF expression template from a set of API hosts

`POST /zones/{zone_id}/api_gateway/expression-template/fallthrough`

Operation ID: `api-shield-expression-templates-fallthrough`

Creates an expression template fallthrough rule for API Shield. Used for configuring default behavior when no other expression templates match.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "api-shield-expression-templates-fallthrough", "summary": "Generate fallthrough WAF expression template from a set of API hosts", "description": "Creates an expression template fallthrough rule for API Shield. Used for configuring default behavior when no other expression templates match.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_request_expression_templates_fallthrough"}}}}, "responses": {"200": {"description": "Generate fallthrough WAF expression template response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/api-shield_response_expression_templates_fallthrough"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Generate fallthrough WAF expression template failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/api-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["API Shield WAF Expression Templates"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "api-gateway.expression-template.fallthrough", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
