## How to run assignment
This is a dockerized app, to run this application make sure you have docker installed and running on your local system. Use the `docker-compose.yml` file in the project to run it.

## API Specs
This project has four end-points.

### `POST /signup `
**Path** `dapper-lab/signup` </br>
Endpoint to create an user row in postgres db. The payload should have the following fields, all the fields are required (there is a proper validation for it):

```json
{
  "email": "test@gmail.com",
  "password": "password",
  "firstName": "fName",
  "lastName": "lName"
}
```

where `email` is an unique key in the database.

The response body return a JWT on success which is valid for 1 hr that can be used for other endpoints:

```json
{
  "token": "some_jwt_token" 
}
```

### `POST /login`

**Path** `dapper-lab/login` </br>
Endpoint to log an user in. The payload should have the following fields:
```json
{
  "email": "test@gmail.com",
  "password": "password"
}
```

The response body should return a JWT on success that can be used for other endpoints:

```json
{
  "token": "some_jwt_token"
}
```

### `GET /users`
***Path*** `/auth/dapper-lab/users` </br>
Endpoint to retrieve a json of all users. This endpoint requires a valid `x-authentication-token` header to be passed in with the request.

The response body should look like:
```json
 [
    {
      "email": "test@gmail.com",
      "firstName": "fName",
      "lastName": "lName"
    }
  ]

```

### `PUT /user`
***Path*** `/auth/dapper-lab/user` </br>
Endpoint to update the current user `firstName` or `lastName` only. This endpoint requires a valid `x-authentication-token` header to be passed in and it should only update the user of the JWT being passed in. The payload can have the following fields:

```json
{
  "firstName": "NewFirstName",
  "lastName": "NewLastName"
}
```

The response body is empty.

## What I could have improved, if I had more time
1) A robost error handling, I should have created an error struct, and I would manage to get more human-readable errors.
2) A signup in proper transaction as we are touching two tables in the same method.

## Note
Please feel free to reach out to me if you want to discuss any part of the code.

## Q and A
1) How long did this assignment take? 
 </br>***Ans*** - Around 6 hours
2) What was the hardest part? 
 </br> ***Ans*** - All the challenges were equally fun to work with but if I have to pick one than it would be the Designing the middleware
3) Did you learn anything new?
   </br> ***Ans*** - Yeah, everyday you learn something new. During my interview process, I had chance to look at the Dapper labs and its product. 
4) Is there anything you would have liked to implement but didn't have the time to?
   </br> ***Ans*** I already answered it in `What I could have Improved`
5) What are the security holes (if any) in your system? If there are any, how would you fix them?
   </br> ***Ans*** - As I am submitting this task right after my Red hot rotation, I have tested it with quite few scenarios and it is working fine. For past one week, I was fix any bug on prod on real time (because of red hot rotation) so, I didn't had time to think a lot about the Security flaws in my code.
6) Do you feel that your skills were well tested?
</br> ***Ans*** - Absolutely.