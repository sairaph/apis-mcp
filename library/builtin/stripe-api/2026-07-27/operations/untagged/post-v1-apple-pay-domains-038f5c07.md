---
title: POST /v1/apple_pay/domains
page_id: operation-post-v1-apple-pay-domains-cb6bf7ef
path: operations/untagged
description: <p>Create an apple pay domain.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/apple_pay/domains
operation_ids:
    - PostApplePayDomains
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# POST /v1/apple_pay/domains

`POST /v1/apple_pay/domains`

Operation ID: `PostApplePayDomains`

<p>Create an apple pay domain.</p>

## Definition

```yaml
{"description": "<p>Create an apple pay domain.</p>", "operationId": "PostApplePayDomains", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["domain_name"], "type": "object", "properties": {"domain_name": {"type": "string"}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/apple_pay_domain"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
