---
title: treasury_financial_accounts_resource_platform_restrictions
page_id: schema-treasury-financial-accounts-resource-platform-restrictions-a068daed
path: schemas
description: Restrictions that a Connect Platform has placed on this FinancialAccount.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# treasury_financial_accounts_resource_platform_restrictions

Restrictions that a Connect Platform has placed on this FinancialAccount.

```yaml
{"title": "TreasuryFinancialAccountsResourcePlatformRestrictions", "type": "object", "properties": {"inbound_flows": {"type": "string", "description": "Restricts all inbound money movement.", "nullable": true, "enum": ["restricted", "unrestricted"]}, "outbound_flows": {"type": "string", "description": "Restricts all outbound money movement.", "nullable": true, "enum": ["restricted", "unrestricted"]}}, "description": "Restrictions that a Connect Platform has placed on this FinancialAccount.", "x-expandableFields": []}
```
