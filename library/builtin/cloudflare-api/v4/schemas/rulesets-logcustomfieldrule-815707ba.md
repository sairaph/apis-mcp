---
title: rulesets_LogCustomFieldRule
page_id: schema-rulesets-logcustomfieldrule-815707ba
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_LogCustomFieldRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["log_custom_field"]}, "action_parameters": {"minProperties": 1, "properties": {"cookie_fields": {"$ref": "#/components/schemas/rulesets_LogCustomFieldCookieFields"}, "raw_response_fields": {"$ref": "#/components/schemas/rulesets_LogCustomFieldRawResponseFields"}, "request_fields": {"$ref": "#/components/schemas/rulesets_LogCustomFieldRequestFields"}, "response_fields": {"$ref": "#/components/schemas/rulesets_LogCustomFieldResponseFields"}, "transformed_request_fields": {"$ref": "#/components/schemas/rulesets_LogCustomFieldTransformedRequestFields"}}}, "description": {"example": "Log additional custom fields."}}, "title": "Log Custom Field Rule"}]}
```
