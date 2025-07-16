/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

import path from "node:path";
import semver from "semver";
import * as fs from "./fs.js";
import * as vi from "./versioninfo.js";
import { SDK } from "../custom/globals.js";
import { isOnline } from "./index.js";

export type ModusAppInfo = {
  name: string;
  sdk: SDK;
  sdkVersion: string;
};

export async function getAppInfo(appPath: string): Promise<ModusAppInfo> {
  const appInfo = await getInfoFromApp(appPath);

  if (appInfo.sdkVersion == "latest") {
    const online = await isOnline();
    const version = online ? await vi.getLatestSdkVersion(appInfo.sdk, false) : await vi.getLatestInstalledSdkVersion(appInfo.sdk, false);
    if (version) {
      appInfo.sdkVersion = version;
    } else {
      throw new Error(`Could not determine the latest version of the ${appInfo.sdk} SDK`);
    }
  }

  return appInfo;
}

async function getInfoFromApp(appPath: string): Promise<ModusAppInfo> {
  if (await fs.exists(path.join(appPath, "package.json"))) {
    return await getAssemblyScriptAppInfo(appPath);
  }

  if (await fs.exists(path.join(appPath, "go.mod"))) {
    return await getGoAppInfo(appPath);
  }

  if (await fs.exists(path.join(appPath, "moon.mod.json")) ||
      await fs.exists(path.join(appPath, "moon.pkg.json"))) {
    return await getMoonBitAppInfo(appPath);
  }

  throw new Error("Could not determine which Modus SDK to use.");
}

async function getAssemblyScriptAppInfo(appPath: string): Promise<ModusAppInfo> {
  const sdk = SDK.AssemblyScript;

  let name: string;
  try {
    const appPackage = JSON.parse(await fs.readFile(path.join(appPath, "package.json"), "utf8"));
    name = appPackage.name;
  } catch {
    throw new Error("Failed to read name from package.json");
  }

  let sdkVersion: string | undefined;
  try {
    if (await fs.exists(path.join(appPath, "package-lock.json"))) {
      const lockfile = JSON.parse(await fs.readFile(path.join(appPath, "package-lock.json"), "utf8"));
      sdkVersion = lockfile.packages["node_modules/@hypermode/modus-sdk-as"].version;
    }
  } catch {
    /* empty */
  }

  if (!sdkVersion || !semver.valid(sdkVersion)) {
    try {
      const appPackage = JSON.parse(await fs.readFile(path.join(appPath, "package.json"), "utf8"));
      const constraint = appPackage.dependencies["@hypermode/modus-sdk-as"];
      if (semver.valid(constraint)) {
        sdkVersion = constraint;
      } else if (semver.validRange(constraint)) {
        const versions = (await vi.getInstalledSdkVersions(sdk)).map((v) => v.slice(1));
        const maxSatisfying = semver.maxSatisfying(versions, constraint, { includePrerelease: true });
        if (maxSatisfying) {
          sdkVersion = maxSatisfying;
        }
      }
    } catch {
      /* empty */
    }
  }

  if (!sdkVersion || !semver.valid(sdkVersion)) {
    sdkVersion = "latest";
  } else {
    sdkVersion = "v" + sdkVersion;
  }

  return { name, sdk, sdkVersion };
}

async function getGoAppInfo(appPath: string): Promise<ModusAppInfo> {
  const sdk = SDK.Go;

  const data = await fs.readFile(path.join(appPath, "go.mod"), "utf8");
  const lines = data.split("\n");

  const moduleLine = lines.find((line) => line.startsWith("module"));
  if (!moduleLine) {
    throw new Error("Could not determine the module name from go.mod");
  }
  const name = moduleLine.split(" ")[1];

  const modName = "github.com/gmlewis/modus/sdk/go";
  const versionLine = lines.find((line) => line.includes(modName))?.trim();
  let sdkVersion: string | undefined;
  try {
    if (versionLine) {
      const parts = versionLine.split(" ");
      if (parts[0] == "require") {
        sdkVersion = parts[2];
      } else {
        sdkVersion = parts[1];
      }
    }
  } catch {
    /* empty */
  }
  if (!sdkVersion || sdkVersion == "v0.0.0" || !sdkVersion.startsWith("v")) {
    sdkVersion = "latest";
  }

  return { name, sdk, sdkVersion };
}

async function getMoonBitAppInfo(appPath: string): Promise<ModusAppInfo> {
  const sdk = SDK.MoonBit;

  const data = await fs.readFile(path.join(appPath, "moon.mod.json"), "utf8");
  const fields = JSON.parse(data);

  const moduleName = fields.name;
  if (!moduleName) {
    throw new Error("Could not determine the module 'name' from moon.mod.json");
  }
  const name = moduleName.split("/").pop();  // Return module name after last '/'.

  const versionField = fields.deps?.["gmlewis/modus"];
  let sdkVersion: string | undefined;
  try {
    if (versionField?.path) {
      // retrieve the `version` field from the `moon.mod.json` file in the SDK itself:
      const data = await fs.readFile(path.join(appPath, versionField.path, "moon.mod.json"), "utf8");
      const fields = JSON.parse(data);
      if (fields?.version) {
        sdkVersion = "v"+fields.version;
      }
    } else if (versionField && typeof versionField === "string") {
      sdkVersion = "v"+versionField;
    }
  } catch {
    /* empty */
  }
  if (!sdkVersion || sdkVersion == "v0.0.0" || !sdkVersion.startsWith("v")) {
    sdkVersion = "latest";
  }

  return { name, sdk, sdkVersion };
}
