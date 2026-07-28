---
title: bank_connections_resource_account_status_details_api_resource_active_status_details
page_id: schema-bank-connections-resource-account-status-details-api-resource-active-sta-dff73b47
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# bank_connections_resource_account_status_details_api_resource_active_status_details

```yaml
{"title": "BankConnectionsResourceAccountStatusDetailsAPIResourceActiveStatusDetails", "required": ["action", "cause", "expected_deactivation_date"], "type": "object", "properties": {"action": {"type": "string", "description": "The action (if any) to proactively relink the Account.", "enum": ["none", "relink_required"]}, "cause": {"type": "string", "description": "The underlying cause of the Account becoming inactive.", "enum": ["access_expired", "institution_requirement", "unspecified"]}, "expected_deactivation_date": {"type": "integer", "description": "When the Account is expected to become inactive, if applicable.", "format": "unix-time"}}, "description": "", "x-expandableFields": []}
```
