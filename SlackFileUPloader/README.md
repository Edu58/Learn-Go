# SLACK FILE UPLOADER BOT
You must first have a slack app which you can create from `api.slack.com` and enable the required Oauth Scopes in the oauth & permissions tab. i.e `channels:join, channels:manage, channels:read chat:write files:read files:write`
1. Create a .env in the root folder
2. Add BOT_OAUTH_TOKEN with the value of OAuth Tokens for Your Workspace that you will get from api.slack.com/apps/<your_app_id>. The final result should be somthing like `BOT_OAUTH_TOKEN=blablablabla`
3. CHANNEL_ID with the value of the channel id that you can easily find by opening your slack app, right click on a channel -> view channel details -> scroll to the bottom.
4. Also add the bot to the channel where you want to send the files to. It should be a member.