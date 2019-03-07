# mtdBot #

Version: 3.0.0

DiscordBot for Minetest.

# License of source code: AGPL v.3
Copyright (C) 2019 Litvinoff Pavel <googolgl@gmail.com>

# Notice:
Transmitting messages from the general chat Minetest in the Discord channel and back.
Support command:

***!status*** - Server’s Minetest version, time the server is running in seconds list of connected players and the message of the day.

***!msg <playername> <message>*** - Send a private message <message> to <player>.

***!ban <playername>*** - Ban IP of player.

***!unban <playername/IP address>*** - Remove ban of player with the specified name or IP address.

***!kick <playername> <reason>*** - Kicks the player with the name <player name>. Optionally a <reason> can be provided in text-form. This text is also shown to the kicked player.

***!grant <playername> <priv1> <priv2> ...*** - Gives the <privilege> to <player>.

***!revoke <playername> <priv1> <priv2> ...*** - Takes away a <privilege> from <player>.

***!privs <playername>*** - Show privs of player.

# Installing
Download the latest released version and build it:
```
git clone https://github.com/googolgl/mtdBot.git
cd mtdBot
go build -o mtdBot
```
Install mod "lp_api" in the minetest.

```
1. Copy ~/mtdBot/minetest/mods/lp_api > ~/minetest/mods/lp_api
2. Add this mods to trusted_mods.
--- Open : /minetest/minetest.confg
--- Add : secure.http_mods = lp_api
3. Activate lp_api mod in menu
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
Edit config.ini

Run on the server:
```
mtdBot
```

# Links:
GitHub:
- https://github.com/googolgl/mtdBot.git

Discord:
- https://discordapp.com/
