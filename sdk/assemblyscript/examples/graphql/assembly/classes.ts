/*
 * This example is part of the Modus project, licensed under the Apache License 2.0.
 * You may modify and use this example in accordance with the license.
 * See the LICENSE file that accompanied this code for further details.
 */

// These classes are used by the example functions in the index.ts file.

@json
export class Country {
  code: string = "";
  name: string = "";
  capital: string = "";
}


@json
export class CountriesResponse {
  countries: Country[] = [];
}


@json
export class CountryResponse {
  country: Country | null = null;
}
