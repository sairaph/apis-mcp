---
title: access_saml_group_rule
page_id: schema-access-saml-group-rule-054f4cf6
path: schemas
description: |-
    Matches a SAML group.
    Requires a SAML identity provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_saml_group_rule

Matches a SAML group.
Requires a SAML identity provider.

```yaml
{"description": "Matches a SAML group.\nRequires a SAML identity provider.", "type": "object", "properties": {"saml": {"type": "object", "properties": {"attribute_name": {"description": "The name of the SAML attribute.", "type": "string", "example": "group"}, "attribute_value": {"description": "The SAML attribute value to look for.", "type": "string", "example": "devs@cloudflare.com"}, "identity_provider_id": {"description": "The ID of your SAML identity provider.", "type": "string", "example": "ea85612a-29c8-46c2-bacb-669d65136971"}}, "required": ["attribute_name", "attribute_value", "identity_provider_id"]}}, "required": ["saml"], "title": "SAML group"}
```
