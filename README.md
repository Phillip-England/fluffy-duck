# Cloning the Repository
Tune this command to clone the repository. Make sure to include your actual path before running as this snippet does not contain a valid path.
```bash
git clone https://github.com/phillip-england/go-web-quickstart <your-project-dir>
```

# Installing Tailwind
This project relies on tailwind for styling. Run the following command from your project root to install it.
```bash
npm install
```

# Environment Variables
This project relies on a number of local environment variables. Here is a list of the variables you will need to set. They are written in the following format: VARIABLE_NAME=description of variable
```bash
TEAM_USERNAME=a username for your team members
TEAM_PASSWORD=a password for your team members
TEAM_SESSION_TOKEN=a session token for your team members
ADMIN_USERNAME=a username for your admins
ADMIN_PASSWORD=a password for your admins
ADMIN_SESSION_TOKEN=a session token for your admins
FONTAWESOME_SRC=a fontawesome src value as provided by fontaweomse.com for icons
POSTGRES_CONNECTION_STRING=a postgres database connection string
```

# Bundling
This project uses Bun to bundle our client-side javascript into a single file. The package.json file contains many core scripts for building and running the app. To bundle, simply run:
```bash
npm run build
```
This will put the terminal in watch mode and will continue to rebundle after any changes are made to client side code.

# Executing Tailwind
To build a valid output.css file from tailwind, simply run:
```bash
npm run tailwind
```
This will place the terminal in watch mode and tailwind will continue to rebuild output.css on changes.

# Starting the Server
To start the server, run:
```bash
npm run dev
```
This uses nodemon under the hood and restarts our gin server when we make changes to our server code.

# Overview
When developing in the application, you should have 3 seperate terminals handling bundling, tailwind, and changes to the server. This will ensure your javascript and css are bundled into a single file and that your server does not need to be contantly restarted.

