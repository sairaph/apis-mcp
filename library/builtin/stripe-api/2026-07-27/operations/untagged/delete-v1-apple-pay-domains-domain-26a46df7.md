---
title: DELETE /v1/apple_pay/domains/{domain}
page_id: operation-delete-v1-apple-pay-domains-domain-7da241a2
path: operations/untagged
description: <p>Delete an apple pay domain.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/apple_pay/domains/{domain}
operation_ids:
    - DeleteApplePayDomainsDomain
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# DELETE /v1/apple_pay/domains/{domain}

`DELETE /v1/apple_pay/domains/{domain}`

Operation ID: `DeleteApplePayDomainsDomain`

<p>Delete an apple pay domain.</p>

## Definition

```yaml
{"description": "<p>Delete an apple pay domain.</p>", "operationId": "DeleteApplePayDomainsDomain", "parameters": [{"name": "domain", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_apple_pay_domain"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
