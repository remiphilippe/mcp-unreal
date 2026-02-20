<!-- Source: https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-command-line-arguments-reference -->

## Flags

Unreal Engine command-line flags are passed to the executable to specify certain behavior. For example, to run the Unreal Editor with the `DumpAssetRegistry` flag, use

```cpp
UnrealEditor.exe -DumpAssetRegistry
```

Command-line arguments are not case-sensitive.

### Table of Flags

| **Argument** | **Description** |
| --- | --- |
| `AllowCommandletRendering` | Allow commandlet rendering |
| `AllowSoftwareRendering` | Allows the D3D11 and D3D12 RHI to fall back to software rendering |
| `AsyncLoadingThread` | Enable async loading thread for package streaming |
| `AttachRenderDoc` | Attach RenderDoc to process |
| `AudioMixer` | Force load the audio mixer |
| `AutoQuit` | Close application when analysis fails to start or completes successfully |
| `BENCHMARK` | Set benchmarking |
| `BUILDMACHINE` | Set this machine as a build machine |
| `CleanCrashReports` | Clean crash report folder located in `../Saved/Crashes` |
| `CookOnTheFly` | Cook-on-the-fly server |
| `crashreports` | Always display crash reports |
| `d3d11` | Use D3D11 RHI |
| `d3d12` | Use D3D12 RHI |
| `d3ddebug` | Use a debug device for D3D |
| `Deterministic` | Shortcut for `-UseFixedTimeStep -FixedSeed` |
| `DisablePython` | Disable python scripting |
| `DumpAssetRegistry` | Dump extended info about asset registry to log |
| `dumpconfig` | Dump all configuration settings to log |
| `DumpRPCs` | Dump all RPCs and a full parameter list to log |
| `FATALSCRIPTWARNINGS` | Treat script warnings as fatal errors |
| `FixedSeed` | Use 0 as the seed for `FRandomStream` |
| `forcelogflush` | Force a log flush after each line |
| `fullcrashdump` | Create full memory minidumps for crashes |
| `fullscreen` | Use fullscreen mode |
| `game` | Run in game mode |
| `gpucrashdebugging` | Enable all possible GPU crash debug modes |
| `iterate` | Use iterative cooking |
| `LLM` | Enable low-level memory tracker |
| `LOG` | Open a new log window |
| `MAXQUALITYMODE` | Set all `r.Shadow` console variables to maximum settings |
| `Messaging` | Explicitly enable the messaging module |
| `NoAsyncLoadingThread` | Disable async loading thread for package streaming |
| `NOAUTOINIUPDATE` | Disable automatic updating of configuration (`.ini`) files |
| `NoCache` | Disable cache |
| `NOCONSOLE` | Disable the console |
| `NoGamepad` | Force disable gamepads |
| `nohmd` | Disable HMD device |
| `NoMCP` | Disable MCP backend |
| `NoPak` | Disable PAK file |
| `noraytracing` | Disable ray tracing |
| `norenderthread` | Disable the render thread |
| `NoShaderCompile` | Do not compile shaders |
| `NOSPLASH` | Disable splash screen |
| `nothreading` | Disable multithreading |
| `novsync` | Set `r.vsync 0` |
| `nullrhi` | Use null rendering hardware interface to run UE headless |
| `onethread` | Use single-threading |
| `OpenGL` | Use OpenGL |
| `parallelrendering` | Enable parallel rendering |
| `REGENERATEINIS` | Regenerate configuration files |
| `RenderOffScreen` | Render off screen |
| `rhithread` | Enable RHI thread |
| `RHIValidation` | Enable RHI validation |
| `SILENT` | Disable all log text output |
| `sm5` | Force use SM5 |
| `sm6` | Force use SM6 |
| `stdout` | Use stdout for log output |
| `unattended` | Run in unattended mode (disable UI pop-ups and dialogs) |
| `USEALLAVAILABLECORES` | Use all available cores |
| `UseFixedTimeStep` | Use a fixed time step |
| `UseIoStore` | Force use Io store |
| `UsePaks` | Use PAK files |
| `Verbose` | Use verbose logging |
| `vr` | Use VR mode |
| `vsync` | Set `r.vsync 1` |
| `vulkan` | Use Vulkan RHI |
| `vulkandebug` | Enable Vulkan validation (errors and warnings) |
| `waitforattach` | Halt startup and wait for debugger to attach before continuing |
| `WARNINGSASERRORS` | Treat warnings as errors |
| `Windowed` | Use windowed mode |

## Keyword Arguments

The syntax to run keyword arguments is `-<Keyword>=<Value>`.

For example, to listen for incoming online beacon connections on port `8888`:

```cpp
UnrealEditor.exe -BeaconPort=8888
```

### List of Keyword Arguments

| **Argument** | **Description** |
| --- | --- |
| `ABSLOG` | Absolute log filename |
| `BeaconPort` | Override default port for online beacon host |
| `BENCHMARKSECONDS` | Add a seconds setting for benchmarking |
| `Builder` | World partition builder commandlet builder class name |
| `CULTURE` | Text localization manager uses the provided culture |
| `DeviceProfile` | Override device profile name |
| `EXEC` | Execute the specified exec file |
| `ExecCmds` | Execute the specified console commands (Usage: `-ExecCmds="Cmd1,Cmd2"`) |
| `filehostip` | File host IP |
| `FPS` | Override fixed tick rate frames per second |
| `graphicsadapter` | Select graphics adapter |
| `LOG` | Specify log file name |
| `map` | Use the provided map |
| `MaxGPUCount` | Maximum number of GPUs |
| `MULTIHOME` | Multihome IP address |
| `NetDriverOverrides` | Override network drivers |
| `Port` | Default server port |
| `Project` | Path to `.uproject` file |
| `Res` | Set the window resolution (Usage: `-Res=1280x768`) |
| `ResX` | Specify window width resolution |
| `ResY` | Specify window height resolution |
| `SECONDS` | Maximum tick time in seconds |
| `TargetPlatform` | Target platform |
| `WinX` | Set the initial horizontal window position |
| `WinY` | Set the initial vertical window position |

For the complete list of all flags and keyword arguments, see the [full reference page](https://dev.epicgames.com/documentation/en-us/unreal-engine/unreal-engine-command-line-arguments-reference).
