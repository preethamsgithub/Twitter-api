INTRODUCTION

The purpose of this assignment is to help you learn how to work with external APIs, particularly the Twitter API, using a real-world programming environment. It teaches how to perform essential tasks such as authentication, making API requests, handling errors, and securely managing sensitive data (like API keys and tokens) in a development setting.

SETUP INSTRUCTIONS
•	Login to your twitter account 
•	Signup for the developers account using your gmail
•	Go to your projects and then to apps
•	Copy your generated access key, secret access key, access token, secret access token and save it
•	Go to VS code create a folder for you project and also create 2 different files one for post tweet and the other for delete post 
•	Set up your keys and tokens in env variables instead of hardcoding it.
•	Compile the files in terminal and make sure GO is installed


Program Details

Authentication Using OAuth:
•	The program first sets up OAuth 1.0a authentication by using the API Key, API Secret Key, Access Token, and Access Token Secret stored in environment variables.
•	These credentials are passed with every API request to authenticate the request as coming from the correct user.
 HTTP Request to the Twitter API:
•	The program constructs an HTTP POST request to the Twitter API endpoint:
https://api.twitter.com/1.1/statuses/update.json
•	The tweet text is passed as a parameter called status in the request body (for example: "Hello from the Twitter API!").
Headers and Signing the Request:
•	The request must include proper headers, which are signed using the OAuth keys. This ensures that Twitter knows the request is securely authenticated.
•	Libraries like go-oauth1 are used to help create and sign the request.
Identify the Tweet ID:

To delete a tweet, you need the tweet ID. This is a unique identifier for each tweet, which can be obtained from the response of the POST request or by listing tweets.
Make a Delete Request:

The program constructs an HTTP POST request to the following API endpoint:
https://api.twitter.com/1.1/statuses/destroy/:id.json
Replace :id with the actual tweet ID of the tweet you want to delete.
Authentication and Request Signing:

Similar to posting a tweet, the delete request is signed using OAuth 1.0a credentials to ensure the request is authorized.
Handling the Response:

If the request is successful (HTTP status code 200), the tweet is deleted. If there’s an error (e.g., trying to delete a non-existent tweet), an appropriate error message is returned.


