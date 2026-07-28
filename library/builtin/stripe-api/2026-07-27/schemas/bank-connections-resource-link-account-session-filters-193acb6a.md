---
title: bank_connections_resource_link_account_session_filters
page_id: schema-bank-connections-resource-link-account-session-filters-193acb6a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# bank_connections_resource_link_account_session_filters

```yaml
{"title": "BankConnectionsResourceLinkAccountSessionFilters", "type": "object", "properties": {"account_subcategories": {"type": "array", "description": "Restricts the Session to subcategories of accounts that can be linked. Valid subcategories are: `checking`, `savings`, `mortgage`, `line_of_credit`, `credit_card`.", "nullable": true, "items": {"type": "string", "enum": ["checking", "credit_card", "line_of_credit", "mortgage", "savings"]}}, "countries": {"type": "array", "description": "List of countries from which to filter accounts.", "nullable": true, "items": {"maxLength": 5000, "type": "string"}}}, "description": "", "x-expandableFields": []}
```
