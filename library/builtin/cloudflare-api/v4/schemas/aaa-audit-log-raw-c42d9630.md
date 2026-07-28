---
title: aaa_audit-log-raw
page_id: schema-aaa-audit-log-raw-c42d9630
path: schemas
description: Provides raw information about the request and response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_audit-log-raw

Provides raw information about the request and response.

```yaml
{"description": "Provides raw information about the request and response.", "type": "object", "properties": {"cf_ray_id": {"description": "The Cloudflare Ray ID for the request.", "type": "string", "example": "8e9b1c60ef9e1c9a"}, "method": {"description": "The HTTP method of the request.", "type": "string", "example": "POST"}, "status_code": {"description": "The HTTP response status code returned by the API.", "type": "integer", "example": 200}, "uri": {"description": "The URI of the request.", "type": "string", "example": "/accounts/4bb334f7c94c4a29a045f03944f072e5/members"}, "user_agent": {"description": "The client's user agent string sent with the request.", "type": "string", "example": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) Safari/605.1.15"}}}
```
