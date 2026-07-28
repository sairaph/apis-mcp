---
title: email-security_CreateAllowPolicy
page_id: schema-email-security-createallowpolicy-01957bfb
path: schemas
description: Create an allow policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_CreateAllowPolicy

Create an allow policy.

```yaml
{"description": "Create an allow policy.", "allOf": [{"$ref": "#/components/schemas/email-security_AllowPolicy"}], "required": ["is_exempt_recipient", "is_trusted_sender", "is_acceptable_sender", "pattern", "is_regex", "verify_sender", "pattern_type"]}
```
