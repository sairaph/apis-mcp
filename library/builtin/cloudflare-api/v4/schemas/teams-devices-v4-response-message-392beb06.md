---
title: teams-devices_v4_response_message
page_id: schema-teams-devices-v4-response-message-392beb06
path: schemas
description: A message which can be returned in either the 'errors' or 'messages' fields in a v4 API response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_v4_response_message

A message which can be returned in either the 'errors' or 'messages' fields in a v4 API response.

```yaml
{"description": "A message which can be returned in either the 'errors' or 'messages' fields in a v4 API response.", "type": "object", "properties": {"code": {"type": "integer"}, "message": {"type": "string"}}, "required": ["code", "message"]}
```
