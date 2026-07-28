---
title: confirmation_tokens_resource_mandate_data_resource_customer_acceptance_resource_online
page_id: schema-confirmation-tokens-resource-mandate-data-resource-customer-acceptance-r-299ecc12
path: schemas
description: This hash contains details about the online acceptance.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# confirmation_tokens_resource_mandate_data_resource_customer_acceptance_resource_online

This hash contains details about the online acceptance.

```yaml
{"title": "ConfirmationTokensResourceMandateDataResourceCustomerAcceptanceResourceOnline", "type": "object", "properties": {"ip_address": {"maxLength": 5000, "type": "string", "description": "The IP address from which the Mandate was accepted by the customer.", "nullable": true}, "user_agent": {"maxLength": 5000, "type": "string", "description": "The user agent of the browser from which the Mandate was accepted by the customer.", "nullable": true}}, "description": "This hash contains details about the online acceptance.", "x-expandableFields": []}
```
