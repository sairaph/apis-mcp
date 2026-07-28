---
title: zones_proxy_read_timeout
page_id: schema-zones-proxy-read-timeout-9d60d57c
path: schemas
description: Maximum time between two read operations from origin.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_proxy_read_timeout

Maximum time between two read operations from origin.

```yaml
{"description": "Maximum time between two read operations from origin.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "proxy_read_timeout", "enum": ["proxy_read_timeout"]}, "value": {"$ref": "#/components/schemas/zones_proxy_read_timeout_value"}}}], "title": "Proxy Read Timeout"}
```
