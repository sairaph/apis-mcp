---
title: forwarded_response_details
page_id: schema-forwarded-response-details-43e0a868
path: schemas
description: Details about the response from the destination endpoint.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# forwarded_response_details

Details about the response from the destination endpoint.

```yaml
{"title": "ForwardedResponseDetails", "required": ["body", "headers", "status"], "type": "object", "properties": {"body": {"maxLength": 5000, "type": "string", "description": "The response body from the destination endpoint to Stripe."}, "headers": {"type": "array", "description": "HTTP headers that the destination endpoint returned.", "items": {"$ref": "#/components/schemas/forwarded_request_header"}}, "status": {"type": "integer", "description": "The HTTP status code that the destination endpoint returned."}}, "description": "Details about the response from the destination endpoint.", "x-expandableFields": ["headers"]}
```
