---
title: confirmation_tokens_resource_mandate_data_resource_customer_acceptance
page_id: schema-confirmation-tokens-resource-mandate-data-resource-customer-acceptance-5f4ec9d5
path: schemas
description: This hash contains details about the customer acceptance of the Mandate.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# confirmation_tokens_resource_mandate_data_resource_customer_acceptance

This hash contains details about the customer acceptance of the Mandate.

```yaml
{"title": "ConfirmationTokensResourceMandateDataResourceCustomerAcceptance", "required": ["type"], "type": "object", "properties": {"online": {"description": "If this is a Mandate accepted online, this hash contains details about the online acceptance.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/confirmation_tokens_resource_mandate_data_resource_customer_acceptance_resource_online"}]}, "type": {"maxLength": 5000, "type": "string", "description": "The type of customer acceptance information included with the Mandate."}}, "description": "This hash contains details about the customer acceptance of the Mandate.", "x-expandableFields": ["online"]}
```
