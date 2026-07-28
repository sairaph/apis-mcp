---
title: List zone feedback reports
page_id: operation-get-zones-zone-id-bot-management-feedback-c4fc9961
path: operations/feedback
description: Returns all feedback reports previously submitted for the specified zone. Feedback reports help improve detection by sharing samples of traffic that were misclassified as bots or humans.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/bot_management/feedback
operation_ids:
    - bot-management-zone-feedback-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List zone feedback reports

`GET /zones/{zone_id}/bot_management/feedback`

Operation ID: `bot-management-zone-feedback-list`

Returns all feedback reports previously submitted for the specified zone. Feedback reports help improve detection by sharing samples of traffic that were misclassified as bots or humans.

## Definition

```yaml
{"operationId": "bot-management-zone-feedback-list", "summary": "List zone feedback reports", "description": "Returns all feedback reports previously submitted for the specified zone. Feedback reports help improve detection by sharing samples of traffic that were misclassified as bots or humans.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bot-management_identifier"}}], "responses": {"200": {"description": "List of feedback reports", "content": {"application/json": {"examples": {"sample": {"summary": "Example list response", "value": [{"created_at": "2025-10-01T12:00:00Z", "description": "Legitimate checkout traffic was blocked as bots", "expression": "(http.host eq 'shop.example.com' and http.request.uri.path starts_with '/checkout') and cf.bot_management.score lt 5", "first_request_seen_at": "2025-09-30T08:00:00Z", "last_request_seen_at": "2025-09-30T09:00:00Z", "requests": 1200, "requests_by_attribute": {"topIPs": [{"metric": "203.0.113.10", "requests": 180}, {"metric": "203.0.113.11", "requests": 150}], "topPaths": [{"metric": "/checkout", "requests": 1000}]}, "requests_by_score": {"1": 200, "2": 300, "3": 400, "4": 300}, "requests_by_score_src": {"heuristics": 200, "machine_learning": 1000}, "subtype": "Spamming", "type": "false_positive"}]}}, "schema": {"type": "array", "items": {"$ref": "#/components/schemas/bot-management_feedback_report"}}}}}, "4XX": {"description": "Feedback list failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bot-management_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Feedback"], "x-api-token-group": ["Bot Management Feedback Report Write", "Bot Management Feedback Report Read", "Bot Management Write", "Bot Management Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
