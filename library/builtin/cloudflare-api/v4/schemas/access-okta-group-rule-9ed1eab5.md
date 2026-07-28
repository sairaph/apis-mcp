---
title: access_okta_group_rule
page_id: schema-access-okta-group-rule-9ed1eab5
path: schemas
description: |-
    Matches an Okta group.
    Requires an Okta identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_okta_group_rule

Matches an Okta group.
Requires an Okta identity provider.

```yaml
{"description": "Matches an Okta group.\nRequires an Okta identity provider.", "type": "object", "properties": {"okta": {"type": "object", "properties": {"identity_provider_id": {"description": "The ID of your Okta identity provider.", "type": "string", "example": "ea85612a-29c8-46c2-bacb-669d65136971"}, "name": {"description": "The name of the Okta group.", "type": "string", "example": "devs"}}, "required": ["name", "identity_provider_id"]}}, "required": ["okta"], "title": "Okta group"}
```
