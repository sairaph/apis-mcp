---
title: access_auth_context_rule
page_id: schema-access-auth-context-rule-d92cd2be
path: schemas
description: |-
    Matches an Azure Authentication Context.
    Requires an Azure identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_auth_context_rule

Matches an Azure Authentication Context.
Requires an Azure identity provider.

```yaml
{"description": "Matches an Azure Authentication Context.\nRequires an Azure identity provider.", "type": "object", "properties": {"auth_context": {"type": "object", "properties": {"ac_id": {"description": "The ACID of an Authentication context.", "type": "string", "example": "ea85612a-29c8-46c2-bacb-669d65136971"}, "id": {"description": "The ID of an Authentication context.", "type": "string", "example": "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"}, "identity_provider_id": {"description": "The ID of your Azure identity provider.", "type": "string", "example": "ea85612a-29c8-46c2-bacb-669d65136971"}}, "required": ["id", "identity_provider_id", "ac_id"]}}, "required": ["auth_context"], "title": "Authentication Context"}
```
