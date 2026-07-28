---
title: issuing_cardholder_id_document
page_id: schema-issuing-cardholder-id-document-72a81736
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_cardholder_id_document

```yaml
{"title": "IssuingCardholderIdDocument", "type": "object", "properties": {"back": {"description": "The back of a document returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}, "front": {"description": "The front of a document returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}}, "description": "", "x-expandableFields": ["back", "front"]}
```
