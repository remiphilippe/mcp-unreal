# FAssetData

**Parent**: (struct — no parent class)
**Module**: AssetRegistry

Lightweight asset registry metadata struct for querying and referencing assets without loading them into memory. The Asset Registry populates `FAssetData` entries by scanning the project's asset files on startup. Use `IAssetRegistry::GetAssetsByClass`, `GetAssetsByPath`, or `GetAssets` to retrieve collections. `FAssetData` is the preferred way to browse the content browser, build asset pickers, and validate asset references without triggering asset loads.

## Key Properties

- `ObjectPath` — `FName` of the full object path including package and object name (e.g. `/Game/Characters/Hero.Hero`); deprecated in UE 5.1+ in favour of `GetSoftObjectPath()`
- `PackageName` — `FName` of the package (`.uasset` file path without extension) that contains this asset (e.g. `/Game/Characters/Hero`)
- `PackagePath` — `FName` of the directory portion of the package path (e.g. `/Game/Characters`)
- `AssetName` — `FName` of the asset's object name within the package (e.g. `Hero`)
- `AssetClass` — Deprecated in UE 5.7; use `AssetClassPath` instead; was a `FName` of the native class name
- `AssetClassPath` — `FTopLevelAssetPath` containing the module-qualified class path (e.g. `{"/Script/Engine", "SkeletalMesh"}`); the correct way to identify asset type in UE 5.1+
- `TagsAndValues` — `FAssetDataTagMapSharedView` (immutable shared map) of asset registry tag name–value pairs extracted from the asset's `GetAssetRegistryTags()` implementation; contains metadata like thumbnail info, import settings, and custom per-class tags

## Key Functions

- `GetAsset()` — Loads and returns the `UObject*` for this asset; triggers a synchronous load if the asset is not already in memory; prefer `TSoftObjectPtr` for lazy loading
- `GetClass()` — Resolves and returns the `UClass*` corresponding to `AssetClassPath`; may return null if the class is not loaded
- `IsValid()` — Returns `true` if the struct contains valid (non-empty) package and asset name data; does not check whether the asset exists on disk
- `IsUAsset()` — Returns `true` if the asset is the primary (top-level) export of its package, i.e. `AssetName == FPackageName::GetShortName(PackageName)`
- `GetFullName()` — Returns an `FString` in the form `ClassName /Package/Path.AssetName`; equivalent to `UObject::GetFullName()` for the unloaded asset
- `GetSoftObjectPath()` — Returns an `FSoftObjectPath` pointing to this asset; preferred over `ObjectPath` for asset references in UE 5.x
- `GetPrimaryAssetId()` — Returns the `FPrimaryAssetId` (type + name) if this asset implements `GetPrimaryAssetId()` via the Asset Manager; returns an invalid id if not a primary asset
- `GetTagValue(FName Tag, FString& OutValue)` — Looks up `Tag` in `TagsAndValues`; writes the string value to `OutValue` and returns `true` if found
- `IsRedirector()` — Returns `true` if the asset is a `UObjectRedirector`; redirectors are stale references left after asset moves or renames
- `IsAssetLoaded()` — Returns `true` if the referenced `UObject` is currently resident in memory without triggering a load
- `ToSoftObjectPath()` — Returns an `FSoftObjectPath` for this asset; alias for `GetSoftObjectPath()`
- `GetExportTextName()` — Returns the asset path in `ExportText` format (`ClassName'/Package/Path.AssetName'`) used by copy/paste and property serialisation
- `PrintAssetData()` — Logs all fields of the struct (package name, class, tags) to the output log; useful for debugging Asset Registry queries
