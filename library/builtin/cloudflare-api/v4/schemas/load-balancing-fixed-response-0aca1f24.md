---
title: load-balancing_fixed_response
page_id: schema-load-balancing-fixed-response-0aca1f24
path: schemas
description: A collection of fields used to directly respond to the client instead of routing to a pool. When supplied on a rule, that rule stops further rule evaluation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_fixed_response

A collection of fields used to directly respond to the client instead of routing to a pool. When supplied on a rule, that rule stops further rule evaluation.

```yaml
{"description": "A collection of fields used to directly respond to the client instead of routing to a pool. When supplied on a rule, that rule stops further rule evaluation.", "type": "object", "properties": {"content_type": {"description": "The http 'Content-Type' header to include in the response.", "type": "string", "example": "application/json", "maxLength": 32, "x-auditable": true}, "location": {"description": "The http 'Location' header to include in the response.", "type": "string", "example": "www.example.com", "maxLength": 2048, "x-auditable": true}, "message_body": {"description": "Text to include as the http body.", "type": "string", "example": "Testing Hello", "maxLength": 1024, "x-auditable": true}, "status_code": {"description": "The http status code to respond with.", "type": "integer", "x-auditable": true}}}
```
