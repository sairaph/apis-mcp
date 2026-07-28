---
title: access_cloudflare_account_member_rule
page_id: schema-access-cloudflare-account-member-rule-3b22b3f4
path: schemas
description: |-
    Matches users who are members of a specific Cloudflare account.
    Requires a Cloudflare identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_cloudflare_account_member_rule

Matches users who are members of a specific Cloudflare account.
Requires a Cloudflare identity provider.

```yaml
{"description": "Matches users who are members of a specific Cloudflare account.\nRequires a Cloudflare identity provider.", "type": "object", "properties": {"cloudflare_account_member": {"type": "object", "properties": {"account_id": {"$ref": "#/components/schemas/access_identifier"}}}}, "required": ["cloudflare_account_member"], "title": "Cloudflare Account Member"}
```
