---
title: Bulk update events
page_id: operation-patch-accounts-account-id-cloudforce-one-events-update-bulk-c68a3d73
path: operations/event
description: Updates multiple events with the same field values. Maximum 100 events per request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/update/bulk
operation_ids:
    - patch_EventUpdateBulk
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Bulk update events

`PATCH /accounts/{account_id}/cloudforce-one/events/update/bulk`

Operation ID: `patch_EventUpdateBulk`

Updates multiple events with the same field values. Maximum 100 events per request.

## Definition

```yaml
{"operationId": "patch_EventUpdateBulk", "summary": "Bulk update events", "description": "Updates multiple events with the same field values. Maximum 100 events per request.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"datasetId": {"description": "Dataset ID containing the events to update. Required to prevent cross-account modifications.", "type": "string", "example": "9b769969-a211-466c-8ac3-cb91266a066a"}, "eventIds": {"description": "List of event UUIDs to update (1-100)", "type": "array", "items": {"example": "12345678-1234-1234-1234-1234567890ab", "type": "string"}, "example": ["uuid-1", "uuid-2", "uuid-3"], "maxItems": 100, "minItems": 1}, "updates": {"description": "Fields to update on all specified events. All fields including 'insight' are supported, except 'date' which requires shard migration.", "type": "object", "properties": {"attacker": {"type": "string", "example": "Flying Yeti", "nullable": true}, "attackerCountry": {"type": "string", "example": "CN"}, "category": {"type": "string", "example": "Domain Resolution"}, "createdAt": {"type": "string", "format": "date-time", "example": "2025-12-19T00:00:00Z"}, "event": {"type": "string", "example": "An attacker registered the domain domain.com"}, "indicator": {"type": "string", "example": "domain2.com"}, "indicatorType": {"type": "string", "example": "domain"}, "insight": {"type": "string", "example": "This event indicates a potential phishing campaign"}, "raw": {"type": "object", "properties": {"data": {"type": "object", "additionalProperties": true, "nullable": true}, "source": {"type": "string", "example": "example.com"}, "tlp": {"type": "string", "example": "amber"}}}, "targetCountry": {"type": "string", "example": "US"}, "targetIndustry": {"type": "string", "example": "Insurance"}, "tlp": {"type": "string", "example": "amber"}}}}, "required": ["eventIds", "datasetId", "updates"]}}}}, "responses": {"200": {"description": "Returns the count of updated events and any failures.", "content": {"application/json": {"schema": {"type": "object", "properties": {"failedCount": {"type": "number", "example": 1}, "failures": {"description": "List of events that failed to update with error messages", "type": "array", "items": {"properties": {"error": {"type": "string", "example": "Event not found"}, "eventId": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}}, "required": ["eventId", "error"], "type": "object"}}, "updatedCount": {"type": "number", "example": 5}}, "required": ["updatedCount", "failedCount"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write"]}
```
