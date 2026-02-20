<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USkyLightComponent -->

|  |  |
| --- | --- |
| _Name_ | USkyLightComponent |
| _Type_ | class |
| _Header File_ | /Engine/Source/Runtime/Engine/Classes/Components/SkyLightComponent.h |
| _Include Path_ | #include "Components/SkyLightComponent.h" |

## Syntax

```cpp

UCLASS (Blueprintable, ClassGroup=Lights,

       HideCategories=(Trigger, Activation, "Components|Activation", Physics),

       Meta=(BlueprintSpawnableComponent), MinimalAPI)

class USkyLightComponent : public ULightComponentBase
```

## Inheritance Hierarchy

- [UObjectBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObjectBase) → [UObjectBaseUtility](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObjectBaseUtility) → [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) → [UActorComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UActorComponent) → [USceneComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USceneComponent) → [ULightComponentBase](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/ULightComponentBase) → **USkyLightComponent**

## Implements Interfaces

- [IAsyncPhysicsStateProcessor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/IAsyncPhysicsStateProcessor)
- [IInterface\_AssetUserData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/IInterface_AssetUserData)

## Constructors

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| USkyLightComponent<br>(<br>const [FObjectInitializer](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FObjectInitializer)& ObjectInitializer<br>) |  | Components/SkyLightComponent.h |  |

## Constants

| Name | Type | Remarks | Include Path |
| --- | --- | --- | --- |
| SkyCapturesToUpdate | [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [USkyLightComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USkyLightComponent) \\* > | List of sky captures that need to be recaptured. | Components/SkyLightComponent.h |
| SkyCapturesToUpdateBlendDestinations | [TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [USkyLightComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USkyLightComponent) \\* > |  | Components/SkyLightComponent.h |
| SkyCapturesToUpdateLock | FTransactionallySafeCriticalSection |  | Components/SkyLightComponent.h |

## Variables

### Public

| Name | Type | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/unreal-engine-uproperties#propertyspecifiers) |
| --- | --- | --- | --- | --- |
| bCaptureEmissiveOnly | bool | Only capture emissive materials. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=Light<br>- AdvancedDisplay |
| bCloudAmbientOcclusion | uint32 | Whether the cloud should occlude sky contribution within the atmosphere (progressively fading multiple scattering out) or not. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=AtmosphereAndCloud |
| bLowerHemisphereIsBlack | bool | Whether all distant lighting from the lower hemisphere should be set to LowerHemisphereColor. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=Light<br>- AdvancedDisplay<br>- Meta=(DisplayName="Lower Hemisphere Is Solid Color") |
| bRealTimeCapture | bool | When enabled, the sky will be captured and convolved to achieve dynamic diffuse and specular environment lighting. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=Light |
| CloudAmbientOcclusionApertureScale | float | Controls the cone aperture angle over which the sky occlusion due to volumetric clouds is evaluated. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=AtmosphereAndCloud<br>- Meta=(UIMin="0.0", UIMax="0.1", ClampMin="0.0", ClampMax="1.0", SliderExponent=2.0) |
| CloudAmbientOcclusionExtent | float | The world space radius of the cloud ambient occlusion map around the camera in kilometers. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=AtmosphereAndCloud<br>- Meta=(UIMin="1", ClampMin="1") |
| CloudAmbientOcclusionMapResolutionScale | float | Scale the cloud ambient occlusion map resolution, base resolution is 512. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=AtmosphereAndCloud<br>- Meta=(UIMin="0.25", UIMax="8", ClampMin="0.25", SliderExponent=1.0) |
| CloudAmbientOcclusionStrength | float | The strength of the ambient occlusion, higher value will block more light. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Interp<br>- Category=AtmosphereAndCloud<br>- Meta=(UIMin="0", UIMax="1", ClampMin="0", SliderExponent=1.0) |
| Contrast | float | Contrast S-curve applied to the computed AO. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=DistanceFieldAmbientOcclusion<br>- Meta=(UIMin="0", UIMax="1", DisplayName="Occlusion Contrast") |
| Cubemap | [TObjectPtr](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TObjectPtr) < class [UTextureCube](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UTextureCube) > | Cubemap to use for sky lighting if SourceType is set to SLS\_SpecifiedCubemap. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=Light |
| CubemapResolution | int32 | Maximum resolution for the very top processed cubemap mip. Must be a power of 2. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=Light |
| LowerHemisphereColor | [FLinearColor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FLinearColor) |  | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=Light<br>- AdvancedDisplay |
| MinOcclusion | float | Controls the darkest that a fully occluded area can get. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=DistanceFieldAmbientOcclusion<br>- Meta=(UIMin="0", UIMax="1") |
| OcclusionCombineMode | [TEnumAsByte](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TEnumAsByte) < enum [EOcclusionCombineMode](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/EOcclusionCombineMode) > | Controls how occlusion from Distance Field Ambient Occlusion is combined with Screen Space Ambient Occlusion. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=DistanceFieldAmbientOcclusion |
| OcclusionExponent | float | Exponent applied to the computed AO. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=DistanceFieldAmbientOcclusion<br>- Meta=(UIMin=".6", UIMax="1.6") |
| OcclusionMaxDistance | float | Max distance that the occlusion of one point will affect another. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=DistanceFieldAmbientOcclusion<br>- Meta=(UIMin="200", UIMax="1500") |
| OcclusionTint | [FColor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FColor) | Tint color on occluded areas, artistic control. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=DistanceFieldAmbientOcclusion |
| SkyDistanceThreshold | float | Distance from the sky light at which any geometry should be treated as part of the sky. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=Light |
| SourceCubemapAngle | float | Angle to rotate the source cubemap when SourceType is set to SLS\_SpecifiedCubemap. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=Light<br>- Meta=(UIMin="0", UIMax="360") |
| SourceType | [TEnumAsByte](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TEnumAsByte) < enum [ESkyLightSourceType](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/ESkyLightSourceType) > | Indicates where to get the light contribution from. | Components/SkyLightComponent.h | - EditAnywhere<br>- BlueprintReadOnly<br>- Category=Light |

## Functions

### Public

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| void ApplyComponentInstanceData<br>(<br>[FPrecomputedSkyLightInstanceData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FPrecomputedSkyLightInstanceData)\\* ComponentInstanceData<br>) |  | Components/SkyLightComponent.h |  |
| void CaptureEmissiveRadianceEnvironmentCubeMap<br>(<br>FSHVectorRGB3& OutIrradianceMap,<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [FFloat16Color](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FFloat16Color) >& OutRadianceMap<br>) const | Computes a radiance map using only emissive contribution from the sky light. | Components/SkyLightComponent.h |  |
| [FSkyLightSceneProxy](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FSkyLightSceneProxy) \\* CreateSceneProxy() |  | Components/SkyLightComponent.h |  |
| FSHVectorRGB3 GetIrradianceEnvironmentMap() |  | Components/SkyLightComponent.h |  |
| const [FTexture](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/RenderCore/FTexture) \\* GetProcessedSkyTexture() |  | Components/SkyLightComponent.h |  |
| bool IsOcclusionSupported() | Whether sky occlusion is supported by current feature level | Components/SkyLightComponent.h |  |
| bool IsRealTimeCaptureEnabled() |  | Components/SkyLightComponent.h |  |
| void [RecaptureSky](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USkyLightComponent/RecaptureSky) () | Recaptures the scene for the skylight. | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="Rendering\|Components\|SkyLight" |
| void SanitizeCubemapSize() |  | Components/SkyLightComponent.h |  |
| void SetBlendDestinationCaptureIsDirty() |  | Components/SkyLightComponent.h |  |
| void SetCaptureIsDirty() | Indicates that the capture needs to recapture the scene, adds it to the recapture queue. | Components/SkyLightComponent.h |  |
| void SetCubemap<br>(<br>[UTextureCube](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UTextureCube)\\* NewCubemap<br>) | Sets the cubemap used when SourceType is set to SpecifiedCubemap, and causes a skylight update on the next tick. | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="SkyLight" |
| void [SetCubemapBlend](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USkyLightComponent/SetCubemapBlend)<br>(<br>[UTextureCube](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UTextureCube)\\* SourceCubemap,<br>[UTextureCube](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UTextureCube)\\* DestinationCubemap,<br>float InBlendFraction<br>) | Creates sky lighting from a blend between two cubemaps, which is only valid when SourceType is set to SpecifiedCubemap. | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="SkyLight" |
| void SetIndirectLightingIntensity<br>(<br>float NewIntensity<br>) |  | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="Rendering\|Components\|Light" |
| void SetIntensity<br>(<br>float NewIntensity<br>) | Set brightness of the light | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="Rendering\|Components\|SkyLight" |
| void SetLightColor<br>(<br>[FLinearColor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FLinearColor) NewLightColor<br>) | Set color of the light | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="Rendering\|Components\|SkyLight" |
| void SetLowerHemisphereColor<br>(<br>const [FLinearColor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FLinearColor)& InLowerHemisphereColor<br>) |  | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="Rendering\|Components\|SkyLight" |
| void SetMinOcclusion<br>(<br>float InMinOcclusion<br>) |  | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="Rendering\|Components\|SkyLight" |
| void SetOcclusionContrast<br>(<br>float InOcclusionContrast<br>) |  | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="Rendering\|Components\|SkyLight" |
| void SetOcclusionExponent<br>(<br>float InOcclusionExponent<br>) |  | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="Rendering\|Components\|SkyLight" |
| void SetOcclusionTint<br>(<br>const [FColor](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FColor)& InTint<br>) |  | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="Rendering\|Components\|SkyLight" |
| void SetRealTimeCapture<br>(<br>bool bInRealTimeCapture<br>) |  | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="Rendering\|Components\|SkyLight" |
| void SetRealTimeCaptureEnabled<br>(<br>bool bNewRealTimeCaptureEnabled<br>) |  | Components/SkyLightComponent.h |  |
| void [SetSourceCubemapAngle](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USkyLightComponent/SetSourceCubemapAngle)<br>(<br>float NewValue<br>) | Sets the angle of the cubemap used when SourceType is set to SpecifiedCubemap and it is non static. | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="Rendering\|Components\|SkyLight" |
| void SetVolumetricScatteringIntensity<br>(<br>float NewIntensity<br>) |  | Components/SkyLightComponent.h | - BlueprintCallable<br>- Category="Rendering\|Components\|Light" |

#### Overridden from [UActorComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UActorComponent)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void [CheckForErrors](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USkyLightComponent/CheckForErrors) () | Function that gets called from within Map\_Check to allow this actor component to check itself for any potential errors and register them with map check dialog. | Components/SkyLightComponent.h |  |
| virtual [TStructOnScope](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/TStructOnScope) < [FActorComponentInstanceData](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FActorComponentInstanceData) \> GetComponentInstanceData() | Called before we throw away components during RerunConstructionScripts, to cache any data we wish to persist across that operation | Components/SkyLightComponent.h |  |

#### Overridden from [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void BeginDestroy() |  | Components/SkyLightComponent.h |  |
| virtual bool CanEditChange<br>(<br>const [FProperty](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FProperty)\\* InProperty<br>) const |  | Components/SkyLightComponent.h |  |
| virtual bool IsReadyForFinishDestroy() |  | Components/SkyLightComponent.h |  |
| virtual void PostEditChangeProperty<br>(<br>[FPropertyChangedEvent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FPropertyChangedEvent)& PropertyChangedEvent<br>) |  | Components/SkyLightComponent.h |  |
| virtual void PostInitProperties() |  | Components/SkyLightComponent.h |  |
| virtual void PostLoad() | [UObject](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/UObject) Interface | Components/SkyLightComponent.h |  |
| virtual void PreEditChange<br>(<br>[FProperty](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/CoreUObject/FProperty)\\* PropertyAboutToChange<br>) |  | Components/SkyLightComponent.h |  |
| virtual void Serialize<br>(<br>[FArchive](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/FArchive)& Ar<br>) |  | Components/SkyLightComponent.h |  |

### Protected

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| void UpdateLimitedRenderingStateFast() | Fast path for updating light properties that doesn't require a re-register, Which would otherwise cause the scene's static draw lists to be recreated. | Components/SkyLightComponent.h |  |
| void UpdateOcclusionRenderingStateFast() |  | Components/SkyLightComponent.h |  |

#### Overridden from [USceneComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USceneComponent)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void OnVisibilityChanged() | Overridable internal function to respond to changes in the visibility of the component. | Components/SkyLightComponent.h |  |

#### Overridden from [UActorComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UActorComponent)

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| virtual void [CreateRenderState\_Concurrent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USkyLightComponent/CreateRenderState_Concurrent)<br>(<br>[FRegisterComponentContext](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/FRegisterComponentContext)\\* Context<br>) | Used to create any rendering thread information for this component | Components/SkyLightComponent.h |  |
| virtual void [DestroyRenderState\_Concurrent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USkyLightComponent/DestroyRenderState_Concurrent) () | Used to shut down any rendering thread structure for this component | Components/SkyLightComponent.h |  |
| virtual void [SendRenderTransform\_Concurrent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USkyLightComponent/SendRenderTransform_Concurrent) () | Called to send a transform update for this component to the rendering thread | Components/SkyLightComponent.h |  |

### Static

| Name | Remarks | Include Path | [Unreal Specifiers](https://dev.epicgames.com/documentation/unreal-engine/ufunctions-in-unreal-engine#functionspecifiers) |
| --- | --- | --- | --- |
| static void UpdateSkyCaptureContents<br>(<br>[UWorld](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UWorld)\\* WorldToUpdate<br>) | Called each tick to recapture and queued sky captures. | Components/SkyLightComponent.h |  |
| static void UpdateSkyCaptureContentsArray<br>(<br>[UWorld](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/UWorld)\\* WorldToUpdate,<br>[TArray](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Core/TArray) < [USkyLightComponent](https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Runtime/Engine/USkyLightComponent)\\* >& ComponentArray,<br>bool bBlendSources<br>) |  | Components/SkyLightComponent.h |  |

* * *