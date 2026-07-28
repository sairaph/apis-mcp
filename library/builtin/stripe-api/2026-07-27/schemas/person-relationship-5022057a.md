---
title: person_relationship
page_id: schema-person-relationship-5022057a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# person_relationship

```yaml
{"title": "PersonRelationship", "type": "object", "properties": {"authorizer": {"type": "boolean", "description": "Whether the person is the authorizer of the account's representative.", "nullable": true}, "director": {"type": "boolean", "description": "Whether the person is a director of the account's legal entity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.", "nullable": true}, "executive": {"type": "boolean", "description": "Whether the person has significant responsibility to control, manage, or direct the organization.", "nullable": true}, "legal_guardian": {"type": "boolean", "description": "Whether the person is the legal guardian of the account's representative.", "nullable": true}, "owner": {"type": "boolean", "description": "Whether the person is an owner of the account’s legal entity.", "nullable": true}, "percent_ownership": {"type": "number", "description": "The percent owned by the person of the account's legal entity.", "nullable": true}, "representative": {"type": "boolean", "description": "Whether the person is authorized as the primary representative of the account. This is the person nominated by the business to provide information about themselves, and general information about the account. There can only be one representative at any given time. At the time the account is created, this person should be set to the person responsible for opening the account.", "nullable": true}, "title": {"maxLength": 5000, "type": "string", "description": "The person's title (e.g., CEO, Support Engineer).", "nullable": true}}, "description": "", "x-expandableFields": []}
```
