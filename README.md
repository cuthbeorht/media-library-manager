# Media Library Manager

Media Library Manager is an application that will take in a Dropbox account and root path, and start recursively walking through all the files inside.

## Clone the project

```bash
git clone https://github.com/cuthbeorht/media-library-manager.git
cd media-library-manager
```

## Setup configurations in .ENV file

If using Dropbox as where you want to search for media, create an application using your Dropbox account. Start [here](https://www.dropbox.com/developers/apps?_tk=pilot_lp&_ad=topbar4&_camp=myapps).

### Create

Create a `.env` file in the root of the project and add the following info:

```.env
CLIENT_ID=<client_id from Dropbox application>
CLIENT_SECRET=<secret from Dropbox application>
SCOPES=files.metadata.read,account_info.read,sharing.read
AUTH_URL=https://www.dropbox.com/oauth2/authorize
TOKEN_URL=https://api.dropboxapi.com/oauth2/token
```

