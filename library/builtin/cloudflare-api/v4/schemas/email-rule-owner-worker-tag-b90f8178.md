---
title: email_rule_owner_worker_tag
page_id: schema-email-rule-owner-worker-tag-b90f8178
path: schemas
description: |-
    Public tag (script_tag) of the Worker that owns this rule. Required when
    `source` is `wrangler`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_rule_owner_worker_tag

Public tag (script_tag) of the Worker that owns this rule. Required when
`source` is `wrangler`.

```yaml
{"description": "Public tag (script_tag) of the Worker that owns this rule. Required when\n`source` is `wrangler`.\n", "type": "string", "example": "a7e6fb77503c41d8a7f3113c6918f10c", "maxLength": 32, "writeOnly": true, "x-auditable": true}
```
