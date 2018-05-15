# Install the Heroku CLI

## Download and install the Heroku CLI.

### If you haven't already, log in to your Heroku account and follow the prompts to create a new SSH public key.
```
$ heroku login
```
Create a new Git repository

Initialize a git repository in a new or existing directory
```
$ cd my-project/
$ git init
$ heroku git:remote -a multi-langs
```
Deploy your application

Commit your code to the repository and deploy it to Heroku using Git.
```
$ git add .
$ git commit -am "make it better"
$ git push heroku master
```
Existing Git repository

For existing repositories, simply add the heroku remote
```
$ heroku git:remote -a multi-langs
```

## Connect Database
```
jdbc:postgresql://<host>:<port>/<dbname>?sslmode=require&user=<username>&password=<password>
```
```
jdbc:postgresql://host:port/database?user=username&password=secret&ssl=true&sslfactory=org.postgresql.ssl.NonValidatingFactory
```
```
jdbc:postgresql://ec2-54-225-88-191.compute-1.amazonaws.com:5432/d31ibrodp1403o?sslmode=require&user=enxzlztruhxwlg&password=bf1f4ccc6c1e2393695541093e70c2c17b994deda7163463734994f232c178e8
```
```
jdbc:postgresql://ec2-54-225-88-191.compute-1.amazonaws.com:5432/d31ibrodp1403o?user=enxzlztruhxwlg&password=bf1f4ccc6c1e2393695541093e70c2c17b994deda7163463734994f232c178e8&ssl=true&sslfactory=org.postgresql.ssl.NonValidatingFactory
```
It'Work
> https://stackoverflow.com/questions/17377118/heroku-database-connection-properties
```
jdbc:postgresql://ec2-54-225-88-191.compute-1.amazonaws.com:5432/d31ibrodp1403o?&ssl=true&sslfactory=org.postgresql.ssl.NonValidatingFactory
```

## Build Image
```
docker build -t multi-langs .
```

## API

### Attributes
#### GET

> Request 
```
http://localhost:9000/api/v1/attributes
```
> Parameters
- app=hotel-connect
- lang=th

> Path
```
http://localhost:9000/api/v1/attributes/:id
```

> Header
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJvbmdiYW5nIiwiYWRtaW4iOnRydWUsImV4cCI6MTUxNzkyNzE0M30.MkLPK_UzhcGTVBW-ThLReYdxVbjNtThBKdcXjQMkq-k"
```

> Response
```json
[
    {
        "id": "2f032d6d-8d11-4f14-9579-24f052f470d2",
        "app": "hotel-connect",
        "en": {
            "label_name": "Hotel Connect"
        },
        "th": {
            "label_name": "เชื่อมต่อโรงแรม"
        }
    },
    {
        ...
    }
]
```

#### POST

> Request
```
http://localhost:9000/api/v1/attributes
```
> Header
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJvbmdiYW5nIiwiYWRtaW4iOnRydWUsImV4cCI6MTUxNzkyNzE0M30.MkLPK_UzhcGTVBW-ThLReYdxVbjNtThBKdcXjQMkq-k"
```
> Body
```json
{
    "app": "hotel-connect",
    "th": {
        "label_name": "เชื่อมต่อโรงแรม"
    },
    "en": {
        "label_name": "Hotel Connect"
    }
}
```

> Response
```json
{
    "id": "2f032d6d-8d11-4f14-9579-24f052f470d2",
    "app": "hotel-connect",
    "th": {
        "label_name": "เชื่อมต่อโรงแรม"
    },
    "en": {
        "label_name": "Hotel Connect"
    }
}
```

#### PUT
> Request
```
http://localhost:9000/api/v1/attributes/:id
```
> Header
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJvbmdiYW5nIiwiYWRtaW4iOnRydWUsImV4cCI6MTUxNzkyNzE0M30.MkLPK_UzhcGTVBW-ThLReYdxVbjNtThBKdcXjQMkq-k"
```
> Body
```json
{
	"app": "hotel-connect",
	"th": {
	 		"label_name": "เชื่อมต่อโรงแรม"
 	},
	"en": {
	 		"label_name": "Hotel Connect"
	},
	"chn": {
		"label_name": "酒店連接"
	}
}
```
> Response
```json
{
    "id": "2f032d6d-8d11-4f14-9579-24f052f470d2",
    "app": "hotel-connect",
    "chn": {
        "label_name": "酒店連接"
    },
    "en": {
        "label_name": "Hotel Connect"
    },
    "th": {
        "label_name": "เชื่อมต่อโรงแรม"
    }
}
```

#### DELETE
> Request
```
http://localhost:9000/api/v1/attributes/:id
```
> Header
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJvbmdiYW5nIiwiYWRtaW4iOnRydWUsImV4cCI6MTUxNzkyNzE0M30.MkLPK_UzhcGTVBW-ThLReYdxVbjNtThBKdcXjQMkq-k"
```
> Body
```json
{
    "id": "2f032d6d-8d11-4f14-9579-24f052f470d2"
}
```

### Language

#### GET
> Request
```
http://localhost:9000/api/v1/language
```
> Parameters
- app=hotel-connect
- lang=th
- id=b49d93d0-85ec-4a62-b904-a548e5be5671

> Path
```
http://localhost:9000/api/v1/attributes/:id
```

> Header
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJvbmdiYW5nIiwiYWRtaW4iOnRydWUsImV4cCI6MTUxNzkyNzE0M30.MkLPK_UzhcGTVBW-ThLReYdxVbjNtThBKdcXjQMkq-k"
```

> Response
```json
{
    "app": "hotel-connect",
    "id": "b49d93d0-85ec-4a62-b904-a548e5be5671",
    "th": {
        "image": "https://findicons.com/files/icons/2838/flat_round_world_flag_icon_set/512/thailand.png",
        "name": "ไทย"
    }
}
```

#### POST
> Request
```
http://localhost:9000/api/v1/language
```

> Header
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHJvbmdiYW5nIiwiYWRtaW4iOnRydWUsImV4cCI6MTUxNzkyNzE0M30.MkLPK_UzhcGTVBW-ThLReYdxVbjNtThBKdcXjQMkq-k"
```

> Body
```json
{
	"app": "hotel-connect",
	"th": {
		"name": "ไทย",
		"image": "https://findicons.com/files/icons/2838/flat_round_world_flag_icon_set/512/thailand.png"
	}
}
```
> Response
```json
{
    "app": "hotel-connect",
    "id": "b49d93d0-85ec-4a62-b904-a548e5be5671",
    "th": {
        "image": "https://findicons.com/files/icons/2838/flat_round_world_flag_icon_set/512/thailand.png",
        "name": "ไทย"
    }
}
```
