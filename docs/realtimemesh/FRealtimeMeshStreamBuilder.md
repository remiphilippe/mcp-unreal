# FRealtimeMeshStreamBuilder

**Module**: RealtimeMeshComponent

FRealtimeMeshStreamBuilder is the primary interface for constructing vertex and index data for RealtimeMesh sections. It provides a fluent API for adding vertices, normals, tangents, UVs, colors, and triangle indices. Data is organized into streams that can be independently enabled.

## Enabling Streams

Before adding data, enable the streams you need:

```cpp
FRealtimeMeshStreamBuilder Builder;
Builder.EnableVertices();    // Position data (required)
Builder.EnableTriangles();   // Index data (required)
Builder.EnableNormals();     // Normal vectors
Builder.EnableTangents();    // Tangent vectors
Builder.EnableUVs(2);        // UV channels (parameter = number of channels)
Builder.EnableColors();      // Vertex colors
```

## Key Functions

- `AddVertex(FVector3f Position)` — Adds a vertex and returns its index.
- `SetNormal(int32 Index, FVector3f Normal)` — Sets the normal for a vertex.
- `SetTangent(int32 Index, FVector3f Tangent)` — Sets the tangent for a vertex.
- `SetUV(int32 Index, FVector2f UV, int32 Channel)` — Sets UV coordinates for a vertex on a channel.
- `SetColor(int32 Index, FColor Color)` — Sets vertex color.
- `AddTriangle(int32 V0, int32 V1, int32 V2)` — Adds a triangle from three vertex indices (CCW winding).
- `GetVertexCount()` — Returns number of vertices added.
- `GetTriangleCount()` — Returns number of triangles added.

## Complete Example

```cpp
FRealtimeMeshStreamBuilder Builder;
Builder.EnableVertices();
Builder.EnableTriangles();
Builder.EnableNormals();
Builder.EnableUVs();

// Create a quad
int32 V0 = Builder.AddVertex(FVector3f(0, 0, 0));
int32 V1 = Builder.AddVertex(FVector3f(100, 0, 0));
int32 V2 = Builder.AddVertex(FVector3f(100, 100, 0));
int32 V3 = Builder.AddVertex(FVector3f(0, 100, 0));

// Set normals (all facing up)
for (int32 i = V0; i <= V3; i++)
    Builder.SetNormal(i, FVector3f(0, 0, 1));

// Set UVs
Builder.SetUV(V0, FVector2f(0, 0));
Builder.SetUV(V1, FVector2f(1, 0));
Builder.SetUV(V2, FVector2f(1, 1));
Builder.SetUV(V3, FVector2f(0, 1));

// Two triangles for the quad (CCW winding)
Builder.AddTriangle(V0, V1, V2);
Builder.AddTriangle(V0, V2, V3);

// Apply to mesh section
Mesh->UpdateSectionMesh(GroupKey, SectionKey, Builder);
```
