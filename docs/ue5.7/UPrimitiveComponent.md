# UPrimitiveComponent

**Parent**: USceneComponent
**Module**: Engine

UPrimitiveComponent is the base class for all components that have a visible representation or participate in the physics and collision system. This includes static meshes, skeletal meshes, shapes, brushes, and procedural geometry. It manages the render proxy, physics body instance, overlap/hit events, and material slots.

## Key Properties

- `bGenerateOverlapEvents` — If true, this component fires overlap begin/end events when it intersects other primitives.
- `bCastDynamicShadow` — Whether this component casts dynamic shadows.
- `CollisionProfileName` — The FName collision preset (e.g., "BlockAll", "OverlapAllDynamic", "NoCollision") applied to the body instance.
- `BodyInstance` — The FBodyInstance holding all physics simulation state, mass, constraints, and per-body collision settings.
- `bSimulatePhysics` — Shorthand accessor for BodyInstance.bSimulatePhysics; enables rigid body simulation.
- `bVisible` — Whether this component is visible in game (does not affect collision).
- `CastShadow` — Master switch controlling whether this component contributes any shadows.

## Key Functions

- `SetCollisionEnabled(ECollisionEnabled::Type NewType)` — Sets collision mode: NoCollision, QueryOnly, PhysicsOnly, or QueryAndPhysics.
- `SetCollisionProfileName(FName ProfileName)` — Applies a collision preset by name, updating object type and response channels.
- `SetSimulatePhysics(bool bSimulate)` — Enables or disables rigid body physics simulation on this component.
- `AddForce(FVector Force, FName BoneName, bool bAccelChange)` — Applies a continuous force to the physics body (world space).
- `AddImpulse(FVector Impulse, FName BoneName, bool bVelChange)` — Applies an instantaneous impulse to the physics body.
- `SetVisibility(bool bNewVisibility, bool bPropagateToChildren)` — Shows or hides this component and optionally its children.
- `GetOverlappingActors(TArray<AActor*>& OverlappingActors, TSubclassOf<AActor> ClassFilter)` — Fills an array with all actors currently overlapping this component.
- `GetOverlappingComponents(TArray<UPrimitiveComponent*>& OutOverlappingComponents)` — Fills an array with all components currently overlapping this component.
- `OnComponentBeginOverlap` — Multicast delegate fired when another primitive begins overlapping this component.
- `OnComponentHit` — Multicast delegate fired on blocking collision hit; provides hit result and impulse data.
- `SetMaterial(int32 ElementIndex, UMaterialInterface* Material)` — Assigns a material to a specific mesh slot.
- `GetMaterial(int32 ElementIndex)` — Returns the UMaterialInterface at the given slot index.
- `GetNumMaterials()` — Returns the number of material slots on this component.
- `SetRenderCustomDepth(bool bValue)` — Enables custom depth rendering for post-process outline effects.
