---
title: registrar-api-sandbox_registration_update_request
page_id: schema-registrar-api-sandbox-registration-update-request-0c00c1e4
path: schemas
description: |-
    Request to update an existing domain registration.

    This endpoint currently supports updating `auto_renew` only.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_registration_update_request

Request to update an existing domain registration.

This endpoint currently supports updating `auto_renew` only.

```yaml
{"description": "Request to update an existing domain registration.\n\nThis endpoint currently supports updating `auto_renew` only.\n", "type": "object", "properties": {"auto_renew": {"description": "Enable or disable automatic renewal.\nSetting this field to `true` authorizes Cloudflare to charge the\naccount's default payment method up to 30 days before domain expiry\nto renew the domain automatically. Renewal pricing may change over\ntime based on registry pricing.\n", "type": "boolean", "example": true}}, "minProperties": 1}
```
