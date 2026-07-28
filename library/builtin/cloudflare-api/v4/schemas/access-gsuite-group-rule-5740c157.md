---
title: access_gsuite_group_rule
page_id: schema-access-gsuite-group-rule-5740c157
path: schemas
description: |-
    Matches a group in Google Workspace.
    Requires a Google Workspace identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_gsuite_group_rule

Matches a group in Google Workspace.
Requires a Google Workspace identity provider.

```yaml
{"description": "Matches a group in Google Workspace.\nRequires a Google Workspace identity provider.", "type": "object", "properties": {"gsuite": {"type": "object", "properties": {"email": {"description": "The email of the Google Workspace group.", "type": "string", "example": "devs@cloudflare.com"}, "identity_provider_id": {"description": "The ID of your Google Workspace identity provider.", "type": "string", "example": "ea85612a-29c8-46c2-bacb-669d65136971"}}, "required": ["email", "identity_provider_id"]}}, "required": ["gsuite"], "title": "Google Workspace group"}
```
