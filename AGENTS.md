## Project Overview  
CLI tool for real-time audio DSP using PipeWire.  
  
## Architecture Rules  
- Use hexagonal architecture  
- DSP must be isolated from I/O  
- PipeWire only inside adapters  
- CLI must not contain DSP logic  
  
## DSP Rules (STRICT)  
- Processing must be in-place when possible  
- Avoid allocations in hot path  
- Use float32 buffers  
- Deterministic behavior only  
  
## Performance Constraints  
- No allocations inside audio callback  
- No locks in processing loop  
- Avoid interface dispatch in hot path  
- Prefer simple loops over abstractions  
  
## Project Structure  
- cmd/ → CLI  
- internal/domain → DSP core  
- internal/app → orchestration (Engine)  
- internal/ports → interfaces  
- internal/adapters → implementations (pipewire, file, fake)  
  
## Testing Rules  
- TDD is required  
- DSP must have unit tests  
- No real audio backend in tests  
- Use fake adapters for integration tests  
  
## Coding Guidelines  
- Keep functions small and explicit  
- Avoid unnecessary abstractions  
- Prefer clarity over cleverness  
- Ask before introducing new patterns  
  
## Forbidden  
- Mixing DSP with I/O  
- Allocating per buffer  
- Adding hidden global state  
- Breaking architecture boundaries  
  
