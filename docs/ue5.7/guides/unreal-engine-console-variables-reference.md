<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-console-variables-reference -->

This is a reference page for console variables (CVars) available in Unreal Engine. Console variables allow you to change engine behavior at runtime through the console, configuration files, or command-line arguments.

## How to Use Console Variables

Console variables can be set in multiple ways:

- **In-game console**: Press the tilde key (`~`) to open the console, then type the variable name and value (e.g., `r.ScreenPercentage 80`)
- **Configuration files**: Add entries to `.ini` files under the appropriate section (e.g., `[/Script/Engine.RendererSettings]`)
- **Command-line**: Pass as arguments when launching the engine (e.g., `-ExecCmds="r.ScreenPercentage 80"`)
- **C++ code**: Use `IConsoleManager::Get().FindConsoleVariable()` or the `TAutoConsoleVariable` template

## Common Console Variable Categories

### Rendering (r.)

| Variable | Description |
| --- | --- |
| `r.ScreenPercentage` | Controls the resolution percentage for rendering (default 100) |
| `r.VSync` | Enable/disable vertical sync (0=off, 1=on) |
| `r.FullScreen` | Fullscreen mode (0=windowed, 1=fullscreen) |
| `r.Shadow.MaxResolution` | Maximum shadow map resolution |
| `r.Shadow.MaxCSMResolution` | Maximum cascaded shadow map resolution |
| `r.MotionBlurQuality` | Motion blur quality (0=off, 1=low, 2=medium, 3=high, 4=very high) |
| `r.BloomQuality` | Bloom quality (0=off through 5=max) |
| `r.AmbientOcclusionLevels` | Number of ambient occlusion levels |
| `r.DepthOfFieldQuality` | Depth of field quality |
| `r.PostProcessAAQuality` | Post process anti-aliasing quality |
| `r.Tonemapper.Quality` | Tonemapper quality |
| `r.SSR.Quality` | Screen space reflections quality |
| `r.Lumen.DiffuseIndirect.Allow` | Enable/disable Lumen diffuse indirect lighting |
| `r.Lumen.Reflections.Allow` | Enable/disable Lumen reflections |
| `r.Nanite.MaxPixelsPerEdge` | Nanite target edge length in pixels |
| `r.RayTracing` | Enable/disable hardware ray tracing |
| `r.RDG.Debug` | RDG debug mode |
| `r.RDG.ImmediateMode` | Execute RDG passes immediately for debugging |

### Scalability (sg.)

| Variable | Description |
| --- | --- |
| `sg.ResolutionQuality` | Resolution quality scalability group |
| `sg.ViewDistanceQuality` | View distance quality scalability group |
| `sg.AntiAliasingQuality` | Anti-aliasing quality scalability group |
| `sg.ShadowQuality` | Shadow quality scalability group |
| `sg.GlobalIlluminationQuality` | Global illumination quality scalability group |
| `sg.ReflectionQuality` | Reflection quality scalability group |
| `sg.PostProcessQuality` | Post-process quality scalability group |
| `sg.TextureQuality` | Texture quality scalability group |
| `sg.EffectsQuality` | Effects quality scalability group |
| `sg.FoliageQuality` | Foliage quality scalability group |
| `sg.ShadingQuality` | Shading quality scalability group |

### Physics (p.)

| Variable | Description |
| --- | --- |
| `p.MaxPhysicsDeltaTime` | Maximum physics delta time |
| `p.PhysXDefaultSimFilterShader` | PhysX default simulation filter shader |
| `p.EnableStabilization` | Enable physics stabilization |
| `p.Chaos.Solver.Iterations` | Number of Chaos solver iterations |

### Audio (au.)

| Variable | Description |
| --- | --- |
| `au.MaxChannels` | Maximum number of audio channels |
| `au.DisableAudio` | Disable audio entirely |

### Network (net.)

| Variable | Description |
| --- | --- |
| `net.MaxRepArraySize` | Maximum replicated array size |
| `net.MaxRepArrayMemory` | Maximum memory for replicated arrays |
| `net.DormancyEnable` | Enable network dormancy |

### Streaming (s.)

| Variable | Description |
| --- | --- |
| `s.AsyncLoadingThreadEnabled` | Enable async loading thread |
| `s.ForceFlushStreamingOnLevel` | Force flush streaming on level change |

### Garbage Collection (gc.)

| Variable | Description |
| --- | --- |
| `gc.TimeBetweenPurgingPendingKillObjects` | Time between purging pending kill objects |
| `gc.MaxObjectsNotConsideredByGC` | Maximum objects not considered by GC |

For the complete list of all console variables, see the [full reference page](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-console-variables-reference).

For information on how to create and use console variables in C++, see the [Console Variables and Commands](https://dev.epicgames.com/documentation/en-us/unreal-engine/console-variables-cplusplus-in-unreal-engine) documentation.
