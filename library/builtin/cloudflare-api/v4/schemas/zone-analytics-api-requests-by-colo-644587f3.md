---
title: zone-analytics-api_requests_by_colo
page_id: schema-zone-analytics-api-requests-by-colo-644587f3
path: schemas
description: Breakdown of totals for requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_requests_by_colo

Breakdown of totals for requests.

```yaml
{"description": "Breakdown of totals for requests.", "type": "object", "properties": {"all": {"description": "Total number of requests served.", "type": "integer"}, "cached": {"description": "Total number of cached requests served.", "type": "integer"}, "country": {"description": "Key/value pairs where the key is a two-digit country code and the value is the number of requests served to that country.", "type": "object", "example": {"AG": 37298, "GI": 293846, "US": 4181364}, "additionalProperties": true}, "http_status": {"description": "A variable list of key/value pairs where the key is a HTTP status code and the value is the number of requests with that code served.", "type": "object", "example": {"200": 13496983, "301": 283, "400": 187936, "402": 1828, "404": 1293}}, "uncached": {"description": "Total number of requests served from the origin.", "type": "integer"}}}
```
