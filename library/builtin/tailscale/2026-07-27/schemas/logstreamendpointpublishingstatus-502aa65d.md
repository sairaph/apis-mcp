---
title: LogstreamEndpointPublishingStatus
page_id: schema-logstreamendpointpublishingstatus-502aa65d
path: schemas
description: Latest status of log stream publishing for a specific type of log.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# LogstreamEndpointPublishingStatus

Latest status of log stream publishing for a specific type of log.

```yaml
type: object
description: Latest status of log stream publishing for a specific type of log.
example:
    lastActivity: '2024-06-10T15:42:13.984555636Z'
    lastError: ''
    maxBodySize: 524288
    numBytesSent: 17238983
    numEntriesSent: 8363
    numFailedRequests: 5434
    numSpoofedEntries: 0
    numTotalRequests: 10610
    rateBytesSent: 3.524073767296142
    rateEntriesSent: 0.008564949767446907
    rateFailedRequests: 4.1431119220540763e-157
    rateTotalRequests: 0.0037038341100629453
properties:
    lastActivity:
        type: string
        description: Timestamp of the most recent publishing activity, in RFC 3339 format.
        example: '2024-06-10T15:42:13.984555636Z'
    lastError:
        type: string
        description: The most recent error (if any).
        example: Something went wrong.
    maxBodySize:
        type: integer
        description: The size of the largest single request body.
        example: 524288
    numBytesSent:
        type: integer
        description: Total bytes published across all requests.
        example: 17238983
    numEntriesSent:
        type: integer
        description: The total number of entries published.
        example: 8363
    numSpoofedEntries:
        type: integer
        description: The number of spoofed entries published. A spoofed entry is one that failed to validate because we did not see receive a matching flow log from the other side of the connection.
        example: 0
    numTotalRequests:
        type: integer
        description: The total number of requests made to the streaming endpoint.
        example: 10610
    numFailedRequests:
        type: integer
        description: The total number of requests to the streaming endpoint that have failed.
        example: 5434
    rateBytesSent:
        type: number
        description: The exponentially weighted moving average rate at which data is being streamed to the endpoint, in bytes per second.
        example: 3.524073767296142
    rateEntriesSent:
        type: number
        description: The exponentially weighted moving average rate at which entries are being sent to the endpoint, in entries per second.
        example: 0.008564949767446907
    rateTotalRequests:
        type: number
        description: The exponentially weighted moving average rate at which requests are being made to the endpoint, in requests per second.
        example: 0.0037038341100629453
    rateFailedRequests:
        type: number
        description: The exponentially weighted moving average rate at which requests are failing, in requests per second.
        example: 4.1431119220540763e-157
required:
    - lastActivity
    - lastError
    - maxBodySize
    - numBytesSent
    - numEntriesSent
    - numSpoofedEntries
    - numTotalRequests
    - numFailedRequests
    - rateBytesSent
    - rateEntriesSent
    - rateTotalRequests
    - rateFailedRequests
```
