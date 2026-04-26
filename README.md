# DSP CLI for PipeWire

Low-latency audio DSP engine (EQ, filters) for Linux PipeWire.  

## Goals  
- Real-time audio processing
- Lowest possible latency
- Config-driven DSP chain
- Strong TDD discipline

## Example

```bash  
dsp ./eq.yaml
```

```yaml
audio:
  sample_rate: 48000        # Hz
  buffer_size: 256          # Frames per buffer (lower = lower latency)
  channels: 2               # 1 = mono, 2 = stereo

pipeline:
  enabled: true
  stages:
    - name: input
      type: capture
      device: "default"

    - name: equalizer
      type: eq
      enabled: true
      bands:
        - freq: 60
          gain: 0.0
          q: 1.0
        - freq: 230
          gain: 0.0
          q: 1.0
        - freq: 910
          gain: 0.0
          q: 1.0
        - freq: 3600
          gain: 0.0
          q: 1.0
        - freq: 14000
          gain: 0.0
          q: 1.0

    - name: compressor
      type: compressor
      enabled: false
      threshold: -18.0       # dB
      ratio: 2.5
      attack: 10             # ms
      release: 100           # ms
      makeup_gain: 0.0       # dB

    - name: limiter
      type: limiter
      enabled: true
      threshold: -1.0        # dB
      release: 50            # ms

    - name: output
      type: playback
      device: "default"

latency:
  target_ms: 10              # Desired latency
  max_ms: 20

logging:
  level: info                # debug, info, warn, error
```

## Features (planned)

- Gain
- Equalizer (biquad filters)
- Compressor / limiter
- PipeWire integration

## Architecture

Hexagonal architecture:

```
CLI → App → Domain (DSP) → Ports → Adapters
```

- DSP core is pure and testable
- PipeWire is isolated in adapters
- Config drives DSP chain

See: `docs/architecture.md`

## Project Structure

```
cmd/dsp           # CLI entrypoint  
internal/  
  domain/         # DSP core (pure logic)  
  app/            # orchestration (engine)  
  ports/          # interfaces  
  adapters/       # pipewire, file, fake  
docs/             # documentation
```

## Status

Early stage:

- [ ]  DSP core
- [ ]  Config loader
- [ ]  PipeWire adapter

---

## Philosophy

- Real-time safe code
- Zero allocations in audio path
- Simple, composable DSP units
