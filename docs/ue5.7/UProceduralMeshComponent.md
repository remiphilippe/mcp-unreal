# UProceduralMeshComponent

**Parent**: UMeshComponent
**Module**: ProceduralMeshComponent

Runtime mesh generation component for creating arbitrary procedural geometry from vertex and index data at runtime. Supports multiple independently visible mesh sections, collision generation (both simple convex and complex triangle mesh), and standard UE material assignment. Ideal for dynamic terrain, generated geometry, and data-driven meshes. For high-performance or streaming meshes consider the RealtimeMesh plugin instead.

## Key Properties

- `bUseAsyncCooking` — When true, collision cooking (PhysX/Chaos convex and trimesh generation) runs asynchronously on a background thread to avoid hitches; collision becomes available once cooking completes
- `bUseComplexAsSimpleCollision` — When true, the complex (per-triangle) collision mesh is also used for simple collision queries (sweeps, overlaps); more accurate but slower than convex hulls
- `ProcMeshBodySetup` — The `UBodySetup*` generated from the current mesh data; holds the cooked physics collision shapes
- `LocalBounds` — Cached `FBoxSphereBounds` in component local space; recomputed whenever mesh sections change

## Key Functions

- `CreateMeshSection(int32 SectionIndex, TArray<FVector> Vertices, TArray<int32> Triangles, TArray<FVector> Normals, TArray<FVector2D> UV0, TArray<FColor> VertexColors, TArray<FProcMeshTangent> Tangents, bool bCreateCollision)` — Creates or replaces a mesh section at `SectionIndex` with the provided geometry data; set `bCreateCollision` to generate a physics collision shape for this section
- `CreateMeshSection_LinearColor(int32 SectionIndex, TArray<FVector> Vertices, TArray<int32> Triangles, TArray<FVector> Normals, TArray<FVector2D> UV0, TArray<FLinearColor> VertexColors, TArray<FProcMeshTangent> Tangents, bool bCreateCollision)` — Variant of `CreateMeshSection` accepting `FLinearColor` for HDR vertex colour support
- `UpdateMeshSection(int32 SectionIndex, TArray<FVector> Vertices, TArray<FVector> Normals, TArray<FVector2D> UV0, TArray<FColor> VertexColors, TArray<FProcMeshTangent> Tangents)` — Updates vertex data for an existing section without recreating the index buffer; faster than `CreateMeshSection` when topology is unchanged
- `ClearMeshSection(int32 SectionIndex)` — Removes all geometry data from the section at `SectionIndex` and marks it empty; does not remove the section slot
- `ClearAllMeshSections()` — Removes all sections and frees all geometry data; resets the component to an empty state
- `SetMeshSectionVisible(int32 SectionIndex, bool bVisible)` — Shows or hides a mesh section without destroying its data; useful for LOD-style switching or debugging
- `GetNumSections()` — Returns the number of allocated section slots as an `int32`; some slots may be empty after `ClearMeshSection`
- `AddCollisionConvexMesh(TArray<FVector> ConvexVerts)` — Adds a convex hull collision shape from the given vertex list, independent of rendered sections; supports multiple convex hulls for complex simple collision
- `ClearCollisionConvexMeshes()` — Removes all manually added convex collision hulls
- `SetMaterial(int32 ElementIndex, UMaterialInterface* Material)` — Assigns a material to the given section index; each mesh section maps to one material element
- `ContainsPhysicsTriMeshData(bool InUseAllTriData)` — Returns `true` if the component has triangle mesh collision data suitable for complex collision queries; used internally by the physics system
