# Golang CRUD: Data CV

Implement basic CRUD with ORM.  
Golang, MySQL

## How to run
1. Create database
2. Change .env values
3. Run `go build main.go`
4. Run `go run main.go`
5. Open `localhost:8080`

## Database Schema
### CV
| Field  | Data Type | Description |
| ------------- | ------------- | -------- |
| name  | string  | |           
| age  | int  | |
| summary  | string | Short description about self |
| photo  | string | Photo url |

### Education
| Field  | Data Type | Description |
| ------------- | ------------- | -------- |
| cv_id  | int  | FK |           
| title  | string  | |
| major  | string |  |
| school  | string |  |
| description  | string |  |
| start_year  | string |  |
| end_year  | string |  |

### Experience
| Field  | Data Type | Description |
| ------------- | ------------- | -------- |
| cv_id  | int  | FK |           
| title  | string  | |
| company  | string |  |
| description  | string |  |
| start_year  | string |  |
| end_year  | string |  |

### Contact
| Field  | Data Type | Description |
| ------------- | ------------- | -------- |
| cv_id  | int  | FK |           
| type  | string  | Contact type e.g: phone, email, etc |
| value  | string |  |

## Example
### Create New CV
`
{
    "name": "John",
    "age": 21,
    "summary": "I am a developer",
    "photo": "www.google.com",
    "educations": [
        {
            "title": "Undergraduate Student",
            "major": "Computer Science",
            "school": "Harvard",
            "description": "I am a passionate software engineer",
            "start_year": 2019
        }
    ],
    "experiences": [
        {
            "title": "Software Engineer",
            "company": "Yukbisnis",
            "description": "My role here as backend engineer",
            "start_year": 2022
        }
    ],
    "contacts": [
        {
            "type": "phone",
            "value": "08123456"
        }
    ]
}
`

## Author
Thomas Ferdinand Martin