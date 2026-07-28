---
title: teams-devices_client_certificate_input_request
page_id: schema-teams-devices-client-certificate-input-request-785a4589
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_client_certificate_input_request

```yaml
{"type": "object", "properties": {"certificate_id": {"description": "UUID of Cloudflare managed certificate.", "type": "string", "example": "b14ddcc4-bcd2-4df4-bd4f-eb27d5a50c30", "maxLength": 36}, "cn": {"description": "Common Name that is protected by the certificate.", "type": "string", "example": "example.com"}}, "required": ["certificate_id", "cn"], "title": "Client Certificate"}
```
