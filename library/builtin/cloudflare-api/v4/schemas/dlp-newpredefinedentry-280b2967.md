---
title: dlp_NewPredefinedEntry
page_id: schema-dlp-newpredefinedentry-280b2967
path: schemas
description: |-
    Used to create a new predefined or integration entry.

    Predefined or integration entries can not be updated via the API so
    these fields will update the entry's settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_NewPredefinedEntry

Used to create a new predefined or integration entry.

Predefined or integration entries can not be updated via the API so
these fields will update the entry's settings.

```yaml
{"description": "Used to create a new predefined or integration entry.\n\nPredefined or integration entries can not be updated via the API so\nthese fields will update the entry's settings.", "type": "object", "properties": {"enabled": {"type": "boolean"}, "entry_id": {"type": "string", "format": "uuid"}, "profile_id": {"description": "This field is not used as the owning profile.\nFor predefined entries it is already set to a predefined profile.", "type": "string", "format": "uuid", "nullable": true}}, "required": ["enabled", "entry_id"]}
```
