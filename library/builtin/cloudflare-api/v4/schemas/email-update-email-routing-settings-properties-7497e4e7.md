---
title: email_update_email_routing_settings_properties
page_id: schema-email-update-email-routing-settings-properties-7497e4e7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_update_email_routing_settings_properties

```yaml
{"type": "object", "properties": {"enabled": {"description": "State of your zone Email Routing settings. No-op on this endpoint - use `POST`/`DELETE /zones/{zone_id}/email/routing/dns`.", "type": "boolean", "example": true, "enum": [true, false]}, "skip_wizard": {"description": "Flag to check if the user skipped the configuration wizard.", "type": "boolean", "example": true, "enum": [true, false]}, "support_subaddress": {"description": "Whether subaddressing (plus-addressing) is honored when matching incoming mail against routing rules.", "type": "boolean", "example": true, "enum": [true, false]}}}
```
