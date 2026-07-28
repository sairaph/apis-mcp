---
title: access_github_organization_rule
page_id: schema-access-github-organization-rule-e6c4ba54
path: schemas
description: |-
    Matches a Github organization.
    Requires a Github identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_github_organization_rule

Matches a Github organization.
Requires a Github identity provider.

```yaml
{"description": "Matches a Github organization.\nRequires a Github identity provider.", "type": "object", "properties": {"github-organization": {"type": "object", "properties": {"identity_provider_id": {"description": "The ID of your Github identity provider.", "type": "string", "example": "ea85612a-29c8-46c2-bacb-669d65136971"}, "name": {"description": "The name of the organization.", "type": "string", "example": "cloudflare"}, "team": {"description": "The name of the team", "type": "string", "example": "api-team"}}, "required": ["name", "identity_provider_id"]}}, "required": ["github-organization"], "title": "Github organization"}
```
