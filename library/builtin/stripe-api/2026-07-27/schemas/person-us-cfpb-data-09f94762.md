---
title: person_us_cfpb_data
page_id: schema-person-us-cfpb-data-09f94762
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# person_us_cfpb_data

```yaml
{"title": "PersonUSCfpbData", "type": "object", "properties": {"ethnicity_details": {"description": "The persons ethnicity details", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/person_ethnicity_details"}]}, "race_details": {"description": "The persons race details", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/person_race_details"}]}, "self_identified_gender": {"maxLength": 5000, "type": "string", "description": "The persons self-identified gender", "nullable": true}}, "description": "", "x-expandableFields": ["ethnicity_details", "race_details"]}
```
