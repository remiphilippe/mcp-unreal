# UStaticMeshComponent

**Parent**: UMeshComponent
**Module**: Engine

UStaticMeshComponent renders a static mesh asset in the world. It is the most common way to display 3D geometry. Supports per-instance materials, LODs, collision from the mesh asset, and instanced static mesh rendering for performance.

## Key Properties

- `StaticMesh` — The UStaticMesh asset to render
- `OverrideMaterials` — Array of material overrides per material slot
- `bCastDynamicShadow` — Whether this mesh casts dynamic shadows
- `bCastStaticShadow` — Whether this mesh casts static shadows (for lightmaps)
- `CollisionProfileName` — Collision preset name (BlockAll, OverlapAll, NoCollision, etc.)

## Key Functions

- `SetStaticMesh(UStaticMesh* NewMesh)` — Changes the displayed mesh at runtime.
- `SetMaterial(int32 ElementIndex, UMaterialInterface* Material)` — Sets a material on a specific slot.
- `GetMaterial(int32 ElementIndex)` — Returns the material on a slot.
- `GetNumMaterials()` — Returns the number of material slots.
- `SetCollisionEnabled(ECollisionEnabled::Type NewType)` — Enables/disables collision.
- `SetCollisionProfileName(FName InCollisionProfileName)` — Sets the collision preset.
- `SetSimulatePhysics(bool bSimulate)` — Enables/disables physics simulation.
- `SetRenderCustomDepth(bool bValue)` — Enables custom depth for post-process effects.
