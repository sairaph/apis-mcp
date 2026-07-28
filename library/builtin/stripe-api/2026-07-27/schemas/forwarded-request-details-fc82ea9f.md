---
title: forwarded_request_details
page_id: schema-forwarded-request-details-fc82ea9f
path: schemas
description: Details about the request forwarded to the destination endpoint.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# forwarded_request_details

Details about the request forwarded to the destination endpoint.

```yaml
{"title": "ForwardedRequestDetails", "required": ["body", "headers", "http_method"], "type": "object", "properties": {"body": {"maxLength": 5000, "type": "string", "description": "The body payload to send to the destination endpoint."}, "headers": {"type": "array", "description": "The headers to include in the forwarded request. Can be omitted if no additional headers (excluding Stripe-generated ones such as the Content-Type header) should be included.", "items": {"$ref": "#/components/schemas/forwarded_request_header"}}, "http_method": {"type": "string", "description": "The HTTP method used to call the destination endpoint.", "enum": ["POST"]}}, "description": "Details about the request forwarded to the destination endpoint.", "x-expandableFields": ["headers"]}
```
