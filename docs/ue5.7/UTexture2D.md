# UTexture2D

**Parent**: UTexture
**Module**: Engine

Two-dimensional texture asset used in materials, UI, and render targets. Stores compressed or uncompressed pixel data that the GPU samples during rendering. Supports mipmapping, multiple compression formats, and platform-specific streaming. Transient instances can be created at runtime for procedural or dynamic content.

## Key Properties

- `SizeX` — `int32` width of the texture in pixels at mip level 0
- `SizeY` — `int32` height of the texture in pixels at mip level 0
- `PixelFormat` — `EPixelFormat` enum describing the GPU texture format: `PF_B8G8R8A8` (uncompressed BGRA), `PF_DXT1` (BC1, no alpha), `PF_DXT5` (BC3, with alpha), `PF_BC5` (normal maps), `PF_R16F`, `PF_FloatRGBA`, and many others
- `LODBias` — `int32` offset applied to the mip level selection; positive values force lower-resolution mips to be used
- `CompressionSettings` — `TextureCompressionSettings` enum controlling how the source image is compressed: `TC_Default`, `TC_Normalmap`, `TC_Masks`, `TC_Grayscale`, `TC_HDR`, `TC_EditorIcon`, `TC_Alpha`, `TC_DistanceFieldFont`, `TC_HDRCompressed`, `TC_BC7`
- `Filter` — `TextureFilter` enum controlling texture sampling interpolation: `TF_Nearest`, `TF_Bilinear`, `TF_Trilinear`, `TF_Default`
- `AddressX` — `TextureAddress` enum for horizontal wrapping behavior: `TA_Wrap`, `TA_Clamp`, `TA_Mirror`
- `AddressY` — `TextureAddress` enum for vertical wrapping behavior: `TA_Wrap`, `TA_Clamp`, `TA_Mirror`
- `SRGB` — `bool` indicating whether pixel data is stored in sRGB color space and should be converted to linear during sampling
- `bHasBeenPaintedInEditor` — `bool` set when the texture has been modified by the in-editor mesh paint tool
- `MaxTextureSize` — `int32` clamping the maximum resolution used at runtime, regardless of the source asset resolution

## Key Functions

- `CreateTransient(int32 InSizeX, int32 InSizeY, EPixelFormat InFormat, const FName InName)` — Static factory returning a new `UTexture2D*` with no package/asset association; use for runtime-generated textures
- `GetSizeX()` — Returns `int32` width of the texture at the current mip level
- `GetSizeY()` — Returns `int32` height of the texture at the current mip level
- `GetPixelFormat()` — Returns `EPixelFormat` of the platform texture resource currently loaded in GPU memory
- `GetNumMips()` — Returns `int32` total number of mip levels present in the texture
- `GetSurfaceWidth()` — Returns `float` physical surface width in texels, accounting for any platform alignment padding
- `GetSurfaceHeight()` — Returns `float` physical surface height in texels, accounting for any platform alignment padding
- `UpdateResource()` — Releases the existing render resource and re-creates it from current source data; call after programmatically modifying pixel data
- `GetPlatformData()` — Returns `FTexturePlatformData*` containing the mip chain and raw pixel data for the current target platform
- `GetRunningPlatformData()` — Returns `FTexturePlatformData**` pointer to the platform data actively used by the running game, which may differ from editor-side data during cooking
