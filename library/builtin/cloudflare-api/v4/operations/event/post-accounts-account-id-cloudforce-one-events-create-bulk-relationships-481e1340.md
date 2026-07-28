---
title: Creates bulk DOS event with relationships and indicators
page_id: operation-post-accounts-account-id-cloudforce-one-events-create-bulk-relationships-44887fa1
path: operations/event
description: This method is deprecated. Please use `event_create_bulk` instead
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/create/bulk/relationships
operation_ids:
    - post_DOSEventCreateBulkWithRelationships
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates bulk DOS event with relationships and indicators

`POST /accounts/{account_id}/cloudforce-one/events/create/bulk/relationships`

Operation ID: `post_DOSEventCreateBulkWithRelationships`

This method is deprecated. Please use `event_create_bulk` instead

## Definition

```yaml
{"operationId": "post_DOSEventCreateBulkWithRelationships", "summary": "Creates bulk DOS event with relationships and indicators", "description": "This method is deprecated. Please use `event_create_bulk` instead", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"data": {"type": "array", "items": {"properties": {"accountId": {"type": "number", "example": 123456}, "attacker": {"type": "string", "example": "Flying Yeti", "nullable": true}, "attackerCountry": {"type": "string", "example": "CN"}, "category": {"type": "string", "example": "Domain Resolution"}, "datasetId": {"type": "string", "example": "durableObjectName"}, "date": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "event": {"type": "string", "example": "An attacker registered the domain domain.com"}, "indicator": {"type": "string", "example": "domain.com"}, "indicatorType": {"type": "string", "example": "domain"}, "indicators": {"description": "Array of indicators for this event. Supports multiple indicators per event for complex scenarios.", "type": "array", "items": {"properties": {"indicatorType": {"description": "The type of indicator (e.g., DOMAIN, IP, JA3, HASH)", "type": "string", "example": "domain"}, "value": {"description": "The indicator value (e.g., domain name, IP address, hash)", "type": "string", "example": "malicious.com"}}, "required": ["value", "indicatorType"], "type": "object"}}, "insight": {"type": "string", "example": "This domain was likely registered for phishing purposes"}, "raw": {"type": "object", "properties": {"data": {"type": "object", "additionalProperties": true, "nullable": true}, "source": {"type": "string", "example": "example.com"}, "tlp": {"type": "string", "example": "amber"}}, "required": ["data"]}, "tags": {"type": "array", "items": {"example": "malware", "type": "string"}}, "targetCountry": {"type": "string", "example": "US"}, "targetIndustry": {"type": "string", "example": "Agriculture"}, "tlp": {"type": "string", "example": "amber"}}, "required": ["date", "category", "event", "tlp", "raw"], "type": "object"}}, "datasetId": {"type": "string", "example": "durableObjectName"}}, "required": ["data", "datasetId"]}}}}, "responses": {"200": {"description": "Returns the number of created bulk events with relationships.", "content": {"application/json": {"schema": {"description": "Result of bulk relationship creation operation", "type": "object", "properties": {"createdEventsCount": {"description": "Number of events created", "type": "number"}, "createdIndicatorsCount": {"description": "Number of indicators created", "type": "number"}, "createdRelationshipsCount": {"description": "Number of relationships created", "type": "number"}, "errorCount": {"description": "Number of errors encountered", "type": "number"}, "errors": {"description": "Array of error details", "type": "array", "items": {"properties": {"error": {"description": "Error message", "type": "string"}, "eventIndex": {"description": "Index of the event that caused the error", "type": "number"}}, "required": ["eventIndex", "error"], "type": "object"}}}, "required": ["createdEventsCount", "createdIndicatorsCount", "createdRelationshipsCount", "errorCount"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "deprecated": true, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
