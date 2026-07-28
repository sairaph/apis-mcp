---
title: access_external_evaluation_rule
page_id: schema-access-external-evaluation-rule-f2d60537
path: schemas
description: Create Allow or Block policies which evaluate the user based on custom criteria.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_external_evaluation_rule

Create Allow or Block policies which evaluate the user based on custom criteria.

```yaml
{"description": "Create Allow or Block policies which evaluate the user based on custom criteria.", "type": "object", "properties": {"external_evaluation": {"type": "object", "properties": {"evaluate_url": {"description": "The API endpoint containing your business logic.", "type": "string", "example": "https://eval.example.com"}, "keys_url": {"description": "The API endpoint containing the key that Access uses to verify that the response came from your API.", "type": "string", "example": "https://eval.example.com/keys"}}, "required": ["evaluate_url", "keys_url"]}}, "required": ["external_evaluation"], "title": "External Evaluation"}
```
