---
title: zero-trust-gateway_extended-email-matching
page_id: schema-zero-trust-gateway-extended-email-matching-124a0e8a
path: schemas
description: Configures user email settings for firewall policies. When you enable this, the system standardizes email addresses in the identity portion of the rule to match extended email variants in firewall policies. When you disable this setting, the system matches email addresses exactly as you provide them. Enable this setting if your email uses `.` or `+` modifiers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_extended-email-matching

Configures user email settings for firewall policies. When you enable this, the system standardizes email addresses in the identity portion of the rule to match extended email variants in firewall policies. When you disable this setting, the system matches email addresses exactly as you provide them. Enable this setting if your email uses `.` or `+` modifiers.

```yaml
{"description": "Configures user email settings for firewall policies. When you enable this, the system standardizes email addresses in the identity portion of the rule to match extended email variants in firewall policies. When you disable this setting, the system matches email addresses exactly as you provide them. Enable this setting if your email uses `.` or `+` modifiers.", "type": "object", "properties": {"enabled": {"description": "Specify whether to match all variants of user emails (with + or . modifiers) used as criteria in Firewall policies.", "type": "boolean", "example": true, "nullable": true, "x-auditable": true}, "read_only": {"description": "Indicate that this setting was shared via the Orgs API and read only for the current account.", "type": "boolean", "readOnly": true, "x-auditable": true, "x-stainless-terraform-configurability": "computed"}, "source_account": {"description": "Indicate the account tag of the account that shared this setting.", "type": "string", "readOnly": true, "x-auditable": true, "x-stainless-terraform-configurability": "computed"}, "version": {"description": "Indicate the version number of the setting.", "type": "integer", "example": 1, "readOnly": true, "x-auditable": true, "x-stainless-terraform-configurability": "computed"}}, "nullable": true, "x-stainless-terraform-configurability": "optional"}
```
