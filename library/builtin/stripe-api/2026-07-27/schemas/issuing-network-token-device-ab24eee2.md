---
title: issuing_network_token_device
page_id: schema-issuing-network-token-device-ab24eee2
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_network_token_device

```yaml
{"title": "IssuingNetworkTokenDevice", "type": "object", "properties": {"device_fingerprint": {"maxLength": 5000, "type": "string", "description": "An obfuscated ID derived from the device ID."}, "ip_address": {"maxLength": 5000, "type": "string", "description": "The IP address of the device at provisioning time."}, "location": {"maxLength": 5000, "type": "string", "description": "The geographic latitude/longitude coordinates of the device at provisioning time. The format is [+-]decimal/[+-]decimal."}, "name": {"maxLength": 5000, "type": "string", "description": "The name of the device used for tokenization."}, "phone_number": {"maxLength": 5000, "type": "string", "description": "The phone number of the device used for tokenization."}, "type": {"type": "string", "description": "The type of device used for tokenization.", "enum": ["other", "phone", "watch"]}}, "description": "", "x-expandableFields": []}
```
