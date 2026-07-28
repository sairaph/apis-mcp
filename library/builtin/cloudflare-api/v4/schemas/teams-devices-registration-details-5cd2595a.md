---
title: teams-devices_registration_details
page_id: schema-teams-devices-registration-details-5cd2595a
path: schemas
description: The summary of a registration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_registration_details

The summary of a registration.

```yaml
{"description": "The summary of a registration.", "type": "object", "properties": {"policy": {"description": "A summary of the device profile evaluated for the registration.", "type": "object", "allOf": [{"$ref": "#/components/schemas/teams-devices_policy_summary"}], "nullable": true}}}
```
