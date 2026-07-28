---
title: email_update_destination_address_properties
page_id: schema-email-update-destination-address-properties-11d2e21c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_update_destination_address_properties

```yaml
{"type": "object", "properties": {"status": {"description": "Destination address status. Non-admin callers may only set verified addresses back to unverified; setting to verified requires admin privileges.", "type": "string", "example": "verified", "enum": ["unverified", "verified"], "x-auditable": true}}, "required": ["status"]}
```
