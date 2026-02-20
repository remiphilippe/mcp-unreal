# URealtimeMeshComponent

**Parent**: UMeshComponent
**Module**: RealtimeMeshComponent

URealtimeMeshComponent is the scene component that renders a URealtimeMesh. It functions similarly to UStaticMeshComponent but for runtime-generated meshes. Attach it to an actor and assign a URealtimeMeshSimple to display procedural geometry.

## Key Properties

- `RealtimeMesh` — The URealtimeMesh object containing mesh data
- `bCastDynamicShadow` — Whether the mesh casts dynamic shadows
- `CollisionProfileName` — Collision preset name

## Key Functions

- `InitializeRealtimeMesh(TSubclassOf<URealtimeMesh> MeshClass)` — Creates and initializes a new RealtimeMesh. Call this before setting up mesh data.
- `GetRealtimeMesh()` — Returns the current URealtimeMesh object.
- `GetRealtimeMeshAs<T>()` — Returns the RealtimeMesh cast to a specific subclass (e.g., URealtimeMeshSimple).
- `SetMaterial(int32 ElementIndex, UMaterialInterface* Material)` — Overrides a material slot.
- `SetCollisionEnabled(ECollisionEnabled::Type NewType)` — Enables/disables collision.

## Usage Pattern

```cpp
// In actor constructor or BeginPlay:
URealtimeMeshComponent* MeshComp = CreateDefaultSubobject<URealtimeMeshComponent>(TEXT("RealtimeMesh"));
RootComponent = MeshComp;

// Initialize with URealtimeMeshSimple
MeshComp->InitializeRealtimeMesh<URealtimeMeshSimple>();
URealtimeMeshSimple* Mesh = MeshComp->GetRealtimeMeshAs<URealtimeMeshSimple>();

// Set up materials
Mesh->SetupMaterialSlot(0, TEXT("Default"), DefaultMaterial);

// Create LOD and section group
Mesh->CreateSectionGroup(FRealtimeMeshLODKey(0), FRealtimeMeshSectionGroupKey(0));

// Build mesh data with FRealtimeMeshStreamBuilder and update section
```
