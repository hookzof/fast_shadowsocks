[Russian version](README.md)

# fast_shadowsocks
Fast deployment of the Shadowsocks server.

## Working with the script
### Downloading an executable file
```bash
curl -L -o install https://git.io/JegXf && chmod +x install
```

### Quick installation
```bash
./install
```
A random port from 49152 to 65535 and a 10-digit password are generated.

### Output
Two fields appear **Plain** and **Encoded**.  
The **Plain** field specifies the parameters of connection to the server - `ss://method:password@hostname:port`  

For <a href="https://getoutline.org/" target="_blank">Outline</a> you need a link from the **Encoded** field.

### Server key
Saved after installation in the ss-key.txt file, for output:
`cat ss-key.txt`

### Server configuration file
File path - `/etc/shadowsocks-libev/config.json`  
When changing, you should reboot the server with the command:  
`service shadowsocks-libev restart`
