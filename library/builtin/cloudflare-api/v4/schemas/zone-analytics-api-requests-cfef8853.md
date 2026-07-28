---
title: zone-analytics-api_requests
page_id: schema-zone-analytics-api-requests-cfef8853
path: schemas
description: Breakdown of totals for requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_requests

Breakdown of totals for requests.

```yaml
{"description": "Breakdown of totals for requests.", "type": "object", "properties": {"all": {"description": "Total number of requests served.", "type": "integer"}, "cached": {"description": "Total number of cached requests served.", "type": "integer"}, "content_type": {"description": "A variable list of key/value pairs where the key represents the type of content served, and the value is the number of requests.", "type": "object", "example": {"css": 15343, "gif": 23178, "html": 1234213, "javascript": 318236, "jpeg": 1982048}}, "country": {"description": "A variable list of key/value pairs where the key is a two-digit country code and the value is the number of requests served to that country.", "type": "object", "example": {"AG": 37298, "GI": 293846, "US": 4181364}}, "http_status": {"description": "Key/value pairs where the key is a HTTP status code and the value is the number of requests served with that code.", "type": "object", "example": {"200": 13496983, "301": 283, "400": 187936, "402": 1828, "404": 1293}, "additionalProperties": true}, "ssl": {"description": "A break down of requests served over HTTPS.", "type": "object", "properties": {"encrypted": {"description": "The number of requests served over HTTPS.", "type": "integer"}, "unencrypted": {"description": "The number of requests served over HTTP.", "type": "integer"}}}, "ssl_protocols": {"description": "A breakdown of requests by their SSL protocol.", "type": "object", "properties": {"TLSv1": {"description": "The number of requests served over TLS v1.0.", "type": "integer"}, "TLSv1.1": {"description": "The number of requests served over TLS v1.1.", "type": "integer"}, "TLSv1.2": {"description": "The number of requests served over TLS v1.2.", "type": "integer"}, "TLSv1.3": {"description": "The number of requests served over TLS v1.3.", "type": "integer"}, "none": {"description": "The number of requests served over HTTP.", "type": "integer"}}}, "uncached": {"description": "Total number of requests served from the origin.", "type": "integer"}}}
```
