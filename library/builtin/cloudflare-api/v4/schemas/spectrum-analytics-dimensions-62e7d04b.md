---
title: spectrum-analytics_dimensions
page_id: schema-spectrum-analytics-dimensions-62e7d04b
path: schemas
description: |-
    Can be used to break down the data by given attributes. Options are:

    Dimension                 | Name                            | Example
    --------------------------|---------------------------------|--------------------------
    event                     | Connection Event                | connect, progress, disconnect, originError, clientFiltered
    appID                     | Application ID                  | 40d67c87c6cd4b889a4fd57805225e85
    coloName                  | Colo Name                       | SFO
    ipVersion                 | IP version used by the client   | 4, 6.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# spectrum-analytics_dimensions

Can be used to break down the data by given attributes. Options are:

Dimension                 | Name                            | Example
--------------------------|---------------------------------|--------------------------
event                     | Connection Event                | connect, progress, disconnect, originError, clientFiltered
appID                     | Application ID                  | 40d67c87c6cd4b889a4fd57805225e85
coloName                  | Colo Name                       | SFO
ipVersion                 | IP version used by the client   | 4, 6.

```yaml
{"description": "Can be used to break down the data by given attributes. Options are:\n\nDimension                 | Name                            | Example\n--------------------------|---------------------------------|--------------------------\nevent                     | Connection Event                | connect, progress, disconnect, originError, clientFiltered\nappID                     | Application ID                  | 40d67c87c6cd4b889a4fd57805225e85\ncoloName                  | Colo Name                       | SFO\nipVersion                 | IP version used by the client   | 4, 6.", "type": "array", "items": {"enum": ["event", "appID", "coloName", "ipVersion"], "type": "string"}, "example": ["event", "appID"]}
```
