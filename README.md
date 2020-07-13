# One Song a Day
**HEROKU DEPLOYMENT VERSION**
Tired of listening to the same set of songs everyday?  With this tool, you will be sent a new song to listen to daily.

## Usage
Use heroku to set up your heroku project

Set the following environmental variables in heroku to ensure proper usage
```
CLIENT_ID       // your client id from your created Spotify application
CLIENT_SECRET   // your client secret from your created Spotify application
CURRENT_USER    // your name, to personalize the email
PERSONAL_EMAIL  // email where you would like to recieve these songs
SONG_EMAIL      // email you will use to send the song emails
SONG_EMAIL_PASS // password of the song email
```

Go to http://kaffeine.herokuapp.com and place your heroku url.  For bedtime, search when 9:00 AM in your timezone is in GMT (since our app will run at 8:00 AM we are okay with it sleeping from 9:00 AM - 3:00 PM)
Kaffeine will send requests every 30 minutes from when you submit the form on the kaffeine website.  So if you would like a more persise time range, go to `server.go` and on lines 25-26, hardcode the stricter timeframe based on the Kaffeine call time.  

## Notes
The application is defaulted to send a message at 8:00 AM your local time every 24 hours.


