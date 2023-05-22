# Node App Launcher

# Prerequests (Development Only)

```bash
# Install go-winres. go-winres is required to generate *.syso files which embed Windows Icon for exe file
go install github.com/tc-hib/go-winres@latest
```

# Building

```bash
yarn build:win
yarn build:darwin
```

# Packing
## MacOS
```bash
yarn pack:darwin
```

## Windows

> Optional Optimize Executable Size
> ```bat
> upx node-app-launcher.exe
> ```

```bash
yarn pack:win
```

## Debugging

```bash
# Build exe which can be attached for debugging
yarn build4attach
# Copy this exe to app folder
cp ./node-app-launcher /Applications/node-app-launcher.app/Contents/MacOS/node-app-launcher
# Run copied exe
/Applications/node-app-launcher.app/Contents/MacOS/node-app-launcher
# GoLand Menu/Run/Attach to Processes, Select node-app-launcher

```

