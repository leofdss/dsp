# Architectural Decisions  
  
## Use Hexagonal Architecture

Reason:
- Isolate DSP from PipeWire
- Improve testability

## Use Go (initially)

Reason:
- Fast development
- Strong tooling
## Consider Rust for DSP

Reason:
- Avoid GC in real-time path
- Deterministic performance

## In-place Processing

Reason:
- Reduce allocations
- Lower latency

## Config-driven DSP

Reason:
- Flexible pipeline definition
