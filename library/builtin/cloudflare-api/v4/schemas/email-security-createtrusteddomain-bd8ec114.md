---
title: email-security_CreateTrustedDomain
page_id: schema-email-security-createtrusteddomain-bd8ec114
path: schemas
description: Create a trusted domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-security_CreateTrustedDomain

Create a trusted domain.

```yaml
{"description": "Create a trusted domain.", "allOf": [{"$ref": "#/components/schemas/email-security_TrustedDomain"}], "required": ["pattern", "is_regex", "is_recent", "is_similarity"]}
```
