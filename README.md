# Quick (& Dirty) Dotnet Uninstall Tool

QDUT is a simple tool which will list installed .NET SDKs and Runtimes and provide an easy way to uninstall them on Linux and macOS.

QDUT is a similar the [dotnet-core-uninstall](https://github.com/dotnet/cli-lab) tool, however I have found that it tends to be a while before the latest .NET versions are supported. QDUT will be .NET version agnostic and will look for installed .NET SDKs and Runtimes in the dotnet directory, list them and provide an option to uninstall (delete) them.

Initially this will be a very simple tool, but I plan to add more flags and options to make it more dynamic in what it can do (uninstall SDKs and runtimes for a specific version, for example).
