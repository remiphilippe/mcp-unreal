# URealtimeMeshSimple

**Parent**: URealtimeMesh
**Module**: RealtimeMeshComponent

URealtimeMeshSimple provides a simplified API for creating and updating runtime meshes with full LOD support, collision generation, and section groups. It is the primary class for procedural mesh generation using the RealtimeMesh plugin. Unlike ProceduralMeshComponent, it supports multiple LODs, section groups for material batching, and async mesh updates.

## Key Properties

- `MeshData` — Internal mesh data storage
- `LODCount` — Number of LOD levels configured

## Key Functions

- `SetupMaterialSlot(int32 MaterialSlot, FName MaterialSlotName, UMaterialInterface* Material)` — Configures a material slot before creating sections.
- `CreateSectionGroup(FRealtimeMeshLODKey LODKey, FRealtimeMeshSectionGroupKey GroupKey)` — Creates a section group within a LOD for organizing mesh sections.
- `RemoveSectionGroup(FRealtimeMeshLODKey LODKey, FRealtimeMeshSectionGroupKey GroupKey)` — Removes a section group.
- `CreateSection(FRealtimeMeshLODKey LODKey, FRealtimeMeshSectionGroupKey GroupKey, FRealtimeMeshSectionKey SectionKey, FRealtimeMeshSectionConfig Config)` — Creates a mesh section within a group.
- `UpdateSectionMesh(FRealtimeMeshSectionGroupKey GroupKey, FRealtimeMeshSectionKey SectionKey, FRealtimeMeshStreamBuilder& Builder)` — Updates vertex/index data for a section using a stream builder.
- `SetLODScreenSize(FRealtimeMeshLODKey LODKey, float ScreenSize)` — Sets the screen-size threshold for LOD transitions.
- `GetLocalBounds()` — Returns the local-space bounding box of the mesh.

## Stream Builder Pattern

```cpp
FRealtimeMeshStreamBuilder Builder;
Builder.EnableVertices();
Builder.EnableTriangles();
Builder.EnableNormals();
Builder.EnableTangents();
Builder.EnableUVs();
Builder.EnableColors();

// Add vertices
int32 V0 = Builder.AddVertex(FVector3f(0, 0, 0));
int32 V1 = Builder.AddVertex(FVector3f(100, 0, 0));
int32 V2 = Builder.AddVertex(FVector3f(0, 100, 0));

// Set normals
Builder.SetNormal(V0, FVector3f(0, 0, 1));

// Add triangle
Builder.AddTriangle(V0, V1, V2);
```

## LOD Configuration

RealtimeMesh supports up to 8 LOD levels. Each LOD has independent section groups and sections:

- LOD 0: Full detail (screen size > 0.5)
- LOD 1: Reduced detail (screen size > 0.25)
- LOD 2: Low detail (screen size > 0.1)

Set screen sizes with `SetLODScreenSize()` to control transitions.
