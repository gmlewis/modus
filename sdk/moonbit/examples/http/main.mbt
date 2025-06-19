// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///| This function makes a simple HTTP GET request to example.com, and returns the HTML text of the response.
pub fn get_example_html() -> String raise Error {
  let response = @http.fetch_url("https://example.com/")
  if not(response.ok()) {
    raise @http.HttpError(
      "Failed to fetch url. Received: \{response.status} \{response.status_text}",
    )
  }
  response.text()
}

///| This function makes a request to an API that returns data in JSON format, and returns an object representing the data.
/// It also demonstrates how to check the HTTP response status.
pub fn get_random_quote() -> Quote raise Error {
  let response = @http.fetch_url("https://zenquotes.io/api/random")
  if not(response.ok()) {
    raise @http.HttpError(
      "Failed to fetch quote. Received: \{response.status} \{response.status_text}",
    )
  }
  let quotes : Array[Quote] = response.json()
  quotes[0]
}

///| This function makes a request to an API that returns an image, and returns the image data.
pub fn get_random_image(width : Int, height : Int) -> Image raise Error {
  let url = "https://picsum.photos/\{width}/\{height}"
  let response = @http.fetch_url(url)
  if not(response.ok()) {
    raise @http.HttpError(
      "Failed to fetch image. Received: \{response.status} \{response.status_text}",
    )
  }
  {
    contentType: response.headers.unwrap().get("Content-Type"),
    data: response.to_ints(),
  }
}

///| This function demonstrates a more complex HTTP call.
/// It makes a POST request to the GitHub API to create an issue.
///
/// See https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#create-an-issue
///
/// To use it, you must add a GitHub personal access token to your secrets.
/// Create a fine-grained token at https://github.com/settings/tokens?type=beta with access
/// to write issues to the repository you want to use, then add it to the appropriate secret
/// store for your environment.  (See the modus documentation for details.)
///
/// NOTE: Do not pass the Authorization header in code when creating the request.
/// That would be a security risk, as the token could be exposed in the source code repository.
///
/// Instead, configure the headers in the modus.json manifest as follows:
///
/// 	"hosts": {
/// 	  "github": {
/// 	    "baseUrl": "https://api.github.com/",
/// 	    "headers": {
/// 	      "Authorization": "Bearer {{AUTH_TOKEN}}"
/// 	    }
/// 	  }
/// 	}
///
/// The Modus runtime will retrieve the token from your secrets and add it to the request.
pub fn create_github_issue(
  owner : String,
  repo : String,
  title : String,
  body : String
) -> Issue raise Error {
  let url = "https://api.github.com/repos/\{owner}/\{repo}/issues"
  let request = @http.Request::new(url, options=[
    @http.Method("POST"),
    @http.HeaderMap({
      "Accept": "application/vnd.github+json",
      "X-GitHub-Api-Version": "2022-11-28",
      "Content-Type": "application/json",
    }),
    @http.BodyJson(
      { "title": title.to_json(), "body": body.to_json() }.to_json(),
    ),
  ])
  let response = request.fetch()
  if not(response.ok()) {
    raise @http.HttpError(
      "Failed to create issue. Received: \{response.status} \{response.status_text}",
    )
  }
  response.json()
}
