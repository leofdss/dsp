# Testing Strategy  
  
## Principles  
  
- Test-driven development (TDD)  
- DSP tested in isolation  
- No real audio dependencies  
  
---  
  
## Unit Tests  
  
Test individual processors:  
  
- Gain doubles amplitude  
- EQ attenuates frequencies  
  
---  
  
## Integration Tests  
  
- Use fake audio stream  
- Validate processing pipeline  
  
---  
  
## Example  
  
```go  
func TestGain(t *testing.T) {  
    g := Gain{Factor: 2}  
  
    in := []float32{1, 2, 3}  
    out := g.Process(in)  
  
    expected := []float32{2, 4, 6}  
    require.Equal(t, expected, out)  
}
```

---

## Tools

- go test
- table-driven tests

---

## Rules

- No PipeWire in tests
- No timing-dependent tests