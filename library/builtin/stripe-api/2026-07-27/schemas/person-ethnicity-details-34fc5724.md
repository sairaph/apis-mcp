---
title: person_ethnicity_details
page_id: schema-person-ethnicity-details-34fc5724
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# person_ethnicity_details

```yaml
{"title": "PersonEthnicityDetails", "type": "object", "properties": {"ethnicity": {"type": "array", "description": "The persons ethnicity", "nullable": true, "items": {"type": "string", "enum": ["cuban", "hispanic_or_latino", "mexican", "not_hispanic_or_latino", "other_hispanic_or_latino", "prefer_not_to_answer", "puerto_rican"]}}, "ethnicity_other": {"maxLength": 5000, "type": "string", "description": "Please specify your origin, when other is selected.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
