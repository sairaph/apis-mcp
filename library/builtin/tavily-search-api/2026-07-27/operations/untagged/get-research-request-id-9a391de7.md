---
title: Get research task status and results
page_id: operation-get-research-request-id-7f1bdc7f
path: operations/untagged
description: Retrieve the status and results of a research task using its request ID.
source: https://docs.tavily.com/documentation/api-reference/openapi.json
http_methods:
    - GET
api_endpoints:
    - /research/{request_id}
source_type: openapi
imported_from: https://docs.tavily.com/documentation/api-reference/openapi.json
---

# Get research task status and results

`GET /research/{request_id}`

Retrieve the status and results of a research task using its request ID.

## Definition

```yaml
{"summary": "Get research task status and results", "description": "Retrieve the status and results of a research task using its request ID.", "security": [{"bearerAuth": []}], "x-codeSamples": [{"lang": "python", "label": "Python SDK", "source": "from tavily import TavilyClient\n\ntavily_client = TavilyClient(api_key=\"tvly-YOUR_API_KEY\")\nresponse = tavily_client.get_research(\"123e4567-e89b-12d3-a456-426614174111\")\n\nprint(response)"}, {"lang": "javascript", "label": "JavaScript SDK", "source": "const { tavily } = require(\"@tavily/core\");\n\nconst tvly = tavily({ apiKey: \"tvly-YOUR_API_KEY\" });\nconst response = await tvly.get_research(\"123e4567-e89b-12d3-a456-426614174111\");\n\nconsole.log(response);"}], "parameters": [{"name": "request_id", "in": "path", "required": true, "description": "The unique identifier of the research task.", "schema": {"type": "string"}, "example": "123e4567-e89b-12d3-a456-426614174111"}], "responses": {"200": {"description": "Research task is completed or failed.", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/ResearchTaskCompleted"}, {"$ref": "#/components/schemas/ResearchTaskFailed"}], "discriminator": {"propertyName": "status", "mapping": {"completed": "#/components/schemas/ResearchTaskCompleted", "failed": "#/components/schemas/ResearchTaskFailed"}}}}}}, "202": {"description": "Research task is not yet completed (pending or in_progress).", "content": {"application/json": {"schema": {"type": "object", "properties": {"request_id": {"type": "string", "description": "The unique identifier of the research task.", "example": "123e4567-e89b-12d3-a456-426614174111"}, "status": {"type": "string", "description": "Current status of the research task.", "enum": ["pending", "in_progress"]}, "response_time": {"type": "integer", "description": "Time in seconds it took to complete the request.", "example": 1.23}}, "required": ["request_id", "response_time", "status"]}, "example": {"request_id": "123e4567-e89b-12d3-a456-426614174111", "status": "in_progress", "response_time": 1.23}}}}, "401": {"description": "Unauthorized - Your API key is wrong or missing.", "content": {"application/json": {"schema": {"type": "object", "properties": {"detail": {"type": "object", "properties": {"error": {"type": "string"}}}}}, "example": {"detail": {"error": "Unauthorized: missing or invalid API key."}}}}}, "404": {"description": "Research task not found", "content": {"application/json": {"schema": {"type": "object", "properties": {"detail": {"type": "object", "properties": {"error": {"type": "string"}}}}}, "example": {"detail": {"error": "Research task not found"}}}}}, "500": {"description": "Internal Server Error - We had a problem with our server.", "content": {"application/json": {"schema": {"type": "object", "properties": {"detail": {"type": "object", "properties": {"error": {"type": "string"}}}}}, "example": {"detail": {"error": "Error getting research status"}}}}}}}
```
