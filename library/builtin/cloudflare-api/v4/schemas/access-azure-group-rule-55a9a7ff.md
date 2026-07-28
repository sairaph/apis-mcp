---
title: access_azure_group_rule
page_id: schema-access-azure-group-rule-55a9a7ff
path: schemas
description: |-
    Matches an Azure group.
    Requires an Azure identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_azure_group_rule

Matches an Azure group.
Requires an Azure identity provider.

```yaml
{"description": "Matches an Azure group.\nRequires an Azure identity provider.", "type": "object", "properties": {"azureAD": {"type": "object", "properties": {"id": {"description": "The ID of an Azure group.", "type": "string", "example": "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"}, "identity_provider_id": {"description": "The ID of your Azure identity provider.", "type": "string", "example": "ea85612a-29c8-46c2-bacb-669d65136971"}}, "required": ["id", "identity_provider_id"]}}, "required": ["azureAD"], "title": "Azure group"}
```
