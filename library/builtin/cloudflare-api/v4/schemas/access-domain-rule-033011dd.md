---
title: access_domain_rule
page_id: schema-access-domain-rule-033011dd
path: schemas
description: Match an entire email domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_domain_rule

Match an entire email domain.

```yaml
{"description": "Match an entire email domain.", "type": "object", "properties": {"email_domain": {"type": "object", "properties": {"domain": {"description": "The email domain to match.", "type": "string", "example": "example.com"}}, "required": ["domain"]}}, "required": ["email_domain"], "title": "Email domain"}
```
