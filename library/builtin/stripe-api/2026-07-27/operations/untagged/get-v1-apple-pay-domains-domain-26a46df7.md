---
title: GET /v1/apple_pay/domains/{domain}
page_id: operation-get-v1-apple-pay-domains-domain-a1d1f446
path: operations/untagged
description: <p>Retrieve an apple pay domain.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/apple_pay/domains/{domain}
operation_ids:
    - GetApplePayDomainsDomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# GET /v1/apple_pay/domains/{domain}

`GET /v1/apple_pay/domains/{domain}`

Operation ID: `GetApplePayDomainsDomain`

<p>Retrieve an apple pay domain.</p>

## Definition

```yaml
{"description": "<p>Retrieve an apple pay domain.</p>", "operationId": "GetApplePayDomainsDomain", "parameters": [{"name": "domain", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/apple_pay_domain"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
