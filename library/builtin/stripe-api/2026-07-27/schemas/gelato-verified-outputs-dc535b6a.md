---
title: gelato_verified_outputs
page_id: schema-gelato-verified-outputs-dc535b6a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# gelato_verified_outputs

```yaml
{"title": "GelatoVerifiedOutputs", "type": "object", "properties": {"address": {"description": "The user's verified address.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/address"}]}, "dob": {"description": "The user’s verified date of birth.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/gelato_data_verified_outputs_date"}]}, "email": {"maxLength": 5000, "type": "string", "description": "The user's verified email address", "nullable": true}, "first_name": {"maxLength": 5000, "type": "string", "description": "The user's verified first name.", "nullable": true}, "id_number": {"maxLength": 5000, "type": "string", "description": "The user's verified id number.", "nullable": true}, "id_number_type": {"type": "string", "description": "The user's verified id number type.", "nullable": true, "enum": ["br_cpf", "sg_nric", "us_ssn"]}, "last_name": {"maxLength": 5000, "type": "string", "description": "The user's verified last name.", "nullable": true}, "phone": {"maxLength": 5000, "type": "string", "description": "The user's verified phone number", "nullable": true}, "sex": {"type": "string", "description": "The user's verified sex.", "nullable": true, "enum": ["[redacted]", "female", "male", "unknown"]}, "unparsed_place_of_birth": {"maxLength": 5000, "type": "string", "description": "The user's verified place of birth as it appears in the document.", "nullable": true}, "unparsed_sex": {"maxLength": 5000, "type": "string", "description": "The user's verified sex as it appears in the document.", "nullable": true}}, "description": "", "x-expandableFields": ["address", "dob"]}
```
