---
title: Submit a feedback report
page_id: operation-post-zones-zone-id-bot-management-feedback-03121c04
path: operations/feedback
description: |-
    Submit a feedback report for the specified zone. Use `type` to indicate whether the report is a false positive (good traffic flagged as bot) or a false negative (bot traffic missed). Furthermore, you can also use `expression` as a wirefilter to identify the affected traffic sample.

    See more accepted API fields and expression types at https://developers.cloudflare.com/bots/concepts/feedback-loop/#api-fields and https://developers.cloudflare.com/bots/concepts/feedback-loop/#expression-fields, respectively.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/bot_management/feedback
operation_ids:
    - bot-management-zone-feedback-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Submit a feedback report

`POST /zones/{zone_id}/bot_management/feedback`

Operation ID: `bot-management-zone-feedback-create`

Submit a feedback report for the specified zone. Use `type` to indicate whether the report is a false positive (good traffic flagged as bot) or a false negative (bot traffic missed). Furthermore, you can also use `expression` as a wirefilter to identify the affected traffic sample.

See more accepted API fields and expression types at https://developers.cloudflare.com/bots/concepts/feedback-loop/#api-fields and https://developers.cloudflare.com/bots/concepts/feedback-loop/#expression-fields, respectively.

## Definition

```yaml
{"operationId": "bot-management-zone-feedback-create", "summary": "Submit a feedback report", "description": "Submit a feedback report for the specified zone. Use `type` to indicate whether the report is a false positive (good traffic flagged as bot) or a false negative (bot traffic missed). Furthermore, you can also use `expression` as a wirefilter to identify the affected traffic sample.\n\nSee more accepted API fields and expression types at https://developers.cloudflare.com/bots/concepts/feedback-loop/#api-fields and https://developers.cloudflare.com/bots/concepts/feedback-loop/#expression-fields, respectively.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bot-management_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"false_negative_example": {"summary": "False negative example", "value": {"description": "Automated scraping missed by detections", "expression": "http.host eq 'www.example.com' and http.request.uri.path starts_with '/products' and cf.bot_management.score gt 25", "first_request_seen_at": "2025-09-29T00:00:00Z", "last_request_seen_at": "2025-09-29T06:00:00Z", "requests": 2000, "requests_by_attribute": {"topIPs": [{"metric": "203.0.113.55", "requests": 400}], "topJA3Hashes": [{"metric": "ab12cd34ef56...", "requests": 900}]}, "requests_by_score": {"30": 800, "40": 700, "50": 500}, "requests_by_score_src": {"heuristics": 200, "ml": 1800}, "subtype": "Scraping", "type": "false_negative"}}, "false_positive_example": {"summary": "False positive example", "value": {"description": "Legitimate users flagged as bots during login", "expression": "http.host eq 'app.example.com' and http.request.uri.path starts_with '/login' and cf.bot_management.score lt 5", "first_request_seen_at": "2025-09-30T10:00:00Z", "last_request_seen_at": "2025-09-30T11:00:00Z", "requests": 500, "requests_by_attribute": {"topIPs": [{"metric": "198.51.100.23", "requests": 60}], "topUserAgents": [{"metric": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36", "requests": 220}]}, "requests_by_score": {"1": 120, "2": 150, "3": 130, "4": 100}, "requests_by_score_src": {"heuristics": 50, "ml": 450}, "subtype": "Login Abuse", "type": "false_positive"}}}, "schema": {"$ref": "#/components/schemas/bot-management_feedback_report"}}}}, "responses": {"201": {"description": "Feedback report created"}, "4XX": {"description": "Feedback creation failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bot-management_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Feedback"], "x-api-token-group": ["Bot Management Feedback Report Write", "Bot Management Feedback Report Read", "Bot Management Write", "Bot Management Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
