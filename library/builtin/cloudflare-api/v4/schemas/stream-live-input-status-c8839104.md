---
title: stream_live_input_status
page_id: schema-stream-live-input-status-c8839104
path: schemas
description: The connection status of a live input.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# stream_live_input_status

The connection status of a live input.

```yaml
{"description": "The connection status of a live input.", "type": "string", "enum": [null, "connected", "reconnected", "reconnecting", "client_disconnect", "ttl_exceeded", "failed_to_connect", "failed_to_reconnect", "new_configuration_accepted"], "nullable": true, "x-auditable": true}
```
