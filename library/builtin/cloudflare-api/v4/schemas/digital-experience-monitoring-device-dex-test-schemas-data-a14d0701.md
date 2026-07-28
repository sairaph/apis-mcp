---
title: digital-experience-monitoring_device-dex-test-schemas-data
page_id: schema-digital-experience-monitoring-device-dex-test-schemas-data-a14d0701
path: schemas
description: The configuration object which contains the details for the WARP client to conduct the test.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_device-dex-test-schemas-data

The configuration object which contains the details for the WARP client to conduct the test.

```yaml
{"description": "The configuration object which contains the details for the WARP client to conduct the test.", "type": "object", "properties": {"host": {"description": "The desired endpoint to test.", "type": "string", "example": "https://dash.cloudflare.com", "x-auditable": true}, "kind": {"description": "The type of test.", "type": "string", "example": "http", "enum": ["http", "traceroute"], "x-auditable": true}, "method": {"description": "The HTTP request method type.", "type": "string", "example": "GET", "enum": ["GET"], "x-auditable": true}}, "example": {"host": "https://dash.cloudflare.com", "kind": "http", "method": "GET"}, "required": ["kind", "host"]}
```
