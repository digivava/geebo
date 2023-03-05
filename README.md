# geebo

A sample REST API for brand-new programmers learning Go, and learning how to work on a software engineering team for the first time.

Pretend this application is the backend for a pet-focused social media site, where folks can create profiles for their pets and show them off.

## Dependencies

Make sure you have Go installed by following [these](https://go.dev/doc/install) instructions.

## How To Run

1. Open your terminal, `cd` into where you'd like to save this repo on your computer (e.g. `cd ~/Desktop`) and copy down the repo from GitHub with `git clone `. There will now be a copy of the `geebo` repository on your computer. Go into that directory with `cd geebo`.
2. Start up the application server with `go run .`.
3. You should see a log message from the server indicating that it started up. Now you can try hitting the application's various endpoints using the command line tool `curl`. Example: `curl http://localhost:8000/users` to list all users (note that we have some code to create a few example users automatically when we first start our sample app).

When you're ready to make your own changes, practice creating your own branch with `git checkout -b <branch-name>` , then once you've written your code and tested it, create a pull request in GitHub against the `main` branch when you're ready for review.

## Exercises

1. Your teammates noticed an issue where for some reason, they aren't seeing any server logs when a request comes in to delete a user. Can you fix this?
2. We want users to be able to have usernames. Can you make the `createUsers` endpoint accept a new field called "username"?
3. Change your code to make it so if the user doesn't provide a username, we generate a default one for them that starts with their name, followed by a random number between 1 and 99.
4. Write a unit test for your username generation function.
5. Write a validation function that ensures two things:
   1. that "name" and "email" were provided in the request
   2. that there is not already an existing user with this username
6. Can you figure out how to write a test for this function?