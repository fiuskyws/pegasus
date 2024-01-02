# pegasus
Simple Message Broker service written in Go.

## POC Releases

### Base - POC Release 1
- DONE!
- [x] Pub/Sub Logic
- [x] Server
- [x] Makefile:
    - [x] `run`
    - [x] `test`

### Event Sourcing - POC Release 2 
- In Development
- [ ] Event Sourcing (from Thoth, Blob, etc.)

### Polishing - PoC Final Release (???)
- Not Started

## Backlog
- Message Broker Key Features (asked ChatGPT for it)

#### Messaging Patterns
- [x] Implement Publish-Subscribe (Pub/Sub) pattern
    - Kind of implemented at `POC R1`
- [ ] Implement Point-to-Point (P2P) pattern

#### Message Queues
- [ ] Implement Queue Management
- [ ] Implement Queue Persistence

#### Message Durability
- [ ] Implement Message Persistence

#### Reliability and Fault Tolerance
- [ ] Implement High Availability
- [ ] Implement Replication

#### Scalability
- [ ] Implement Horizontal Scaling
- [ ] Implement Partitioning/Sharding

#### Message Formats
- [ ] Support Multiple Message Formats

#### Message Transformation
- [ ] Implement Message Transformation

#### Security
- [ ] Implement Authentication and Authorization
- [ ] Implement Encryption

#### Monitoring and Management
- [ ] Implement Metrics and Logging
- [ ] Implement Administrative Interfaces

#### Compatibility and Integration
- [ ] Provide Client Libraries
- [ ] Ensure Integration with Other Systems
    - Test with other FWS

#### Message Acknowledgment
- [ ] Implement Acknowledgment Mechanisms

#### Dead Letter Queues (DLQ)
- [ ] Implement Dead Letter Queues

#### Event Sourcing (if applicable)
- [ ] Implement Event Sourcing
    - Test with other FWS(e.g [Thoth](https://github.com/fiuskyws/thoth), [Blob](https://github.com/fiuskyws/blob), etc.)

#### Support for Retries and Redelivery
- [ ] Implement Retry Policies

## Disclaimer

This repository still in research and `PoC's` phase, meaning, it is NOT ready to be used in production environments. (if it'll be)
