import discord

TOKEN = "Njk2MTYwODMyNTM0NDc4OTEw.XokvDg.mgZfq3As01ivc2uqf0DXxO5iSe0"

# login

# create channel

# add users to it

# destroy


class MyClient(discord.Client):
    async def on_ready(self):
        print("Logged on as", self.user)

    async def on_message(self, message):
        # don't respond to ourselves
        if message.author == self.user:
            return

        if message.content == "ping":
            await message.channel.send("pong")
            guild = self.get_guild(696097291945377937)
            await guild.create_text_channel("test-channel")


client = MyClient()
client.run(TOKEN)
