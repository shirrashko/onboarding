{\rtf1\ansi\ansicpg1252\cocoartf2709
\cocoatextscaling0\cocoaplatform0{\fonttbl\f0\fswiss\fcharset0 Helvetica;}
{\colortbl;\red255\green255\blue255;}
{\*\expandedcolortbl;;}
\paperw11900\paperh16840\margl1440\margr1440\vieww34360\viewh19240\viewkind0
\pard\tx720\tx1440\tx2160\tx2880\tx3600\tx4320\tx5040\tx5760\tx6480\tx7200\tx7920\tx8640\pardirnatural\partightenfactor0

\f0\fs24 \cf0 # User Profile Service - Stage 2: Get some REST\
\
This project focuses on building a service with a RESTful API for managing user profiles. It is part of a larger exercise and involves designing data models, creating a REST API, and implementing endpoints.\
\
## Table of Contents\
\
- [Prerequisites](#prerequisites)\
- [Actions](#actions)\
- [Knowledge](#knowledge)\
- [Steps](#steps)\
- [Design](#design)\
- [API Documentation](#api-documentation)\
- [Implementation](#implementation)\
- [Testing](#testing)\
\
## Prerequisites\
\
Before starting this project, you should have completed Stage 1 and have a basic understanding of backend concepts, the HTTP protocol, and REST API principles.\
\
## Actions\
\
- Complete Stage 1 before proceeding with this stage.\
\
## Knowledge\
\
This project assumes familiarity with:\
\
- Backend basic concepts\
- HTTP protocol\
- REST APIs concepts and best practices\
\
## Steps\
\
1. Design data models and structures for user profiles, ensuring each user has a unique ID for read and update operations.\
2. Design a RESTful API for the user profile service.\
3. Discuss the API design with your mentor before implementation using tools like draw.io.\
4. Document the API endpoints comprehensively for future developers who will implement clients.\
5. Implement the API endpoints, adhering to proper abstraction and separation of business logic from the API layer.\
6. Write tests to ensure the functionality of your code.\
7. Collaborate with your mentor to determine the scope and nature of tests to be written.\
\
## Design\
\
### Data Models\
\
- Each user has a unique ID for read and update operations.\
- Profile data includes:\
  - Username (unique per user)\
  - Full name (as one string)\
  - Bio (short text describing the user)\
  - Profile picture URL\
\
### API Design\
\
**Create User Profile:**\
- HTTP Method: POST\
- Path: `/profiles`\
- Request Body: User profile data\
- Response: Created user profile with unique ID\
\
**Get User Profile:**\
- HTTP Method: GET\
- Path: `/profiles/\{userID\}`\
- Response: User profile data\
\
**Update User Profile:**\
- HTTP Method: PUT or PATCH\
- Path: `/profiles/\{userID\}`\
- Request Body: Updated user profile data\
- Response: Updated user profile data\
\
## API Documentation\
\
- Provide detailed documentation for each API endpoint, including expected request and response formats, parameters, and headers.\
\
## Implementation\
\
- Implement the API endpoints following best practices.\
- Ensure proper separation of business logic components from the API layer.\
- The business logic layer should not return HTTP errors but should return errors that the API layer can translate to proper HTTP errors.\
\
## Testing\
\
- Write tests to validate the functionality of the implemented endpoints.\
- Collaborate with your mentor to decide on the scope and type of tests to be written.\
\
## Conclusion\
\
This project aims to create a user profile service with a RESTful API, emphasizing proper design, documentation, implementation, and testing practices. By following these steps, you will build a functional and maintainable solution.\
\
}