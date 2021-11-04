# 👨🏻‍💻 - Go Url Shortener

This is a small project made with a friend from the University to learn Go and MongoDB.

# 🤔 - What is a URL Shortener?

According to Google a URL Shortener it's a type of software service that allows the use of short URLs which redirect you to the original URL. The trick of this service is to make the URL more manageable, easy to share and remember, especially with long urls.

# 📡 - Technologies used

### ♟ - Backend Stack
- ### [Go](https://golang.org/project)
- ### [Fiber](https://docs.gofiber.io/)
- ### [MongoDB](https://www.mongodb.com/)

### 💻 - Frontend Stack
- ### [HTML5](https://developer.mozilla.org/docs/Glossary/HTML5)
- ### [CSS3](https://developer.mozilla.org/docs/Web/CSS)
- ### [JavaScript](https://developer.mozilla.org/docs/Web/JavaScript)

# 🧭 - Enviorment Variables

```
DB_COLLECTION=Collection Name
DB_NAME=Database Name
DB_PASSWORD=Database password
DB_URI=Database Uri (This has to include your cluster name, database name and it's respective password)
PORT=This one is optional
```
# 👨🏻‍🔧 - How to deploy it?

1. Go to [Heroku](https://dashboard.heroku.com/login).

2. Clone the repository in your Desktop.

3. After cloning the repository in your Desktop follow the steps from Heroku and deploy it.

4. Run the procfile using the following command (it's assumed that you have Heroku CLI installed and you included the Golang WebPack) `heroku ps:scale web=1 --app app_name` replace app_name with the name of your app. You can consult it running `heroku apps` in your CLI.

5. Enjoy!
