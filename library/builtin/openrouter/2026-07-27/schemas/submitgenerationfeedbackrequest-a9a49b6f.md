---
title: SubmitGenerationFeedbackRequest
page_id: schema-submitgenerationfeedbackrequest-a9a49b6f
path: schemas
description: Structured feedback about a specific generation
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# SubmitGenerationFeedbackRequest

Structured feedback about a specific generation

```yaml
{"description": "Structured feedback about a specific generation", "example": {"category": "incorrect_response", "comment": "The model repeated the same paragraph three times.", "generation_id": "gen-3bhGkxlo4XFrqiabUM7NDtwDzWwG"}, "properties": {"category": {"description": "The category of feedback being reported", "enum": ["latency", "incoherence", "incorrect_response", "formatting", "billing", "api_error", "other"], "example": "incorrect_response", "type": "string", "x-speakeasy-unknown-values": "allow"}, "comment": {"description": "An optional free-text comment describing the feedback", "example": "The model repeated the same paragraph three times.", "maxLength": 1000, "type": "string"}, "generation_id": {"description": "The generation to submit feedback on", "example": "gen-3bhGkxlo4XFrqiabUM7NDtwDzWwG", "minLength": 1, "type": "string"}}, "required": ["generation_id", "category"], "type": "object"}
```
