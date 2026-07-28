---
title: access_decision
page_id: schema-access-decision-4a762825
path: schemas
description: The action Access will take if a user matches this policy. Infrastructure application policies can only use the Allow action.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_decision

The action Access will take if a user matches this policy. Infrastructure application policies can only use the Allow action.

```yaml
{"description": "The action Access will take if a user matches this policy. Infrastructure application policies can only use the Allow action.", "type": "string", "example": "allow", "enum": ["allow", "deny", "non_identity", "bypass"]}
```
