---
title: Create a relationship between two events
page_id: operation-post-accounts-account-id-cloudforce-one-events-relationships-create-fa7ac5d7
path: operations/event
description: Creates a directed relationship between two events. The relationship is from parent to child with a specified type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/relationships/create
operation_ids:
    - post_CreateEventRelationship
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a relationship between two events

`POST /accounts/{account_id}/cloudforce-one/events/relationships/create`

Operation ID: `post_CreateEventRelationship`

Creates a directed relationship between two events. The relationship is from parent to child with a specified type.

## Definition

```yaml
{"operationId": "post_CreateEventRelationship", "summary": "Create a relationship between two events", "description": "Creates a directed relationship between two events. The relationship is from parent to child with a specified type.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"childIds": {"description": "Array of UUIDs for child events. Single child = 1:1 relationship, multiple = 1:many relationships", "type": "array", "items": {"format": "uuid", "type": "string"}, "minItems": 1}, "datasetId": {"description": "Dataset identifier where the events are stored", "type": "string"}, "parentId": {"description": "UUID of the parent event that will be the source of the relationship", "type": "string", "format": "uuid"}, "relationshipType": {"description": "Type of relationship to create between parent and child events", "type": "string", "enum": ["related_to", "caused_by", "attributed_to"]}}, "required": ["parentId", "childIds", "relationshipType", "datasetId"]}}}}, "responses": {"200": {"description": "Relationship created successfully", "content": {"application/json": {"schema": {"type": "object", "properties": {"childIds": {"description": "Array of child event UUIDs that were processed", "type": "array", "items": {"format": "uuid", "type": "string"}}, "errors": {"description": "Array of errors for relationships that failed to be created (only present if some relationships failed)", "type": "array", "items": {"properties": {"childId": {"description": "UUID of the child event that failed to create a relationship", "type": "string", "format": "uuid"}, "error": {"description": "Error message describing why the relationship creation failed", "type": "string"}, "errorType": {"description": "Type/category of the error that occurred", "type": "string"}}, "required": ["childId", "error"], "type": "object"}}, "message": {"description": "Human-readable message describing the operation result", "type": "string"}, "relationships": {"description": "Array of successfully created relationship objects", "type": "array", "items": {"properties": {"childDatasetId": {"description": "Dataset ID where the child event resides", "type": "string"}, "childId": {"description": "UUID of the child event in the relationship", "type": "string", "format": "uuid"}, "parentDatasetId": {"description": "Dataset ID where the parent event resides", "type": "string"}, "parentId": {"description": "UUID of the parent event in the relationship", "type": "string", "format": "uuid"}, "relationshipType": {"description": "Type of relationship between the events", "type": "string", "enum": ["related_to", "caused_by", "attributed_to"]}}, "required": ["parentId", "childId", "relationshipType", "parentDatasetId", "childDatasetId"], "type": "object"}}, "relationshipsCreated": {"description": "Number of relationships that were successfully created", "type": "number"}, "success": {"description": "Whether the relationship creation operation completed successfully", "type": "boolean"}}, "required": ["success", "message", "relationships"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
