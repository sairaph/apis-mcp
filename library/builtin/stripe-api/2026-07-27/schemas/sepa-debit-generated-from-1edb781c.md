---
title: sepa_debit_generated_from
page_id: schema-sepa-debit-generated-from-1edb781c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# sepa_debit_generated_from

```yaml
{"title": "sepa_debit_generated_from", "type": "object", "properties": {"charge": {"description": "The ID of the Charge that generated this PaymentMethod, if any.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/charge"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/charge"}]}}, "setup_attempt": {"description": "The ID of the SetupAttempt that generated this PaymentMethod, if any.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/setup_attempt"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/setup_attempt"}]}}}, "description": "", "x-expandableFields": ["charge", "setup_attempt"]}
```
