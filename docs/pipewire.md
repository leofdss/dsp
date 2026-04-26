# PipeWire Integration  
  
## Goal

Provide real-time audio input/output with minimal latency.

## Strategy
  
- Use PipeWire as audio backend
- Implement as adapter
- Keep it isolated from DSP core

## Processing Model
  
PipeWire callback:
  
1. Receive audio buffer
2. Convert to []float32
3. Call DSP processor
4. Write output buffer

## Constraints
  
- No allocations in callback
- No blocking operations
- No logging in hot path

## Risks
  
- GC pauses (Go)
- Buffer underruns (xruns)
- Thread scheduling issues

## Future Improvements
  
- Zero-copy buffers
- SIMD optimizations
- Possible Rust/C DSP core
