---
title: email-security_CreateBlockedSender
page_id: schema-email-security-createblockedsender-24216d4a
path: schemas
description: Create a blocked sender pattern.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_CreateBlockedSender

Create a blocked sender pattern.

```yaml
{"description": "Create a blocked sender pattern.", "allOf": [{"$ref": "#/components/schemas/email-security_BlockedSender"}], "required": ["pattern", "is_regex", "pattern_type"]}
```
