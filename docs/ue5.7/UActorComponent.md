# UActorComponent

**Parent**: UObject
**Module**: Engine

UActorComponent is the base class for all components that can be attached to actors. Components define reusable behavior and data. Unlike USceneComponent, a plain UActorComponent has no transform or spatial presence — use USceneComponent or its subclasses for positioned components.

## Key Properties

- `bAutoActivate` — Whether the component auto-activates when created
- `bIsActive` — Current activation state
- `PrimaryComponentTick` — Tick configuration (bCanEverTick, TickInterval, TickGroup)
- `ComponentTags` — Array of FName tags for identification

## Key Functions

- `BeginPlay()` — Called when the owning actor begins play. Override for initialization.
- `TickComponent(float DeltaTime, ELevelTick TickType, FActorComponentTickFunction* ThisTickFunction)` — Called every frame if ticking enabled.
- `EndPlay(EEndPlayReason::Type Reason)` — Called when component or owning actor is destroyed.
- `Activate(bool bReset)` — Activates the component.
- `Deactivate()` — Deactivates the component.
- `DestroyComponent()` — Destroys this component.
- `GetOwner()` — Returns the AActor that owns this component.
- `RegisterComponent()` — Registers component with the world (required after manual creation).
- `UnregisterComponent()` — Unregisters component from the world.
- `SetComponentTickEnabled(bool bEnabled)` — Enables/disables ticking at runtime.
