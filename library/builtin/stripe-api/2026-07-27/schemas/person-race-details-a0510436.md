---
title: person_race_details
page_id: schema-person-race-details-a0510436
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# person_race_details

```yaml
{"title": "PersonRaceDetails", "type": "object", "properties": {"race": {"type": "array", "description": "The persons race.", "nullable": true, "items": {"type": "string", "enum": ["african_american", "american_indian_or_alaska_native", "asian", "asian_indian", "black_or_african_american", "chinese", "ethiopian", "filipino", "guamanian_or_chamorro", "haitian", "jamaican", "japanese", "korean", "native_hawaiian", "native_hawaiian_or_other_pacific_islander", "nigerian", "other_asian", "other_black_or_african_american", "other_pacific_islander", "prefer_not_to_answer", "samoan", "somali", "vietnamese", "white"]}}, "race_other": {"maxLength": 5000, "type": "string", "description": "Please specify your race, when other is selected.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
