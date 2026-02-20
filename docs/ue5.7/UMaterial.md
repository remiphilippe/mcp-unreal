# UMaterial

**Parent**: UMaterialInterface
**Module**: Engine

Material asset defining the visual appearance of surfaces in Unreal Engine. Encapsulates the full set of shading instructions compiled into GPU shader programs. Controls how light interacts with a surface through shading models, blend modes, and material domains.

## Key Properties

- `ShadingModel` — Enum (`EMaterialShadingModel`) selecting the lighting equation: `MSM_DefaultLit`, `MSM_Unlit`, `MSM_Subsurface`, `MSM_PreintegratedSkin`, `MSM_ClearCoat`, `MSM_SubsurfaceProfile`, `MSM_TwoSidedFoliage`, `MSM_Hair`, `MSM_Cloth`, `MSM_Eye`, `MSM_SingleLayerWater`, `MSM_ThinTranslucent`
- `BlendMode` — Enum (`EBlendMode`) controlling how the material composites over other surfaces: `BLEND_Opaque`, `BLEND_Masked`, `BLEND_Translucent`, `BLEND_Additive`, `BLEND_Modulate`, `BLEND_AlphaComposite`, `BLEND_AlphaHoldout`
- `MaterialDomain` — Enum (`EMaterialDomain`) specifying usage context: `MD_Surface`, `MD_DeferredDecal`, `MD_LightFunction`, `MD_PostProcess`, `MD_UI`, `MD_Volume`, `MD_RuntimeVirtualTexture`
- `OpacityMaskClipValue` — `float` threshold for Masked blend mode; pixels below this opacity value are fully clipped
- `bTwoSided` — `bool` disabling backface culling so the material renders on both faces of a polygon
- `bUsedWithSkeletalMesh` — `bool` usage flag required for the material to compile shaders compatible with skeletal mesh rendering
- `bDisableDepthTest` — `bool` allowing the material to render regardless of depth buffer occlusion; typically used for UI or special effects
- `Expressions` — `TArray<UMaterialExpression*>` containing all material expression nodes in the material graph

## Key Functions

- `GetShadingModels()` — Returns `FMaterialShadingModelField` representing the set of active shading models used by this material
- `SetShadingModel(EMaterialShadingModel NewModel)` — Sets the primary shading model; triggers shader recompilation in the editor
- `GetBlendMode()` — Returns `EBlendMode` indicating how this material blends over the scene
- `GetMaterialDomain()` — Returns `EMaterialDomain` indicating the rendering context this material is designed for
- `IsUsedWithSkeletalMesh()` — Returns `bool` indicating whether `bUsedWithSkeletalMesh` is set and shaders have been compiled for skeletal mesh usage
- `GetPhysicalMaterialMask()` — Returns `UPhysicalMaterialMask*` used to assign different physical materials to regions of the surface based on a mask texture
- `CompileProperty(FMaterialCompiler* Compiler, EMaterialProperty Property)` — Compiles a specific material output property into shader code; used internally by the shader compiler
- `GetDefaultMaterialInstance()` — Returns `FMaterialRenderProxy*` for the default (non-instanced) render-time representation of this material
