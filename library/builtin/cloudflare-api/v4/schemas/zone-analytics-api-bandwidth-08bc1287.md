---
title: zone-analytics-api_bandwidth
page_id: schema-zone-analytics-api-bandwidth-08bc1287
path: schemas
description: Breakdown of totals for bandwidth in the form of bytes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_bandwidth

Breakdown of totals for bandwidth in the form of bytes.

```yaml
{"description": "Breakdown of totals for bandwidth in the form of bytes.", "type": "object", "properties": {"all": {"description": "The total number of bytes served within the time frame.", "type": "integer"}, "cached": {"description": "The number of bytes that were cached (and served) by Cloudflare.", "type": "integer"}, "content_type": {"description": "A variable list of key/value pairs where the key represents the type of content served, and the value is the number in bytes served.", "type": "object", "example": {"css": 237421, "gif": 1234242, "html": 1231290, "javascript": 123245, "jpeg": 784278}}, "country": {"description": "A variable list of key/value pairs where the key is a two-digit country code and the value is the number of bytes served to that country.", "type": "object", "example": {"AG": 2342483, "GI": 984753, "US": 123145433}}, "ssl": {"description": "A break down of bytes served over HTTPS.", "type": "object", "properties": {"encrypted": {"description": "The number of bytes served over HTTPS.", "type": "integer"}, "unencrypted": {"description": "The number of bytes served over HTTP.", "type": "integer"}}}, "ssl_protocols": {"description": "A breakdown of requests by their SSL protocol.", "type": "object", "properties": {"TLSv1": {"description": "The number of requests served over TLS v1.0.", "type": "integer"}, "TLSv1.1": {"description": "The number of requests served over TLS v1.1.", "type": "integer"}, "TLSv1.2": {"description": "The number of requests served over TLS v1.2.", "type": "integer"}, "TLSv1.3": {"description": "The number of requests served over TLS v1.3.", "type": "integer"}, "none": {"description": "The number of requests served over HTTP.", "type": "integer"}}}, "uncached": {"description": "The number of bytes that were fetched and served from the origin server.", "type": "integer"}}}
```
