---
title: v2.error.purpose_of_funds_description_must_be_empty_for_non_other_purpose_of_funds
page_id: schema-v2-error-purpose-of-funds-description-must-be-empty-for-non-other-purpos-4fc67e3c
path: schemas
description: PurposeOfFundsDescription is not empty while PurposeOfFunds is not OTHER.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# v2.error.purpose_of_funds_description_must_be_empty_for_non_other_purpose_of_funds

PurposeOfFundsDescription is not empty while PurposeOfFunds is not OTHER.

```yaml
{"required": ["error"], "type": "object", "properties": {"error": {"required": ["code", "message"], "type": "object", "properties": {"code": {"type": "string", "description": "Short code to identify the error, should not be handled programmatically", "enum": ["purpose_of_funds_description_must_be_empty_for_non_other_purpose_of_funds"]}, "message": {"type": "string", "description": "A human-readable message providing more details about the error"}}, "description": "Information about the error that occurred"}}, "description": "PurposeOfFundsDescription is not empty while PurposeOfFunds is not OTHER."}
```
