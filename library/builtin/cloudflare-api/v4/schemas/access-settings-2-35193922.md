---
title: access_settings-2
page_id: schema-access-settings-2-35193922
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_settings-2

```yaml
{"type": "object", "properties": {"china_network": {"description": "Request client certificates for this hostname in China. Can only be set to true if this zone is china network enabled.", "type": "boolean", "example": false}, "client_certificate_forwarding": {"description": "Client Certificate Forwarding is a feature that takes the client cert provided by the eyeball to the edge, and forwards it to the origin as a HTTP header to allow logging on the origin.", "type": "boolean", "example": true}, "hostname": {"description": "The hostname that these settings apply to.", "type": "string", "example": "admin.example.com"}}, "additionalProperties": false, "required": ["hostname", "china_network", "client_certificate_forwarding"], "title": "Hostname Settings"}
```
