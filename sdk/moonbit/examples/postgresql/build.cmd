@echo off

:: This build script works best for examples that are in this repository.
:: If you are using this as a template for your own project, you may need to modify this script,
:: to invoke the modus-moonbit-build tool with the correct path to your project.

SET "PROJECTDIR=%~dp0"
pushd ..\..\tools\modus-moonbit-build > nul
go run . "%PROJECTDIR%"
set "exit_code=%ERRORLEVEL%"
popd > nul
exit /b %exit_code%
