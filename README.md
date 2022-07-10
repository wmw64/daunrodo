# 🔗 daunrōdo - self-hosted media downloader
 > Daunrōdā, ダウンローダー - means <i>downloader</i> in Japanese

Self-hosted web daemon solution tool to download original media files by extracting their direct URLs from various internet resources and social networks using plugable crawlers via API calls. More features and resources are coming by as I add them. Project started as a personal tool mainly, decided to make it public and is in early alpha stage.

Stop screenshoting, just ```daunrodo``` it!

# 🤩 Features
- 💪 **Plugin based**. Fork and easily add your own crawler through Go ```interface{}```

# 📦 Installation
Best used with Docker and Traefik
```console
git clone https://github.com/wmw64/daunrodo.git && cd daunrodo
sudo docker-compose up -d
sudo docker-compose logs -f
```

# 🔬 Basic usage 
Just add your link to the daunrodo as a path. Example: ```instagram.com/p/CfwlfpcL-li/``` -> ```daunrodo.yourdomain.org/instagram.com/p/CfwlfpcL-li/```

# 🤝 Contributing
Contributions, issues and feature requests are welcome! 👍 <br>
Feel free to check [open issues](https://github.com/rekoda-project/rekoda/issues).

## 🌟 Show your support 
Give a ⭐️ if this project helped you!

# 📝 ToDo
- [x] Instagram crawler
- [ ] Download multiple files in one request by packing it in ZIP file
- [ ] Album image hosting downloader (cyberdrop.me, gofile.io, etc)
- [ ] CLI tool to download media from terminal

# 🧠 What I Learned
- Uncle Bob's clean architecture
- Dependency injection
- Swagger

# 📑 License 
(c) 2022 Ivan Smyshlyaev. [MIT License](https://tldrlegal.com/license/mit-license)
