---
title: access_linked_app_token_rule
page_id: schema-access-linked-app-token-rule-2cdcaac7
path: schemas
description: Matches OAuth 2.0 access tokens issued by the specified Access OIDC SaaS application. Only compatible with non_identity and bypass decisions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_linked_app_token_rule

Matches OAuth 2.0 access tokens issued by the specified Access OIDC SaaS application. Only compatible with non_identity and bypass decisions.

```yaml
{"description": "Matches OAuth 2.0 access tokens issued by the specified Access OIDC SaaS application. Only compatible with non_identity and bypass decisions.", "type": "object", "properties": {"linked_app_token": {"type": "object", "properties": {"app_uid": {"description": "The ID of an Access OIDC SaaS application", "type": "string", "example": "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"}}, "required": ["app_uid"]}}, "required": ["linked_app_token"], "title": "Linked App Token"}
```
