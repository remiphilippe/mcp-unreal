# AWorldSettings

**Parent**: AInfo
**Module**: Engine

Per-level configuration actor that stores global settings for a `ULevel`. Every level has exactly one `AWorldSettings` actor placed at `Actors[0]`. It controls the default game mode, physics constants (gravity, kill zone), navigation system configuration, global audio settings, world composition, lighting build options, and frame rate bounds. Access it via `ULevel::GetWorldSettings()` or `UWorld::GetWorldSettings()`.

## Key Properties

- `DefaultGameMode` — `TSubclassOf<AGameModeBase>` specifying the game mode used when no override is set and this is the persistent level; only meaningful for persistent levels
- `GameModeOverride` — `TSubclassOf<AGameModeBase>` that overrides the project default game mode for this specific level when set; takes precedence over `DefaultGameMode`
- `bEnableWorldBoundsChecks` — When true, actors that fall below `KillZ` are destroyed by `KilledBy()`
- `KillZ` — World-space Z coordinate below which actors are killed when `bEnableWorldBoundsChecks` is true; default `-HALF_WORLD_MAX`
- `WorldToMeters` — Scale factor defining how many Unreal Units equal one metre; default `100.0`; affects physics, audio attenuation, and VR interpupillary distance
- `bGlobalGravitySet` — When true, `GlobalGravityZ` overrides the physics engine's default gravity for this level
- `GlobalGravityZ` — Override gravity in cm/s² (negative = downward); only applied when `bGlobalGravitySet` is true
- `bEnableNavigationSystem` — When false, disables the navigation system for this level; useful for levels with no AI
- `NavigationSystemConfig` — `UNavigationSystemConfig*` asset specifying which navigation system class and settings to use
- `DefaultAmbientZoneSettings` — `FAudioVolumeSettings` applied globally when no `AAudioVolume` overrides are active
- `bEnableWorldComposition` — Enables the legacy World Composition system for tiled worlds; prefer World Partition in UE 5.x
- `bEnableHierarchicalLOD` — Enables Hierarchical Level of Detail (HLOD) generation and runtime HLOD cluster swapping
- `DefaultReverbSettings` — `FReverbSettings` applied globally when no reverb volume overrides are active
- `DefaultBaseSoundMix` — `USoundMix*` set as the base (lowest priority) sound mix for this level
- `bForceNoPrecomputedLighting` — When true, disables all precomputed (baked) lighting; forces fully dynamic lighting; useful during development
- `LightmassSettings` — `FLightmassWorldInfoSettings` struct controlling Lightmass global illumination quality, bounce count, and sky light settings
- `MinUndilatedFrameTime` — Minimum frame delta time in seconds before time dilation is considered (prevents extreme time dilation from very slow frames)
- `MaxUndilatedFrameTime` — Maximum frame delta time in seconds capped before simulation to prevent large physics integration steps

## Key Functions

- `GetDefaultGameMode()` — Returns the `TSubclassOf<AGameModeBase>` for the default game mode; may be null if not set
- `GetGameModeOverride()` — Returns the `TSubclassOf<AGameModeBase>` override; null if no override is set
- `GetGravityZ()` — Returns the effective gravity Z value in cm/s²; returns `GlobalGravityZ` if `bGlobalGravitySet`, otherwise returns the physics engine default
- `GetWorldToMeters()` — Returns the `WorldToMeters` scale as a `float`
- `NotifyBeginPlay()` — Called by the engine when the level begins play; broadcasts to all registered components and the Level Script Actor
- `NotifyMatchStarted()` — Called by `AGameModeBase` when the match officially starts; used to trigger level-specific start logic
- `GetAISystemClassName()` — Returns the `FName` of the AI system class to instantiate for this world; defaults to the project-level AI system class
- `OnRep_DefaultGameMode()` — `RepNotify` called on clients when `DefaultGameMode` replicates; triggers game mode class refresh on the client world
- `GetWorldSettings()` — Static helper that returns the `AWorldSettings*` for a given `UObject`'s world; shorthand for `Object->GetWorld()->GetWorldSettings()`
