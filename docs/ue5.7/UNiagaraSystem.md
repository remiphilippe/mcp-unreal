# UNiagaraSystem

**Parent**: UObject
**Module**: Niagara

Niagara particle system asset that acts as the top-level container for one or more emitters. A `UNiagaraSystem` defines system-level simulation scripts (spawn and update), exposes user parameters, and manages the collection of `FNiagaraEmitterHandle` entries. Spawn it in the world via `UNiagaraFunctionLibrary::SpawnSystemAtLocation` or attach it through a `UNiagaraComponent`.

## Key Properties

- `EmitterHandles` — Array of `FNiagaraEmitterHandle` structs; each handle references a `UNiagaraEmitter` asset and carries per-system overrides and enabled state
- `bFixedBounds` — When true, uses `FixedBounds` instead of computing dynamic bounds each frame; improves performance for effects with predictable extents
- `FixedBounds` — `FBox` used as the system's world-space bounds when `bFixedBounds` is enabled
- `SystemSpawnScript` — `UNiagaraScript` that runs once when the system spawns; initialises system-level attributes and variables
- `SystemUpdateScript` — `UNiagaraScript` that runs every tick to update system-level attributes and drive emitter parameters
- `WarmupTime` — Duration in seconds to simulate before the system becomes visible; useful for effects that should appear already in progress
- `WarmupTickCount` — Number of ticks used to advance simulation during warmup; higher values increase startup cost but improve accuracy
- `WarmupTickDelta` — Time step in seconds for each warmup tick
- `bAutoDeactivate` — When true, the system automatically deactivates and (optionally) destroys its component after all emitters complete
- `MaxPoolSize` — Maximum number of `UNiagaraComponent` instances kept in the object pool for this system; 0 disables pooling
- `PoolPrimeSize` — Number of pooled components pre-allocated at load time to avoid runtime allocation spikes

## Key Functions

- `GetEmitterHandles()` — Returns a const reference to the `TArray<FNiagaraEmitterHandle>` of all emitter handles in this system
- `GetNumEmitters()` — Returns the number of emitter handles as an `int32`
- `GetEmitterHandle(int32 Index)` — Returns a reference to the `FNiagaraEmitterHandle` at the given index
- `AddEmitterHandle(UNiagaraEmitter& Emitter)` — Adds a new emitter handle wrapping the given emitter asset and returns the new handle
- `RemoveEmitterHandle(int32 Index)` — Removes the emitter handle at the given index from the system
- `MoveEmitterHandleToIndex(FNiagaraEmitterHandle& EmitterHandle, int32 NewIndex)` — Reorders an emitter handle within the array
- `Compile(bool bForce)` — Triggers compilation of all system and emitter scripts; pass `true` to recompile even if scripts appear up-to-date
- `GetExposedParameters()` — Returns the `FNiagaraUserRedirectionParameterStore` containing all user-exposed parameters that can be set from Blueprint or C++
- `GetSystemCompiledData()` — Returns the `FNiagaraSystemCompiledData` struct with compiled script data and layout information
- `HasAnyGPUEmitters()` — Returns `true` if any emitter in this system uses GPU compute simulation
- `NeedsWarmup()` — Returns `true` if `WarmupTime` is greater than zero and the system requires pre-simulation on spawn
