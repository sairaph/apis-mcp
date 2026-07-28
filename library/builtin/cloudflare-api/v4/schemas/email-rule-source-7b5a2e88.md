---
title: email_rule_source
page_id: schema-email-rule-source-7b5a2e88
path: schemas
description: |-
    Who manages the rule. `api` covers dashboard, generic API, and Terraform;
    `wrangler` means the rule is managed by a Worker's wrangler.jsonc. Defaults
    to `api` when omitted on write.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_rule_source

Who manages the rule. `api` covers dashboard, generic API, and Terraform;
`wrangler` means the rule is managed by a Worker's wrangler.jsonc. Defaults
to `api` when omitted on write.

```yaml
{"description": "Who manages the rule. `api` covers dashboard, generic API, and Terraform;\n`wrangler` means the rule is managed by a Worker's wrangler.jsonc. Defaults\nto `api` when omitted on write.\n", "type": "string", "example": "api", "default": "api", "enum": ["api", "wrangler"], "x-auditable": true}
```
