# UMaterialInstanceDynamic

**Parent**: UMaterialInstanceConstant
**Module**: Engine

Dynamic material instance that allows runtime modification of material parameters without creating new assets. Created from a base `UMaterialInterface` at runtime, enabling per-object or per-frame changes to scalar, vector, and texture parameters. Essential for gameplay-driven visual effects such as health-based color changes, dissolve effects, and animated material properties.

## Key Properties

- `Parent` — `UMaterialInterface*` pointing to the base material or material instance from which this dynamic instance inherits its parameter defaults and shader

## Key Functions

- `Create(UMaterialInterface* Parent, UObject* Outer)` — Static factory function returning a new `UMaterialInstanceDynamic*`; `Outer` is typically the owning actor or component
- `SetScalarParameterValue(FName ParameterName, float Value)` — Sets a named scalar (float) parameter at runtime; triggers no shader recompilation
- `SetVectorParameterValue(FName ParameterName, FLinearColor Value)` — Sets a named vector parameter (RGBA linear color) at runtime
- `SetTextureParameterValue(FName ParameterName, UTexture* Value)` — Assigns a texture asset to a named texture parameter slot at runtime
- `GetScalarParameterValue(FName ParameterName)` — Returns the current `float` value of a named scalar parameter; checks instance overrides first, then falls back to parent hierarchy
- `GetVectorParameterValue(FName ParameterName, FLinearColor& OutValue)` — Retrieves the current `FLinearColor` value of a named vector parameter into `OutValue`; returns `bool` indicating whether the parameter was found
- `GetTextureParameterValue(FName ParameterName, UTexture*& OutValue)` — Retrieves the current texture assigned to a named parameter into `OutValue`; returns `bool` indicating success
- `CopyMaterialUniformParameters(UMaterialInterface* Source)` — Copies all scalar, vector, and texture parameter values from `Source` into this instance; useful for cloning appearance at runtime
- `SetScalarParameterByIndex(int32 ParameterIndex, float Value)` — Sets a scalar parameter by its cached index rather than name for lower-overhead per-frame updates
- `K2_CopyMaterialInstanceParameters(UMaterialInterface* Source, bool bQuickParametersOnly)` — Blueprint-callable version of parameter copying; `bQuickParametersOnly` limits the copy to uniform parameters for better performance
