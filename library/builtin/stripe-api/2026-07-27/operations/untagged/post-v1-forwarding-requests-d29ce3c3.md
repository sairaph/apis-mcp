---
title: Create a ForwardingRequest
page_id: operation-post-v1-forwarding-requests-2c6a2efb
path: operations/untagged
description: <p>Creates a ForwardingRequest object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/forwarding/requests
operation_ids:
    - PostForwardingRequests
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a ForwardingRequest

`POST /v1/forwarding/requests`

Operation ID: `PostForwardingRequests`

<p>Creates a ForwardingRequest object.</p>

## Definition

```yaml
{"summary": "Create a ForwardingRequest", "description": "<p>Creates a ForwardingRequest object.</p>", "operationId": "PostForwardingRequests", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["payment_method", "replacements", "url"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "payment_method": {"maxLength": 5000, "type": "string", "description": "The PaymentMethod to insert into the forwarded request. Forwarding previously consumed PaymentMethods is allowed."}, "replacements": {"type": "array", "description": "The field kinds to be replaced in the forwarded request.", "items": {"type": "string", "enum": ["card_cvc", "card_expiry", "card_number", "cardholder_name", "request_signature"], "x-stripeBypassValidation": true}}, "request": {"title": "request_param", "type": "object", "properties": {"body": {"maxLength": 5000, "type": "string"}, "headers": {"type": "array", "items": {"title": "header_param", "required": ["name", "value"], "type": "object", "properties": {"name": {"maxLength": 5000, "type": "string"}, "value": {"maxLength": 5000, "type": "string"}}}}}, "description": "The request body and headers to be sent to the destination endpoint."}, "url": {"maxLength": 5000, "type": "string", "description": "The destination URL for the forwarded request. Must be supported by the config."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "replacements": {"style": "deepObject", "explode": true}, "request": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/forwarding.request"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
