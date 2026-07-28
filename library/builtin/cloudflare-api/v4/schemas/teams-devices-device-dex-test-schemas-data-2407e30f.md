---
title: teams-devices_device-dex-test-schemas-data
page_id: schema-teams-devices-device-dex-test-schemas-data-2407e30f
path: schemas
description: The configuration object which contains the details for the WARP client to conduct the test.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_device-dex-test-schemas-data

The configuration object which contains the details for the WARP client to conduct the test.

```yaml
{"description": "The configuration object which contains the details for the WARP client to conduct the test.", "type": "object", "properties": {"host": {"description": "The desired endpoint to test.", "type": "string", "example": "https://dash.cloudflare.com"}, "kind": {"description": "The type of test.", "type": "string", "example": "http"}, "method": {"description": "The HTTP request method type.", "type": "string", "example": "GET"}}, "example": {"host": "https://dash.cloudflare.com", "kind": "http", "method": "GET"}}
```
