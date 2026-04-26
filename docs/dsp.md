# DSP Design  
  
## Data Model  
  
- Sample: `float32` 
- Buffer: `[]float32` 
  
## Processing Model  
  
- In-place processing preferred  
- Avoid memory allocations  
- Deterministic output  
  
## Processor Interface  
  
```go  
type Processor interface {  
    Process(in []float32) []float32  
}
```

Future optimization:

```go
ProcessInPlace(buf []float32)
```

## DSP Chain

Processors are composed sequentially:

```
input → processor1 → processor2 → output
```

## Planned Processors

### Gain

- Multiply signal amplitude

### EQ

- Biquad filters
- Multiple bands

### Compressor (future)

- Dynamic range control

## Example

Gain:

```go
out[i] = in[i] * factor
```

## Performance Rules

- No allocations per buffer
- No copying unless necessary
- Tight loops only
