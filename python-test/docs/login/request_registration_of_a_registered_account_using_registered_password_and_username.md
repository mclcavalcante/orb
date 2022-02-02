## Scenario: Request registration of a registered account using registered password and username 
## Steps: 


1 - Request an account registration using an already registered email, same registered password and fullname field filled

- REST API Method: POST
- endpoint: /users
- body: `{"email":"already_registered_email", "password":"registered_password", "metadata":{"fullName":"name"}}`

## Expected Result:

- The request must fail with conflict (error 409), response message must be "email already taken"
- No changes should be made to the previously registered account
  (name, company and password must be the ones registered for the first time)
