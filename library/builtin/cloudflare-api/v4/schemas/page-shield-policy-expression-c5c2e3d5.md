---
title: page-shield_policy-expression
page_id: schema-page-shield-policy-expression-c5c2e3d5
path: schemas
description: The expression which must match for the policy to be applied, using the Cloudflare Firewall rule expression syntax
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# page-shield_policy-expression

The expression which must match for the policy to be applied, using the Cloudflare Firewall rule expression syntax

```yaml
{"description": "The expression which must match for the policy to be applied, using the Cloudflare Firewall rule expression syntax", "type": "string", "example": "ends_with(http.request.uri.path, \"/checkout\")", "x-auditable": true}
```
