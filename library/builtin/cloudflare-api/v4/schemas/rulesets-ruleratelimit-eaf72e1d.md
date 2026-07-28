---
title: rulesets_RuleRatelimit
page_id: schema-rulesets-ruleratelimit-eaf72e1d
path: schemas
description: An object configuring the rule's rate limit behavior.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RuleRatelimit

An object configuring the rule's rate limit behavior.

```yaml
{"description": "An object configuring the rule's rate limit behavior.", "type": "object", "properties": {"characteristics": {"description": "Characteristics of the request on which the rate limit counter will be incremented.", "type": "array", "items": {"description": "The characteristic of the request.", "example": "cf.colo.id", "minLength": 1, "title": "Characteristic", "type": "string"}, "minItems": 1, "title": "Characteristics", "uniqueItems": true}, "counting_expression": {"description": "An expression that defines when the rate limit counter should be incremented. It defaults to the same as the rule's expression.", "type": "string", "example": "http.request.body.raw eq \"abcd\"", "minLength": 1, "title": "Counting Expression"}, "mitigation_timeout": {"description": "Period of time in seconds after which the action will be disabled following its first execution.", "type": "integer", "example": 600, "title": "Mitigation Timeout"}, "period": {"description": "Period in seconds over which the counter is being incremented.", "type": "integer", "example": 60, "minimum": 0, "title": "Period"}, "requests_per_period": {"description": "The threshold of requests per period after which the action will be executed for the first time.", "type": "integer", "example": 1000, "minimum": 1, "title": "Requests per Period"}, "requests_to_origin": {"description": "Whether counting is only performed when an origin is reached.", "type": "boolean", "example": true, "default": false, "title": "Requests to Origin"}, "score_per_period": {"description": "The score threshold per period for which the action will be executed the first time.", "type": "integer", "example": 400, "title": "Score per Period"}, "score_response_header_name": {"description": "A response header name provided by the origin, which contains the score to increment rate limit counter with.", "type": "string", "example": "my-score", "minLength": 1, "title": "Score Response Header Name"}}, "required": ["characteristics", "period"], "title": "Rate Limit"}
```
