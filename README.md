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