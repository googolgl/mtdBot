# mtdBot #

Version: 3.0.0

DiscordBot for Minetest.

# License of source code: AGPL v.3
Copyright (C) 2019 <googolgl@gmail.com>

# Notice:
Transmitting messages from the general chat Minetest in the Discord channel and back.
Support command:

***!status*** - Server’s Minetest version, time the server is running in seconds list of connected players and the message of the day.<br>
***!msg playername message*** - Send a private message message to playername.<br>
***!ban playername*** - Ban IP of player.<br>
***!unban playername/IP address*** - Remove ban of player with the specified name or IP address.<br>
***!kick playername reason*** - Kicks the player with the name playername. Optionally a reason can be provided in text-form. This text is also shown to the kicked player.<br>
***!grant playername priv1 priv2 ...*** - Gives the privilege to playername.<br>
***!revoke playername priv1 priv2 ...*** - Takes away a privilege from playername.<br>
***!privs playername*** - Show privs of player.

# Installing
Download the latest released version and build it:
```
git clone https://github.com/googolgl/mtdBot.git
cd mtdBot
go build
```
Install mod "lp_api" in the minetest.

```
1. > git clone https://github.com/googolgl/lp_api
2. Copy ~/lp_api > ~/minetest/mods/lp_api
3. Add this mods to trusted_mods.
--- Open : /minetest/minetest.confg
--- Add : secure.http_mods = lp_api
4. Activate lp_api mod in menu
```

# Usage
Add your bot.
```
Open https://discordapp.com/developers/applications/me
--- Click "New App"
--- Add "App Name" (Example: minetest-bot)
--- Push "Create App"
--- Push "Create a Bot User"
--- "Token:click to reveal", this is your TokenBot
--- Push "Save changes"
```
Edit config.yaml

Run on the server:
```
mtdBot
```

# Links:
GitHub:
- https://github.com/googolgl/mtdBot.git

Discord:
- https://discordapp.com/
